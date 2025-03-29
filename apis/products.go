package apis

import (
	"encoding/json"
	"fmt"
	jazidb "httpproj1/db"
	"httpproj1/logger"
	"httpproj1/models"
	"net/http"
)

func GetRoot(w http.ResponseWriter, r *http.Request) {
	globalDatabase := jazidb.GetDatabase()
	myslog := logger.GetLogger()

	rows, err := globalDatabase.Query(`select * from Product`)
	if err != nil {
		myslog.Error(err.Error())
	} else {
		fmt.Println(rows)
	}

	var products []models.Product
	var prd models.Product

	for rows.Next() {
		err := rows.Scan(&prd.Id, &prd.Title, &prd.Price)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println(product1.Id, product1.Title, product1.Price)
		products = append(products, prd)
	}
	jsoned_product, _ := json.Marshal(products)

	// io.WriteString(w, "This is jazi server\n")

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsoned_product) //better , io
	// io.WriteString(w, string(jsoned_product))

}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	globalDatabase := jazidb.GetDatabase()
	myslog := logger.GetLogger()

	body := json.NewDecoder(r.Body)
	fmt.Println(body)
	p := &models.Product{}

	err := body.Decode(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	res, err := globalDatabase.Query(`insert into Product VALUES ($1,$2,$3)`, p.Id, p.Title, p.Price)
	if err != nil {
		myslog.Error(err.Error())
	} else {
		fmt.Println(res)
	}

}
