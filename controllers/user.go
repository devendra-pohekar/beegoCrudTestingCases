package controllers

import (
	"crud/helpers"
	"crud/models"
	requestStruct "crud/requstStruct"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	beego "github.com/beego/beego/v2/server/web"

	"github.com/dgrijalva/jwt-go"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

type Users struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegister struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type JwtClaim struct {
	Email  string `json:"user_email"`
	UserID int    `json:"user_id"`
	jwt.StandardClaims
}

var secretKey = []byte("devendra_secretkey")

// Login...
// @Title Login User
// @Description This is a Login API for User
// @Param body body requestStruct.LoginUser false
// @Success 200{object}models.UserMasterTable
// @Failure 403
// @router /login [post]
func (c *UserController) Login() {
	var user requestStruct.LoginUser
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &user); err != nil {
		c.CustomAbort(http.StatusBadRequest, "Invalid JSON format")
		return
	}
	loginUserData, err := models.LoginUsers(user)
	if err != nil {
		helpers.ApiFailedResponse(c.Ctx.ResponseWriter, "Invalid Email And Password ! Try Again")
		return
	}

	if loginUserData.IsVerified == 0 {
		helpers.ApiFailedResponse(c.Ctx.ResponseWriter, "Please Verified Email Address")
		return
	}

	tokenExpire := time.Now().Add(1 * time.Hour)
	c.SetSession("user_login", loginUserData.Email)
	claims := &JwtClaim{Email: loginUserData.Email, UserID: loginUserData.UserId, StandardClaims: jwt.StandardClaims{
		ExpiresAt: tokenExpire.Unix(),
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		c.CustomAbort(http.StatusBadRequest, fmt.Sprintf("Error signing token: %s", err.Error()))
		return
	}

	data := map[string]interface{}{"User_Data": token.Claims, "Token": tokenString}
	session_user_ := c.GetSession("user_login")
	c.Data["json"] = map[string]interface{}{"data": data, "session": session_user_}
	c.ServeJSON()
}

// RegisterUser...
// @Title Register User
// @Description This api used to register the new user
// @Param body body  requestStruct.InsertUser false "sample of swagger register user details field"
// @Success 200 {object} models.RegisterUserTable
// @Failure 403
// @router /add_user [post]
func (u *UserController) RegisterUser() {
	var user requestStruct.InsertUser
	if err := u.ParseForm(&user); err != nil {
		helpers.ApiFailedResponse(u.Ctx.ResponseWriter, "Parsing Data Error")
		return
	}

	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	user_email, _, _, _ := models.VerifyEmail(user.Email)

	if user_email == user.Email {
		helpers.ApiFailedResponse(u.Ctx.ResponseWriter, "Use Another Email ! This Email Address Already Exists")
		return
	}
	result, _ := models.RegisterUser(user)
	if result != nil {
		helpers.ApiSuccessResponse(u.Ctx.ResponseWriter, "", "Register Successfully User Please Login Now", "", "")
		return
	}
	helpers.ApiFailedResponse(u.Ctx.ResponseWriter, "Please Try Again")
}

// func (u *UserController) LoginUser() {
// 	var user requestStruct.LoginUser
// 	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
// 	result := models.LoginUser(user)
// 	if result != 0 {
// 		u.Data["json"] = result
// 		// helpers.ApiSuccessResponse(u.Controller, "", "Login Successfully User")
// 	} else {
// 		// helpers.ApiFailedResponse(u.Controller, "Invalid Email and Password Please Try Again")
// 	}
// }

// SendMailForm
// @Title Send Mail on the User Register Email Address for email verification process
// @Description In Email verification of user ,in this process we send an email on the register mail address for verification .the given email id is valid or not
// @Param body body requestStruct.SendMailUser false "In this process we take email address and send email on the register email address with code "
// @Success 200 {object} models.UserMasterTable
// @Failure 403
// @router /send_otp [post]
func (c *UserController) SendMailForm() {
	var requestData requestStruct.SendMailUser

	if err := c.ParseForm(&requestData); err != nil {
		helpers.ApiFailedResponse(c.Ctx.ResponseWriter, "Parsing Data Error")
		return
	}

	json.Unmarshal(c.Ctx.Input.RequestBody, &requestData)
	email, user_first_name, is_verified, user_id := models.VerifyEmail(requestData.Email)
	if email == requestData.Email && is_verified == 0 {
		result, _ := helpers.SendOTpOnMail(requestData.Email, user_first_name)
		models.FirstOTPUpdate(email, user_first_name, result, user_id)
		helpers.ApiSuccessResponse(c.Ctx.ResponseWriter, "", "Verification Mail Send On The Given User Email Address ,Please verified first", "", "")
		return
	}

	if email == requestData.Email && is_verified == 1 {
		helpers.ApiFailedResponse(c.Ctx.ResponseWriter, "User Already Verified")
		return

	}
	helpers.ApiFailedResponse(c.Ctx.ResponseWriter, "Please Provide Valid Email Address ! , Try Again")

}

// VerifyEmail
// @Title Verify User Email address after Registration
// @Description This function work after Registration of user to check the user email address is valid or not .Here we verify the code that we send already on user email address.if user email is verified than user can perform operation after login
// @Param body body requestStruct.EmailVerfiy false "we verify the code than user send "
// @Success 200 {object} models.UserMasterTable
// @Failure 403
// @router /verify_email [post]
func (c *UserController) VerifyEmail() {
	var requestData requestStruct.EmailVerfiy

	if err := c.ParseForm(&requestData); err != nil {
		helpers.ApiFailedResponse(c.Ctx.ResponseWriter, "Parsing Data Error")
		return
	}
	json.Unmarshal(c.Ctx.Input.RequestBody, &requestData)
	if requestData.OTP == "" {
		helpers.ApiFailedResponse(c.Ctx.ResponseWriter, "OTP Should Not be Empty")
	}
	user_email, user_id := models.VerifyOTP(requestData.OTP)
	if user_email != "" && user_id != 0 {
		models.UpdateVerifiedStatus(user_email, user_id)
		helpers.ApiSuccessResponse(c.Ctx.ResponseWriter, "", "User Verified Successfully ", "", "")
	}

}

// SendMailForForgotPassword
// @Title Send Mail for ForgotPassword
// @Description This function work to send an Email on Register User Email Address with code
// @Param body body  requestStruct.SendMailForgotPassword false "here users email address send as parameter for email sending"
// @Success 200 {object} models.UserMasterTable
// @Failure 403
// @router /send_otp_forgot [post]
func (c *UserController) SendMailForForgotPassword() {
	var requestData requestStruct.SendMailForgotPassword
	if err := c.ParseForm(&requestData); err != nil {
		helpers.ApiFailedResponse(c.Ctx.ResponseWriter, "Parsing Data Error")
		return
	}

	json.Unmarshal(c.Ctx.Input.RequestBody, &requestData)

	email, user_first_name, is_verified, user_id := models.VerifyEmail(requestData.Email)
	if email == requestData.Email && is_verified == 1 {
		result, _ := helpers.SendOTpOnMail(requestData.Email, user_first_name)
		models.FirstOTPUpdate(email, user_first_name, result, user_id)
		helpers.ApiSuccessResponse(c.Ctx.ResponseWriter, email, "OTP Verification Mail Send On The Register User Email Address ,Please verified OTP", "", "")
		return
	}

	helpers.ApiFailedResponse(c.Ctx.ResponseWriter, "Please Provide Valid Email Address ! , Try Again")

}

// ForgotPasswordUpdate
// @Title SetNewPassword
// @Description This function used for to update or to set new password.When ths user send the otp and newpassword
// @Param body body requestStruct.ForgotPassword false "here user send the otp and newpassword ,than this function update the password as newpassword for user"
// @Success 200 {object} models.UserMasterTable
// @Failure 403
// @router /verify_otp_forgot [post]
func (c *UserController) ForgotPasswordUpdate() {
	var requestData requestStruct.ForgotPassword
	if err := c.ParseForm(&requestData); err != nil {
		helpers.ApiFailedResponse(c.Ctx.ResponseWriter, "Parsing Data Error")
		return
	}
	json.Unmarshal(c.Ctx.Input.RequestBody, &requestData)
	if requestData.OTP == "" || requestData.NewPassword == "" {
		helpers.ApiFailedResponse(c.Ctx.ResponseWriter, "OTP And Password Should Not be Empty ")
		return
	}
	user_email, user_id := models.VerifyOTP(requestData.OTP)
	if user_email != "" && user_id != 0 {
		models.UpdatePassword(user_email, user_id, requestData.NewPassword)
		helpers.ApiSuccessResponse(c.Ctx.ResponseWriter, "", "Password Change Successfully ", "", "")
		return
	}
	helpers.ApiFailedResponse(c.Ctx.ResponseWriter, "OTP IS Expired PLEASE GO ON FORGOT PASSWORD SECTION")

}

// LogoutUser
// @Title Logout User
// @Description This function used for to logout user
// @Success 200 {object} models.UserMasterTable
// @Failure 403
// @router /logout_user [post]
func (u *UserController) LogoutUser() {
	err := u.DestroySession()
	if err != nil {
		helpers.ApiFailedResponse(u.Ctx.ResponseWriter, "Not Logout User")
		return
	}
	helpers.ApiSuccessResponse(u.Ctx.ResponseWriter, "", "User Logout Successfully", "", "")
}
