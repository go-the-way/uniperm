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

package permission

type (
	TreeReq struct{ PermissionId []uint }
	IdReq   struct {
		Id uint `validate:"min(1,权限id不能为空)" json:"id"`
	}
	GetReq IdReq
	AddReq struct {
		Name     string `validate:"minlength(1,权限名称不能为空) maxlength(50,权限名称长度不能超过50)" json:"name"`
		Route    string `validate:"minlength(1,权限路由不能为空) maxlength(200,权限路由长度不能超过200)" json:"route"`
		IsButton byte   `validate:"enum(1|2,是否状态不合法)" json:"is_button"`
		ParentId uint   `json:"parent_id"`
		Callback func(req AddReq)
	}
	UpdateReq struct {
		IdReq    `validate:"valid(T)"`
		Name     string `validate:"minlength(1,权限名称不能为空) maxlength(50,权限名称长度不能超过50)" json:"name"`
		Route    string `validate:"minlength(1,权限路由不能为空) maxlength(200,权限路由长度不能超过200)" json:"route"`
		Callback func(req UpdateReq)
	}
	DeleteReq struct {
		IdReq
		Callback func(req DeleteReq)
	}
)
