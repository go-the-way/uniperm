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

package role

import "github.com/go-the-way/uniperm/internal/services/base"

func (req *AddReq) Check() (err error) { return }

func (req *UpdateReq) Check() (err error) { return base.CheckRoleExist(req.Id) }

func (req *DeleteReq) Check() (err error) {
	if err = base.CheckRoleExist(req.Id); err != nil {
		return
	}
	if err = base.CheckRoleRefUser(req.Id); err != nil {
		return
	}
	return base.CheckRoleRefPermission(req.Id)
}

func (req *EnableReq) Check() (err error)     { return base.CheckRoleExist(req.Id) }
func (req *DisableReq) Check() (err error)    { return base.CheckRoleExist(req.Id) }
func (req *GetPermReq) Check() (err error)    { return base.CheckRoleExist(req.Id) }
func (req *UpdatePermReq) Check() (err error) { return base.CheckRoleExist(req.Id) }
