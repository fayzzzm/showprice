package handlers

import (
	"net/http"
)

// Pay Handler
func Pay(w http.ResponseWriter, r *http.Request) {
	// client := db.Connect()

	// err := r.ParseForm()
	// //------------------\\
	// account := r.FormValue("account")
	// service := r.FormValue("service")
	// //------------------\\
	// query := `
	// 	SELECT * FROM service
	// 	WHERE service_description = $1
	// `
	// row := client.QueryRow(query, service)

	// var (
	// 	id, price                       int
	// 	serviceDescription, serviceName string
	// )

	// fmt.Println(service)
	// err = row.Scan(&id, &serviceDescription, &serviceName, &price)
	// //fmt.Println(id, service_description)
	// query = `
	// 	INSERT INTO transaction
	// 	VALUES($1, $2, $3, $4, $5)
	// `
	// fmt.Println(id, serviceDescription, serviceName, price)
	// _, err = client.Exec(query, account, id, serviceDescription, serviceName, price)

	// if err != nil {
	// 	panic(err.Error())
	// }

	// query = `
	// 	SELECT * FROM billing
	// 	WHERE account = $1
	// `
	// userBilling := client.QueryRow(query, account)

	// var (
	// 	balance, credit, total int
	// )

	// err = userBilling.Scan(&account, &balance, &credit, &total)
	// fmt.Print(credit)

	// if price > balance {
	// 	credit += price
	// } else {
	// 	balance -= price
	// 	query = `
	// 		INSERT INTO transaction
	// 		VALUES($1, $2, 'Погашение', $3, $4)
	// 	`
	// 	fmt.Println("Am here bro")
	// 	_, err = client.Exec(query, account, id, serviceName, price)
	// }

	// total = balance - credit
	// query = `
	// 	UPDATE billing
	// 	SET balance = $1, credit = $2, total = $3
	// 	WHERE account = $4
	// `

	// _, err = client.Exec(query, balance, credit, total, account)

	// redirect := "/login?account=" + account
	// http.Redirect(w, r, redirect, 301)
}
