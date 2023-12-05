package requestStruct

type FileType struct {
	FileType string `json:"file_type" form:"file_type"`
	Limit    int    `json:"limit" form:"limit"`
}
