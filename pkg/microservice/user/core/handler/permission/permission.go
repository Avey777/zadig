/*
Copyright 2023 The KodeRover Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package permission

import (
	"github.com/gin-gonic/gin"
	"github.com/koderover/zadig/v2/pkg/microservice/user/core/service/permission"

	internalhandler "github.com/koderover/zadig/v2/pkg/shared/handler"
)

func GetUserRules(c *gin.Context) {
	ctx := internalhandler.NewContext(c)
	defer func() { internalhandler.JSONResponse(c, ctx) }()

	ctx.Resp, ctx.RespErr = permission.GetUserRules(ctx.UserID, ctx.Logger)
}

func GetUserRulesByProject(c *gin.Context) {
	ctx := internalhandler.NewContext(c)
	defer func() { internalhandler.JSONResponse(c, ctx) }()

	ctx.Resp, ctx.RespErr = permission.GetUserPermissionByProject(ctx.UserID, c.Param("name"), ctx.Logger)
}
