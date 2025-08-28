// Copyright 2025 uniperm Author. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package permission

import "github.com/go-the-way/unilog"

func (req *AddReq) LogName() (name string)    { return "新增权限" }
func (req *UpdateReq) LogName() (name string) { return "修改权限" }
func (req *DeleteReq) LogName() (name string) { return "删除权限" }

func (req *AddReq) LogFields() (fields unilog.FieldSlice)    { return unilog.GetFields(req) }
func (req *UpdateReq) LogFields() (fields unilog.FieldSlice) { return unilog.GetFields(req) }
func (req *DeleteReq) LogFields() (fields unilog.FieldSlice) { return unilog.GetFields(req) }
