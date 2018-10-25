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

func newQueueSlice(size int) Queue {
	queue := QueueSlice{content: make([]interface{}, 0), maxLen: size}
	return queue
}

type QueueSlice struct {
	content   []interface{}
	readHead  int
	writeHead int
	maxLen    int
}

func (q QueueSlice) Content() []interface{} {
	return q.content
}

func (q QueueSlice) Count() int {
	return len(q.content)
}

func (q QueueSlice) IsFull() bool {
	return q.Count() >= q.maxLen
}

func (q QueueSlice) Push(e interface{}) bool {
	if q.IsFull() {
		return false
	}
	q.content = append(q.content, e)
	return true
}

func (q QueueSlice) Pop() (interface{}, bool) {
	if q.Count() <= 0 {
		return nil, false
	}
	result := q.content[0]
	q.content = q.content[1:]
	return result, true
}

func (q QueueSlice) Peek() (interface{}, bool) {
	if q.Count() <= 0 {
		return nil, true
	}
	return q.content[0], false
}
