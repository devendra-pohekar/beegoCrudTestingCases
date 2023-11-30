package test

import (
	"bytes"
	"crud/controllers"
	"crud/middelware"
	"log"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	_ "github.com/lib/pq"

	beego "github.com/beego/beego/v2/server/web"

	"github.com/beego/beego/v2/client/orm"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=postgres password=Dev@123 host=localhost port=5432 dbname=golang_practice sslmode=disable")
	orm.RunSyncdb("default", false, true)

}

func RequestTestingFunction(token, endPoint string, method string, jsonStr []byte, controllerFunction string, controller beego.ControllerInterface) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, endPoint, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal(err)
	}

	res := httptest.NewRecorder()
	router := beego.NewControllerRegister()
	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", token)
		router.InsertFilter(endPoint, beego.BeforeRouter, middelware.Auth, beego.WithCaseSensitive(true))

	}
	router.Add(endPoint, controller, beego.WithRouterMethods(controller, controllerFunction))
	router.ServeHTTP(res, req)
	return res
}

func TestRegisterUser(t *testing.T) {
	controller := &controllers.UserController{}
	token := ""
	endPoint := "/v1/user/add_user"
	var jsonStr = []byte(`{"first_name":"Testing FirstName", "last_name":"Testing LastName", "email":"sampletesting123@gmail.com", "mobile":"1234324343543", "password":"1234567"}`)
	result := RequestTestingFunction(token, endPoint, "POST", jsonStr, "post:RegisterUser", controller)
	log.Print(result.Body)

}

func TestLoginUser(t *testing.T) {
	controller := &controllers.UserController{}
	endPoint := "/v1/user/login"
	token := ""
	var jsonStr = []byte(`{"email":"devendrapohekar.siliconithub@gmail.com","password":"Dev@123"}`)
	result := RequestTestingFunction(token, endPoint, "POST", jsonStr, "post:Login", controller)
	log.Println(result.Body)

}

func TestSendOTPonMails(t *testing.T) {
	controller := &controllers.UserController{}
	endPoint := "/v1/user/send_otp"
	token := ""
	var jsonStr = []byte(`{"email":"devendrapohekar.siliconithub@gmail.com", "user_name":"Devendra Pohekar"}`)

	result := RequestTestingFunction(token, endPoint, "POST", jsonStr, "post:SendMailForm", controller)
	log.Println(result.Body)
}

func TestVerifyEmailOTP(t *testing.T) {
	controller := &controllers.UserController{}
	endPoint := "/v1/user/verify_email"
	token := ""
	var jsonStr = []byte(`{"email":"devendrapohekar.siliconithub@gmail.com", "otp":"3D61qW1u"}`)
	result := RequestTestingFunction(token, endPoint, "POST", jsonStr, "post:VerifyEmail", controller)
	log.Println(result.Body)

}

func TestForgotPasswordSendMail(t *testing.T) {
	controller := &controllers.UserController{}
	endPoint := "/v1/user/send_otp_forgot"
	token := ""
	var jsonStr = []byte(`{"email":"devendrapohekar.siliconithub@gmail.com"}`)
	result := RequestTestingFunction(token, endPoint, "POST", jsonStr, "post:SendMailForForgotPassword", controller)
	log.Println(result.Body)

}

func TestForgotPassword(t *testing.T) {
	controller := &controllers.UserController{}
	endPoint := "/v1/user/verify_otp_forgot"
	token := ""
	var jsonStr = []byte(`{"otp":"Kmqqlvyb","new_password":"Devendra@123"}`)
	result := RequestTestingFunction(token, endPoint, "POST", jsonStr, "post:ForgotPasswordUpdate", controller)
	log.Println(result.Body)

}
