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

import "github.com/rwscode/uniperm/models"

type (
	GetPageResp struct {
		Total int64         `json:"total"`
		List  []models.User `json:"list"`
	}
	GetResp     struct{ models.User }
	GetPermResp struct {
		SuperAdmin bool               `json:"super_admin"` // 超级管理员
		Routes     []GetPermRespRoute `json:"routes,omitempty"`
	}
	GetPermRespRoute struct {
		Id       uint               `json:"id"`
		Path     string             `json:"path"`
		Children []GetPermRespRoute `json:"children,omitempty"`
	}
	GetPermButtonResp struct {
		SuperAdmin bool     `json:"super_admin"`       // 超级管理员
		Buttons    []string `json:"buttons,omitempty"` // Button key
	}
	LoginResp struct {
		Token string `json:"token"` // Login token
	}
)
