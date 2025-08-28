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

package base

import (
	"github.com/go-the-way/uniperm/internal/db"
	"github.com/go-the-way/uniperm/internal/models"
)

// TreeResp represents a response structure for a permission tree.
type TreeResp struct {
	List []TreeRespPerm `json:"list"` // List contains the hierarchical list of permissions.
}

// TreeRespPerm represents a single permission node in the tree response.
type TreeRespPerm struct {
	models.Permission
	Children []TreeRespPerm `json:"children,omitempty"` // Children contains the sub-permissions of the current permission.
	Check    string         `json:"check"`              // Check indicates whether the permission is selected ("true" or "false").
}

// PermTree constructs a permission tree from the database, marking specified permissions as checked.
func PermTree(permissionId []uint) (resp TreeResp, err error) {
	// Create a map to track checked permissions
	var checkedMap = map[uint]struct{}{}
	for _, perm := range permissionId {
		checkedMap[perm] = struct{}{}
	}

	// Retrieve all permissions from the database
	var perms []models.Permission
	if err = db.GetDB().Model(&models.Permission{}).Find(&perms).Error; err != nil {
		return
	}

	// Initialize the response list
	resp.List = make([]TreeRespPerm, 0)

	// Build the tree starting with top-level permissions (ParentId == 0)
	for _, perm := range perms {
		if perm.ParentId == 0 {
			resp.List = append(resp.List, TreeRespPerm{
				Permission: perm,
				Children:   children(perm.Id, checkedMap, perms),
				Check:      getChecked(checkedMap, perm.Id),
			})
		}
	}
	return
}

// getChecked determines if a permission is checked based on the provided map.
func getChecked(checkedMap map[uint]struct{}, permissionId uint) string {
	if _, checked := checkedMap[permissionId]; checked {
		return "true"
	}
	return "false"
}

// children recursively builds the child nodes for a given parent permission.
func children(parentId uint, checkedMap map[uint]struct{}, perms []models.Permission) []TreeRespPerm {
	var trpS []TreeRespPerm
	for _, perm := range perms {
		if perm.ParentId == parentId {
			trpS = append(trpS, TreeRespPerm{
				Permission: perm,
				Children:   children(perm.Id, checkedMap, perms),
				Check:      getChecked(checkedMap, perm.Id),
			})
		}
	}
	return trpS
}
