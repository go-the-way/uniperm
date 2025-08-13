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

package role

import (
	"github.com/go-the-way/uniperm/internal/models"
	"github.com/go-the-way/uniperm/internal/pkg"
)

func (req *AddReq) transform() *models.Role {
	return &models.Role{
		Name:        req.Name,
		Description: req.Description,
		Type:        req.Type,
		State:       models.RoleStateEnable,
		CreateTime:  pkg.TimeNowStr(),
		UpdateTime:  pkg.TimeNowStr(),
	}
}

func (req *UpdateReq) transform() map[string]any {
	return map[string]any{
		"name":        req.Name,
		"description": req.Description,
		"type":        req.Type,
		"update_time": pkg.TimeNowStr(),
	}
}
