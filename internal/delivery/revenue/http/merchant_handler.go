package http

import (
	"net/http"
	"strconv"

	cWrapper "github.com/fajarabdillahfn/merchants/common/wrapper"
)

func (d *Delivery) GetMerchantRevenue(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userId, _ := strconv.Atoi(r.Header.Get("accessUserId"))

	data, err := d.revenueUC.TotalMerchantRevenue(ctx, userId)
	if err != nil {
		cWrapper.ErrorJSON(w, err)
		return
	}

	err = cWrapper.WriteJSON(w, http.StatusOK, data, "data")
	if err != nil {
		cWrapper.ErrorJSON(w, err)
		return
	}
}
