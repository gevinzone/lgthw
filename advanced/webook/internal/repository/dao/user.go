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

package dao

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	ErrUserDuplicateEmail = errors.New("邮箱冲突")
	ErrUserDuplicate      = errors.New("用户冲突")
	ErrUserNotFound       = gorm.ErrRecordNotFound
)

type UserDAO interface {
	FindByEmail(ctx context.Context, email string) (User, error)
	FindById(ctx context.Context, id int64) (User, error)
	FindByPhone(ctx context.Context, phone string) (User, error)
	Insert(ctx context.Context, u User) (User, error)
}

type GormUserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) UserDAO {
	return &GormUserDAO{db: db}
}

func (dao *GormUserDAO) FindByEmail(ctx context.Context, email string) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("email=?", email).First(&u).Error
	return u, err
}

func (dao *GormUserDAO) FindById(ctx context.Context, id int64) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("id=?", id).First(&u).Error
	return u, err
}

func (dao *GormUserDAO) FindByPhone(ctx context.Context, phone string) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("phone=?", phone).First(&u).Error
	return u, err
}

func (dao *GormUserDAO) Insert(ctx context.Context, u User) (User, error) {
	now := time.Now().UnixMilli()
	u.Ctime = now
	u.Utime = now
	err := dao.db.WithContext(ctx).Create(&u).Error
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		const uniqueConflictsErrNo uint16 = 1062
		if mysqlErr.Number == uniqueConflictsErrNo {
			// 邮箱冲突
			return u, ErrUserDuplicateEmail
		}
	}
	return u, err
}

type ProfileDAO interface {
	FindByUserId(ctx context.Context, id int64) (Profile, error)
	Insert(ctx context.Context, p Profile) error
	Update(ctx context.Context, p Profile) error
}

type GORMProfileDAO struct {
	db *gorm.DB
}

func NewProfileDAO(db *gorm.DB) ProfileDAO {
	return &GORMProfileDAO{db: db}
}

func (dao *GORMProfileDAO) FindByUserId(ctx context.Context, id int64) (Profile, error) {
	var p Profile
	err := dao.db.WithContext(ctx).Where("user_id=?", id).First(&p).Error
	return p, err
}

func (dao *GORMProfileDAO) Insert(ctx context.Context, p Profile) error {
	now := time.Now().UnixMilli()
	p.Ctime = now
	p.Utime = now
	return dao.db.WithContext(ctx).Create(&p).Error
}

func (dao *GORMProfileDAO) Update(ctx context.Context, p Profile) error {
	p.Utime = time.Now().UnixMilli()
	return dao.db.WithContext(ctx).Model(&p).Where("user_id=?", p.UserId).Updates(p).Error
}

type UserWithProfileDAO interface {
	Create(ctx context.Context, u User) error
	FindProfileByEmail(ctx context.Context, email string) (User, Profile, error)
}

type GORMUserWithProfileDAO struct {
	db         *gorm.DB
	userDAO    UserDAO
	profileDAO ProfileDAO
}

func NewUserWithProfileDAO(db *gorm.DB, userDAO UserDAO, profileDAO ProfileDAO) UserWithProfileDAO {
	return &GORMUserWithProfileDAO{
		db:         db,
		userDAO:    userDAO,
		profileDAO: profileDAO,
	}
}

func (dao *GORMUserWithProfileDAO) Create(ctx context.Context, u User) error {
	return dao.db.Transaction(func(tx *gorm.DB) error {
		var (
			user User
			err  error
		)
		userDAO := NewUserDAO(tx)
		if user, err = userDAO.Insert(ctx, u); err != nil {
			return err
		}

		profileDAO := NewProfileDAO(tx)
		err = profileDAO.Insert(ctx, Profile{UserId: user.Id, Birthday: time.Now().UnixMilli()})
		return err
	})
}

func (dao *GORMUserWithProfileDAO) FindProfileByEmail(ctx context.Context, email string) (User, Profile, error) {
	var (
		u   User
		p   Profile
		err error
	)
	er := dao.db.Transaction(func(tx *gorm.DB) error {
		userDAO := NewUserDAO(tx)
		profileDAO := NewProfileDAO(tx)
		if u, err = userDAO.FindByEmail(ctx, email); err != nil {
			return err
		}
		if p, err = profileDAO.FindByUserId(ctx, u.Id); err != nil {
			return err
		}

		return nil
	})
	return u, p, er
}

// User 直接对应数据库表结构
// 有些人叫做 entity，有些人叫做 model，有些人叫做 PO(persistent object)
type User struct {
	Id       int64          `gorm:"primaryKey, autoIncrement"`
	Email    sql.NullString `gorm:"unique"`
	Phone    sql.NullString `gorm:"unique"`
	Password string
	Ctime    int64
	Utime    int64
}

type Profile struct {
	Id       int64 `gorm:"primaryKey, autoIncrement"`
	UserId   int64
	Nickname string
	Biology  string
	Birthday int64
	Ctime    int64
	Utime    int64
}
