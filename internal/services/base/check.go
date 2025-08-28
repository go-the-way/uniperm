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
	"errors"
	"fmt"

	"github.com/go-the-way/uniperm/internal/db"
	"github.com/go-the-way/uniperm/internal/models"
)

// CheckAll executes a series of check functions and returns the first error encountered.
func CheckAll(fns ...func() (err error)) (err error) {
	for _, fn := range fns {
		if fn != nil {
			if err = fn(); err != nil {
				break
			}
		}
	}
	return
}

// CheckUsernameExists verifies if a username already exists in the database.
func CheckUsernameExists(username string) (err error) {
	var cc int64
	if err = db.GetDB().Model(new(models.User)).Where("username=?", username).Count(&cc).Error; err != nil {
		return
	}
	if cc > 0 {
		return errors.New(fmt.Sprintf("User [%s] already exists", username))
	}
	return
}

// CheckUserExists checks if a user with the specified ID exists in the database.
func CheckUserExists(userId uint) (err error) {
	var cc int64
	if err = db.GetDB().Model(new(models.User)).Where("id=?", userId).Count(&cc).Error; err != nil {
		return
	}
	if cc <= 0 {
		return errors.New(fmt.Sprintf("User [%d] does not exist", userId))
	}
	return
}

// CheckUserIsSuper ensures the user is not a super admin (user ID 1).
func CheckUserIsSuper(userId uint) (err error) {
	if userId == 1 {
		return errors.New("super admin does not support this operation")
	}
	return
}

// CheckRoleExist verifies if a role with the specified ID exists in the database.
func CheckRoleExist(roleId uint) (err error) {
	var cc int64
	if err = db.GetDB().Model(new(models.Role)).Where("id=?", roleId).Count(&cc).Error; err != nil {
		return
	}
	if cc <= 0 {
		return errors.New(fmt.Sprintf("Role [%d] does not exist", roleId))
	}
	return
}

// CheckRoleRefUser checks if the specified role is referenced by any users.
func CheckRoleRefUser(roleId uint) (err error) {
	var cc int64
	if err = db.GetDB().Model(new(models.User)).Where("role_id=?", roleId).Count(&cc).Error; err != nil {
		return
	}
	if cc > 0 {
		return errors.New(fmt.Sprintf("Role [%d] is referenced by users", roleId))
	}
	return
}

// CheckRoleRefPermission checks if the specified role is associated with any permissions.
func CheckRoleRefPermission(roleId uint) (err error) {
	var cc int64
	if err = db.GetDB().Model(new(models.RolePermission)).Where("role_id=?", roleId).Count(&cc).Error; err != nil {
		return
	}
	if cc > 0 {
		return errors.New(fmt.Sprintf("Role [%d] is associated with permissions", roleId))
	}
	return
}

// CheckPermissionExist verifies if a permission with the specified ID exists in the database.
func CheckPermissionExist(permissionId uint) (err error) {
	var cc int64
	if err = db.GetDB().Model(new(models.Permission)).Where("id=?", permissionId).Count(&cc).Error; err != nil {
		return
	}
	if cc <= 0 {
		return errors.New(fmt.Sprintf("Permission [%d] does not exist", permissionId))
	}
	return
}

// CheckPermissionIsNotButton ensures the specified permission is not a button-type permission.
func CheckPermissionIsNotButton(permissionId uint) (err error) {
	type perm struct {
		Id       uint
		IsButton byte
	}
	var pm perm
	if err = db.GetDB().Model(new(models.Permission)).Where("id=?", permissionId).Scan(&pm).Error; err != nil {
		return
	}
	if pm.Id <= 0 {
		return errors.New(fmt.Sprintf("Permission [%d] does not exist", permissionId))
	}
	if pm.IsButton == models.PermissionIsButtonYes {
		return errors.New(fmt.Sprintf("Permission [%d] is a button permission", permissionId))
	}
	return
}

// CheckPermissionHaveNoSubPerms checks if the specified permission has no sub-permissions.
func CheckPermissionHaveNoSubPerms(permissionId uint) (err error) {
	var cc int64
	if err = db.GetDB().Model(new(models.Permission)).Where("parent_id=?", permissionId).Count(&cc).Error; err != nil {
		return
	}
	if cc > 0 {
		return errors.New(fmt.Sprintf("Permission [%d] has sub-permissions [count: %d]", permissionId, cc))
	}
	return
}

// CheckPermissionRefRole checks if the specified permission is referenced by any roles.
func CheckPermissionRefRole(permissionId uint) (err error) {
	var cc int64
	if err = db.GetDB().Model(new(models.RolePermission)).Where("permission_id=?", permissionId).Count(&cc).Error; err != nil {
		return
	}
	if cc > 0 {
		return errors.New(fmt.Sprintf("Permission [%d] is referenced by roles [count: %d]", permissionId, cc))
	}
	return
}
