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

package user

import (
	"context"
	domain "github.com/gevinzone/lgthw/advanced/gin/basic/internal/domain/user"
	"github.com/gevinzone/lgthw/advanced/gin/basic/internal/repository/user/dao"
)

var (
	ErrDuplicateEmail = dao.ErrDuplicateEmail
	ErrUserNotFound   = dao.ErrRecordNotFound
)

type User struct {
	dao dao.UserDao
}

func NewUser(dao dao.UserDao) *User {
	return &User{dao: dao}
}

func (repo *User) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	userDao, err := repo.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return repo.toDomain(userDao), nil
}

func (repo *User) Create(ctx context.Context, u domain.User) error {
	return repo.dao.Insert(ctx, &dao.User{
		Email:    u.Email,
		Password: u.Password,
	})
}

func (repo *User) toDomain(userDao dao.User) domain.User {
	return domain.User{
		Id:       userDao.Id,
		Email:    userDao.Email,
		Password: userDao.Password,
	}
}
