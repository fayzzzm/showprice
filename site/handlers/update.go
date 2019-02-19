package handlers

import (
	"billing/site/db"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func Update(w http.ResponseWriter, r *http.Request) {
	client := db.Connect()

	// Parsing our request
	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	// Reading each FormValue from parsed request
	account := (r.FormValue("account"))
	name := (r.FormValue("name"))
	surname := (r.FormValue("surname"))
	phone := (r.FormValue("phone"))
	id := (r.FormValue("id"))
	oldUser := (r.FormValue("old-user"))
	// end

	if !checkForm(&account, &name, &surname, &phone, &id) {
		http.Error(w, "Value is missed", 400)
		return
	}

	query := `
		UPDATE customer 
		SET account = $1, full_name = $2, phone = $3
	 	WHERE merchant_id = $4 AND account = $5
	`

	if _, err := client.Exec(query, account, name+surname, phone, id, oldUser); err != nil {
		log.Fatalf("Account or phone consists at Database\n")
	}
	fmt.Println(account, name+surname, phone, id, oldUser)
	fmt.Fprintf(w, "The new account is: %s", account)
}

func checkForm(values ...interface{}) bool {
	for index, value := range values {

		if value == "" {
			fmt.Println(index, value)
			return false
		}
		value = strings.TrimSpace(value.(string))
	}

	return true
}
