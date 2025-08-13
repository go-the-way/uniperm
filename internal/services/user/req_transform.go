// Copyright 2025 uniperm Author. All Rights Reserved.
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
	"github.com/go-the-way/uniperm/internal/models"
	"github.com/go-the-way/uniperm/internal/pkg"
)

func (req *AddReq) transform() *models.User {
	return &models.User{
		Username:    req.Username,
		Password:    pkg.MD5(req.Password),
		BusinessId1: req.BusinessId1,
		BusinessId2: req.BusinessId2,
		BusinessId3: req.BusinessId3,
		RoleId:      req.RoleId,
		State:       models.RoleStateEnable,
		Remark1:     req.Remark1,
		Remark2:     req.Remark2,
		Remark3:     req.Remark3,
		CreateTime:  pkg.TimeNowStr(),
		UpdateTime:  pkg.TimeNowStr(),
	}
}

func (req *UpdateReq) transform() map[string]any {
	return map[string]any{
		"business_id1": req.BusinessId1,
		"business_id2": req.BusinessId2,
		"business_id3": req.BusinessId3,
		"remark1":      req.Remark1,
		"remark2":      req.Remark2,
		"remark3":      req.Remark3,
		"update_time":  pkg.TimeNowStr(),
	}
}
