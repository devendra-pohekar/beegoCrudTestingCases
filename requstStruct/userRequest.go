package requestStruct

type InsertUser struct {
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	Mobile    string `json:"mobile" form:"mobile"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type MobileNumber struct {
	MobileNumber string `json:"mobile_number"`
}

type SendMailUser struct {
	Email    string `json:"email" form :"email"`
	UserName string `json:"user_name" form:"user_name"`
}

type EmailVerfiy struct {
	Email string `json:"email" form:"email"`
	OTP   string `json:"otp" form:"otp"`
}

type SendMailForgotPassword struct {
	Email string `json:"email" form :"email"`
}

type ForgotPassword struct {
	OTP         string `json:"otp" form:"otp"`
	NewPassword string `json:"new_password" form :"new_password"`
}
