package handlers

// Imported Packages
import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Delete Handler
func Delete(w http.ResponseWriter, r *http.Request) {
	// Connect to db
	//client := db.Connect()

	data := mux.Vars(r)
	account := data["account"]
	id := data["id"]

	if account == "" || id == "" {
		http.Error(w, "missing value", 400)
	}

	if _, err := strconv.Atoi(id); err != nil {
		http.Error(w, "ID should be an integer, not string", 400)
	}

	//fmt.Fprint(w, account, id)
	query := `
		DELETE FROM CUSTOMER
		WHERE account = $1 AND merchant_id = $2
	`
	if _, err := client.Exec(query, account, id); err != nil {
		log.Fatal("Something wrong")
	}
}
