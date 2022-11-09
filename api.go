package main

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
)

type APIFunc func(context.Context, http.ResponseWriter, *http.Request) error

type PriceResponse struct {
	Price float64 `json:"price"`
	Ticker string `json:"ticker"`
}

type JSONAPIServer struct {
	listenAddress string
	svc PriceFetcher
}

func NewJSONAPIService(listenAdr string, a PriceFetcher ) *JSONAPIServer {
	return &JSONAPIServer{
		listenAddress: listenAdr,
		svc: a,
	}
}


func makeHTTPHandlerFunc(apiFn APIFunc) http.HandlerFunc {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "requestID", rand.Intn(1000000))

	return func(w http.ResponseWriter, r *http.Request) {
		if err := apiFn(ctx, w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()}  )
		}
	}
}

func (s *JSONAPIServer) Run() {
	http.HandleFunc("/", makeHTTPHandlerFunc(s.handleFetchPrice))
	http.ListenAndServe(s.listenAddress, nil)
}

func (s *JSONAPIServer) handleFetchPrice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ticker := r.URL.Query().Get("ticker")

	price, err := s.svc.FetchPrice(ctx,ticker)
	if err != nil {
		return err
	}
	priceResp := PriceResponse {
		Price : price,
		Ticker : ticker,
	}

	return writeJSON(w, http.StatusOK, &priceResp)
}


func writeJSON(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	return json.NewEncoder(w).Encode(v)
}