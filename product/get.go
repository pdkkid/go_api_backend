package product

import (
	"encoding/json"
	"net/http"

	"github.com/gocql/gocql"
	"github.com/pdkkid/go_api_backend/cassandra"
)

// Get Provides all product via /product enpoint get request
func Get(w http.ResponseWriter, r *http.Request) {
	var productList []Product
	m := map[string]interface{}{}
	query := "select productid,product_active,product_description,product_name,product_price from products"
	iterable := cassandra.Session.Query(query).Iter()
	for iterable.MapScan(m) {
		productList = append(productList, Product{
			ID:       m["productid"].(gocql.UUID),
			Active:   m["product_active"].(bool),
			Descript: m["product_description"].(string),
			Name:     m["product_name"].(string),
			Price:    m["product_price"].(float64),
		})
		m = map[string]interface{}{}
	}
	setupResponse(&w, r)
	json.NewEncoder(w).Encode(AllProductsResponse{Products: productList})
}

//I Love CORS
func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
