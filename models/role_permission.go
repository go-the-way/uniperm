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
	RolePermission        = UnipermRolePermission
	UnipermRolePermission struct {
		Id           uint `gorm:"column:id;type:uint;primaryKey;autoIncrement:true;comment:id" json:"id"`                    // id
		RoleId       uint `gorm:"column:role_id;type:uint;not null;default:0;comment:角色id;index" json:"role_id"`             // 角色id
		PermissionId uint `gorm:"column:permission_id;type:uint;not null;default:0;comment:权限id;index" json:"permission_id"` // 权限id
	}
)
