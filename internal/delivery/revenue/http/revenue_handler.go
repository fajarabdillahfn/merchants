package http

import (
	"errors"
	"net/http"
	"strconv"

	cWrapper "github.com/fajarabdillahfn/merchants/common/wrapper"
	"github.com/gorilla/mux"
)

func (d *Delivery) GetRevenue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dataType := vars["dataType"]

	if dataType != "merchant" && dataType != "outlet" {
		cWrapper.ErrorJSON(w, errors.New("invalid param, should be [merchant] or [outlet]"), http.StatusBadRequest)
		return
	}

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	ctx := r.Context()

	userId, _ := strconv.Atoi(r.Header.Get("accessUserId"))

	data, err := d.revenueUC.TotalMerchantRevenue(ctx, userId, dataType)
	if err != nil {
		cWrapper.ErrorJSON(w, err)
		return
	}

	cWrapper.RevenuePaginationJSON(w, page, limit, data)
}
