package revenue

import (
	"context"
	"log"

	"github.com/fajarabdillahfn/merchants/internal/models"
)

func (r *repository) GetOutletRevenuePerDay(ctx context.Context, startDate string, endDate string, outletId int) (revenue float64, err error) {
	var transaction models.Transaction
	tableName := transaction.TableName()

	check := r.conn.WithContext(ctx).Table(tableName).
		Where("merchant_id", outletId).
		Where("created_at > ?", startDate).
		Where("created_at < ?", endDate).
		Find(&transaction)
	if check.RowsAffected < 1 {
		return 0, nil
	}

	res := r.conn.WithContext(ctx).Raw("SELECT SUM(bill_total) FROM Transactions WHERE merchant_id = ? AND created_at > ? AND created_at < ?",
		outletId,
		startDate,
		endDate).Scan(&revenue)

	if res.Error != nil {
		log.Println("ERROR get merchant revenue: " + res.Error.Error())
		return 0, res.Error
	}
	return
}
