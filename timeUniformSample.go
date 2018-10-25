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
	"fmt"
	"sync"
	"time"

	metrics "github.com/rcrowley/go-metrics"
)

type TimeUniformSample struct {
	mutex    sync.Mutex
	values   Queue
	lifetime time.Duration
}

// // Clear clears all samples.
// func (s *TimeUniformSample) Clear() {
// 	s.mutex.Lock()
// 	defer s.mutex.Unlock()
// 	s.values = Queue{}
// }

// Count returns the number of samples recorded, which may exceed the
// reservoir size.
func (s *TimeUniformSample) Count() int64 {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return int64(s.values.Count())
}

// Max returns the maximum value in the sample, which may not be the maximum
// value ever to be part of the sample.
func (s *TimeUniformSample) Max() int64 {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return metrics.SampleMax(s.rawValues())
}

// Mean returns the mean of the values in the sample.
func (s *TimeUniformSample) Mean() float64 {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return metrics.SampleMean(s.rawValues())
}

// Min returns the minimum value in the sample, which may not be the minimum
// value ever to be part of the sample.
func (s *TimeUniformSample) Min() int64 {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return metrics.SampleMin(s.rawValues())
}

// Percentile returns an arbitrary percentile of values in the sample.
func (s *TimeUniformSample) Percentile(p float64) float64 {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return metrics.SamplePercentile(s.rawValues(), p)
}

// Percentiles returns a slice of arbitrary percentiles of values in the
// sample.
func (s *TimeUniformSample) Percentiles(ps []float64) []float64 {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return metrics.SamplePercentiles(s.rawValues(), ps)
}

// Size returns the size of the sample, which is at most the reservoir size.
func (s *TimeUniformSample) Size() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return len(s.rawValues())
}

// Snapshot returns a read-only copy of the sample.
func (s *TimeUniformSample) Snapshot() metrics.Sample {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	values := make([]int64, s.values.Count())
	copy(values, s.rawValues())
	return metrics.NewSampleSnapshot(int64(s.values.Count()), values)
}

// StdDev returns the standard deviation of the values in the sample.
func (s *TimeUniformSample) StdDev() float64 {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return metrics.SampleStdDev(s.rawValues())
}

// Sum returns the sum of the values in the sample.
func (s *TimeUniformSample) Sum() int64 {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return metrics.SampleSum(s.rawValues())
}

// Update samples a new value.
func (s *TimeUniformSample) Update(v int64) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.cleanOldValues()
	if s.values.IsFull() {
		s.values.Pop()
	}
	ok := s.values.Push(&sampleValue{v: v, time: time.Now()})
	if !ok {
		panic(fmt.Sprint("------"))
	}
}

func (s *TimeUniformSample) cleanOldValues() {
	now := time.Now()
	var needPop bool
	for more := true; more; more = needPop {
		sample, empty := s.values.Peek()
		if empty {
			needPop = false
		} else {
			needPop = (now.Sub(sample.(*sampleValue).time).Seconds() > s.lifetime.Seconds())
			if needPop {
				s.values.Pop()
			}
		}
	}
}

// Values returns a copy of the values in the sample.
func (s *TimeUniformSample) Values() []int64 {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.rawValues()
}

func (s *TimeUniformSample) rawValues() []int64 {
	values := make([]int64, s.values.Count())
	idx := 0
	for _, value := range s.values.Content() {
		if value != nil {
			values[idx] = value.(*sampleValue).v
			idx++
		}
	}
	return values
}

// Variance returns the variance of the values in the sample.
func (s *TimeUniformSample) Variance() float64 {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return metrics.SampleVariance(s.rawValues())
}

type sampleValue struct {
	v    int64
	time time.Time
}
