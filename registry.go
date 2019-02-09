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
	"sync"
)

type Registry struct {
	types map[fmt.Stringer]Unit
}

func NewRegistry() *Registry {
	return &Registry{
		types: map[fmt.Stringer]Unit{},
	}
}

func (r *Registry) RegisterProfileType(typ fmt.Stringer, unit Unit) {
	if _, ok := r.types[typ]; ok {
		panic("already registered")
	}
	r.types[typ] = unit
}

func (r *Registry) Export(p *Profile) Export {
	var e []ExportedMeasurement
	for typ, unit := range r.types {
		v, ok := p.m.Load(typ)
		if ok {
			v.(*sync.Map).Range(func(k, v interface{}) bool {
				em := ExportedMeasurement{
					Name: typ.String(),
					Unit: unit,
				}
				pcs := k.([32]uintptr)
				var l int
				for l = 0; l < len(pcs[:]); l++ {
					if pcs[l] == 0 {
						break
					}
				}
				em.PCS = pcs[:l]
				em.Amount = *(*int64)(v.(*Measurement))
				e = append(e, em)
				return true // want more
			})
		}
	}
	return e
}
