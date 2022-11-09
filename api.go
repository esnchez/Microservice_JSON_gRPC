package main

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/esnchez/types"

)

type APIFunc func(context.Context, http.ResponseWriter, *http.Request) error

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

//this handler manages the response we will output if the service call fails and returns error
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

//api call
func (s *JSONAPIServer) handleFetchPrice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ticker := r.URL.Query().Get("ticker")

	//service call
	price, err := s.svc.FetchPrice(ctx,ticker)
	if err != nil {
		return err
	}
	priceResp := types.PriceResponse {
		Price : price,
		Ticker : ticker,
	}

	//the api call if succeeds returns PriceResponse json formatted + statusOK
	return writeJSON(w, http.StatusOK, &priceResp)
}


func writeJSON(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	return json.NewEncoder(w).Encode(v)
}