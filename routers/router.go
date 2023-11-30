package routers

import (
	"crud/controllers"
	"crud/middelware"

	beego "github.com/beego/beego/v2/server/web"
)

func RoutersFunction() {

	userController := &controllers.UserController{}
	homeSettingController := &controllers.HomeSettingController{}

	user := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSRouter("/add_user", userController, "post:RegisterUser"),

			beego.NSRouter("/login", userController, "post:Login"),
			beego.NSRouter("/send_otp", userController, "post:SendMailForm"),
			beego.NSRouter("/verify_email", userController, "post:VerifyEmail"),
			beego.NSRouter("/send_otp_forgot", userController, "post:SendMailForForgotPassword"),
			beego.NSRouter("/verify_otp_forgot", userController, "post:ForgotPasswordUpdate"),
		),
		beego.NSNamespace("/homepage",
			beego.NSBefore(middelware.Auth),
			beego.NSRouter("/register_settings", homeSettingController, "post:RegisterSettings"),
			beego.NSRouter("/update_settings", homeSettingController, "post:UpdateSettings"),
			beego.NSRouter("/fetch_settings", homeSettingController, "post:FetchSettings"),
			beego.NSRouter("/delete_settings", homeSettingController, "post:DeleteSetting"),
		),
	)

	beego.AddNamespace(user)

}
