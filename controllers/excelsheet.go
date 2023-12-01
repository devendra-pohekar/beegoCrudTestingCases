package controllers

import (
	"fmt"
	"time"
)

// type ExcelController struct {
// 	beego.Controller
// }

// func (c *ExcelController) InsertData() {
// 	// Sample data, replace this with the data you want to insert
// 	data := [][]interface{}{
// 		{"Name", "Age", "City"},
// 		{"John Doe", 30, "New York"},
// 		{"Jane Doe", 25, "San Francisco"},
// 		// Add more rows as needed
// 	}

// 	// Create a new Excel file
// 	xlsx := excelize.NewFile()

// 	// Create a new sheet in the Excel file
// 	sheetName := "Sheet1"
// 	xlsx.NewSheet(sheetName)

// 	// Insert data into the sheet
// 	for rowIndex, row := range data {
// 		for colIndex, cellValue := range row {
// 			cellAddress, _ := excelize.CoordinatesToCellName(colIndex+1, rowIndex+1)
// 			xlsx.SetCellValue(sheetName, cellAddress, cellValue)
// 		}
// 	}

// 	// Save the Excel file
// 	filePath := "path/to/your/excel/file.xlsx" // Provide the desired file path
// 	if err := xlsx.SaveAs(filePath); err != nil {
// 		fmt.Println("Error saving Excel file:", err)
// 		c.Data["json"] = map[string]interface{}{"error": err.Error()}
// 	} else {
// 		fmt.Println("Excel file saved successfully:", filePath)
// 		c.Data["json"] = map[string]interface{}{"success": true, "filePath": filePath}
// 	}

// 	c.ServeJSON()
// }

// import (
// 	"fmt"
// 	"log"
// 	"regexp"
// 	"strconv"
// 	"strings"
// 	"time"

// 	"github.com/360EntSecGroup-Skylar/excelize"
// )

func formatValue(value interface{}) interface{} {
	switch v := value.(type) {
	case time.Time:
		// Format time as needed
		return v.Format("2006-01-02 15:04:05")
	default:
		return v
	}
}

func TransformToKeyValuePairs(data interface{}) ([]map[string]interface{}, error) {

	responseData, ok := data.([]struct {
		Section     string    `json:"section"`
		DataType    string    `json:"data_type"`
		SettingData string    `json:"setting_data"`
		CreatedDate time.Time `json:"created_date"`
		UpdatedDate time.Time `json:"updated_date"`
		CreatedBy   string    `json:"created_by"`
	})

	if !ok {
		return nil, fmt.Errorf("invalid data type")
	}

	result := make([]map[string]interface{}, len(responseData))
	for i, item := range responseData {
		result[i] = map[string]interface{}{
			"section":      item.Section,
			"data_type":    item.DataType,
			"setting_data": item.SettingData,
			"created_date": item.CreatedDate,
			"updated_date": item.UpdatedDate,
			"created_by":   item.CreatedBy,
		}
	}

	return result, nil
}

// func CreateExcel(data []map[string]interface{}, headers []string, folderPath, fileNamePrefix string) (string, error) {
// 	file := excelize.NewFile()
// 	sheet := "Sheet1"
// 	file.NewSheet(sheet)

// 	// Set header row
// 	for colNum, header := range headers {
// 		// cell := excelize.ToAlphaString(colNum+1) + "1"
// 		cell := fmt.Sprintf("%c%d", 'A'+colNum, 1)
// 		file.SetCellValue(sheet, cell, header)
// 	}

// 	// Set data rows
// 	for rowNum, rowData := range data {
// 		for colNum, key := range headers {
// 			cell := fmt.Sprintf("%c%d", 'A'+colNum, rowNum+2)
// 			// cell := excelize.ToAlphaString(colNum+1) + strconv.Itoa(rowNum+2)
// 			if value, ok := rowData[key]; ok {
// 				file.SetCellValue(sheet, cell, formatValue(value))
// 			}
// 		}
// 	}

// 	if folderPath == "" {
// 		folderPath = "FILES/XLSX"
// 	}

// 	if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
// 		return "", fmt.Errorf("failed to create folder: %v", err)
// 	}

// 	fileName := fmt.Sprintf("%s_%s.xlsx", fileNamePrefix, time.Now().Format("20060102150405"))
// 	filePath := filepath.Join(folderPath, fileName)
// 	if err := file.SaveAs(filePath); err != nil {
// 		return "", err
// 	}
// 	return filePath, nil
// }

// func CreateCSV(data []map[string]interface{}, headers []string, folderPath, fileNamePrefix string) (string, error) {
// 	if folderPath == "" {
// 		folderPath = "FILES/CSV"
// 	}

// 	if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
// 		return "", fmt.Errorf("failed to create folder: %v", err)
// 	}

// 	fileName := fmt.Sprintf("%s_%s.csv", fileNamePrefix, time.Now().Format("20060102150405"))
// 	filePath := filepath.Join(folderPath, fileName)
// 	file, err := os.Create(filePath)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to create CSV file: %v", err)
// 	}
// 	defer file.Close()

// 	csvWriter := csv.NewWriter(file)
// 	defer csvWriter.Flush()

// 	// Write header row
// 	if err := csvWriter.Write(headers); err != nil {
// 		return "", fmt.Errorf("failed to write CSV header: %v", err)
// 	}

// 	// Write data rows
// 	for _, rowData := range data {
// 		var row []string
// 		for _, key := range headers {
// 			if value, ok := rowData[key]; ok {
// 				row = append(row, FormateCSVDate(value))
// 			} else {
// 				row = append(row, "") // Handle missing data
// 			}
// 		}
// 		if err := csvWriter.Write(row); err != nil {
// 			return "", fmt.Errorf("failed to write CSV row: %v", err)
// 		}
// 	}

// 	return filePath, nil
// }
