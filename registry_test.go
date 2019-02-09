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
	"testing"

	"github.com/stretchr/testify/assert"
)

type CounterA struct{}

func (CounterA) String() string {
	return "CounterA"
}

type CounterB struct{}

func (CounterB) String() string {
	return "CounterB"
}

func TestRegistry(t *testing.T) {
	r := NewRegistry()
	assert.NotEqual(t, CounterA{}, CounterB{})
	r.RegisterProfileType(CounterB{}, Count())
	r.RegisterProfileType(CounterA{}, Count())

	p := NewProfile()
	for i := 0; i < 3; i++ {
		p.Add(CounterA{}, 3)
	}
	p.Add(CounterB{}, 7)

	e := r.Export(p)
	e.Sort()

	assert.Equal(t, `CounterA 9
  github.com/tbg/uprof.TestRegistry
  testing.tRunner
CounterB 7
  github.com/tbg/uprof.TestRegistry
  testing.tRunner
`, e.String())

}
