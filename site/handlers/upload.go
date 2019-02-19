package handlers

// Imported Packages
import (
	"billing/site/db"
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/tealeg/xlsx"
)

var client *sql.DB

// Upload Handler
func Upload(w http.ResponseWriter, r *http.Request) {
	client = db.Connect()

	file := "site/excel/" + r.FormValue("file")
	xlFile, err := xlsx.OpenFile(file)

	if err != nil {
		panic(err)
	}

	for _, sheet := range xlFile.Sheets {
		for index, row := range sheet.Rows {
			if index == 0 {
				// We should skip header
				continue
			}
			var user []string
			for _, cell := range row.Cells {
				parsed := strings.TrimSpace(cell.String())
				user = append(user, parsed)
			}
			query := `
					SELECT COUNT(account) from customer
					WHERE account = $1
					LIMIT 1
				`
			// If row hasn't found
			if _, err := client.Query(query, user[0]); err != nil {
				err = addProduct(user[0], user[1], user[2], user[3])
				fmt.Println(err, "Added")
			}
		}
	}

	http.Redirect(w, r, "/forms", 200)
}

// AddProduct Function
func addProduct(account, name, surname, phone string) (err error) {
	// query := `
	// 	INSERT INTO customers(accout, name, surname, phone)
	// 	VALUES($1, $2, $3, $4)
	// `
	// _, err = client.Exec(query, model, name, company, price)

	return
}
