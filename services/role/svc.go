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

var (
	Service SVC = &service{}

	GetPage    = Service.GetPage
	Get        = Service.Get
	GetPerm    = Service.GetPerm
	UpdatePerm = Service.UpdatePerm
	Add        = Service.Add
	Update     = Service.Update
	Delete     = Service.Delete
	Enable     = Service.Enable
	Disable    = Service.Disable
)

type SVC interface {
	GetPage(req GetPageReq) (resp GetPageResp, err error)
	Get(req GetReq) (resp GetResp, err error)
	GetPerm(req GetPermReq) (resp GetPermResp, err error)
	UpdatePerm(req UpdatePermReq) (err error)
	Add(req AddReq) (err error)
	Update(req UpdateReq) (err error)
	Delete(req DeleteReq) (err error)
	Enable(req EnableReq) (err error)
	Disable(req DisableReq) (err error)
}
