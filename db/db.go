/*
   Copyright [2023] [veerdone]

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package db

import (
	"time"
)

type DB struct {
	inner      map[string]*Entity
	expireChan chan string
	closeChan  chan struct{}
	size       int
}

func NewDB() *DB {
	return &DB{
		inner:      make(map[string]*Entity),
		expireChan: make(chan string, 1000),
		closeChan:  make(chan struct{}),
	}
}

func (d *DB) Get(key string) *Entity {
	entity := d.inner[key]
	if entity == nil {
		return nil
	}
	if entity.ExpireAt != 0 && entity.ExpireAt < time.Now().UnixMilli() {
		d.Del(key)
		entity = nil
	}

	return entity
}

func (d *DB) Set(key string, entity *Entity) {
	d.inner[key] = entity
}

func (d *DB) Del(key string) {
	if d.Get(key) != nil {
		d.size -= 1
		d.inner[key] = nil
		delete(d.inner, key)
	}
}

func (d *DB) DelExpireKey() {
	for {
		select {
		case key := <- d.expireChan:
			d.Del(key)
		default:
			return
		}
	}
}

func (d *DB) Size() int {
	return d.size
}

func (d *DB) PurgePeriodically() {
	for {
		select {
		case <-d.closeChan:
			return
		default:
			now := time.Now().UnixMilli()
			for key, value := range d.inner {
				if value.ExpireAt < now {
					d.expireChan <- key
				}
			}
			time.Sleep(time.Millisecond * 10)
		}
	}
}

func (d *DB) Close() {
	d.closeChan <- struct{}{}
	close(d.closeChan)
	close(d.expireChan)
}
