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

//go:build wireinject

package main

import (
	"github.com/gevinzone/basic-go/live/webook/internal/repository"
	"github.com/gevinzone/basic-go/live/webook/internal/repository/cache"
	"github.com/gevinzone/basic-go/live/webook/internal/repository/dao"
	"github.com/gevinzone/basic-go/live/webook/internal/service"
	"github.com/gevinzone/basic-go/live/webook/internal/web"
	"github.com/gevinzone/basic-go/live/webook/ioc"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitWebServer() *gin.Engine {
	wire.Build(
		ioc.InitDB, ioc.InitRedis,
		dao.NewUserDAO, dao.NewProfileDAO, dao.NewUserWithProfileDAO,
		cache.NewUserCache, cache.NewCodeCache,
		repository.NewUserRepository, repository.NewCodeRepository,
		ioc.InitSmsService,
		service.NewUserService, service.NewCodeService,
		web.NewUserHandler,
		ioc.InitWebServer,
		ioc.InitMiddlewares,
	)
	return new(gin.Engine)
}
