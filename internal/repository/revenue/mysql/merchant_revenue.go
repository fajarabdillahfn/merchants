package revenue

import (
	"context"
	"log"

	"github.com/fajarabdillahfn/merchants/internal/models"
)

func (r *repository) GetMerchantsByUserId(ctx context.Context, userId int) (merchants []models.Merchant, err error) {
	tableName := models.Merchant.TableName(models.Merchant{})

	dataPrep := r.conn.WithContext(ctx).Table(tableName)

	res := dataPrep.Find(&merchants, "user_id = ?", userId)
	if res.Error != nil {
		log.Println("ERROR get merchant: " + res.Error.Error())
		return nil, res.Error
	}

	return
}

func (r *repository) GetMerchantRevenuePerDay(ctx context.Context, startDate string, endDate string, merchantId int) (revenue float64, err error) {
	var transaction models.Transaction
	tableName := transaction.TableName()

	check := r.conn.WithContext(ctx).Table(tableName).
	Where("merchant_id", merchantId).
	Where("created_at > ?", startDate).
	Where("created_at < ?", endDate).
	Find(&transaction)
	if check.RowsAffected < 1 {
		return 0, nil
	}

	res := r.conn.WithContext(ctx).Raw("SELECT SUM(bill_total) FROM Transactions WHERE merchant_id = ? AND created_at > ? AND created_at < ?",
		merchantId,
		startDate,
		endDate).Scan(&revenue)

	if res.Error != nil {
		log.Println("ERROR get merchant revenue: " + res.Error.Error())
		return 0, res.Error
	}
	return
}
