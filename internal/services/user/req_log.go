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

package user

import ul "github.com/go-the-way/unilog"

func (req *AddReq) LogName() (name string)            { return "新增用户" }
func (req *UpdateReq) LogName() (name string)         { return "修改用户" }
func (req *UpdatePasswordReq) LogName() (name string) { return "修改用户密码" }
func (req *UpdateRoleReq) LogName() (name string)     { return "修改用户角色" }
func (req *DeleteReq) LogName() (name string)         { return "删除用户" }
func (req *EnableReq) LogName() (name string)         { return "启用用户" }
func (req *DisableReq) LogName() (name string)        { return "禁用用户" }
func (req *LoginReq) LogName() (name string)          { return "用户登录" }

func (req *AddReq) LogFields() (fields ul.FieldSlice)            { return ul.GetFields(req) }
func (req *UpdateReq) LogFields() (fields ul.FieldSlice)         { return ul.GetFields(req) }
func (req *UpdatePasswordReq) LogFields() (fields ul.FieldSlice) { return ul.GetFields(req) }
func (req *UpdateRoleReq) LogFields() (fields ul.FieldSlice)     { return ul.GetFields(req) }
func (req *DeleteReq) LogFields() (fields ul.FieldSlice)         { return ul.GetFields(req) }
func (req *EnableReq) LogFields() (fields ul.FieldSlice)         { return ul.GetFields(req) }
func (req *DisableReq) LogFields() (fields ul.FieldSlice)        { return ul.GetFields(req) }
func (req *LoginReq) LogFields() (fields ul.FieldSlice)          { return ul.GetFields(req) }
