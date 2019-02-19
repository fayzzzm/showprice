package handlers

// Imported Packages
import (
	"billing/site/db"
	"fmt"
	"net/http"
	"strings"
)

//Create Handler
func Create(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		client := db.Connect()

		// Parsing our request
		if err := r.ParseForm(); err != nil {
			panic(err)
		}

		// Reading each FormValue from parsed request
		account := strings.TrimSpace(r.FormValue("account"))
		name := strings.TrimSpace(r.FormValue("name"))
		surname := strings.TrimSpace(r.FormValue("surname"))
		phone := strings.TrimSpace(r.FormValue("phone"))
		id := strings.TrimSpace(r.FormValue("id"))
		// end

		query := `
			SELECT * FROM customer
			WHERE (account = $1 AND merchant_id = $2) OR phone = $3
		`

		if row, err := client.Query(query, account, id, phone); err == nil {
			//http.Redirect(w, r, "/create", 301)
			fmt.Println(row)
			http.Error(w, "The data consists", 400)
		}

		query = `
			INSERT INTO customer(account, full_name, phone, merchant_id) 
			VALUES($1, $2, $3, $4)
		`

		// No connection to db
		if _, err := client.Exec(query, account, name+surname, phone, id); err != nil {
			panic(err)
		}

		// Redirect to show or index.html
		http.Redirect(w, r, "/", 301)
	} else {
		http.ServeFile(w, r, "site/templates/forms.html")
	}
}
