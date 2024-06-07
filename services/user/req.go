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
	"github.com/rwscode/uniperm/models"
	"github.com/rwscode/uniperm/services/base"
)

type (
	GetPageReq struct {
		base.PageReq

		OrderBy string `form:"order_by" json:"order_by"` // 排序

		Id          uint   `form:"id"`           // 用户id
		Username    string `form:"username"`     // 用户名
		BusinessId1 string `form:"business_id1"` // 业务id1
		BusinessId2 string `form:"business_id2"` // 业务id2
		BusinessId3 string `form:"business_id3"` // 业务id3
		RoleId      uint   `form:"role_id"`      // 角色id
		State       byte   `form:"state"`        // 状态：1启用 2禁用
		Remark1     string `form:"remark1"`      // 备注1
		Remark2     string `form:"remark2"`      // 备注2
		Remark3     string `form:"remark3"`      // 备注3
		LoginIp     string `form:"login_ip"`     // 登录ip
		CreateTime1 string `form:"create_time1"` // 创建时间
		CreateTime2 string `form:"create_time2"` // 创建时间
		UpdateTime1 string `form:"update_time1"` // 修改时间
		UpdateTime2 string `form:"update_time2"` // 修改时间
		LoginTime1  string `form:"login_time1"`  // 登录时间
		LoginTime2  string `form:"login_time2"`  // 登录时间
	}
	IdReq struct {
		Id uint `validate:"min(1,用户id不能为空)" json:"id"`
	}
	GetReq           IdReq
	GetPermReq       = models.User
	GetPermButtonReq struct {
		models.User
		Id uint `validate:"min(1,权限id不能为空)" form:"id"`
	}
	AddReq struct {
		Username    string `validate:"minlength(1,用户名不能为空) maxlength(20,用户名长度不能超过20)" json:"username"` // 用户名
		Password    string `validate:"minlength(1,密码不能为空) maxlength(20,密码长度不能超过20)" json:"password"`   // 密码
		BusinessId1 string `validate:"maxlength(50,业务id1长度不能超过50)" json:"business_id1"`                // 业务id1
		BusinessId2 string `validate:"maxlength(50,业务id2长度不能超过50)" json:"business_id2"`                // 业务id2
		BusinessId3 string `validate:"maxlength(50,业务id3长度不能超过50)" json:"business_id3"`                // 业务id3
		RoleId      uint   `json:"role_id"`                                                            // 角色id
		Remark1     string `validate:"maxlength(200,备注1长度不能超过200)" json:"remark1"`                     // 备注1
		Remark2     string `validate:"maxlength(200,备注2长度不能超过200)" json:"remark2"`                     // 备注2
		Remark3     string `validate:"maxlength(200,备注3长度不能超过200)" json:"remark3"`                     // 备注3
		Callback    func(req AddReq)
	}
	UpdateReq struct {
		IdReq       `validate:"valid(T)"`
		BusinessId1 string `validate:"maxlength(50,业务id1长度不能超过50)" json:"business_id1"` // 业务id1
		BusinessId2 string `validate:"maxlength(50,业务id2长度不能超过50)" json:"business_id2"` // 业务id2
		BusinessId3 string `validate:"maxlength(50,业务id3长度不能超过50)" json:"business_id3"` // 业务id3
		Remark1     string `validate:"maxlength(200,备注1长度不能超过200)" json:"remark1"`      // 备注1
		Remark2     string `validate:"maxlength(200,备注2长度不能超过200)" json:"remark2"`      // 备注2
		Remark3     string `validate:"maxlength(200,备注3长度不能超过200)" json:"remark3"`      // 备注3
		Callback    func(req UpdateReq)
	}
	UpdatePasswordReq struct {
		IdReq    `validate:"valid(T)"`
		Password string `validate:"minlength(1,密码不能为空) maxlength(20,密码长度不能超过20)" json:"password"` // 密码
		Callback func(req UpdatePasswordReq)
	}
	UpdateRoleReq struct {
		IdReq    `validate:"valid(T)"`
		RoleId   uint `validate:"min(1,角色id不能为空)" json:"role_id"` // 角色id
		Callback func(req UpdateRoleReq)
	}
	DeleteReq struct {
		IdReq
		Callback func(req DeleteReq)
	}
	EnableReq struct {
		IdReq
		Callback func(req EnableReq)
	}
	DisableReq struct {
		IdReq
		Callback func(req DisableReq)
	}
	LoginReqCallback struct {
		NotFound      func(req LoginReq)
		Disabled      func(req LoginReq, user models.User)
		RoleDisabled  func(req LoginReq, user models.User)
		PasswordWrong func(req LoginReq, user models.User)
		Success       func(req LoginReq, user models.User, resp *LoginResp)
	}
	LoginReq struct {
		Username string `validate:"minlength(1,用户名不能为空) maxlength(20,用户名长度不能超过20)" json:"username"`
		Password string `validate:"minlength(1,密码不能为空) maxlength(20,密码长度不能超过20)" json:"password"`
		ClientIp string `validate:"maxlength(20,客户端ip长度不能超过20)" json:"client_ip"`
		Callback LoginReqCallback
	}
	LogoutReq struct {
		Token    string `validate:"minlength(1,token不能为空)" form:"token" json:"token"`
		Callback func(req LogoutReq) (err error)
	}
)
