package controllers

import (
	"crud/helpers"
	"crud/models"
	requestStruct "crud/requstStruct"
	"encoding/json"
	"log"

	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type HTMLData struct {
	HTML string `json:"html"`
}
type HomeSettingController struct {
	beego.Controller
}

func (u *HomeSettingController) RegisterSettings() {
	var settings requestStruct.HomeSeetingInsert
	var filePath string

	if err := u.ParseForm(&settings); err != nil {
		helpers.ApiFailedResponse(u.Ctx.ResponseWriter, "Parsing Data Error")
		return
	}
	json.Unmarshal(u.Ctx.Input.RequestBody, &settings)

	data_types := strings.ToUpper(settings.DataType)
	// uploadDir := os.Getenv("uploadHomePageImages")
	uploadDir := "uploads/Home/files/images"
	if data_types == "LOGO" {
		// uploadDir = os.Getenv("uploadHomePageLogos")
		uploadDir = "uploads/Home/files/logo"
	} else if data_types != "BANNER" {
		filePath = ""
	}
	if data_types == "LOGO" || data_types == "BANNER" {
		file, fileHeader, err := u.GetFile("setting_data")
		if err != nil {
			helpers.ApiFailedResponse(u.Ctx.ResponseWriter, "File Getting Error")
			return
		}

		filePath, err = helpers.UploadFile(file, fileHeader, uploadDir)
		if err != nil {
			helpers.ApiFailedResponse(u.Ctx.ResponseWriter, err.Error())
			return
		}
	}

	tokenData := helpers.GetTokenClaims(u.Ctx)
	userID := tokenData["User_id"]
	result, _ := models.RegisterSetting(settings, userID.(float64), filePath)
	if result != 0 {
		helpers.ApiSuccessResponse(u.Ctx.ResponseWriter, "", "Home Page Settings Register Successfully", "", "")
		return
	}

	helpers.ApiFailedResponse(u.Ctx.ResponseWriter, "Please Try Again")
}

func (u *HomeSettingController) UpdateSettings() {
	var settings requestStruct.HomeSeetingUpdate
	var filePath string

	if err := u.ParseForm(&settings); err != nil {
		helpers.ApiFailedResponse(u.Ctx.ResponseWriter, "Parsing Data Error")
		return
	}

	json.Unmarshal(u.Ctx.Input.RequestBody, &settings)
	data_types := strings.ToUpper(settings.DataType)

	// uploadDir := os.Getenv("uploadHomePageImages")
	uploadDir := "uploads/Home/files/images"

	if data_types == "LOGO" {
		// uploadDir = os.Getenv("uploadHomePageLogos")
		uploadDir = "uploads/Home/files/logo"

	} else if data_types != "BANNER" {
		filePath = ""
	}

	if data_types == "LOGO" || data_types == "BANNER" {
		file, fileHeader, err := u.GetFile("setting_data")
		if err != nil {
			helpers.ApiFailedResponse(u.Ctx.ResponseWriter, "File Getting Error")
			return
		}

		filePath, err = helpers.UploadFile(file, fileHeader, uploadDir)
		if err != nil {
			helpers.ApiFailedResponse(u.Ctx.ResponseWriter, "File Uploading Error")
			return
		}
	}

	tokenData := helpers.GetTokenClaims(u.Ctx)
	userID := tokenData["User_id"]
	result, _ := models.UpdateSetting(settings, filePath, userID.(float64))

	if result != 0 {
		helpers.ApiSuccessResponse(u.Ctx.ResponseWriter, "", "Home Page Settings Updated  Successfully", "", "")
		return
	}

	helpers.ApiFailedResponse(u.Ctx.ResponseWriter, "Please Try Again")
}

func (u *HomeSettingController) FetchSettings() {
	var search requestStruct.HomeSeetingSearch
	if err := u.ParseForm(&search); err != nil {
		helpers.ApiFailedResponse(u.Ctx.ResponseWriter, "Parsing Data Error")
		return
	}
	json.Unmarshal(u.Ctx.Input.RequestBody, &search)

	result, _ := models.FetchSetting()

	if result != nil {
		u.Data["json"] = result
		helpers.ApiSuccessResponse(u.Ctx.ResponseWriter, result, "Home Setting Found Successfully", "", "")
		return
	}
	helpers.ApiFailedResponse(u.Ctx.ResponseWriter, "Not Found Data Please Try Again")
}

func (u *HomeSettingController) DeleteSetting() {
	var home_settings requestStruct.HomeSeetingDelete
	if err := u.ParseForm(&home_settings); err != nil {
		helpers.ApiFailedResponse(u.Ctx.ResponseWriter, "Parsing Data Error")
		return
	}
	json.Unmarshal(u.Ctx.Input.RequestBody, &home_settings)
	log.Print("===================", home_settings.SettingId)
	result := models.HomePageSettingExistsDelete(home_settings)
	if result != 0 {
		helpers.ApiSuccessResponse(u.Ctx.ResponseWriter, "", "Home Page Setting  Deleted Successfully", "", "")
		return
	}
	helpers.ApiFailedResponse(u.Ctx.ResponseWriter, "Please Try Again")
}
