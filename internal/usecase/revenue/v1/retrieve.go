package v1

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/fajarabdillahfn/merchants/internal/models"
)

func (u *useCase) TotalMerchantRevenue(ctx context.Context, userId int, dataType string) (revenues []models.Revenue, err error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*5000)
	defer cancel()

	merchants, err := u.RevenueRepo.GetMerchantsByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	for i := 1; i < 31; i++ {
		var date string
		if i < 10 {
			date = strconv.Itoa(i)
			date = "0" + date
			date = fmt.Sprintf("2021-11-%s", date)
		} else {
			date = strconv.Itoa(i)
			date = fmt.Sprintf("2021-11-%s", date)
		}

		startDate := date + " 00:00:00"
		endDate := date + " 23:59:59"

		var details []models.MerchantDetails
		for _, merchant := range merchants {
			switch dataType {
			case "outlet":
				for _, outlet := range merchant.Outlets {
					revenue, err := u.RevenueRepo.GetOutletRevenuePerDay(ctx, startDate, endDate, outlet.Id)
					if err != nil {
						return nil, err
					}

					details = append(details, models.MerchantDetails{
						MerchantName: merchant.MerchantName,
						OutletName:   outlet.OutletName,
						DailyRevenue: revenue,
					})
				}

			case "merchant":
				revenue, err := u.RevenueRepo.GetMerchantRevenuePerDay(ctx, startDate, endDate, merchant.Id)
				if err != nil {
					return nil, err
				}

				details = append(details, models.MerchantDetails{
					MerchantName: merchant.MerchantName,
					DailyRevenue: revenue,
				})

			default:
				return nil, errors.New("invalid data wanted")
			}

		}

		revenues = append(revenues, models.Revenue{
			Date:    date,
			Revenue: details,
		})
	}

	return
}