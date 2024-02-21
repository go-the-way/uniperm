// Copyright 2024 uniperm Author. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pkg

import (
	"fmt"
	"time"
)

type number interface{ uint | byte }

func IfFunc(ok bool, fn func()) {
	if ok {
		fn()
	}
}

func IfGt0Func[T number](n T, fn func())   { IfFunc(n > 0, fn) }
func IfNotEmptyFunc(str string, fn func()) { IfFunc(str != "", fn) }
func TimeNow() time.Time                   { return time.Now() }
func TimeNowStr() string                   { return TimeNow().Format("2006-01-02 15:04:05") }
func TimeNowNumStr() string                { return TimeNow().Format("20060102150405") }
func TimeNowStamp() string                 { return fmt.Sprintf("%d", TimeNow().Unix()) }
func TimeNowStampLong() string             { return fmt.Sprintf("%d", TimeNow().UnixMilli()) }
