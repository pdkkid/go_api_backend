package product

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gocql/gocql"
	"github.com/pdkkid/go_api_backend/cassandra"
)

// Post function to handle a post request to the products endpoint
func Post(w http.ResponseWriter, r *http.Request) {
	var errs []string
	var gocqlUUID gocql.UUID

	product, errs := FormToProduct(r)

	// did it create the new product?
	var created bool = false

	if len(errs) == 0 {
		fmt.Println("trying to create a new product")
		// we need a unique UUID
		gocqlUUID = gocql.TimeUUID()
		// lets do this
		if err := cassandra.Session.Query(`
		insert into products (productid, product_name, product_description, product_price, product_active) values (?, ?, ?, ?, ?)`,
			gocqlUUID, product.Name, product.Descript, product.Price, product.Active).Exec(); err != nil {
			errs = append(errs, err.Error())
		} else {
			created = true
		}
	}
	//return json info including ID/Errors
	if created {
		fmt.Println("product_id", gocqlUUID)
		json.NewEncoder(w).Encode(NewProductResponse{ID: gocqlUUID})
	} else {
		fmt.Println("errors", errs)
		json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
	}
}
