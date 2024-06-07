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
	"github.com/rwscode/uniperm/deps/db"
	"github.com/rwscode/uniperm/models"
	"github.com/rwscode/uniperm/services/base"
)

type service struct{}

func (s *service) Tree(req TreeReq) (resp TreeResp, err error) {
	return base.PermTree(req.PermissionId)
}

func (s *service) Get(req GetReq) (resp GetResp, err error) {
	return base.Return(resp, db.GetDb().Model(new(models.Permission)).Where("id=?", req.Id).First(&resp).Error)
}

func (s *service) Add(req AddReq) (err error) {
	return base.Callback1(db.GetDb().Create(req.Transform()).Error, req, req.Callback)
}

func (s *service) Update(req UpdateReq) (err error) {
	return base.Callback1(db.GetDb().Model(&models.Permission{Id: req.Id}).Select("name", "route").Updates(req.Transform()).Error, req, req.Callback)
}

func (s *service) Delete(req DeleteReq) (err error) {
	return base.Callback1(db.GetDb().Delete(&models.Permission{Id: req.Id}).Error, req, req.Callback)
}
