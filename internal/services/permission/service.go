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

package permission

import (
	"github.com/go-the-way/unilog"
	"github.com/go-the-way/uniperm/internal/db"
	"github.com/go-the-way/uniperm/internal/models"
	"github.com/go-the-way/uniperm/internal/services/base"
)

type service struct{}

func (s *service) Tree(req TreeReq) (resp TreeResp, err error) {
	return base.PermTree(req.PermissionId)
}

func (s *service) Get(req GetReq) (resp GetResp, err error) {
	return base.Return(resp, db.GetDB().Model(new(models.Permission)).Where("id=?", req.Id).First(&resp).Error)
}

func (s *service) Add(req AddReq) (err error) {
	return base.Callback1(db.GetDB().Create(req.transform()).Error, &req, unilog.Callback[*AddReq](unilog.Type1Admin()), base.TransformCallback(&req, req.Callback))
}

func (s *service) Update(req UpdateReq) (err error) {
	return base.Callback1(db.GetDB().Model(&models.Permission{Id: req.Id}).Updates(req.transform()).Error, &req, unilog.Callback[*UpdateReq](unilog.Type1Admin()), base.TransformCallback(&req, req.Callback))
}

func (s *service) Delete(req DeleteReq) (err error) {
	return base.Callback1(db.GetDB().Delete(&models.Permission{Id: req.Id}).Error, &req, unilog.Callback[*DeleteReq](unilog.Type1Admin()), base.TransformCallback(&req, req.Callback))
}
