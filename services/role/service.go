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

package role

import (
	"errors"
	"fmt"
	"github.com/go-the-way/uniperm/deps/db"
	"github.com/go-the-way/uniperm/deps/pkg"
	"github.com/go-the-way/uniperm/models"
	"github.com/go-the-way/uniperm/services/base"
)

type service struct{}

func (s *service) GetPage(req GetPageReq) (resp GetPageResp, err error) {
	q := db.GetDb().Model(new(models.Role))
	pkg.IfGt0Func(req.Id, func() { q.Where("id=?", req.Id) })
	pkg.IfNotEmptyFunc(req.Name, func() { q.Where("name like concat('%',?,'%')", req.Name) })
	pkg.IfNotEmptyFunc(req.Type, func() { q.Where("type=?", req.Type) })
	pkg.IfNotEmptyFunc(req.Description, func() { q.Where("description like concat('%',?,'%')", req.Description) })
	pkg.IfGt0Func(req.State, func() { q.Where("state=?", req.State) })
	pkg.IfNotEmptyFunc(req.CreateTime1, func() { q.Where("create_time>=concat(?,' 00:00:00')", req.CreateTime1) })
	pkg.IfNotEmptyFunc(req.CreateTime2, func() { q.Where("create_time<=concat(?,' 23:59:59')", req.CreateTime2) })
	pkg.IfNotEmptyFunc(req.UpdateTime1, func() { q.Where("update_time>=concat(?,' 00:00:00')", req.UpdateTime1) })
	pkg.IfNotEmptyFunc(req.UpdateTime2, func() { q.Where("update_time<=concat(?,' 23:59:59')", req.UpdateTime2) })
	if req.OrderBy != "" {
		q.Order(req.OrderBy)
	}
	resp.List = make([]models.Role, 0)
	base.Return(resp, db.GetPagination()(q, req.Page, req.Limit, &resp.Total, &resp.List))
	return
}

func (s *service) Get(req GetReq) (resp GetResp, err error) {
	var list []models.Role
	if err = db.GetDb().Model(new(models.Role)).Where("id=?", req.Id).Find(&list).Error; err != nil {
		return
	}
	if len(list) == 0 {
		err = errors.New(fmt.Sprintf("角色[%d]不存在", req.Id))
		return
	}
	resp.Role = list[0]
	return
}

func (s *service) Add(req AddReq) error {
	return base.Callback1(db.GetDb().Create(req.Transform()).Error, req, req.Callback)
}

func (s *service) Update(req UpdateReq) (err error) {
	return base.Callback1(db.GetDb().Model(&models.Role{Id: req.Id}).Omit("create_time", "state").Updates(req.Transform()).Error, req, req.Callback)
}

func (s *service) Delete(req DeleteReq) (err error) {
	return base.Callback1(db.GetDb().Delete(&models.Role{Id: req.Id}).Error, req, req.Callback)
}

func (s *service) Enable(req EnableReq) (err error) {
	return base.Callback1(s.updateState(req.Id, models.RoleStateEnable), req, req.Callback)
}

func (s *service) Disable(req DisableReq) (err error) {
	return base.Callback1(s.updateState(req.Id, models.RoleStateDisable), req, req.Callback)
}

func (s *service) updateState(id uint, state byte) (err error) {
	return db.GetDb().Model(&models.Role{Id: id}).Updates(models.Role{State: state, UpdateTime: pkg.TimeNowStr()}).Error
}

func (s *service) GetPerm(req GetPermReq) (resp GetPermResp, err error) {
	var perms []uint
	if err = db.GetDb().Model(new(models.RolePermission)).Where("role_id=?", req.Id).Select("permission_id").Find(&perms).Error; err != nil {
		return
	}
	return base.PermTree(perms)
}

func (s *service) UpdatePerm(req UpdatePermReq) (err error) {
	var rps = make([]models.RolePermission, 0)
	for _, perm := range req.Permissions {
		rps = append(rps, models.RolePermission{RoleId: req.Id, PermissionId: perm})
	}
	if len(rps) <= 0 {
		return errors.New("需要分配的权限为空")
	}
	tx := db.GetDb().Begin()
	if err = tx.Model(new(models.RolePermission)).Where("role_id=?", req.Id).Error; err != nil {
		_ = tx.Rollback().Error
		return
	}
	if err = tx.Model(new(models.RolePermission)).Where("role_id=?", req.Id).Delete(models.RolePermission{}).Error; err != nil {
		_ = tx.Rollback().Error
		return
	}
	if err = tx.CreateInBatches(rps, len(rps)).Error; err != nil {
		_ = tx.Rollback().Error
		return
	}
	if err = tx.Model(&models.Role{Id: req.Id}).UpdateColumn("update_time", pkg.TimeNowStr()).Error; err != nil {
		_ = tx.Rollback().Error
		return
	}
	_ = tx.Commit().Error
	return base.Callback1(err, req, req.Callback)
}
