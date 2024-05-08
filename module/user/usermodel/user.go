package usermodel

import (
	"github.com/go-playground/validator/v10"
	"github.com/orgball2608/helmet-shop-be/common"
)

const EntityName = "User"

type User struct {
	common.SQLModel `json:",inline"`
	Email           string        `json:"email" gorm:"column:email;"`
	Password        string        `json:"-" gorm:"column:password;"`
	Salt            string        `json:"-" gorm:"column:salt;"`
	LastName        string        `json:"last_name" gorm:"column:last_name;"`
	FirstName       string        `json:"first_name" gorm:"column:first_name;"`
	Phone           string        `json:"phone" gorm:"column:phone;"`
	Role            string        `json:"role" gorm:"column:role;"`
	CartId          int           `json:"-" gorm:"column:cart_id;"`
	CartUID         *common.UID   `json:"cart_id" gorm:"-"`
	Avatar          *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) Mask() {
	u.GenUID(common.DbTypeUser)
	uid := common.NewUID(uint32(u.CartId), int(common.DbTypeCart), 1)
	u.CartUID = &uid
}

func (u *User) GetUserEmail() string {
	return u.Email
}
func (u *User) GetUserRole() string {
	return u.Role
}

func (u *User) GetUserId() int {
	return u.Id
}

func (u *User) GetCartId() int {
	return u.CartId
}

type UserCreate struct {
	common.SQLModel `json:",inline"`
	Email           string        `json:"email" form:"email" gorm:"column:email;"`
	Password        string        `json:"password" form:"password"  gorm:"column:password;"`
	LastName        string        `json:"last_name" form:"last_name" gorm:"column:last_name;"`
	FirstName       string        `json:"first_name" form:"first_name" gorm:"column:first_name;"`
	Phone           string        `json:"phone" gorm:"column:phone;"`
	Role            string        `json:"role" gorm:"column:role;"`
	Salt            string        `json:"-" gorm:"column:salt;"`
	Avatar          *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
	CartId          int           `json:"-" gorm:"column:cart_id;"`
}

func (res *UserCreate) Validate() error {
	validate := validator.New()

	if err := validate.Var(res.Email, "required"); err != nil {
		return ErrEmailIsRequired
	}

	if err := validate.Var(res.Email, "email"); err != nil {
		return InvalidEmailFormat
	}

	if err := validate.Var(res.Password, "required"); err != nil {
		return ErrPasswordIsRequired
	}

	if err := validate.Var(res.Password, "min=8,max=20"); err != nil {
		return InvalidPasswordFormat
	}

	if err := validate.Var(res.FirstName, "required"); err != nil {
		return ErrFirstNameIsRequired
	}

	if err := validate.Var(res.FirstName, "min=2,max=20"); err != nil {
		return InvalidFirstNameFormat
	}

	if err := validate.Var(res.LastName, "required"); err != nil {
		return ErrLastNameIsRequired
	}

	if err := validate.Var(res.LastName, "min=2,max=20"); err != nil {
		return InvalidLastNameFormat
	}

	if err := validate.Var(res.Phone, "omitempty,min=10,max=12"); err != nil {
		return InvalidPhoneFormat
	}

	return nil
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

func (u *UserCreate) Mask() {
	u.GenUID(common.DbTypeUser)
}

func (u *UserCreate) GetUserId() int {
	return u.Id
}

type UserLogin struct {
	Email    string `json:"email" form:"email"  gorm:"column:email;"`
	Password string `json:"password" form:"password" gorm:"column:password;"`
}

func (res *UserLogin) Validate() error {
	validate := validator.New()

	if err := validate.Var(res.Email, "required"); err != nil {
		return ErrEmailIsRequired
	}

	if err := validate.Var(res.Email, "email"); err != nil {
		return InvalidEmailFormat
	}

	if err := validate.Var(res.Password, "required"); err != nil {
		return ErrPasswordIsRequired
	}

	if err := validate.Var(res.Password, "min=8,max=20"); err != nil {
		return InvalidPasswordFormat
	}
	return nil
}

func (UserLogin) TableName() string {
	return User{}.TableName()
}

type UserUpdate struct {
	Email     string        `json:"email" form:"email" gorm:"column:email;"`
	LastName  string        `json:"last_name" form:"last_name" gorm:"column:last_name;"`
	FirstName string        `json:"first_name" form:"first_name" gorm:"column:first_name;"`
	Phone     string        `json:"phone" gorm:"column:phone;"`
	Role      string        `json:"-" gorm:"column:role;"`
	Avatar    *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (UserUpdate) TableName() string {
	return User{}.TableName()
}

func (res *UserUpdate) Validate() error {
	validate := validator.New()

	if err := validate.Var(res.Email, "omitempty,email"); err != nil {
		return InvalidEmailFormat
	}

	if err := validate.Var(res.FirstName, "omitempty,min=2,max=20"); err != nil {
		return InvalidFirstNameFormat
	}

	if err := validate.Var(res.LastName, "omitempty,min=2,max=20"); err != nil {
		return InvalidLastNameFormat
	}

	if err := validate.Var(res.Phone, "omitempty,min=10,max=12"); err != nil {
		return InvalidPhoneFormat
	}

	if res.Role != "" && res.Role != "user" && res.Role != "admin" {
		return InvalidRoleFormat
	}

	return nil
}

type PasswordUpdate struct {
	Password    string `json:"password" form:"password" gorm:"-"`
	NewPassword string `json:"new_password" form:"password" gorm:"column:password;"`
}

func (PasswordUpdate) TableName() string {
	return User{}.TableName()
}

func (res *PasswordUpdate) Validate() error {
	validate := validator.New()

	if err := validate.Var(res.Password, "required"); err != nil {
		return ErrPasswordIsRequired
	}

	if err := validate.Var(res.Password, "min=8,max=20"); err != nil {
		return InvalidPasswordFormat
	}

	if err := validate.Var(res.NewPassword, "required"); err != nil {
		return ErrNewPasswordIsRequired
	}

	if err := validate.Var(res.NewPassword, "min=8,max=20"); err != nil {
		return InvalidNewPasswordFormat
	}
	return nil
}
