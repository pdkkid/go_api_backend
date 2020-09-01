package product

import (
	"net/http"
	"strconv"
)

// FormToProduct -- Populate a product struct with form data
// vars:
// r -- request reader to fetch form data or url params
// returns:
// Product struct on success
// String arr of errors on failure
func FormToProduct(r *http.Request) (Product, []string) {
	var product Product
	var errStr, priceStr, activeStr string
	var errs []string
	var err error

	// Handle grabbing form data and putting into struct
	product.Name, errStr = processFormField(r, "name")
	errs = appendError(errs, errStr)
	product.Descript, errStr = processFormField(r, "description")
	errs = appendError(errs, errStr)

	//Some data validation is required for non-string input to db
	priceStr, errStr = processFormField(r, "price")
	if len(errStr) != 0 {
		errs = append(errs, errStr)
	} else {
		product.Price, err = strconv.ParseFloat(priceStr, 64)
		if err != nil {
			errs = append(errs, "Parameter 'price' is not a float")
		}
	}

	activeStr, errStr = processFormField(r, "active")
	if len(errStr) != 0 {
		errs = append(errs, errStr)
	} else {
		product.Active, err = strconv.ParseBool(activeStr)
		if err != nil {
			errs = append(errs, "Parameter 'active' is not a valid boolean")
		}
	}

	return product, errs
}

func appendError(errs []string, errStr string) []string {
	if len(errStr) > 0 {
		errs = append(errs, errStr)
	}
	return errs
}

// get the form data from client & handle data separation
func processFormField(r *http.Request, field string) (string, string) {
	fieldData := r.PostFormValue(field)
	if len(fieldData) == 0 {
		return "", "Missing '" + field + "' parameter, cannot continue"
	}
	return fieldData, ""
}
