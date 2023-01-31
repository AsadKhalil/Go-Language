package api

import (
	"net/http"
)

func Router() *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc("/login", handleLogin)
	r.HandleFunc("/signup", handleSignUp)
	r.HandleFunc("/products", handleProducts)
	r.HandleFunc("/buy", handlePurchase)
	r.HandleFunc("/credit-cards", handleAddCreditCard)
	return r
}
