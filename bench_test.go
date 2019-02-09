// Copyright 2019 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package uprof

import (
	"fmt"
	"testing"
)

var metrics []*metric

type metric string

func (m *metric) String() string { return string(*m) }

func init() {
	for i := 0; i < 1000; i++ {
		m := metric(fmt.Sprintf("metric#%d", i))
		metrics = append(metrics, &m)
	}
}

func BenchmarkProfileAdd(b *testing.B) {
	p := NewProfile()
	for i := 0; i < b.N; i++ {
		p.Add(metrics[0], int64(i))
	}
}

func BenchmarkProfileAllocs(b *testing.B) {
	for _, nMetrics := range []int{0, 1, 10, 100, 1000} {
		r := NewRegistry()
		for i := 0; i < nMetrics; i++ {
			r.RegisterProfileType(metrics[i], Count())
		}
		b.Run(fmt.Sprintf("metrics=%d", nMetrics), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				p := NewProfile()
				for i := 0; i < nMetrics; i++ {
					p.Add(metrics[i%nMetrics], int64(100+i))
				}
			}
		})
	}
}
