package models

import (
	"crud/helpers"
	requestStruct "crud/requstStruct"
	"os"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

func RegisterSetting(c requestStruct.HomeSeetingInsert, user_id float64, file_path interface{}) (int, error) {
	db := orm.NewOrm()
	if file_path == "" {

		file_path = c.SettingData

	}

	res := HomePagesSettingTable{
		Section:     c.Section,
		DataType:    c.DataType,
		UniqueCode:  "",
		SettingData: file_path.(string),
		CreatedBy:   int(user_id),
		UpdatedBy:   0,
		// UpdatedDate: time.Now(),
		CreatedDate: time.Now(),
	}

	_, err := db.Insert(&res)
	if err != nil {
		return 0, err
	}
	lastInsertID := res.PageSettingId
	UpdateUniqueCode(lastInsertID)
	return lastInsertID, nil
}

func UpdateUniqueCode(user_id int) (int64, error) {
	db := orm.NewOrm()

	unique_codes := helpers.UniqueCode(user_id, os.Getenv("homePageModule"))
	home_page_setting := HomePagesSettingTable{PageSettingId: user_id}
	if db.Read(&home_page_setting) == nil {
		home_page_setting.UniqueCode = unique_codes
		if num, err := db.Update(&home_page_setting); err == nil {
			return num, nil
		}
	}
	return 1, nil
}

func UpdateSetting(c requestStruct.HomeSeetingUpdate, file_path interface{}, user_id float64) (int64, error) {
	db := orm.NewOrm()
	page_setting_id := c.SettingId
	homePageSetting, setting_data_type, err := FetchPageSettingByID(page_setting_id)
	if err != nil {
		return 0, err
	}

	if file_path == "" {
		file_path = c.SettingData
	}
	setting_dataType := strings.ToUpper(setting_data_type)
	if setting_dataType == "LOGO" || setting_dataType == "BANNER" {
		file_name, file_directory := helpers.SplitFilePath(homePageSetting)
		helpers.RemoveFile(file_name, file_directory)

	}
	homePageData := HomePagesSettingTable{PageSettingId: page_setting_id,
		UpdatedBy:   int(user_id),
		UpdatedDate: time.Now(),
		DataType:    c.DataType,
		Section:     c.Section,
		SettingData: file_path.(string),
	}
	if num, err := db.Update(&homePageData, "updated_by", "updated_date", "data_type", "section", "setting_data"); err == nil {
		return num, nil
	}
	return 1, nil

}

func FetchPageSettingByID(pageSettingID int) (string, string, error) {
	db := orm.NewOrm()
	var pageSetting HomePagesSettingTable
	err := db.Raw(`SELECT  setting_data,data_type FROM home_pages_setting_table WHERE page_setting_id = ?`, pageSettingID).QueryRow(&pageSetting)
	if err != nil {
		return "Some errro occured in fetch page setting by ID function", "some errror", err
	}
	return pageSetting.SettingData, pageSetting.DataType, nil
}

func DeleteSetting(page_setting_id int) int {
	db := orm.NewOrm()
	setting := HomePagesSettingTable{PageSettingId: page_setting_id}
	if _, err := db.Delete(&setting); err == nil {
		return 1
	}
	return 0

}

func HomePageSettingExistsDelete(u requestStruct.HomeSeetingDelete) int {
	page_setting_id := u.SettingId

	page_setting_data, page_setting_type, err := FetchPageSettingByID(page_setting_id)
	if err != nil {
		return 0
	}

	if strings.ToUpper(page_setting_type) == "LOGO" || strings.ToUpper(page_setting_type) == "BANNER" {
		file_name, file_directory := helpers.SplitFilePath(page_setting_data)
		helpers.RemoveFile(file_name, file_directory)
	}

	DeleteSetting(page_setting_id)
	return 1

}

func FetchSetting() (interface{}, error) {
	db := orm.NewOrm()
	var homeResponse []struct {
		Section     string    `json:"section"`
		DataType    string    `json:"data_type"`
		SettingData string    `json:"setting_data"`
		CreatedDate time.Time `json:"created_date"`
		UpdatedDate time.Time `json:"updated_date"`
		CreatedBy   string    `json:"created_by"`
	}
	_, err := db.Raw(`SELECT hpst.section, hpst.data_type, hpst.setting_data,hpst.created_date, hpst.updated_date ,concat(umt.first_name,' ',umt.last_name) as created_by  FROM home_pages_setting_table as hpst LEFT JOIN user_master_table as umt ON umt.user_id = hpst.created_by`).QueryRows(&homeResponse)

	if err != nil {
		return nil, err
	}
	if len(homeResponse) == 0 {
		return "Not Found Cars", nil
	}
	return homeResponse, nil
}

// func FetchSetting() (interface{}, error) {
// 	db := orm.NewOrm()
// 	var homeResponse []struct {
// 		Section     string    `json:"section"`
// 		DataType    string    `json:"data_type"`
// 		SettingData string    `json:"setting_data"`
// 		CreatedDate time.Time `json:"created_date"`
// 		UpdatedDate time.Time `json:"updated_date"`
// 		CreatedBy   string    `json:"created_by"`
// 	}

// 	// Execute the SQL query and check for errors
// 	_, err := db.Raw(`
// 		SELECT hpst.section, hpst.data_type, hpst.setting_data, hpst.created_date, hpst.updated_date,umt.first_name
// 		FROM home_pages_setting_table AS hpst
// 		LEFT JOIN user_master_table AS umt ON umt.user_id = hpst.created_by
// 	`).QueryRows(&homeResponse)

// 	// Check for query execution errors
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Check if no data is found
// 	if len(homeResponse) == 0 {
// 		return "Not Found Cars", nil
// 	}

// 	// Return the data if successful
// 	return homeResponse, nil
// }
