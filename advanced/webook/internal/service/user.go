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

package service

import (
	"context"
	"errors"
	"github.com/gevinzone/basic-go/live/webook/internal/domain"
	"github.com/gevinzone/basic-go/live/webook/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var ErrUserDuplicateEmail = repository.ErrUserDuplicateEmail
var ErrInvalidUserOrPassword = errors.New("账号/邮箱或密码不对")

type UserService interface {
	Login(ctx context.Context, email, password string) (domain.User, error)
	SignUp(ctx context.Context, u domain.User) error
	FindOrCreate(ctx context.Context, phone string) (domain.User, error)
	//Profile(ctx context.Context,
	//	id int64) (domain.User, error)
	EditProfile(ctx context.Context, p domain.Profile) error
	GetProfileByEmail(ctx context.Context, email string) (domain.Profile, error)
}

type DefaultUserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &DefaultUserService{
		repo: repo,
	}
}

func (svc *DefaultUserService) Login(ctx context.Context, email, password string) (domain.User, error) {
	u, err := svc.repo.FindByEmail(ctx, email)
	var user domain.User
	if errors.Is(err, repository.ErrUserNotFound) {
		return user, ErrInvalidUserOrPassword
	}
	if err != nil {
		return user, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return user, ErrInvalidUserOrPassword
	}
	return u, nil
}

func (svc *DefaultUserService) SignUp(ctx context.Context, u domain.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return svc.repo.Create(ctx, u)
}

func (svc *DefaultUserService) EditProfile(ctx context.Context, p domain.Profile) error {
	return svc.repo.UpdateProfile(ctx, p)
}

func (svc *DefaultUserService) GetProfileByEmail(ctx context.Context, email string) (domain.Profile, error) {
	return svc.repo.FindProfileByEmail(ctx, email)
}

func (svc *DefaultUserService) FindOrCreate(ctx context.Context, phone string) (domain.User, error) {
	u, err := svc.repo.FindByPhone(ctx, phone)
	if err != repository.ErrUserNotFound {
		return u, err
	}
	u = domain.User{
		Phone: phone,
	}
	err = svc.repo.Create(ctx, u)
	if err != nil && !errors.Is(err, repository.ErrUserDuplicate) {
		return u, err
	}
	// 因为这里会遇到主从延迟的问题
	return svc.repo.FindByPhone(ctx, phone)
}
