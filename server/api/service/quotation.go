package service

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/FreitasGabriel/client-server-api/server/api/repository"
	"github.com/FreitasGabriel/client-server-api/server/internal"
	"github.com/FreitasGabriel/client-server-api/server/tools/model"
	"gorm.io/gorm"
)

func GetQuotation(ctx context.Context, db *gorm.DB) (*model.QuotationAPI, error) {
	ctxAPI, cancelAPI := internal.GetContext(ctx, 200*time.Millisecond)
	defer cancelAPI()

	var quotation model.QuotationAPI
	const quotationURL = "https://economia.awesomeapi.com.br/json/last/USD-BRL"

	req, err := http.NewRequestWithContext(ctxAPI, "GET", quotationURL, nil)
	if err != nil {
		log.Println("error to mount request", err)
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("error to execute request", err)
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("error to read info", err)
		return nil, err
	}

	err = json.Unmarshal(body, &quotation)
	if err != nil {
		log.Println("error to unmarshal data", err)
		return nil, err
	}

	repository.CreateQuotation(ctx, db, &quotation)

	select {
	case <-ctxAPI.Done():
		log.Println("request canceled by client or timeout excced")
		return nil, ctxAPI.Err()
	default:
		return &quotation, nil
	}
}
