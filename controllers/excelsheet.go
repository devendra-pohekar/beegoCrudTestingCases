package controllers

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func CreateExcel(data []map[string]interface{}) error {

	file := excelize.NewFile()
	sheet := "Sheet1"
	file.NewSheet(sheet)
	headers := []string{"section", "data_type", "setting_data", "created_date", "updated_date", "created_by"}
	for colNum, header := range headers {
		cell := excelize.ToAlphaString(colNum+1) + "1"
		file.SetCellValue(sheet, cell, header)
	}

	for rowNum, rowData := range data {
		for colNum, key := range headers {
			cell := excelize.ToAlphaString(colNum+1) + strconv.Itoa(rowNum+2)
			if value, ok := rowData[key]; ok {
				file.SetCellValue(sheet, cell, value)
			}
		}
	}

	err := file.SaveAs("data.xlsx")
	if err != nil {
		return err
	}

	return nil
}

// func ConvertData(data interface{}) []map[string]interface{} {
// 	dataString, ok := data.(string)
// 	if !ok {
// 		fmt.Println("Invalid input type. Expected string.")
// 		return nil
// 	}
// 	entries := strings.FieldsFunc(dataString, func(r rune) bool {
// 		return r == '{' || r == '}'
// 	})
// 	re := regexp.MustCompile(`(\S+) (\S+)`)
// 	var result []map[string]interface{}
// 	for _, entry := range entries {
// 		matches := re.FindAllStringSubmatch(entry, -1)
// 		entryMap := make(map[string]interface{})
// 		for _, match := range matches {
// 			key := match[1]
// 			value := match[2]

// 			if t, err := time.Parse(time.RFC3339, value); err == nil {
// 				entryMap[key] = t
// 			} else {
// 				entryMap[key] = value
// 			}
// 		}
// 		result = append(result, entryMap)
// 	}

// 	return result
// }

func ConvertData(data interface{}) []map[string]interface{} {
	dataSlice, ok := data.([]string)
	if !ok {
		fmt.Println("Invalid input type. Expected []string.")
		return nil
	}

	var result []map[string]interface{}

	for _, dataString := range dataSlice {
		entries := strings.FieldsFunc(dataString, func(r rune) bool {
			return r == '{' || r == '}'
		})
		re := regexp.MustCompile(`(\S+) (\S+)`)
		var entryMap map[string]interface{}
		for _, entry := range entries {
			matches := re.FindAllStringSubmatch(entry, -1)
			entryMap = make(map[string]interface{})
			for _, match := range matches {
				key := match[1]
				value := match[2]

				if t, err := time.Parse(time.RFC3339, value); err == nil {
					entryMap[key] = t
				} else {
					entryMap[key] = value
				}
			}
			result = append(result, entryMap)
		}
	}

	return result
}
