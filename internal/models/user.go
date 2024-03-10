package models

import (
	"gorm.io/gorm"

	pwd "github.com/webgamedevelop/webgame-api/internal/pkg/password"
)

type UserLoginRequest struct {
	Name     string `form:"name" binding:"required,min=3,max=20" json:"name"`
	Password string `form:"password" binding:"required,max=16" json:"password"`
}

type UserUpdateRequest struct {
	Name  string `form:"-" json:"-"`
	Email string `form:"email" binding:"email,max=50" json:"email"`
	Phone string `form:"phone" binding:"min=11,max=13" json:"phone"`
}

type UserChangePasswordRequest struct {
	Name            string `form:"-" json:"-"`
	Password        string `form:"password" binding:"required,max=16" json:"password,omitempty"`
	ConfirmPassword string `binding:"required,eqfield=Password" json:"confirmPassword,omitempty"`
}

type User struct {
	gorm.Model      `json:"-" form:"-"`
	Name            string `gorm:"type:varchar(20);unique;not null" form:"name" binding:"required,min=3,max=20" json:"name"`
	Email           string `gorm:"type:varchar(50);unique" form:"email" binding:"required,email,max=50" json:"email"`
	Phone           string `gorm:"type:varchar(13);unique;not null" form:"phone" binding:"required,min=11,max=13" json:"phone"`
	Password        string `gorm:"type:varchar(60);not null" form:"password" binding:"required,max=16" json:"password,omitempty"`
	ConfirmPassword string `gorm:"-" form:"confirmPassword" binding:"required,eqfield=Password" json:"confirmPassword,omitempty"`
	Init            bool   `json:"-" form:"-"`
}

func clearPassword(u *User) {
	u.Password = ""
	u.ConfirmPassword = ""
	u.Init = false
}

func InitAdminUser(username, email, phone, password string) error {
	hashed, err := pwd.Generate([]byte(password))
	if err != nil {
		return err
	}
	user := &User{
		Name:     username,
		Email:    email,
		Phone:    phone,
		Password: string(hashed),
		Init:     true,
	}
	if err := db.FirstOrCreate(user, &User{Name: username}).Error; err != nil {
		return err
	}
	return nil
}

func CreateUser(user *User) (*User, error) {
	var (
		err    error
		hashed []byte
	)

	if hashed, err = pwd.Generate([]byte(user.Password)); err != nil {
		return nil, err
	}

	user.Password = string(hashed)

	tx := db.Begin()
	// rollback when panic or err
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			return
		}
		if err != nil {
			tx.Rollback()
			return
		}
	}()

	if err = tx.Error; err != nil {
		return nil, err
	}

	if err = tx.Create(user).Error; err != nil {
		return nil, err
	}

	if err = tx.Commit().Error; err != nil {
		return nil, err
	}

	clearPassword(user)
	return user, nil
}

func GetUser(name string) (*User, error) {
	var user User
	if err := db.First(&user, &User{Name: name}).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func CompareUser(name, password string) error {
	var (
		user *User
		err  error
	)
	if user, err = GetUser(name); err != nil {
		return err
	}
	return pwd.Compare([]byte(user.Password), []byte(password))
}

func UpdateUser(request *UserUpdateRequest) (*User, error) {
	var (
		user User
		err  error
	)

	tx := db.Begin()
	// rollback when panic or err
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			return
		}
		if err != nil {
			tx.Rollback()
			return
		}
	}()

	if err = tx.Error; err != nil {
		return nil, err
	}

	if err = tx.Set("gorm:query_option", "FOR UPDATE").First(&user, &User{Name: request.Name}).Error; err != nil {
		return nil, err
	}

	if request.Phone != "" {
		user.Phone = request.Phone
	}

	if request.Email != "" {
		user.Email = request.Email
	}

	if err = tx.Save(&user).Error; err != nil {
		return nil, err
	}

	if err = tx.Commit().Error; err != nil {
		return nil, err
	}

	clearPassword(&user)
	return &user, nil
}

func ChangePassword(request *UserChangePasswordRequest) (*User, error) {
	var (
		user   User
		hashed []byte
		err    error
	)

	tx := db.Begin()
	// rollback when panic or err
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			return
		}
		if err != nil {
			tx.Rollback()
			return
		}
	}()

	if err = tx.Error; err != nil {
		return nil, err
	}

	if err = tx.Set("gorm:query_option", "FOR UPDATE").First(&user, &User{Name: request.Name}).Error; err != nil {
		return nil, err
	}

	if hashed, err = pwd.Generate([]byte(request.Password)); err != nil {
		return nil, err
	}

	user.Password = string(hashed)

	if err = tx.Save(&user).Error; err != nil {
		return nil, err
	}

	if err = tx.Commit().Error; err != nil {
		return nil, err
	}

	clearPassword(&user)
	return &user, nil
}
