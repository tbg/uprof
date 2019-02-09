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
	"strconv"

	"github.com/cockroachdb/cockroach/pkg/util/humanizeutil"
)

type UnitFunc func(int64) string

func (f UnitFunc) ToString(n int64) string {
	return f(n)
}

func Count() Unit {
	return UnitFunc(func(n int64) string {
		return strconv.FormatInt(n, 10)
	})
}

func Bytes() Unit {
	return UnitFunc(func(n int64) string {
		return humanizeutil.IBytes(n)
	})
}

type Unit interface {
	ToString(int64) string
}
