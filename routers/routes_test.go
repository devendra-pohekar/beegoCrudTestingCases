package routers_test

// import (
// 	"bytes"
// 	"crud/controllers"
// 	"log"
// 	"net/http"
// 	"net/http/httptest"
// 	"path/filepath"
// 	"runtime"
// 	"testing"

// 	_ "github.com/lib/pq"

// 	beego "github.com/beego/beego/v2/server/web"

// 	// "github.com/beego/beego/v2/client/orm"
// 	"github.com/beego/beego/v2/client/orm"
// 	// . "github.com/smartystreets/goconvey/convey"
// )

// func init() {
// 	_, file, _, _ := runtime.Caller(0)
// 	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
// 	beego.TestBeegoInit(apppath)
// 	orm.RegisterDriver("postgres", orm.DRPostgres)
// 	orm.RegisterDataBase("default", "postgres", "user=postgres password=Dev@123 host=localhost port=5432 dbname=golang_practice sslmode=disable")
// 	orm.RunSyncdb("default", false, true)

// }

// func RequestTestingFunction(endPoint string, method string, jsonStr []byte, controllerFunction string, controller beego.ControllerInterface) *httptest.ResponseRecorder {
// 	req, err := http.NewRequest(method, endPoint, bytes.NewBuffer(jsonStr))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	req.Header.Set("Content-Type", "application/json")
// 	res := httptest.NewRecorder()
// 	router := beego.NewControllerRegister()
// 	router.Add(endPoint, controller, beego.WithRouterMethods(controller, controllerFunction))
// 	router.ServeHTTP(res, req)
// 	return res
// }

// func TestRegister(t *testing.T) {
// 	controller := &controllers.UserController{}
// 	endPoint := "/v1/user/add_user"
// 	var jsonStr = []byte(`{"first_name":"Dwarkesh", "last_name":"Patel", "email":"dwarkesh0007@gmail.com", "mobile":"1234324343543", "password":"1234567"}`)
// 	result := RequestTestingFunction(endPoint, "POST", jsonStr, "post:RegisterUser", controller)

// }

// func TestLoginUser(t *testing.T) {
// 	controller := &controllers.UserController{}
// 	endPoint := "/v1/user/login"
// 	var jsonStr = []byte(`{"email":"dev@11111@gmail.com","password":"1234567"}`)
// 	result := RequestTestingFunction(endPoint, "POST", jsonStr, "post:Login", controller)
// 	log.Println(result.Body)

// }

// func TestSendOTPonMails(t *testing.T) {
// 	controller := &controllers.UserController{}
// 	endPoint := "/v1/user/send_otp"
// 	var jsonStr = []byte(`{"email":"devendrapohekar.siliconithub@gmail.com", "user_name":"Devendra Pohekar"}`)

// 	result := RequestTestingFunction(endPoint, "POST", jsonStr, "post:SendMailForm", controller)
// 	log.Println(result.Body)
// }

// func TestVerifyEmailOTP(t *testing.T) {
// 	controller := &controllers.UserController{}
// 	endPoint := "/v1/user/verify_email"
// 	var jsonStr = []byte(`{"email":"devendrapohekar.siliconithub@gmail.com", "otp":"3D61qW1u"}`)
// 	result := RequestTestingFunction(endPoint, "POST", jsonStr, "post:VerifyEmail", controller)
// 	log.Println(result.Body)

// }

// func TestForgotPasswordSendMail(t *testing.T) {
// 	controller := &controllers.UserController{}
// 	endPoint := "/v1/user/send_otp_forgot"
// 	var jsonStr = []byte(`{"email":"devendrapohekar.siliconithub@gmail.com"}`)
// 	result := RequestTestingFunction(endPoint, "POST", jsonStr, "post:SendMailForForgotPassword", controller)
// 	log.Println(result.Body)

// }

// func TestForgotPassword(t *testing.T) {
// 	controller := &controllers.UserController{}
// 	endPoint := "/v1/user/verify_otp_forgot"
// 	var jsonStr = []byte(`{"otp":"Kmqqlvyb","new_password":"Devendra@123"}`)
// 	result := RequestTestingFunction(endPoint, "POST", jsonStr, "post:ForgotPasswordUpdate", controller)
// 	log.Println(result.Body)

// }
