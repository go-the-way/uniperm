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

package models

type (
	Permission        = UnipermPermission
	UnipermPermission struct {
		Id       uint   `gorm:"column:id;type:uint;primaryKey;autoIncrement:true;comment:权限id" json:"id"`             // 权限id
		Name     string `gorm:"column:name;type:varchar(50);not null;comment:权限名称" json:"name"`                       // 权限名称
		Route    string `gorm:"column:route;type:varchar(200);not null;comment:权限路由" json:"route"`                    // 权限路由
		ParentId uint   `gorm:"column:parent_id;type:uint;not null;default:0;comment:上级权限id;index" json:"parent_id"`  // 上级权限id
		IsButton byte   `gorm:"column:is_button;type:tinyint;not null;default:2;comment:是否按钮 1是 2否" json:"is_button"` // 是否按钮 1是 2否
	}
)

const (
	_ = iota
	PermissionIsButtonYes
	PermissionIsButtonNo
)
