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

import (
	"github.com/go-the-way/uniperm/models"
	"github.com/go-the-way/uniperm/services/base"
)

func (r *AddReq) Check() (err error) {
	if r.ParentId > 0 {
		return base.CheckPermissionIsNotButton(r.ParentId)
	}
	return
}

func (r *UpdateReq) Check() (err error) {
	return base.CheckPermissionExist(r.Id)
}

func (r *DeleteReq) Check() (err error) {
	return base.CheckAll(
		func() (err error) { return base.CheckPermissionExist(r.Id) },
		func() (err error) { return base.CheckPermissionHaveNoSubPerms(r.Id) },
		func() (err error) { return base.CheckPermissionRefRole(r.Id) },
	)
}

func (r *AddReq) Transform() *models.Permission {
	return &models.Permission{Name: r.Name, Route: r.Route, ParentId: r.ParentId, IsButton: r.IsButton}
}

func (r *UpdateReq) Transform() *models.Permission {
	return &models.Permission{Id: r.Id, Name: r.Name, Route: r.Route}
}
