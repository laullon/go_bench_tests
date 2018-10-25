// Copyright 2018 VMware, Inc.
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

func newQueueFix(size int) Queue {
	queue := QueueFix{content: make([]interface{}, size)}
	return queue
}

type QueueFix struct {
	content   []interface{}
	readHead  int
	writeHead int
	len       int
}

func (q QueueFix) Content() []interface{} {
	return q.content
}

func (q QueueFix) Count() int {
	return q.len
}

func (q QueueFix) IsFull() bool {
	return q.len >= len(q.content)
}

func (q QueueFix) Push(e interface{}) bool {
	if q.len >= len(q.content) {
		return false
	}
	q.content[q.writeHead] = e
	q.writeHead = (q.writeHead + 1) % len(q.content)
	q.len++
	return true
}

func (q QueueFix) Pop() (interface{}, bool) {
	if q.len <= 0 {
		return nil, false
	}
	result := q.content[q.readHead]
	q.content[q.readHead] = nil
	q.readHead = (q.readHead + 1) % len(q.content)
	q.len--
	return result, true
}

func (q QueueFix) Peek() (interface{}, bool) {
	if q.len <= 0 {
		return nil, true
	}
	result := q.content[q.readHead]
	return result, false
}
