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

package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:8888")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// set key 1
	_, err = conn.Write([]byte{0, 0, 0, 13, 3, 115, 101, 116, 0, 0, 0, 5, 107, 101, 121, 32, 49})
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 32)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("set ok, data: ", string(buf[:n]))

	// get key
	_, err = conn.Write([]byte{0, 0, 0, 11, 3, 103, 101, 116, 0, 0, 0, 3, 107, 101, 121})
	if err != nil {
		log.Fatal(err)
	}
	n, err = conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("get ok, data: ", string(buf[:n]))

	// incr key
	_, err = conn.Write([]byte{0, 0, 0, 12, 4, 105, 110, 99, 114, 0, 0, 0, 3, 107, 101, 121})
	if err != nil {
		log.Fatal(err)
	}
	n, err = conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("incr ok, data: ", string(buf[:n]))

	// decr key
	_, err = conn.Write([]byte{0, 0, 0, 12, 4, 100, 101, 99, 114, 0, 0, 0, 3, 107, 101, 121})
	if err != nil {
		log.Fatal(err)
	}
	n, err = conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("decr ok, data: ", string(buf[:n]))
}
