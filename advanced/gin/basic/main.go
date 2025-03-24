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

package main

import (
	userRepo "github.com/gevinzone/lgthw/advanced/gin/basic/internal/repository/user"
	userDao "github.com/gevinzone/lgthw/advanced/gin/basic/internal/repository/user/dao"
	svc "github.com/gevinzone/lgthw/advanced/gin/basic/internal/service/basic"
	userSvc "github.com/gevinzone/lgthw/advanced/gin/basic/internal/service/user"
	"github.com/gevinzone/lgthw/advanced/gin/basic/internal/web/basic"
	"github.com/gevinzone/lgthw/advanced/gin/basic/internal/web/middleware"
	userHdl "github.com/gevinzone/lgthw/advanced/gin/basic/internal/web/user"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//server := gin.Default()
	//helloSvc := svc.NewHelloService()
	//helloHandler := basic.NewHelloHandler(helloSvc)
	//helloHandler.RegisterRoutes(server)

	server := initWebServer()
	db := initDB()

	initHelloHandler(server)
	initUserHdl(db, server)

	err := server.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func initWebServer() *gin.Engine {
	server := gin.Default()

	login := &middleware.LoginMiddlewareBuilder{}
	// 存储数据的，也就是你 userId 存哪里
	// 直接存 cookie
	store := cookie.NewStore([]byte("secret"))
	server.Use(sessions.Sessions("ssid", store), login.CheckLogin())

	return server
}

func initHelloHandler(server *gin.Engine) {
	helloSvc := svc.NewHelloService()
	helloHandler := basic.NewHelloHandler(helloSvc)
	helloHandler.RegisterRoutes(server)
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"))
	if err != nil {
		panic(err)
	}

	err = userDao.InitTables(db)
	if err != nil {
		panic(err)
	}
	return db
}

func initUserHdl(db *gorm.DB, server *gin.Engine) {
	ud := userDao.NewUserDao(db)
	ur := userRepo.NewUser(*ud)
	us := userSvc.NewUserService(ur)
	hdl := userHdl.NewUserHandler(us)
	hdl.RegisterRoutes(server)
}
