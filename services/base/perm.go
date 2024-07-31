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

package base

import (
	"github.com/go-the-way/uniperm/deps/db"
	"github.com/go-the-way/uniperm/models"
)

type (
	TreeResp struct {
		List []TreeRespPerm `json:"list"`
	}
	TreeRespPerm struct {
		models.Permission
		Children []TreeRespPerm `json:"children,omitempty"`
		Check    string         `json:"check"`
	}
)

func PermTree(permissionId []uint) (resp TreeResp, err error) {
	var checkedMap = map[uint]struct{}{}
	for _, perm := range permissionId {
		checkedMap[perm] = struct{}{}
	}
	var perms []models.Permission
	if err = db.GetDb().Model(&models.Permission{}).Find(&perms).Error; err != nil {
		return
	}
	resp.List = make([]TreeRespPerm, 0)
	for _, perm := range perms {
		if perm.ParentId == 0 {
			resp.List = append(resp.List, TreeRespPerm{perm, children(perm.Id, checkedMap, perms), getChecked(checkedMap, perm.Id)})
		}
	}
	return
}

func getChecked(checkedMap map[uint]struct{}, permissionId uint) string {
	if _, checked := checkedMap[permissionId]; checked {
		return "true"
	}
	return "false"
}

func children(parentId uint, checkedMap map[uint]struct{}, perms []models.Permission) []TreeRespPerm {
	var trpS []TreeRespPerm
	for _, perm := range perms {
		if perm.ParentId == parentId {
			trpS = append(trpS, TreeRespPerm{perm, children(perm.Id, checkedMap, perms), getChecked(checkedMap, perm.Id)})
		}
	}
	return trpS
}
