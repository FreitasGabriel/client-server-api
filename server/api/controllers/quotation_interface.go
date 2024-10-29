package controllers

import (
	"net/http"

	"gorm.io/gorm"
)

func NewQuotationInterface(database *gorm.DB) QuotationControllerInterface {
	return &quotationControllerInterface{
		database: database,
	}
}

type quotationControllerInterface struct {
	database *gorm.DB
}

type QuotationControllerInterface interface {
	MakeQuotation(w http.ResponseWriter, r *http.Request)
}
