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
	"errors"
	"fmt"

	"github.com/rwscode/uniperm/deps/db"
	"github.com/rwscode/uniperm/models"
)

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

func CheckUsernameExists(username string) (err error) {
	var cc int64
	if err = db.GetDb().Model(new(models.User)).Where("username=?", username).Count(&cc).Error; err != nil {
		return
	}
	if cc > 0 {
		return errors.New(fmt.Sprintf("用户[%s]已存在", username))
	}
	return
}

func CheckUserExists(userId uint) (err error) {
	var cc int64
	if err = db.GetDb().Model(new(models.User)).Where("id=?", userId).Count(&cc).Error; err != nil {
		return
	}
	if cc <= 0 {
		return errors.New(fmt.Sprintf("用户[%d]不存在", userId))
	}
	return
}

func CheckUserIsSuper(userId uint) (err error) {
	if userId == 1 {
		return errors.New("超级管理员不支持当前操作")
	}
	return
}

func CheckRoleExist(roleId uint) (err error) {
	var cc int64
	if err = db.GetDb().Model(new(models.Role)).Where("id=?", roleId).Count(&cc).Error; err != nil {
		return
	}
	if cc <= 0 {
		return errors.New(fmt.Sprintf("角色[%d]不存在", roleId))
	}
	return
}

func CheckRoleRefUser(roleId uint) (err error) {
	var cc int64
	if err = db.GetDb().Model(new(models.User)).Where("role_id=?", roleId).Count(&cc).Error; err != nil {
		return
	}
	if cc > 0 {
		return errors.New(fmt.Sprintf("角色[%d]下有用户", roleId))
	}
	return
}

func CheckRoleRefPermission(roleId uint) (err error) {
	var cc int64
	if err = db.GetDb().Model(new(models.RolePermission)).Where("role_id=?", roleId).Count(&cc).Error; err != nil {
		return
	}
	if cc > 0 {
		return errors.New(fmt.Sprintf("角色[%d]下有关联权限", roleId))
	}
	return
}

func CheckPermissionExist(permissionId uint) (err error) {
	var cc int64
	if err = db.GetDb().Model(new(models.Permission)).Where("id=?", permissionId).Count(&cc).Error; err != nil {
		return
	}
	if cc <= 0 {
		return errors.New(fmt.Sprintf("权限[%d]不存在", permissionId))
	}
	return
}

func CheckPermissionIsNotButton(permissionId uint) (err error) {
	type perm struct {
		Id       uint
		IsButton byte
	}
	var pm perm
	if err = db.GetDb().Model(new(models.Permission)).Where("id=?", permissionId).Scan(&pm).Error; err != nil {
		return
	}
	if pm.Id <= 0 {
		return errors.New(fmt.Sprintf("权限[%d]不存在", permissionId))
	}
	if pm.IsButton == models.PermissionIsButtonYes {
		return errors.New(fmt.Sprintf("权限[%d]是按钮权限", permissionId))
	}
	return
}

func CheckPermissionHaveNoSubPerms(permissionId uint) (err error) {
	var cc int64
	if err = db.GetDb().Model(new(models.Permission)).Where("parent_id=?", permissionId).Count(&cc).Error; err != nil {
		return
	}
	if cc > 0 {
		return errors.New(fmt.Sprintf("权限[%d]有下级权限[数量:%d]", permissionId, cc))
	}
	return
}

func CheckPermissionRefRole(permissionId uint) (err error) {
	var cc int64
	if err = db.GetDb().Model(new(models.RolePermission)).Where("permission_id=?", permissionId).Count(&cc).Error; err != nil {
		return
	}
	if cc > 0 {
		return errors.New(fmt.Sprintf("权限[%d]有角色引用[数量:%d]", permissionId, cc))
	}
	return
}
