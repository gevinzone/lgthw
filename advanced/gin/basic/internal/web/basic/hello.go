// Copyright 2023 igevin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package basic

import (
	"github.com/gevinzone/lgthw/advanced/gin/basic/internal/service/basic"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HelloHandler struct {
	svc basic.Hello
}

func NewHelloHandler(svc basic.Hello) *HelloHandler {
	return &HelloHandler{
		svc: svc,
	}
}

func (h *HelloHandler) RegisterRoutes(server *gin.Engine) {
	g := server.Group("/hello")
	g.GET("/hello", h.Hello)
	g.GET("/hello2", h.Hello2)
}

func (h *HelloHandler) Hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hello World")
}

func (h *HelloHandler) Hello2(ctx *gin.Context) {
	h.svc.Hello(ctx)
}
