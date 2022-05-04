package model

import (
	"blog/utils/errmsg"
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=12" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required, gte=2" label:"角色码"`
	gorm.Model
}

// 对数据库的操作DAO

// 查询用户是否存在
func CheckUser(name string) int {
	var users User
	db.Select("id").Where("username = ?", name).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED //用户已存在
	} else {
		return errmsg.SUCCESS // 用户不存在，用户名可用
	}
}

// 修改上述方法，在api的EditUser中使用
func CheckUpUser(id int, name string) int {
	var user User
	db.Select("id,username").Where("username = ?", name).First(&user)
	if user.ID == uint(id) {
		return errmsg.SUCCESS
	}
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED //用户已存在
	} else {
		return errmsg.SUCCESS // 用户不存在，用户名可用
	}
}

// 新增用户
func CreateUser(data *User) int {
	// 对用户密码加密，也可以使用钩子函数
	data.Password = ScryptPw(data.Password)
	err := db.Create(&data).Error

	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS // 200
}

// 查询单个用户(编辑)
func GetUserInfo(id int) (User, int) {
	var user User
	err = db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, errmsg.ERROR_USER_NOT_EXIST
	}
	return user, errmsg.SUCCESS
}

// 查询用户列表
// 返回User类型的切片
func GetUsers(username string, pageSize int, pageNum int) ([]User, int64) {
	//var user User
	var users []User
	var total int64
	if username == "" {
		db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Count(&total)
		return users, total
	} else {
		db.Where("username LIKE ?", username+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Count(&total)
	}
	//db.Model(&user).Count(&total)
	if err == gorm.ErrRecordNotFound {
		return nil, 0
	}
	return users, total
}

// 编辑用户(不应包含密码修改)
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除用户(软删除)
func DeleteUser(id int) int {
	var user User
	err = db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 钩子函数加密(未使用)
// 钩子函数类似于vue和微信小程序中的生命周期函数，会自动被调用
//func (u *User) BeforeSave(tx *gorm.DB) (err error) {
//	u.Password = ScryptPw(u.Password)
//	return nil
//}

// 用户密码加密
func ScryptPw(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 4, 6, 66, 22, 222, 11}
	// HashPw:[]byte
	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	// fpw:string
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}

// 登录验证数据库中的用户名和密码
func CheckLogin(username string, password string) int {
	var user User
	db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPw(password) != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1 { // 因为是从1开始的
		return errmsg.ERROR_USER_NO_RIGHT
	}
	return errmsg.SUCCESS
}
