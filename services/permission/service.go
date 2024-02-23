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

func (s *service) Add(req AddReq) error { return db.GetDb().Create(req.Transform()).Error }

func (s *service) Update(req UpdateReq) error {
	return db.GetDb().Model(&models.Permission{Id: req.Id}).Select("name", "route").Updates(req.Transform()).Error
}

func (s *service) Del(req DelReq) error {
	return db.GetDb().Delete(&models.Permission{Id: req.Id}).Error
}
