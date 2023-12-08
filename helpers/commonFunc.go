package helpers

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/smtp"
	"os"
	"os/user"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	beego "github.com/beego/beego/v2/server/web"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/dgrijalva/jwt-go"
	"github.com/jung-kurt/gofpdf"
	"github.com/tealeg/xlsx"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func GetTokenClaims(c *context.Context) map[string]interface{} {
	token_claims := c.Input.GetData("LoginUserData")
	user_id := token_claims.(jwt.MapClaims)["user_id"]
	user_email := token_claims.(jwt.MapClaims)["user_email"]
	response := map[string]interface{}{"User_id": user_id, "User_Email": user_email}
	return response
}

func UploadFile(fileToUpload multipart.File, fileHeader *multipart.FileHeader, uploadDir string) (string, error) {
	defer fileToUpload.Close()
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename))
	if err := os.MkdirAll(uploadDir, 0777); err != nil {
		return "", fmt.Errorf("failed to create upload directory: %v", err)
	}
	filePath := filepath.Join(uploadDir, filename)
	out, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create the destination file: %v", err)
	}
	defer out.Close()

	_, err = io.Copy(out, fileToUpload)
	if err != nil {
		return "", fmt.Errorf("failed to copy the file: %v", err)
	}

	return filePath, nil
}

func RemoveFile(fileName, directory string) error {
	err := os.Remove(filepath.Join(directory, fileName))
	if err != nil {
		return err
	}
	return nil
}

func RemoveFileByPath(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		return err
	}

	return nil
}

func SplitFilePath(SplitString string) (string, string) {
	lastIndex := strings.LastIndex(SplitString, "/")

	var fileDirectory string
	var fileName string

	if lastIndex != -1 {
		fileDirectory = SplitString[:lastIndex]
		fileName = SplitString[lastIndex+1:]
	} else {
		fileDirectory = "No '/' found in the string."
		fileName = fileDirectory
	}

	return fileName, fileDirectory
}

func UniqueCode(insertedId int, withString string) string {
	withString = strings.ReplaceAll(withString, " ", "_")
	result := fmt.Sprintf("%s_%d", withString, insertedId)
	return strings.ToUpper(result)
}

func SendOTpOnMail(userEmail string, name string) (string, error) {
	from := "devendra.siliconithub@gmail.com"
	password := "ufax tadd qcoa xbft"
	to := []string{
		userEmail,
	}
	OTP := GenerateUniqueCodeString(8)
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	subject := "Verify your email"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	body := `<table class="body-wrap" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; width: 100%; background-color: #FFC300; margin: 0;" bgcolor="#FF5733">
    <tbody>
        <tr style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;">
            <td style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; vertical-align: top; margin: 0;" valign="top"></td>
            <td class="container" width="600" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; vertical-align: top; display: block !important; max-width: 600px !important; clear: both !important; margin: 0 auto;"
                valign="top">
                <div class="content" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; max-width: 600px; display: block; margin: 0 auto; padding: 20px;">
                    <table class="main" width="100%" cellpadding="0" cellspacing="0" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; border-radius: 3px; background-color: #0000000; margin: 0; border: 1px solid #;"
                        bgcolor="#fff">
                        <tbody>
                            <tr style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;">
                                <td class="" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 16px; vertical-align: top; color: #fff; font-weight: 500; text-align: center; border-radius: 3px 3px 0 0; background-color: #; margin: 0; padding: 20px;"
                                    align="center" bgcolor="#71b6f9" valign="top">
                                    <a href="#" style="font-size:32px;color:#;">www.siliconithub.com</a> <br>
                                    <span style="margin-top: 10px;display: block;">Please Confirm OTP For Email Verification</span>
                                </td>
                            </tr>
                            <tr style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;">
                                <td class="content-wrap" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; vertical-align: top; margin: 0; padding: 20px;" valign="top">
                                    <table width="100%" cellpadding="0" cellspacing="0" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;">
                                        <tbody>
                                            <tr style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;">
                                                <td class="content-block" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; vertical-align: top; margin: 0; padding: 0 0 20px;" valign="top">
                                                    Mr./Ms <strong style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;">
                                   ` + name + `             </td>
                                            </tr>
                                            <tr style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;">
                                                <td class="content-block" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; vertical-align: top; margin: 0; padding: 0 0 20px;" valign="top">
                                                    We are happy you Signed up  for Silicon IT Hub.To start  Exploring The Silicon IT Hub And  Neighborhood ,
                                                    <p style ="color:#C70039">Please Confirm Your Email Address</p>.
                                                </td>
                                            </tr>
                                            <tr style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;">
                                                <td class="content-block" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; vertical-align: top; margin: 0; padding: 0 0 20px;" valign="top">
                                                    <p class="btn-primary" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; color: #FFF; text-decoration: none; line-height: 2em; font-weight: bold; text-align: center; cursor: pointer; display: inline-block; border-radius: 5px; text-transform: capitalize; background-color: #f1556c; margin: 0; border-color: #f1556c; border-style: solid; border-width: 8px 16px;">Verify Email CODE :- ` + OTP + `</p>
                                                </td>
                                            </tr>
                                            <tr style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;">
                                                <td class="content-block" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; vertical-align: top; margin: 0; padding: 0 0 20px;" valign="top">
                                                    Welcome To Silicon IT Hub 
                                                     
                                                </td>
                                               
                                            </tr>
                                        </tbody>
                                    </table>
                                </td>
                            </tr>
                        </tbody>
                    </table>
                    <div class="footer" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; width: 100%; clear: both; color: #999; margin: 0; padding: 20px;">
                        <table width="100%" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;">
                            <tbody>
                                <tr style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;">
                                    <td class="aligncenter content-block" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif;color:"#fff"; box-sizing: border-box; font-size: 12px; vertical-align: top; color: #999; text-align: center; margin: 0; padding: 0 0 20px;" align="center" valign="top"><a href="#"  style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 12px; color: #999; text-decoration: underline; margin: 0;color:#ffff">Silicon IT Hb</a> Thanks & Regards.
                                    </td>
                                </tr>
                            </tbody>
                        </table>
                    </div>
                </div>
            </td>
            <td style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; vertical-align: top; margin: 0;" valign="top"></td>
        </tr>
    </tbody>
</table>`
	message := []byte("Subject: " + subject + "\r\n" + mime + "\r\n" + body)
	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		return "", err
	}

	return OTP, nil
}

func GenerateUniqueCodeString(length int) string {
	charset := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

/*----------------------------------------XLSX file creating functions------------------------------------------------------------*/
func TransformToKeyValuePairs(data interface{}) ([]map[string]interface{}, error) {
	value := reflect.ValueOf(data)
	if value.Kind() != reflect.Slice {
		return nil, fmt.Errorf("input data must be a slice")
	}

	result := make([]map[string]interface{}, value.Len())

	for i := 0; i < value.Len(); i++ {
		item := value.Index(i)
		if item.Kind() != reflect.Struct {
			return nil, fmt.Errorf("items in the slice must be structs")
		}

		fields := make(map[string]interface{})
		for j := 0; j < item.NumField(); j++ {
			field := item.Type().Field(j)
			fieldName := field.Tag.Get("json")
			if fieldName == "" {
				fieldName = strings.ToLower(field.Name)
			}

			fields[fieldName] = item.Field(j).Interface()
		}

		result[i] = fields
	}

	return result, nil
}

func formatValue(value interface{}) interface{} {
	switch v := value.(type) {
	case time.Time:
		return v.Format("2006-01-02 15:04:05")
	default:
		return v
	}
}

func XlsxFileCreater(data []map[string]interface{}, headers []string, folderPath, fileNamePrefix string) (string, error) {
	file := excelize.NewFile()
	sheet := "Sheet1"
	file.NewSheet(sheet)
	for colNum, header := range headers {
		cell := fmt.Sprintf("%c%d", 'A'+colNum, 1)
		file.SetCellValue(sheet, cell, header)
	}

	for rowNum, rowData := range data {
		for colNum, key := range headers {
			cell := fmt.Sprintf("%c%d", 'A'+colNum, rowNum+2)
			if value, ok := rowData[key]; ok {
				file.SetCellValue(sheet, cell, formatValue(value))
			}
		}
	}

	// Set column width based on the maximum content length in each column
	for colNum, key := range headers {
		maxLength := 0
		for rowNum, rowData := range data {
			log.Print(rowNum)
			if value, ok := rowData[key]; ok {
				cellValue := fmt.Sprintf("%v", formatValue(value))
				valueLength := len(cellValue)
				if valueLength > maxLength {
					maxLength = valueLength
				}
			}
		}

		colName := fmt.Sprintf("%c", 'A'+colNum)
		file.SetColWidth(sheet, colName, colName, float64(maxLength)*1.2) // Adjust the multiplier as needed of column
	}
	//if filepath not given than it take
	if folderPath == "" {
		folderPath = "FILES/XLSX"
	}
	//if folder not present in directory it create new folder directory
	if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create folder: %v", err)
	}

	fileName := fmt.Sprintf("%s_%s.xlsx", fileNamePrefix, time.Now().Format("20060102150405"))
	filePath := filepath.Join(folderPath, fileName)
	if err := file.SaveAs(filePath); err != nil {
		return "", err
	}
	return filePath, nil
}
func FormateCSVDate(value interface{}) string {
	switch v := value.(type) {
	case time.Time:
		return v.Format("2006-01-02 15:04:05") // Format the time value as needed
	default:
		return fmt.Sprintf("%v", value)
	}
}

func CreateExcels(data []map[string]interface{}) error {
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
				file.SetCellValue(sheet, cell, formatValue(value))
			}
		}
	}

	err := file.SaveAs("data.xlsx")
	if err != nil {
		return err
	}

	return nil
}

func SumSliceElements(slice []float64) float64 {
	var total float64
	for _, value := range slice {
		total += value
	}
	return total
}

func CreateFile(data []map[string]interface{}, headers []string, folderPath, fileNamePrefix, fileType string) (string, error) {
	TYPE := strings.ToUpper(fileType)
	switch TYPE {
	case "XLSX":
		file := excelize.NewFile()
		sheet := "Sheet1"
		file.NewSheet(sheet)

		// Set header row
		for colNum, header := range headers {
			cell := fmt.Sprintf("%c%d", 'A'+colNum, 1)
			file.SetCellValue(sheet, cell, header)
		}

		// Set data rows
		for rowNum, rowData := range data {
			for colNum, key := range headers {
				cell := fmt.Sprintf("%c%d", 'A'+colNum, rowNum+2)
				if value, ok := rowData[key]; ok {
					file.SetCellValue(sheet, cell, formatValue(value))
				}
			}
		}

		if folderPath == "" {
			folderPath = "FILES/XLSX"
		}

		if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
			return "", fmt.Errorf("failed to create folder: %v", err)
		}

		fileName := fmt.Sprintf("%s_%s.xlsx", fileNamePrefix, time.Now().Format("20060102150405"))
		filePath := filepath.Join(folderPath, fileName)
		if err := file.SaveAs(filePath); err != nil {
			return "", err
		}
		return filePath, nil

	case "CSV":
		if folderPath == "" {
			folderPath = "FILES/CSV"
		}

		if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
			return "", fmt.Errorf("failed to create folder: %v", err)
		}

		fileName := fmt.Sprintf("%s_%s.csv", fileNamePrefix, time.Now().Format("20060102150405"))
		filePath := filepath.Join(folderPath, fileName)
		file, err := os.Create(filePath)
		if err != nil {
			return "", fmt.Errorf("failed to create CSV file: %v", err)
		}
		defer file.Close()

		csvWriter := csv.NewWriter(file)
		defer csvWriter.Flush()

		// Write header row
		if err := csvWriter.Write(headers); err != nil {
			return "", fmt.Errorf("failed to write CSV header: %v", err)
		}

		// Write data rows
		for _, rowData := range data {
			var row []string
			for _, key := range headers {
				if value, ok := rowData[key]; ok {
					row = append(row, FormateCSVDate(value))
				} else {
					row = append(row, "") // Handle missing data
				}
			}
			if err := csvWriter.Write(row); err != nil {
				return "", fmt.Errorf("failed to write CSV row: %v", err)
			}
		}

		return filePath, nil

	case "PDF":
		pdf := gofpdf.New("L", "mm", "A4", "")
		pdf.AddPage()

		// Set font
		fontSize := 10.0
		pdf.SetFont("Arial", "B", fontSize)

		// Calculate total width of the page
		pageWidth, _ := pdf.GetPageSize()

		// Calculate and set maximum column widths based on headers
		colWidths := make([]float64, len(headers))
		totalWidth := pageWidth - 20 // Adjust the margin as needed
		for colNum, header := range headers {
			colWidths[colNum] = pdf.GetStringWidth(header) + 6 // Add padding
		}

		// Normalize column widths based on the total width
		scaleFactor := totalWidth / SumSliceElements(colWidths)
		for colNum := range colWidths {
			colWidths[colNum] *= scaleFactor
		}

		// Add headers
		for colNum, header := range headers {
			pdf.CellFormat(colWidths[colNum], 10, header, "1", 0, "", false, 0, "")
		}

		pdf.Ln(-1)

		// Set font for data
		pdf.SetFont("Arial", "", fontSize)

		// Add data
		for _, rowData := range data {
			// Add data to PDF with adjusted row height
			for colNum, key := range headers {
				if value, ok := rowData[key]; ok {
					cellValue := fmt.Sprintf("%v", formatValue(value))
					pdf.CellFormat(colWidths[colNum], 10, cellValue, "1", 0, "", false, 0, "")
				}
			}

			pdf.Ln(-1)
		}

		// If filepath not given, it takes the default
		if folderPath == "" {
			folderPath = "FILES/PDF"
		}

		// If the folder is not present in the directory, it creates a new folder directory
		if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
			return "", fmt.Errorf("failed to create folder: %v", err)
		}

		// Generate file name
		fileName := fmt.Sprintf("%s_%s.pdf", fileNamePrefix, time.Now().Format("20060102150405"))
		filePath := filepath.Join(folderPath, fileName)

		// Save the PDF file
		if err := pdf.OutputFileAndClose(filePath); err != nil {
			return "", err
		}

		return filePath, nil

	default:
		return "", fmt.Errorf("unsupported file type: %s", fileType)
	}
}

/* ----------------------end XLSX file creating functions---------------------------------------------------------*/

/*pending in download file in system folder after convert any type */

func CreateFiles(w http.ResponseWriter, r *http.Request, data []map[string]interface{}, headers []string, folderPath, fileNamePrefix, fileType string) (string, error) {
	TYPE := strings.ToUpper(fileType)

	// localpath, _ := GetDownloadsFolderPath()
	switch TYPE {
	case "XLSX":
		file := excelize.NewFile()
		sheet := "Sheet1"
		file.NewSheet(sheet)

		// Set header row
		for colNum, header := range headers {
			cell := fmt.Sprintf("%c%d", 'A'+colNum, 1)
			file.SetCellValue(sheet, cell, header)
		}

		// Set data rows
		for rowNum, rowData := range data {
			for colNum, key := range headers {
				cell := fmt.Sprintf("%c%d", 'A'+colNum, rowNum+2)
				if value, ok := rowData[key]; ok {
					file.SetCellValue(sheet, cell, formatValue(value))
				}
			}
		}

		if folderPath == "" {
			folderPath = "FILES/XLSX"
		}

		if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
			return "", fmt.Errorf("failed to create folder: %v", err)
		}

		fileName := fmt.Sprintf("%s_%s.xlsx", fileNamePrefix, time.Now().Format("20060102150405"))
		filePath := filepath.Join(folderPath, fileName)
		if err := file.SaveAs(filePath); err != nil {
			return "", err
		}
		if err := DownloadFile(w, r, filePath, "XLSX"); err != nil {
			return "", err
		}
		return filePath, nil

	case "CSV":
		if folderPath == "" {
			folderPath = "FILES/CSV"
		}

		if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
			return "", fmt.Errorf("failed to create folder: %v", err)
		}

		fileName := fmt.Sprintf("%s_%s.csv", fileNamePrefix, time.Now().Format("20060102150405"))
		filePath := filepath.Join(folderPath, fileName)
		file, err := os.Create(filePath)
		if err != nil {
			return "", fmt.Errorf("failed to create CSV file: %v", err)
		}
		defer file.Close()

		csvWriter := csv.NewWriter(file)
		defer csvWriter.Flush()

		// Write header row
		if err := csvWriter.Write(headers); err != nil {
			return "", fmt.Errorf("failed to write CSV header: %v", err)
		}

		// Write data rows
		for _, rowData := range data {
			var row []string
			for _, key := range headers {
				if value, ok := rowData[key]; ok {
					row = append(row, FormateCSVDate(value))
				} else {
					row = append(row, "") // Handle missing data
				}
			}
			if err := csvWriter.Write(row); err != nil {
				return "", fmt.Errorf("failed to write CSV row: %v", err)
			}
		}

		if err := DownloadFile(w, r, filePath, "CSV"); err != nil {
			return "", err
		}
		return filePath, nil

	case "PDF":
		pdf := gofpdf.New("L", "mm", "A4", "")
		pdf.AddPage()

		// Set font
		fontSize := 10.0
		pdf.SetFont("Arial", "B", fontSize)

		// Calculate total width of the page
		pageWidth, _ := pdf.GetPageSize()

		// Calculate and set maximum column widths based on headers
		colWidths := make([]float64, len(headers))
		totalWidth := pageWidth - 20 // Adjust the margin as needed
		for colNum, header := range headers {
			colWidths[colNum] = pdf.GetStringWidth(header) + 6 // Add padding
		}

		// Normalize column widths based on the total width
		scaleFactor := totalWidth / SumSliceElements(colWidths)
		for colNum := range colWidths {
			colWidths[colNum] *= scaleFactor
		}

		// Add headers
		for colNum, header := range headers {
			pdf.CellFormat(colWidths[colNum], 10, header, "1", 0, "", false, 0, "")
		}

		pdf.Ln(-1)

		// Set font for data
		pdf.SetFont("Arial", "", fontSize)

		// Add data
		for _, rowData := range data {
			// Add data to PDF with adjusted row height
			for colNum, key := range headers {
				if value, ok := rowData[key]; ok {
					cellValue := fmt.Sprintf("%v", formatValue(value))
					pdf.CellFormat(colWidths[colNum], 10, cellValue, "1", 0, "", false, 0, "")
				}
			}

			pdf.Ln(-1)
		}

		// If filepath not given, it takes the default
		if folderPath == "" {
			folderPath = "FILES/PDF"
		}

		// If the folder is not present in the directory, it creates a new folder directory
		if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
			return "", fmt.Errorf("failed to create folder: %v", err)
		}

		// Generate file name
		fileName := fmt.Sprintf("%s_%s.pdf", fileNamePrefix, time.Now().Format("20060102150405"))
		filePath := filepath.Join(folderPath, fileName)

		// Save the PDF file

		if err := pdf.OutputFileAndClose(filePath); err != nil {
			return "", err
		}
		// Download logic
		if err := DownloadFile(w, r, filePath, "PDF"); err != nil {
			return "", err
		}
		return filePath, nil

	default:
		return "", fmt.Errorf("unsupported file type: %s", fileType)
	}
}

func DownloadFiles(w http.ResponseWriter, r *http.Request, filePath string, fileType string) error {
	// Ensure the directory of the file exists
	dir := filepath.Dir(filePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// Directory does not exist, create it
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			http.Error(w, "Failed to create directory", http.StatusInternalServerError)
			return err
		}
	}

	// Open the file
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDONLY, os.ModePerm)
	if err != nil {
		http.Error(w, "Failed to open file", http.StatusInternalServerError)
		return err
	}
	defer file.Close()

	// Get file information
	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, "Failed to get file information", http.StatusInternalServerError)
		return err
	}

	// Set headers for file download
	w.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(filePath))

	switch strings.ToUpper(fileType) {
	case "XLSX":
		w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	case "CSV":
		w.Header().Set("Content-Type", "text/csv")
	case "PDF":
		w.Header().Set("Content-Type", "application/pdf")
	default:
		w.Header().Set("Content-Type", "application/octet-stream")
	}

	// Copy the file to the response writer
	http.ServeContent(w, r, filepath.Base(filePath), fileInfo.ModTime(), file)

	return nil
}

func DownloadFileToLocal(filePath string, fileType string, localPath string) error {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Ensure the directory of the local file exists
	localDir := filepath.Dir(localPath)
	if _, err := os.Stat(localDir); os.IsNotExist(err) {
		// Directory does not exist, create it
		if err := os.MkdirAll(localDir, os.ModePerm); err != nil {
			return err
		}
	}

	// Create or open the local file
	localFile, err := os.Create(localPath)
	if err != nil {
		return err
	}
	defer localFile.Close()

	// Copy the file content to the local file
	_, err = io.Copy(localFile, file)
	if err != nil {
		return err
	}

	return nil
}

func DownloadFile(w http.ResponseWriter, r *http.Request, filePath string, fileType string) error {
	// Define the directory where you want to store files
	baseDir := "/home/devendra/dev_files"

	// Ensure the directory exists
	if _, err := os.Stat(baseDir); os.IsNotExist(err) {
		// Directory does not exist, create it
		if err := os.MkdirAll(baseDir, os.ModePerm); err != nil {
			log.Println("Failed to create directory:", err)
			http.Error(w, "Failed to create directory", http.StatusInternalServerError)
			return err
		}
	}

	// Construct the local path within the dev_files directory
	localPath := filepath.Join(baseDir, filepath.Base(filePath))

	log.Println(localPath, "----error--------------------------")
	err := DownloadFileToLocal(filePath, fileType, localPath)
	if err != nil {
		log.Println("Failed to download file locally:", err)
		http.Error(w, "Failed to download file locally", http.StatusInternalServerError)
		return err
	}

	// Open the local file
	file, err := os.Open(localPath)
	if err != nil {
		log.Println("Failed to open local file:", err)
		http.Error(w, "Failed to open local file", http.StatusInternalServerError)
		return err
	}
	defer file.Close()

	// Get file information
	fileInfo, err := file.Stat()
	if err != nil {
		log.Println("Failed to get file information:", err)
		http.Error(w, "Failed to get file information", http.StatusInternalServerError)
		return err
	}

	// Set headers for file download
	w.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(localPath))

	switch strings.ToUpper(fileType) {
	case "XLSX":
		w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	case "CSV":
		w.Header().Set("Content-Type", "text/csv")
	case "PDF":
		w.Header().Set("Content-Type", "application/pdf")
	default:
		w.Header().Set("Content-Type", "application/octet-stream")
	}

	// Copy the file to the response writer
	http.ServeContent(w, r, filepath.Base(localPath), fileInfo.ModTime(), file)

	return nil
}

func GetDownloadsFolderPath() (string, error) {
	// Get the current user
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}

	// Construct the downloads folder path
	downloadsFolderPath := filepath.Join(currentUser.HomeDir, "dev_files")

	return downloadsFolderPath, nil
}

/*end pending download file in system folder*/

/*-------------------------------XLSX AND CSV FILE READING FUNCTION*/

func ReadXLSXFiles(filePath string) ([]map[string]interface{}, error) {
	xlFile, err := xlsx.OpenFile(filePath)
	if err != nil {
		return nil, err
	}

	var allRows []map[string]interface{}

	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			rowData := make(map[string]interface{})
			for index, cell := range row.Cells {
				rowData[fmt.Sprintf("Column%d", index+1)] = cell.String()
			}

			allRows = append(allRows, rowData)
		}
	}

	return allRows, nil
}

func ReadXLSXFile(filePath string) ([]map[string]interface{}, error) {
	xlFile, err := xlsx.OpenFile(filePath)
	if err != nil {
		return nil, err
	}

	var allRows []map[string]interface{}

	// Assume the first row is the header
	var headerRow []string
	if len(xlFile.Sheets) > 0 && len(xlFile.Sheets[0].Rows) > 0 {
		headerRow = make([]string, len(xlFile.Sheets[0].Rows[0].Cells))
		for index, cell := range xlFile.Sheets[0].Rows[0].Cells {
			headerRow[index] = cell.String()
		}
	}

	for _, sheet := range xlFile.Sheets {
		for rowIndex, row := range sheet.Rows {
			if rowIndex == 0 {
				// Skip the header row
				continue
			}

			rowData := make(map[string]interface{})
			for index, cell := range row.Cells {
				if index < len(headerRow) {
					rowData[headerRow[index]] = cell.String()
				}
			}

			allRows = append(allRows, rowData)
		}
	}

	return allRows, nil
}

func ReadCSVFile(filePath string) ([]map[string]interface{}, error) {
	csvFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	var allRows []map[string]interface{}

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	columnHeaders := records[0]

	for _, dataRow := range records[1:] {
		rowData := make(map[string]interface{})

		for index, value := range dataRow {
			rowData[columnHeaders[index]] = value
		}

		allRows = append(allRows, rowData)
	}

	return allRows, nil
}

/* END XLSX AND CSV FILE READING FUNCTION END-----------------------*/

func SetSessionByKeyValue(key string, value string, w http.ResponseWriter, r *http.Request) {
	session, _ := beego.GlobalSessions.SessionStart(w, r)
	err := session.Set(r.Context(), key, value)
	if err != nil {
		log.Print("error occured in session create time")
	}
	session.SessionRelease(r.Context(), w)
}

func ExtractKeys(data []map[string]interface{}) []string {
	keys := make(map[string]struct{})
	for _, item := range data {
		for key := range item {
			keys[key] = struct{}{}
		}
	}
	var result []string
	for key := range keys {
		result = append(result, key)
	}
	sort.Strings(result)
	return result
}
