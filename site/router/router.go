package router

// Imported Packages
import (
	"billing/site/handlers"

	"github.com/gorilla/mux"
)

var muxRouter *mux.Router

func Init() (err error) {
	// Creating new Handler
	muxRouter = mux.NewRouter()

	// Handlers
	muxRouter.HandleFunc("/upload", handlers.Upload)
	muxRouter.HandleFunc("/create", handlers.Create).Methods("GET")
	muxRouter.HandleFunc("/create", handlers.Create).Methods("POST")
	muxRouter.HandleFunc("/user/{account}/{id}", handlers.Show).Methods("GET")
	muxRouter.HandleFunc("/delete/{account}/{id}", handlers.Delete).Methods("GET")
	muxRouter.HandleFunc("/update", handlers.Update).Methods("POST")

	return
}

func Get() *mux.Router {
	return muxRouter
}
