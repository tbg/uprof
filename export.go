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
	"runtime"
	"sort"
	"strings"
)

type ExportedMeasurement struct {
	Name   string
	Unit   Unit
	Amount int64
	PCS    []uintptr
}

func (em *ExportedMeasurement) String() string {
	var buf strings.Builder
	buf.WriteString(em.Name)
	buf.WriteByte(' ')
	buf.WriteString(em.Unit.ToString(em.Amount))
	buf.WriteByte('\n')
	frames := runtime.CallersFrames(em.PCS)
	for {
		frame, more := frames.Next()
		if !more {
			break
		}
		buf.WriteString("  ")
		buf.WriteString(frame.Function)
		buf.WriteByte('\n')
	}
	return buf.String()
}

type Export []ExportedMeasurement

func (e Export) Sort() {
	sort.Slice(e, func(i, j int) bool {
		if e[i].Name != e[j].Name {
			return e[i].Name < e[j].Name
		}
		if l, r := len(e[i].PCS), len(e[j].PCS); l < r {
			return true
		} else if r < l {
			return false
		}
		for k := range e[i].PCS {
			if l, r := e[i].PCS[k], e[j].PCS[k]; l < r {
				return true
			} else if r < l {
				return false
			}
		}
		return false // equal
	})
}

func (e Export) String() string {
	var buf strings.Builder
	for _, em := range e {
		buf.WriteString(em.String())
	}
	return buf.String()
}
