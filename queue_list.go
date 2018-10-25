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

import (
	"container/list"
)

func newQueueList(size int) Queue {
	queue := QueueList{content: list.New(), maxLen: size}
	return queue
}

type QueueList struct {
	content   *list.List
	readHead  int
	writeHead int
	maxLen    int
}

func (q QueueList) Content() []interface{} {
	c := make([]interface{}, q.content.Len())
	idx := 0
	for e := q.content.Front(); e != nil; e = e.Next() {
		c[idx] = e.Value
		idx++
	}
	return c
}

func (q QueueList) Count() int {
	return q.content.Len()
}

func (q QueueList) IsFull() bool {
	return q.content.Len() >= q.maxLen
}

func (q QueueList) Push(e interface{}) bool {
	if q.IsFull() {
		return false
	}
	q.content.PushBack(e)
	return true
}

func (q QueueList) Pop() (interface{}, bool) {
	if q.content.Len() <= 0 {
		return nil, false
	}
	result := q.content.Front()
	q.content.Remove(result)
	return result.Value, true
}

func (q QueueList) Peek() (interface{}, bool) {
	if q.content.Len() <= 0 {
		return nil, true
	}
	result := q.content.Front()
	return result.Value, false
}
