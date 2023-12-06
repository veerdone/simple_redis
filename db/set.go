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
	"github.com/veerdone/simple_redis/util"
)

type HashSet struct {
	inner map[*[]byte]struct{}
}

func NewHashSet() *HashSet {
	return &HashSet{
		inner: make(map[*[]byte]struct{}),
	}
}

func (h *HashSet) Add(items ...[]byte) {
	for _, item := range items {
		h.inner[&item] = struct{}{}
	}
}

func (h *HashSet) Remove(items ...[]byte) {
	removeItems := make([]*[]byte, 0, len(items))
	for _, item := range items {
		if removeItem := h.find(item); removeItem != nil {
			removeItems = append(removeItems, removeItem)
		}
	}

	for i := range removeItems {
		delete(h.inner, removeItems[i])
	}
}

func (h *HashSet) Contains(item []byte) bool {
	return h.find(item) != nil
}

func (h *HashSet) Values() [][]byte {
	result := make([][]byte, 0, len(h.inner))
	for k := range h.inner {
		result = append(result, *k)
	}

	return result
}

func (h *HashSet) find(item []byte) *[]byte {
	for k := range h.inner {
		if util.EqualBytes(*k, item) {
			return k
		}
	}

	return nil
}

func (h *HashSet) Len() int {
	return len(h.inner)
}