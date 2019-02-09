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
	"sync"
)

type Profile struct {
	m sync.Map // sync.Map[typ]sync.Map[stack]*int64
}

func NewProfile() *Profile {
	return &Profile{}
}

func (p *Profile) Add(typ interface{}, n int64) {
	if p == nil {
		return
	}

	depth := 1

	v, ok := p.m.Load(typ)
	if !ok {
		v, _ = p.m.LoadOrStore(typ, &sync.Map{})
	}
	m := v.(*sync.Map) // map[stack]*Counter

	var stk [32]uintptr // TODO(tbg) escapes, use pool
	if l := runtime.Callers(depth+1, stk[:]); l > 32 {
		// TODO(tbg): handle case of >32 frames.
		panic("unsupported")
	}

	v, ok = m.Load(stk)
	if !ok {
		v, _ = m.LoadOrStore(stk, new(Measurement))
	}
	v.(*Measurement).Add(n)
}
