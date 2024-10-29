package repository

import (
	"context"
	"log"
	"time"

	"github.com/FreitasGabriel/client-server-api/server/internal"
	"github.com/FreitasGabriel/client-server-api/server/tools/model"
	"gorm.io/gorm"
)

func CreateQuotation(ctx context.Context, db *gorm.DB, quotation *model.QuotationAPI) error {
	ctxDB, cancelDB := internal.GetContext(ctx, 10*time.Millisecond)
	defer cancelDB()

	db.AutoMigrate(&model.QuotationPayload{})

	quot := &model.QuotationPayload{
		Code:       quotation.USDBRL.Code,
		Codein:     quotation.USDBRL.Codein,
		Name:       quotation.USDBRL.Name,
		High:       quotation.USDBRL.High,
		Low:        quotation.USDBRL.Low,
		VarBid:     quotation.USDBRL.VarBid,
		PctChange:  quotation.USDBRL.PctChange,
		Bid:        quotation.USDBRL.Bid,
		Ask:        quotation.USDBRL.Ask,
		Timestamp:  quotation.USDBRL.Timestamp,
		CreateDate: quotation.USDBRL.CreateDate,
	}

	db.Create(&quot)
	select {
	case <-ctxDB.Done():
		log.Println("request canceled by client or timeout excced")
		return ctxDB.Err()
	default:
		return nil
	}

}
