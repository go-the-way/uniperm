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

package user

import "github.com/go-the-way/uniperm/internal/services/base"

func (req *AddReq) Check() (err error) {
	return base.CheckAll(
		func() (err error) { return base.CheckUsernameExists(req.Username) },
		func() (err error) {
			if req.RoleId > 0 {
				return base.CheckRoleExist(req.RoleId)
			}
			return
		},
	)
}

func (req *UpdateReq) Check() (err error) { return base.CheckUserExists(req.Id) }

func (req *UpdatePasswordReq) Check() (err error) { return base.CheckUserExists(req.Id) }

func (req *UpdateRoleReq) Check() (err error) {
	return base.CheckAll(
		func() (err error) { return base.CheckUserIsSuper(req.Id) },
		func() (err error) { return (&DeleteReq{IdReq: IdReq{req.Id}}).Check() },
	)
}

func (req *DeleteReq) Check() (err error) {
	return base.CheckAll(
		func() (err error) { return base.CheckUserIsSuper(req.Id) },
		func() (err error) { return base.CheckUserExists(req.Id) },
	)
}

func (req *EnableReq) Check() (err error) { return (&DeleteReq{IdReq: IdReq{req.Id}}).Check() }

func (req *DisableReq) Check() (err error) { return (&DeleteReq{IdReq: IdReq{req.Id}}).Check() }
