package test

import (
	"crud/controllers"
	"fmt"
	"log"
	"testing"
)

func TestRegisterHomeSetting(t *testing.T) {
	controller := &controllers.HomeSettingController{}
	endPoint := "/v1/homepage/register_settings"
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2VtYWlsIjoiZGV2ZW5kcmFwb2hla2FyLnNpbGljb25pdGh1YkBnbWFpbC5jb20iLCJ1c2VyX2lkIjozMCwiZXhwIjoxNzAxMzQ1ODA4fQ.a_cBLylQLG7_7oIzp2nh6ldoPcgQLj4otx3lIUuR9fI"
	token := fmt.Sprintf("Bearer %s", tokenString)
	var jsonStr = []byte(`{"section":"MIDDEL CONTAINERS","data_type":"TEXT","setting_data":"FROM THE TESTING SECTION SETTING DATA"}`)
	result := RequestTestingFunction(token, endPoint, "POST", jsonStr, "post:RegisterSettings", controller)
	log.Print(result.Body)

}

func TestFetchPageSettings(t *testing.T) {
	controller := &controllers.HomeSettingController{}
	endPoint := "/v1/homepage/fetch_settings"
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2VtYWlsIjoiZGV2ZW5kcmFwb2hla2FyLnNpbGljb25pdGh1YkBnbWFpbC5jb20iLCJ1c2VyX2lkIjozMCwiZXhwIjoxNzAxMzQ1ODA4fQ.a_cBLylQLG7_7oIzp2nh6ldoPcgQLj4otx3lIUuR9fI"
	token := fmt.Sprintf("Bearer %s", tokenString)
	result := RequestTestingFunction(token, endPoint, "POST", nil, "post:FetchSettings", controller)
	log.Print(result.Body)

}

func TestUpdatePageSetting(t *testing.T) {
	controller := &controllers.HomeSettingController{}
	endPoint := "/v1/homepage/update_settings"
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2VtYWlsIjoiZGV2ZW5kcmFwb2hla2FyLnNpbGljb25pdGh1YkBnbWFpbC5jb20iLCJ1c2VyX2lkIjozMCwiZXhwIjoxNzAxMzQ1ODA4fQ.a_cBLylQLG7_7oIzp2nh6ldoPcgQLj4otx3lIUuR9fI"
	token := fmt.Sprintf("Bearer %s", tokenString)

	var jsonStr = []byte(`{"section":"MIDDEL CONTAINERS","data_type":"TEXT","setting_data":"FROM THE TESTING SECTION SETTING DATA","setting_id":30}`)

	result := RequestTestingFunction(token, endPoint, "POST", jsonStr, "post:UpdateSettings", controller)
	log.Print(result.Body)

}

func TestDeletePageSetting(t *testing.T) {
	controller := &controllers.HomeSettingController{}
	endPoint := "/v1/homepage/update_settings"
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2VtYWlsIjoiZGV2ZW5kcmFwb2hla2FyLnNpbGljb25pdGh1YkBnbWFpbC5jb20iLCJ1c2VyX2lkIjozMCwiZXhwIjoxNzAxMzQ1ODA4fQ.a_cBLylQLG7_7oIzp2nh6ldoPcgQLj4otx3lIUuR9fI"
	token := fmt.Sprintf("Bearer %s", tokenString)

	var jsonStr = []byte(`{"setting_id":20}`)

	result := RequestTestingFunction(token, endPoint, "POST", jsonStr, "post:DeleteSetting", controller)
	log.Print(result.Body)

}
