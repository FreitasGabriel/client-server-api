package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/FreitasGabriel/client-server-api/server/api/service"
	"github.com/FreitasGabriel/client-server-api/server/tools/model"
)

func (qc quotationControllerInterface) MakeQuotation(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	result, err := service.GetQuotation(ctx, qc.database)
	if err != nil {
		log.Println("error to get quotation", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	formated := &model.FormatedQuotation{
		Dolar: result.USDBRL.Bid,
	}

	json.NewEncoder(w).Encode(formated)
}
