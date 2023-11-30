package requestStruct

type HomeSeetingInsert struct {
	Section     string `json:"section" form:"section"`
	DataType    string `json:"data_type" form:"data_type"`
	SettingData string `json:"setting_data" form:"setting_data"`
}

type HomeSeetingUpdate struct {
	Section     string `json:"section" form:"section"`
	DataType    string `json:"data_type" form:"data_type"`
	SettingData string `json:"setting_data" form:"setting_data"`
	SettingId   int    `json:"setting_id" form:"setting_id"`
}

type HomeSeetingDelete struct {
	Section   string `json:"section" form:"section"`
	SettingId int    `json:"setting_id" form:"setting_id"`
}

type HomeSeetingSearch struct {
	SettingId int `json:"setting_id" form:"setting_id"`
}
