package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/FreitasGabriel/client-server-api/client/internal/files"
)

type Quotation struct {
	Dolar string `json:"dolar"`
}

func main() {
	var quot Quotation
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Println("error to create request", err)
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("error to execute request", err)
		return
	}

	defer res.Body.Close()

	buff, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("error to buffer data", err)
		return
	}

	err = json.Unmarshal(buff, &quot)
	if err != nil {
		log.Println("error to unmarshal data", err)
		return
	}

	err = files.WriteFile(fmt.Sprintf("Dolar: %s", quot.Dolar))
	if err != nil {
		return
	}

	defer res.Body.Close()
}
