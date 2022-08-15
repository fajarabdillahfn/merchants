package wrapper

import (
	"encoding/json"
	"errors"
	"math"
	"net/http"

	"github.com/fajarabdillahfn/merchants/internal/models"
)

type Header struct {
	Page      int `json:"page"`
	TotalPage int `json:"total_page"`
	TotalData int `json:"total_data"`
}
type Response struct {
	Header Header      `json:"header"`
	Data   interface{} `json:"data"`
}

func WriteJSON(w http.ResponseWriter, status int, data interface{}, wrap string) error {
	wrapper := make(map[string]interface{})

	wrapper[wrap] = data

	js, err := json.Marshal(wrapper)
	if err != nil {
		return err
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func ErrorJSON(w http.ResponseWriter, err error, status ...int) {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}

	type jsonError struct {
		Message string `json:"message"`
	}

	theError := jsonError{
		Message: err.Error(),
	}

	WriteJSON(w, statusCode, theError, "error")
}

func RevenuePaginationJSON(w http.ResponseWriter, page int, limit int, data []models.Revenue) {
	var response Response
	response.Header.Page = page
	response.Header.TotalData = len(data)

	total_items := len(data)

	if page < 0 {
		err := errors.New("invalid page value")
		WriteJSON(w, http.StatusBadRequest, err, "error")
		return
	} else if page == 0 {
		page = 1
		response.Header.Page = page
	}

	if limit == 0 {
		limit = 10
	}
	if limit > total_items {
		limit = total_items
	}

	if total_items == 0 {
		response.Header.TotalPage = 1
		WriteJSON(w, http.StatusOK, response, "")
		return
	}

	total_pages := int(math.Ceil(float64(total_items) / float64(limit)))

	start_data := (page - 1) * limit
	end_data := (page) * limit
	if start_data > total_items {
		err := errors.New("invalid page value")
		WriteJSON(w, http.StatusBadRequest, err, "error")
		return
	}

	if end_data > total_items {
		end_data = total_items
	}

	response.Header.TotalPage = total_pages
	response.Data = data[start_data:end_data]

	WriteJSON(w, http.StatusOK, response, "response")
}
