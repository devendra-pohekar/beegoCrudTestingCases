package models

import (
	"crud/helpers"
	requestStruct "crud/requstStruct"
	"errors"
	"fmt"
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
	_, err := db.Raw(`SELECT hpst.section, hpst.data_type, hpst.setting_data,hpst.created_date, hpst.updated_date ,concat(umt.first_name,' ',umt.last_name) as created_by  FROM home_pages_setting_table as hpst LEFT JOIN user_master_table as umt ON umt.user_id = hpst.created_by ORDER BY hpst.created_date DESC`).QueryRows(&homeResponse)

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

func FetchSetting1() []HomePagesSettingTable {
	o := orm.NewOrm()
	var page_settings []HomePagesSettingTable

	qs := o.QueryTable("home_page_setting_table")

	_, err := qs.All(&page_settings)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return page_settings
}
func UpdateSettings(c requestStruct.HomeSeetingUpdate, file_path interface{}, user_id float64) (int64, error) {
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

func RegisterSettingBatch(c requestStruct.HomeSeetingInsert, user_id float64, filePath string, rows []map[string]interface{}) ([]int64, error) {
	db := orm.NewOrm()
	var lastInsertIDs []int64
	for _, row := range rows {

		section, ok := row["section"].(string)

		if !ok {
			section = ""
		}

		dataType, ok := row["data_type"].(string)
		if !ok {
			dataType = ""

		}

		settingData, ok := row["setting_data"].(string)
		if !ok {
			settingData = ""

		}

		// createdDateStr, ok := row["created_date"].(string)
		// if !ok {
		// 	createdDateStr = ""
		// 	log.Print("4444444444444444444444")

		// }

		// updatedDateStr, ok := row["updated_date"].(string)
		// if !ok {
		// 	updatedDateStr = ""
		// 	log.Print("5555555555555555555555")

		// }
		// createdDate, errs := time.Parse("2006-01-02 15:04:05.999999-07:00", createdDateStr)
		// if errs != nil {
		// 	return nil, errs
		// }
		// updatedDate, errs := time.Parse("2006-01-02 15:04:05.999999-07:00", updatedDateStr)
		// if errs != nil {
		// 	return nil, errs
		// }
		// log.Print(createdDate, "==========================", updatedDate)
		res := HomePagesSettingTable{
			Section:     section,
			DataType:    dataType,
			UniqueCode:  "",
			SettingData: settingData,
			CreatedBy:   int(user_id),
			UpdatedBy:   0,
			CreatedDate: time.Now(),
			UpdatedDate: time.Now(),
		}

		_, err := db.Insert(&res)
		if err != nil {
			return nil, err
		}

		lastInsertIDs = append(lastInsertIDs, int64(res.PageSettingId))

		UpdateUniqueCode(res.PageSettingId)

	}
	helpers.RemoveFileByPath(filePath)

	return lastInsertIDs, nil
}

func RegisterSettingBatchcsv(c requestStruct.HomeSeetingInsert, user_id float64, filePath string, rows []map[string]interface{}) ([]int64, error) {
	db := orm.NewOrm()

	var lastInsertIDs []int64

	for _, row := range rows {
		section, ok := row["section"].(string)
		if !ok {
			return nil, errors.New("missing 'section' in row")
		}

		dataType, ok := row["data_type"].(string)
		if !ok {
			return nil, errors.New("missing 'data_type' in row")
		}

		settingData, ok := row["setting_data"].(string)
		if !ok {
			return nil, errors.New("missing 'setting_data' in row")
		}

		res := HomePagesSettingTable{
			Section:     section,
			DataType:    dataType,
			UniqueCode:  "",
			SettingData: settingData,
			CreatedBy:   int(user_id),
			UpdatedBy:   0,
			CreatedDate: time.Now(),
			UpdatedDate: time.Now(),
		}

		_, err := db.Insert(&res)
		if err != nil {
			return nil, err
		}

		lastInsertIDs = append(lastInsertIDs, int64(res.PageSettingId))

		UpdateUniqueCode(res.PageSettingId)

	}
	helpers.RemoveFileByPath(filePath)
	return lastInsertIDs, nil
}

func ExportData(limit int) (interface{}, error) {
	db := orm.NewOrm()
	var homeResponse []struct {
		PageSettingId int       `json:"page_setting_id"`
		Section       string    `json:"section"`
		DataType      string    `json:"data_type"`
		SettingData   string    `json:"setting_data"`
		CreatedDate   time.Time `json:"created_date"`
		UpdatedDate   time.Time `json:"updated_date"`
		CreatedBy     string    `json:"created_by"`
	}

	query := fmt.Sprintf(`SELECT hpst.page_setting_id as page_setting_id, hpst.section, hpst.data_type, hpst.setting_data, hpst.created_date, hpst.updated_date, concat(umt.first_name,' ',umt.last_name) as created_by FROM home_pages_setting_table as hpst LEFT JOIN user_master_table as umt ON umt.user_id = hpst.created_by LIMIT %d`, limit)

	_, err := db.Raw(query).QueryRows(&homeResponse)

	if err != nil {
		return nil, err
	}
	if len(homeResponse) == 0 {
		return "Not Found Cars", nil
	}
	return homeResponse, nil
}

func ExportDatas(limit int) (interface{}, error) {
	db := orm.NewOrm()

	var homeResponse []orm.Params

	query := fmt.Sprintf(`SELECT hpst.page_setting_id as page_setting_id, hpst.section, hpst.data_type, hpst.setting_data, hpst.created_date, hpst.updated_date, concat(umt.first_name,' ',umt.last_name) as created_by FROM home_pages_setting_table as hpst LEFT JOIN user_master_table as umt ON umt.user_id = hpst.created_by LIMIT %d`, limit)

	_, err := db.Raw(query).Values(&homeResponse)

	if err != nil {
		return nil, err
	}
	if len(homeResponse) == 0 {
		return nil, nil
	}
	return homeResponse, nil
}

func RegisterSettingBatchss(c requestStruct.HomeSeetingInsert, user_id float64, filePath string, rows []map[string]interface{}) ([]int64, []int64, error) {
	db := orm.NewOrm()
	var insertIDs, updateIDs []int64

	for _, row := range rows {
		section, ok := row["section"].(string)
		if !ok {
			return nil, nil, errors.New("missing 'section' in row")
		}

		pageSettingID, ok := row["page_setting_id"].(float64)
		if !ok {
			return nil, nil, errors.New("missing 'page_setting_id' in row")
		}

		dataType, ok := row["data_type"].(string)
		if !ok {
			return nil, nil, errors.New("missing 'data_type' in row")
		}

		settingData, ok := row["setting_data"].(string)
		if !ok {
			return nil, nil, errors.New("missing 'setting_data' in row")
		}

		// Check if the record already exists based on pageSettingID
		var existingRecord HomePagesSettingTable
		err := db.QueryTable("home_pages_setting_table").Filter("page_setting_id", pageSettingID).One(&existingRecord)

		if err == nil {
			// If the record exists, update it
			existingRecord.Section = section
			existingRecord.DataType = dataType
			existingRecord.SettingData = settingData
			existingRecord.UpdatedBy = int(user_id)
			existingRecord.UpdatedDate = time.Now()

			_, err := db.Update(&existingRecord)
			if err != nil {
				return nil, nil, err
			}

			updateIDs = append(updateIDs, int64(existingRecord.PageSettingId))
		} else {
			// If the record doesn't exist, insert a new one
			newRecord := HomePagesSettingTable{
				PageSettingId: int(pageSettingID), // Assuming PageSettingId is an int in your struct
				Section:       section,
				DataType:      dataType,
				UniqueCode:    "",
				SettingData:   settingData,
				CreatedBy:     int(user_id),
				UpdatedBy:     0,
				CreatedDate:   time.Now(),
				UpdatedDate:   time.Now(),
			}

			_, err := db.Insert(&newRecord)
			if err != nil {
				return nil, nil, err
			}

			insertIDs = append(insertIDs, int64(newRecord.PageSettingId))

			UpdateUniqueCode(newRecord.PageSettingId)
		}
	}

	helpers.RemoveFileByPath(filePath)

	return insertIDs, updateIDs, nil
}
