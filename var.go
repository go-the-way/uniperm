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

package uniperm

import (
	"github.com/rwscode/uniperm/services/permission"
	"github.com/rwscode/uniperm/services/role"
	"github.com/rwscode/uniperm/services/user"
)

var (
	UserService        = user.Service
	UserGetPage        = user.GetPage
	UserGet            = user.Get
	UserGetPerm        = user.GetPerm
	UserGetPermButton  = user.GetPermButton
	UserAdd            = user.Add
	UserUpdate         = user.Update
	UserUpdatePassword = user.UpdatePassword
	UserUpdateRole     = user.UpdateRole
	UserDel            = user.Del
	UserEnable         = user.Enable
	UserDisable        = user.Disable
	UserLogin          = user.Login
	UserLogout         = user.Logout
)

var (
	RoleService    = role.Service
	RoleGetPage    = role.GetPage
	RoleGet        = role.Get
	RoleGetPerm    = role.GetPerm
	RoleUpdatePerm = role.UpdatePerm
	RoleAdd        = role.Add
	RoleUpdate     = role.Update
	RoleDel        = role.Del
	RoleEnable     = role.Enable
	RoleDisable    = role.Disable
)

var (
	PermissionService = permission.Service
	permissionTree    = permission.Tree
	permissionAdd     = permission.Add
	permissionUpdate  = permission.Update
	permissionDel     = permission.Del
)
