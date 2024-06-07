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

package role

import (
	"github.com/rwscode/uniperm/deps/pkg"
	"github.com/rwscode/uniperm/models"
	"github.com/rwscode/uniperm/services/base"
)

func (r *AddReq) Check() (err error) { return }

func (r *UpdateReq) Check() (err error) { return base.CheckRoleExist(r.Id) }

func (r *DeleteReq) Check() (err error) {
	if err = base.CheckRoleExist(r.Id); err != nil {
		return
	}
	if err = base.CheckRoleRefUser(r.Id); err != nil {
		return
	}
	return base.CheckRoleRefPermission(r.Id)
}

func (r *EnableReq) Check() (err error)     { return base.CheckRoleExist(r.Id) }
func (r *DisableReq) Check() (err error)    { return base.CheckRoleExist(r.Id) }
func (r *GetPermReq) Check() (err error)    { return base.CheckRoleExist(r.Id) }
func (r *UpdatePermReq) Check() (err error) { return base.CheckRoleExist(r.Id) }

func (r *AddReq) Transform() *models.Role {
	return &models.Role{
		Name:        r.Name,
		Description: r.Description,
		Type:        r.Type,
		State:       models.RoleStateEnable,
		CreateTime:  pkg.TimeNowStr(),
		UpdateTime:  pkg.TimeNowStr(),
	}
}

func (r *UpdateReq) Transform() *models.Role {
	m := r.AddReq.Transform()
	m.Id = r.Id
	return m
}
