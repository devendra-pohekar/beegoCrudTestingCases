package controllers

import (
	"crud/helpers"
	"crud/models"
	requestStruct "crud/requstStruct"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
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

	log.Print(result, "========================")
	if result != nil {

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
	result := models.HomePageSettingExistsDelete(home_settings)
	if result != 0 {
		helpers.ApiSuccessResponse(u.Ctx.ResponseWriter, "", "Home Page Setting  Deleted Successfully", "", "")
		return
	}
	helpers.ApiFailedResponse(u.Ctx.ResponseWriter, "Please Try Again")
}

func (c *HomeSettingController) InsertData() {
	res_data, _ := models.FetchSetting()

	res_s, _ := TransformToKeyValuePairs(res_data)
	headers := []string{"section", "data_type", "setting_data", "created_date", "updated_date", "created_by"}
	res_result, _ := helpers.CreateFile(res_s, headers, "", "apps", "xlsx")
	log.Print(res_result)

	// data, ok := res_data.([][]interface{})
	// if !ok {

	// 	c.Data["json"] = map[string]interface{}{"error": "Invalid data type"}
	// 	c.ServeJSON()
	// 	return
	// }
	// xlsx := excelize.NewFile()

	// sheetName := "Sheet1"
	// xlsx.SetSheetName("Sheet1", sheetName)

	// for colIndex, headerValue := range header {
	// 	cellName := fmt.Sprintf("%c%d", 'A'+colIndex, 1)
	// 	xlsx.SetCellValue(sheetName, cellName, headerValue)
	// }

	// for rowIndex, row := range data {
	// 	for colIndex, cellValue := range row {
	// 		cellName := fmt.Sprintf("%c%d", 'A'+colIndex, rowIndex+2)
	// 		xlsx.SetCellValue(sheetName, cellName, cellValue)
	// 	}
	// }

	// filePath := "excel/file.xlsx"
	// if err := xlsx.SaveAs(filePath); err != nil {
	// 	c.Data["json"] = map[string]interface{}{"error": err.Error()}
	// } else {
	// 	c.Data["json"] = map[string]interface{}{"success": true, "filePath": filePath}
	// }

	// c.ServeJSON()
}

func (c *HomeSettingController) InsertData1() {

	data := []struct {
		Section     string    `json:"section"`
		DataType    string    `json:"data_type"`
		SettingData string    `json:"setting_data"`
		CreatedDate time.Time `json:"created_date"`
		UpdatedDate time.Time `json:"updated_date"`
		CreatedBy   string    `json:"created_by"`
	}{
		{
			Section:     "left pannel",
			DataType:    "html",
			SettingData: "<div><p>Welcome to The Website</p></div>",
			CreatedDate: time.Now(),
			UpdatedDate: time.Now(),
			CreatedBy:   "Dwarkesh Patel",
		},
		{
			Section:     "middel container",
			DataType:    "logo",
			SettingData: "uploads/Home/files/logo/1701150314464851586.jpg",
			CreatedDate: time.Now(),
			UpdatedDate: time.Now(),
			CreatedBy:   "Dwarkesh Patel",
		},
	}

	xlsx := excelize.NewFile()

	sheetName := "Sheet1"
	xlsx.SetSheetName("Sheet1", sheetName)

	// Adding headers
	headers := []string{"Section", "DataType", "SettingData", "CreatedDate", "UpdatedDate", "CreatedBy"}
	for colIndex, header := range headers {
		cellName := fmt.Sprintf("%c%d", 'A'+colIndex, 1)
		xlsx.SetCellValue(sheetName, cellName, header)
	}

	// Adding data rows
	for rowIndex, row := range data {
		cellName := fmt.Sprintf("A%d", rowIndex+2)
		xlsx.SetCellValue(sheetName, cellName, row.Section)
		cellName = fmt.Sprintf("B%d", rowIndex+2)
		xlsx.SetCellValue(sheetName, cellName, row.DataType)
		cellName = fmt.Sprintf("C%d", rowIndex+2)
		xlsx.SetCellValue(sheetName, cellName, row.SettingData)
		cellName = fmt.Sprintf("D%d", rowIndex+2)
		xlsx.SetCellValue(sheetName, cellName, row.CreatedDate)
		cellName = fmt.Sprintf("E%d", rowIndex+2)
		xlsx.SetCellValue(sheetName, cellName, row.UpdatedDate)
		cellName = fmt.Sprintf("F%d", rowIndex+2)
		xlsx.SetCellValue(sheetName, cellName, row.CreatedBy)
	}

	// Save the Excel file
	filePath := "excel/file.xlsx" // Provide the desired file path
	if err := xlsx.SaveAs(filePath); err != nil {
		fmt.Println("Error saving Excel file:", err)
		c.Data["json"] = map[string]interface{}{"error": err.Error()}
	} else {
		fmt.Println("Excel file saved successfully:", filePath)
		c.Data["json"] = map[string]interface{}{"success": true, "filePath": filePath}
	}

	c.ServeJSON()
}
