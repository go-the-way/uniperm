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
	Role        = UnipermRole
	UnipermRole struct {
		Id          uint   `gorm:"column:id;type:uint;primaryKey;autoIncrement:true;comment:角色id" json:"id"`                 // 角色id
		Name        string `gorm:"column:name;type:varchar(50);not null;default:'';comment:角色名称" json:"name"`                // 角色名称
		Description string `gorm:"column:description;type:varchar(200);not null;default:'';comment:角色描述" json:"description"` // 角色描述
		Type        string `gorm:"column:type;type:varchar(50);not null;default:'';comment:角色类型;index" json:"type"`          // 角色类型
		State       byte   `gorm:"column:state;type:tinyint;not null;default:1;comment:状态：1启用 2禁用;index" json:"state"`       // 状态：1启用 2禁用
		CreateTime  string `gorm:"column:create_time;type:varchar(20);not null;default:'';comment:创建时间" json:"create_time"`  // 创建时间
		UpdateTime  string `gorm:"column:update_time;type:varchar(20);not null;default:'';comment:修改时间" json:"update_time"`  // 修改时间
	}
)

const (
	_ = iota
	RoleStateEnable
	RoleStateDisable
)
