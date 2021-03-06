package models

import (
	"DataCertProject/db_mysql"
	"DataCertProject/util"
)

type User struct {
	Id       int    `form:"id"`
	Phone    string `form:"phone"`
	Password string `form:"password"`
	Name     string `form:"name"`
	Card     string `form:"card"`
	Sex      string `form:"sex"`
}

/*
 *保存用户信息到数据库
 *
 */

func (u User) SaveUser() (int64, error) {
	//1、密码脱敏处理  sha256加密

	u.Password = util.MD5HashString(u.Password)

	//2、执行数据库操作
	row, err := db_mysql.Db.Exec("insert into user(phone, password) values(?,?)", u.Phone, u.Password)
	if err != nil {
		return -1, err
	}
	id, err := row.RowsAffected()
	if err != nil {
		return -1, err
	}
	return id, nil
}

//查询用户信息
func (u User) QueryUser() (*User, error) {
	//1、密码脱敏处理  sha256加密
	//md5Hash := md5.New()
	//md5Hash.Write([]byte(u.Password))
	//passwordBytes := md5Hash.Sum(nil)
	//u.Password = hex.EncodeToString(passwordBytes)
	u.Password = util.MD5HashString(u.Password)

	row := db_mysql.Db.QueryRow("select phone from user where phone = ? and password = ?", u.Phone, u.Password)
	err := row.Scan(&u.Phone)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

//根据phone查询用户的相关信息
func QueryUserByPhone(phone string) (*User, error) {
	row := db_mysql.Db.QueryRow("select phone, name, card, sex from user where phone = ?", phone)
	var user User
	err := row.Scan(&user.Phone,&user.Name,&user.Card,&user.Sex)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
