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

package user

import (
	"github.com/rwscode/uniperm/deps/pkg"
	"github.com/rwscode/uniperm/models"
	"github.com/rwscode/uniperm/services/base"
)

func (r *AddReq) Check() (err error) {
	var functions []func() error
	functions = append(functions, func() (err error) { return base.CheckUsernameExists(r.Username) })
	if r.RoleId > 0 {
		functions = append(functions, func() (err error) { return base.CheckRoleExist(r.RoleId) })
	}
	return base.CheckAll(functions...)
}

func (r *UpdateReq) Check() (err error) { return base.CheckUserExists(r.Id) }

func (r *UpdatePasswordReq) Check() (err error) { return base.CheckUserExists(r.Id) }

func (r *UpdateRoleReq) Check() (err error) {
	return base.CheckAll(
		func() (err error) { return base.CheckUserIsSuper(r.Id) },
		func() (err error) { return (&DeleteReq{IdReq: IdReq{r.Id}}).Check() },
	)
}

func (r *DeleteReq) Check() (err error) {
	return base.CheckAll(
		func() (err error) { return base.CheckUserIsSuper(r.Id) },
		func() (err error) { return base.CheckUserExists(r.Id) },
	)
}

func (r *EnableReq) Check() (err error) { return (&DeleteReq{IdReq: IdReq{r.Id}}).Check() }

func (r *DisableReq) Check() (err error) { return (&DeleteReq{IdReq: IdReq{r.Id}}).Check() }

func (r *AddReq) Transform() *models.User {
	return &models.User{
		Username:    r.Username,
		Password:    pkg.MD5(r.Password),
		BusinessId1: r.BusinessId1,
		BusinessId2: r.BusinessId2,
		BusinessId3: r.BusinessId3,
		RoleId:      r.RoleId,
		State:       models.RoleStateEnable,
		Remark1:     r.Remark1,
		Remark2:     r.Remark2,
		Remark3:     r.Remark3,
		CreateTime:  pkg.TimeNowStr(),
		UpdateTime:  pkg.TimeNowStr(),
	}
}

func (r *UpdateReq) Transform() *models.User {
	return &models.User{
		BusinessId1: r.BusinessId1,
		BusinessId2: r.BusinessId2,
		BusinessId3: r.BusinessId3,
		Remark1:     r.Remark1,
		Remark2:     r.Remark2,
		Remark3:     r.Remark3,
		UpdateTime:  pkg.TimeNowStr(),
	}
}
