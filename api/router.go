package api

import (
	"github.com/gorilla/mux"
)

func SetupRouter(handler *APIHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/block/{index:[0-9]+}", handler.GetBlockByIndexHandler).Methods("GET")
	r.HandleFunc("/block/hash/{hash:[a-zA-Z0-9]+}", handler.GetBlockByIndexHandler).Methods("GET")
	r.HandleFunc("/balance/{address:[a-zA-Z0-9]+}", handler.GetBalanceHandler).Methods("GET")
	// 其他路由设置
	return r
}
