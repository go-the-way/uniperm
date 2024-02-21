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
	User        = UnipermUser
	UnipermUser struct {
		Id          uint   `gorm:"column:id;type:uint;primaryKey;autoIncrement:true;comment:用户id" json:"id"`                      // 用户id
		UserName    string `gorm:"column:user_name;type:varchar(20);not null;default:'';comment:用户名;unique" json:"user_name"`     // 用户名
		Password    string `gorm:"column:password;type:varchar(32);not null;default:'';comment:密码（MD5加密）" json:"-"`               // 密码（MD5加密）
		BusinessId1 string `gorm:"column:business1;type:varchar(50);not null;default:'';comment:业务id1;index" json:"Business_id1"` // 业务id1
		BusinessId2 string `gorm:"column:business2;type:varchar(50);not null;default:'';comment:业务id2;index" json:"Business_id2"` // 业务id2
		BusinessId3 string `gorm:"column:business3;type:varchar(50);not null;default:'';comment:业务id3;index" json:"Business_id3"` // 业务id3
		RoleId      uint   `gorm:"column:role_id;type:uint;not null;default:0;comment:角色ID;index" json:"role_id"`                 // 角色id
		State       byte   `gorm:"column:state;type:tinyint;not null;default:1;comment:状态：1启用 2禁用" json:"enable"`                 // 状态：1启用 2禁用
		Remark1     string `gorm:"column:remark1;type:varchar(200);not null;default:'';comment:备注1" json:"remark1"`               // 备注1
		Remark2     string `gorm:"column:remark2;type:varchar(200);not null;default:'';comment:备注2" json:"remark2"`               // 备注2
		Remark3     string `gorm:"column:remark3;type:varchar(200);not null;default:'';comment:备注3" json:"remark3"`               // 备注3
		LoginTime   string `gorm:"column:login_time;type:varchar(20);not null;default:'';comment:登录时间" json:"login_time"`         // 登录时间
		LoginIp     string `gorm:"column:login_ip;type:varchar(20);not null;default:'';comment:登录ip" json:"login_ip"`             // 登录ip
		CreateTime  string `gorm:"column:create_time;type:varchar(20);not null;default:'';comment:创建时间" json:"create_time"`       // 创建时间
		UpdateTime  string `gorm:"column:update_time;type:varchar(20);not null;default:'';comment:修改时间" json:"update_time"`       // 修改时间                                                // 修改时间
	}
)

func (u *User) SuperAdmin() bool { return u.Id == 1 }

const (
	_ = iota
	UserStateEnable
	UserStateDisable
)
