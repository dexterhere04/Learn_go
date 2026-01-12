package main

import (
	"fmt"
	"github.com/dexterhere04/Learn_go/internal/handlers"
	"github.com/go-chi/chi"
	logs "github.com/siruspen/logrus"
	"net/http"
)

func main() {
	logs.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)
	fmt.Println("Starting Go API service..")
	fmt.Println(`
	Goooooo
	APIIIIIIIIIII
	`)
	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		logs.Error(err)
	}
}
