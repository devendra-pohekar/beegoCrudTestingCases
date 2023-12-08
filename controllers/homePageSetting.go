package controllers

import (
	"crud/helpers"
	"crud/models"
	requestStruct "crud/requstStruct"
	"encoding/json"
	"fmt"

	"strings"

	_ "github.com/lib/pq" // PostgreSQL driver

	beego "github.com/beego/beego/v2/server/web"
)

type HTMLData struct {
	HTML string `json:"html"`
}
type HomeSettingController struct {
	beego.Controller
}

// RegisterSettings
// @Title After Login User Can Register Home Page settings
// @Description In this function after login can register Home page settings
// @Param	setting_data   formData 	file	false		"body for file"
// @Param	data_type   formData 	string	false		"body for file"
// @Param	section   formData 	string	false		"body for file"
// @Param   Authorization   header  string  true  "Bearer YourAccessToken"
// @Success 200 {object} models.HomePagesSettingTable
// @Failure 403
// @router /register_settings [post]
func (u *HomeSettingController) RegisterSettings() {
	logedIN := u.GetSession("user_login")
	if logedIN == "" {
		helpers.ApiFailedResponse(u.Ctx.ResponseWriter, "Please Login ")
		return
	}

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

// UpdateSettings
// @Title After Login User Can Update Home Page settings
// @Description In this function after login user  can update Home page settings
// @Param	setting_data   formData 	file	false		"body for file"
// @Param	data_type   formData 	string	false		"body for file"
// @Param	section   formData 	string	false		"body for file"
// @Param	setting_id   formData 	int		false		"body for file"
// @Param   Authorization   header  string  true  "Bearer YourAccessToken"
// @Success 200 {object} models.HomePagesSettingTable
// @Failure 403
// @router /update_settings [post]
func (u *HomeSettingController) UpdateSettings() {
	logedIN := u.GetSession("user_login")
	if logedIN == "" {
		helpers.ApiFailedResponse(u.Ctx.ResponseWriter, "Please Login ")
		return
	}
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

// FetchSettings
// @Title After Login User Can Fetch Data Home Page settings
// @Description In this function after login user  can Fetch Data Home page settings
// @Param   Authorization   header  string  true  "Bearer YourAccessToken"
// @Success 200 {object} models.HomePagesSettingTable
// @Failure 403
// @router /fetch_settings [post]
func (u *HomeSettingController) FetchSettings() {
	logedIN := u.GetSession("user_login")
	if logedIN == nil {
		helpers.ApiFailedResponse(u.Ctx.ResponseWriter, "Please Login ")
		return
	}
	var search requestStruct.HomeSeetingSearch
	if err := u.ParseForm(&search); err != nil {
		helpers.ApiFailedResponse(u.Ctx.ResponseWriter, "Parsing Data Error")
		return
	}
	json.Unmarshal(u.Ctx.Input.RequestBody, &search)

	result, _ := models.FetchSetting()

	if result != nil {

		helpers.ApiSuccessResponse(u.Ctx.ResponseWriter, result, "Home Setting Found Successfully", "", "")
		return
	}
	helpers.ApiFailedResponse(u.Ctx.ResponseWriter, "Not Found Data Please Try Again")
}

func (u *HomeSettingController) DeleteSetting() {

	logedIN := u.GetSession("user_login")
	if logedIN == "" {
		helpers.ApiFailedResponse(u.Ctx.ResponseWriter, "Please Login ")
		return
	}
	var home_settings requestStruct.HomeSeetingDelete
	if err := u.ParseForm(&home_settings); err != nil {
		helpers.ApiFailedResponse(u.Ctx.ResponseWriter, "Parsing Data Error")
		return
	}
	json.Unmarshal(u.Ctx.Input.RequestBody, &home_settings)
	result := models.HomePageSettingExistsDelete(home_settings)
	if result != 0 {
		helpers.ApiSuccessResponse(u.Ctx.ResponseWriter, "", "Home Page Setting  Deleted Successfully", "", "")
		return
	}
	helpers.ApiFailedResponse(u.Ctx.ResponseWriter, "Please Try Again")
}

// ExportFile
// @Title After Login User Can Export File in Home Page settings
// @Description In this function after login user  can Export File in Home page settings
// @Param file_type  formData string true "Here only select file within [XLSX,CSV,PDF]"
// @Param limit  formData int true "How Much you want to export data Ex.10"
// @Param   Authorization   header  string  true  "Bearer YourAccessToken"
// @Success 200 {object} models.HomePagesSettingTable
// @Failure 403
// @router /export [post]
func (c *HomeSettingController) ExportFile() {
	// logedIN := c.GetSession("user_login")
	// if logedIN == "" {
	// 	helpers.ApiFailedResponse(c.Ctx.ResponseWriter, "Please Login ")
	// 	return
	// }
	var fileTypes requestStruct.FileType
	if err := c.ParseForm(&fileTypes); err != nil {
		helpers.ApiFailedResponse(c.Ctx.ResponseWriter, "Parsing Data Error")
		return
	}

	json.Unmarshal(c.Ctx.Input.RequestBody, &fileTypes)

	create_file_type := strings.ToUpper(fileTypes.FileType)

	if create_file_type == "" {
		helpers.ApiFailedResponse(c.Ctx.ResponseWriter, "FILE TYPE SHOULD NOT BE EMPTY")
		return
	}

	if create_file_type == "XLSX" || create_file_type == "PDF" || create_file_type == "CSV" {
		res_data, _ := models.ExportData(fileTypes.Limit, fileTypes.SratingFrom)
		res_s, _ := helpers.TransformToKeyValuePairs(res_data)
		header := helpers.ExtractKeys(res_s)

		res_result, _ := helpers.CreateFile(res_s, header, "", "apps", create_file_type)
		if res_result == "" {
			helpers.ApiFailedResponse(c.Ctx.ResponseWriter, "File Not Create ! Try Again")
			return
		}
		helpers.ApiSuccessResponse(c.Ctx.ResponseWriter, res_result, "successfully Created file ", "", "")
		return
	}

	helpers.ApiFailedResponse(c.Ctx.ResponseWriter, "IT ONLY CONVERT WITHIN [PDF CSV,XLSX] FILE FORMAT")
}

// ImportFile
// @Title After Login User Can Import File in Home Page settings
// @Description In this function after login user  can Import File in Home page settings
// @Param import_type  formData file true "Here only select file within [XLSX,CSV]"
// @Param   Authorization   header  string  true  "Bearer YourAccessToken"
// @Success 200 {object} models.HomePagesSettingTable
// @Failure 403
// @router /import [post]
func (c *HomeSettingController) ImportFile() {
	// logedIN := c.GetSession("user_login")
	// if logedIN == "" {
	// 	helpers.ApiFailedResponse(c.Ctx.ResponseWriter, "Please Login ")
	// 	return
	// }
	file, fileHeader, err := c.GetFile("import_type")
	if err != nil {
		c.Ctx.WriteString("Error uploading file")
		return
	}

	uploadDir := "FILES/IMPORT"
	filePath, err := helpers.UploadFile(file, fileHeader, uploadDir)
	if err != nil {
		fmt.Println("Error uploading file:", err)
		helpers.ApiFailedResponse(c.Ctx.ResponseWriter, "Error uploading file")
		return
	}
	defer helpers.RemoveFileByPath(filePath)

	var allRows []map[string]interface{}

	switch {
	case strings.HasSuffix(filePath, ".xlsx"):
		allRows, err = helpers.ReadXLSXFile(filePath)
		if err != nil {
			fmt.Println("Error reading XLSX file:", err)
			helpers.ApiFailedResponse(c.Ctx.ResponseWriter, "Error reading XLSX file")
			return
		}
		result, update_id, _ := models.RegisterSettingBatchsss(requestStruct.HomeSeetingInsert{}, 35, filePath, allRows)
		if result == nil || update_id == nil {
			helpers.ApiSuccessResponse(c.Ctx.ResponseWriter, "", "File Imported Successfully", "", "")
			return
		}

		helpers.ApiFailedResponse(c.Ctx.ResponseWriter, "File Not Imported Please Try Again")

	case strings.HasSuffix(filePath, ".csv"):
		allRows, err = helpers.ReadCSVFile(filePath)
		if err != nil {
			fmt.Println("Error reading CSV file:", err)
			helpers.ApiFailedResponse(c.Ctx.ResponseWriter, "Error reading CSV file")
			return
		}
		result, update_id, _ := models.RegisterSettingBatchsss(requestStruct.HomeSeetingInsert{}, 100, filePath, allRows)
		if result != nil || update_id != nil {
			helpers.ApiSuccessResponse(c.Ctx.ResponseWriter, "", "File Imported Successfully", "", "")
			return
		}

	default:
		fmt.Println("Unsupported file format")
		helpers.ApiFailedResponse(c.Ctx.ResponseWriter, "Unsupported file format")
		return
	}
}
