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

package ioc

import (
	"github.com/gevinzone/basic-go/live/webook/internal/web"
	"github.com/gevinzone/basic-go/live/webook/internal/web/middleware"
	"github.com/gevinzone/basic-go/live/webook/pkg/ginx/middlewares/ratelimit"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"strings"
	"time"
)

func InitWebServer(mdls []gin.HandlerFunc, userHdl *web.UserHandler) *gin.Engine {
	server := gin.Default()
	server.Use(mdls...)
	userHdl.RegisterRoutes(server)
	return server
}

func InitMiddlewares(redisClient redis.Cmdable) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		func(ctx *gin.Context) {
			println("this is the first middleware")
		},
		func(ctx *gin.Context) {
			println("this is the second middleware")
		},
		ratelimit.NewBuilder(redisClient, time.Second, 100).Build(),
		cors.New(cors.Config{
			//AllowOrigins: []string{"*"},
			AllowHeaders: []string{"Content-Type", "Authorization"},
			// 你不加这个，前端是拿不到的
			ExposeHeaders:    []string{"x-jwt-token"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
				if strings.HasPrefix(origin, "http://localhost") {
					return true
				}
				return strings.Contains(origin, "your_domain.com")
			},
			MaxAge: 12 * time.Hour,
		}),
		sessions.Sessions("gevin_session", memstore.NewStore([]byte("this is secret"))),
		middleware.NewLoginJwtMiddlewareBuilder().
			IgnorePaths("/users/signup", "/users/login", "/hello").
			IgnorePaths("/users/login_sms", "/users/login_sms/code/send").
			Build(),
	}
}
