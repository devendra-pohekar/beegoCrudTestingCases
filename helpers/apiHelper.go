package helpers

import (
	"encoding/json"
	"net/http"
)

type ResponseSuccess struct {
	ResStatus int    `json:"status"`
	Message   string `json:"message"`
	// Pagination interface{} `json:"pagination"`
	Result interface{} `json:"data"`
}

type ResponseFailed struct {
	ResStatus int    `json:"status"`
	Message   string `json:"message"`
}

type PaginationRes struct {
	CurrentPage interface{} `json:"current_page"`
	TotalPages  interface{} `json:"total_pages"`
}

func ApiSuccessResponse(w http.ResponseWriter, result interface{}, message string, totalPages, currentPage interface{}) {
	if result == "" {
		result = map[int]int{}
	}
	if totalPages == "" && currentPage == "" {
		totalPages = 1
		currentPage = 1
	}
	// pagination := PaginationRes{CurrentPage: currentPage, TotalPages: totalPages}
	response := ResponseSuccess{
		Message:   message,
		ResStatus: 1,
		// Pagination: pagination,
		Result: result,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonResponse)
}

func ApiFailedResponse(w http.ResponseWriter, message string) {
	response := ResponseFailed{
		Message:   message,
		ResStatus: 0,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)

}
