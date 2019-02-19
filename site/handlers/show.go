package handlers

// Imported Packages
import (
	"billing/site/db"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Show Handler
func Show(w http.ResponseWriter, r *http.Request) {
	// Connect to db
	client := db.Connect()

	//------------------\\
	data := mux.Vars(r)
	account := data["account"]
	id := data["id"]
	//------------------\\

	query := `
		SELECT * FROM customer
		WHERE account = $1 AND merchant_id = $2
		LIMIT 1
	`

	rows, err := client.Query(query, account, id)

	if err != nil {
		http.Error(w, "Error data for getting", 404)
	}

	for rows.Next() {
		var (
			account, fullName, phone, date, info string
			merchantID, id                       int
		)
		// Scan each data of row
		if err := rows.Scan(&id, &account, &fullName, &phone, &date, &merchantID, &info); err != nil {
			fmt.Println(err.Error())
			continue
		}

		jsonData := make(map[string]interface{})
		// Parsing JSON file
		if err := json.Unmarshal([]byte(info), &jsonData); err != nil {
			log.Fatalf("Expected Json file got - > %v", err)
		}
		// Exec data from JSON
		for key, data := range jsonData {
			fmt.Fprint(w, "Key : ", key, "-> value ", data, "\n")

		}
		// Setting Status Added
		w.Header().Add("Status:", "Added")
		// Print out
		fmt.Fprint(w, "Account          - > ", account, "\n",
			"Name and Surname - > ", fullName, "\n",
			"Date of creation - > ", date, "\n",
			"Merchant Id      - > ", merchantID, "\n",
			"Information      - > ", info)
	}

	header := w.Header()
	if _, consist := header["Status"]; !consist {
		http.Error(w, "Bad Request!!", 400)
	}
}
