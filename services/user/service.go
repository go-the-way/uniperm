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

package user

import (
	"errors"
	"fmt"
	"github.com/rwscode/uniperm/deps/db"
	"github.com/rwscode/uniperm/deps/pkg"
	"github.com/rwscode/uniperm/models"
	"github.com/rwscode/uniperm/services/base"
)

type service struct{}

func (s *service) GetPage(req GetPageReq) (resp GetPageResp, err error) {
	q := db.GetDb().Model(new(models.User))
	pkg.IfGt0Func(req.Id, func() { q.Where("id=?", req.Id) })
	pkg.IfNotEmptyFunc(req.Username, func() { q.Where("username like concat('%',?,'%')", req.Username) })
	pkg.IfNotEmptyFunc(req.BusinessId1, func() { q.Where("business_id1=?", req.BusinessId1) })
	pkg.IfNotEmptyFunc(req.BusinessId2, func() { q.Where("business_id2=?", req.BusinessId2) })
	pkg.IfNotEmptyFunc(req.BusinessId3, func() { q.Where("business_id3=?", req.BusinessId3) })
	pkg.IfGt0Func(req.State, func() { q.Where("state=?", req.State) })
	pkg.IfNotEmptyFunc(req.Remark1, func() { q.Where("remark1 like concat('%',?,'%')", req.Remark1) })
	pkg.IfNotEmptyFunc(req.Remark2, func() { q.Where("remark2 like concat('%',?,'%')", req.Remark2) })
	pkg.IfNotEmptyFunc(req.Remark3, func() { q.Where("remark3 like concat('%',?,'%')", req.Remark3) })
	pkg.IfNotEmptyFunc(req.LoginIp, func() { q.Where("login_ip like concat('%',?,'%')", req.LoginIp) })
	pkg.IfNotEmptyFunc(req.CreateTime1, func() { q.Where("create_time>=concat(?,' 00:00:00')", req.CreateTime1) })
	pkg.IfNotEmptyFunc(req.CreateTime2, func() { q.Where("create_time<=concat(?,' 23:59:59')", req.CreateTime2) })
	pkg.IfNotEmptyFunc(req.UpdateTime1, func() { q.Where("update_time>=concat(?,' 00:00:00')", req.UpdateTime1) })
	pkg.IfNotEmptyFunc(req.UpdateTime2, func() { q.Where("update_time<=concat(?,' 23:59:59')", req.UpdateTime2) })
	pkg.IfNotEmptyFunc(req.LoginTime1, func() { q.Where("login_time1>=concat(?,' 00:00:00')", req.LoginTime1) })
	pkg.IfNotEmptyFunc(req.LoginTime2, func() { q.Where("login_time2<=concat(?,' 23:59:59')", req.LoginTime2) })
	if req.OrderBy != "" {
		q.Order(req.OrderBy)
	}
	resp.List = make([]models.User, 0)
	return base.Return(resp, db.GetPagination()(q, req.Page, req.Limit, &resp.Total, &resp.List))
}

func (s *service) Get(req GetReq) (resp GetResp, err error) {
	var list []models.User
	if err = db.GetDb().Model(new(models.Role)).Where("id=?", req.Id).Find(&list).Error; err != nil {
		return
	}
	if len(list) == 0 {
		err = errors.New(fmt.Sprintf("用户[%d]不存在", req.Id))
		return
	}
	resp.User = list[0]
	return
}

func (s *service) Add(req AddReq) (err error) {
	return base.Callback1(db.GetDb().Create(req.Transform()).Error, req, req.Callback)
}

func (s *service) Update(req UpdateReq) (err error) {
	return base.Callback1(db.GetDb().Model(&models.User{Id: req.Id}).Select("business_id1", "business_id2", "business_id3", "remark1", "remark2", "remark3", "update_time").Updates(req.Transform()).Error, req, req.Callback)
}

func (s *service) UpdatePassword(req UpdatePasswordReq) (err error) {
	return base.Callback1(db.GetDb().Model(&models.User{Id: req.Id}).Updates(&models.User{Password: pkg.MD5(req.Password), UpdateTime: pkg.TimeNowStr()}).Error, req, req.Callback)
}

func (s *service) UpdateRole(req UpdateRoleReq) (err error) {
	return base.Callback1(db.GetDb().Model(&models.User{Id: req.Id}).Updates(&models.User{RoleId: req.RoleId, UpdateTime: pkg.TimeNowStr()}).Error, req, req.Callback)
}

func (s *service) Delete(req DeleteReq) (err error) {
	return base.Callback1(db.GetDb().Delete(&models.User{Id: req.Id}).Error, req, req.Callback)
}

func (s *service) Enable(req EnableReq) (err error) {
	return base.Callback1(s.updateState(req.Id, models.UserStateEnable), req.Callback)
}

func (s *service) Disable(req DisableReq) (err error) {
	return base.Callback1(s.updateState(req.Id, models.UserStateDisable), req.Callback)
}

func (s *service) updateState(id uint, state byte) (err error) {
	return db.GetDb().Model(&models.User{Id: id}).Updates(models.User{State: state, UpdateTime: pkg.TimeNowStr()}).Error
}

func (s *service) GetPerm(req GetPermReq) (resp GetPermResp, err error) {
	var rps []models.RolePermission
	if err = db.GetDb().Model(new(models.RolePermission)).Where("role_id=?", req.RoleId).Find(&rps).Error; err != nil {
		return
	}
	var checkedMap = map[uint]bool{}
	for _, rp := range rps {
		checkedMap[rp.PermissionId] = true
	}
	var ps []models.Permission
	if err = db.GetDb().Model(new(models.Permission)).Find(&ps).Error; err != nil {
		return
	}
	resp.SuperAdmin = req.SuperAdmin()
	resp.Routes = make([]GetPermRespRoute, 0)
	for _, p := range ps {
		if p.ParentId == 0 && p.IsButton == models.PermissionIsButtonNo && (req.SuperAdmin() || checkedMap[p.Id]) {
			resp.Routes = append(resp.Routes, GetPermRespRoute{Id: p.Id, Path: p.Route, Children: s.children(p.Id, checkedMap, ps, req.SuperAdmin())})
		}
	}
	return
}

func (s *service) children(parentId uint, checkedMap map[uint]bool, rps []models.Permission, superAdmin bool) []GetPermRespRoute {
	var prrS []GetPermRespRoute
	for _, rp := range rps {
		if rp.ParentId == parentId && rp.IsButton == models.PermissionIsButtonNo && (superAdmin || checkedMap[rp.Id]) {
			prrS = append(prrS, GetPermRespRoute{rp.Id, rp.Route, s.children(rp.Id, checkedMap, rps, superAdmin)})
		}
	}
	return prrS
}

func (s *service) GetPermButton(req GetPermButtonReq) (resp GetPermButtonResp, err error) {
	if superAdmin := req.SuperAdmin(); superAdmin {
		resp.SuperAdmin = superAdmin
		return
	}
	if req.RoleId <= 0 {
		return
	}
	err = db.GetDb().Raw(`
select sp.route
from uniperm_permissions sp,
     uniperm_permissions spp,
     uniperm_role_permissions srp
where sp.id = srp.permission_id
  and sp.parent_id = spp.id
  and sp.is_button = ?
  and srp.role_id = ?
  and spp.id = ?`, models.PermissionIsButtonYes, req.RoleId, req.Id).Find(&resp.Buttons).Error
	return
}

func (s *service) Login(req LoginReq) (resp LoginResp, err error) {
	var user models.User
	if err = db.GetDb().Model(new(models.User)).Where("username=?", req.Username).Scan(&user).Error; err != nil {
		return
	}
	if user.Id == 0 {
		err = errors.New("用户名不存在")
		base.Callback1(nil, req, req.Callback.NotFound)
		return
	}
	if user.State == models.UserStateDisable {
		err = errors.New("用户被禁用，请联系管理员")
		base.Callback2(nil, req, user, req.Callback.Disabled)
		return
	}
	if user.RoleId > 0 {
		var roleState byte
		if err = db.GetDb().Model(new(models.Role)).Where("id=?", user.RoleId).Select("state").Scan(&roleState).Error; err != nil {
			return
		}
		if roleState == models.RoleStateDisable {
			err = errors.New("用户所属角色被禁用，请联系管理员")
			base.Callback2(nil, req, user, req.Callback.RoleDisabled)
			return
		}
	}
	if pwd := pkg.MD5(req.Password); pwd != user.Password {
		err = errors.New("用户名和密码不匹配，请确认后重试")
		base.Callback2(nil, req, user, req.Callback.PasswordWrong)
		return
	}
	if err = db.GetDb().Model(new(models.User)).Where("id=?", user.Id).Select("login_time", "login_ip").Updates(models.User{LoginTime: pkg.TimeNowStr(), LoginIp: req.ClientIp}).Error; err != nil {
		return
	}
	resp.Token = pkg.MD5(pkg.RandStr(32))
	base.Callback3(nil, req, user, &resp, req.Callback.Success)
	return
}

func (s *service) Logout(req LogoutReq) (err error) { return base.Callback1Err(nil, req, req.Callback) }
