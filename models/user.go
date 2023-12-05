package models

import (
	"crud/helpers"
	requestStruct "crud/requstStruct"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

func RegisterUser(u requestStruct.InsertUser) (interface{}, error) {
	db := orm.NewOrm()
	hash_pass := helpers.HashPassword(u.Password)
	res := UserMasterTable{
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Email:       u.Email,
		Password:    hash_pass,
		Mobile:      u.Mobile,
		CreatedDate: time.Now(),
	}
	_, err := db.Insert(&res)
	if err != nil {
		return nil, err
	}
	return res, nil

}

func LoginUsers(u requestStruct.LoginUser) (UserMasterTable, error) {
	db := orm.NewOrm()
	res := UserMasterTable{
		Email: u.Email,
	}
	result := db.Read(&res, "Email")
	if result != nil {
		return UserMasterTable{}, result
	}

	err := helpers.CheckPasswordHash(u.Password, res.Password)
	if err != nil {
		return UserMasterTable{}, err

	}
	return res, nil
}

func VerifyEmail(user_email string) (string, string, int, int) {
	db := orm.NewOrm()
	var user UserMasterTable
	err := db.Raw(`SELECT email,first_name ,is_verified,user_id FROM user_master_table WHERE email = ?`, user_email).QueryRow(&user)
	if err != nil {
		return "errror", "errorrr", 0, 0
	}
	return user.Email, user.FirstName, user.IsVerified, user.UserId
}

func FirstOTPUpdate(email, first_name, OTP string, user_id int) int {
	db := orm.NewOrm()
	emailID := email
	users := UserMasterTable{Email: emailID, FirstName: first_name, UserId: user_id}
	if db.Read(&users) == nil {
		users.OtpCode = strings.ToUpper(OTP)
		num, _ := db.Update(&users)
		if num == 0 {
			return 0
		}
	}
	return 1

}

func VerifyOTP(OTP string) (string, int) {
	db := orm.NewOrm()
	var user UserMasterTable
	err := db.Raw(`SELECT email ,user_id FROM user_master_table WHERE otp_code = ?`, strings.ToUpper(OTP)).QueryRow(&user)
	if err != nil {
		return "error in verifyotp query function", 0
	}

	return user.Email, user.UserId
}

func UpdateVerifiedStatus(email string, user_id int) int {
	db := orm.NewOrm()
	emailID := email

	users := UserMasterTable{Email: emailID, UserId: user_id}
	if db.Read(&users) == nil {
		users.IsVerified = 1
		users.OtpCode = ""
		num, _ := db.Update(&users)
		if num == 0 {
			return 0
		}
	}
	return 1

}

func UpdatePassword(email string, user_id int, updated_password string) int {
	db := orm.NewOrm()
	hash_updated_pass := helpers.HashPassword(updated_password)
	users := UserMasterTable{Email: email, UserId: user_id}
	if db.Read(&users) == nil {
		users.Password = hash_updated_pass
		users.OtpCode = ""
		num, _ := db.Update(&users)
		if num == 0 {
			return 0
		}
	}
	return 1

}
