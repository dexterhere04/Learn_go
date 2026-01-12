package main
import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/dexterhere04/Learn_go/internal/handlers"
	logs "github.com/siruspen/logrus"
)

func main(){
	log.SetReportCaller(true)
	var r *chi.Mux= chi.NewRouter()
	handlers.Handler(r)
	fmt.Println("Starting Go API service..")
	fmt.Println('
	Goooooo
	APIIIIIIIIIII
	')
	err := http.ListenAndServe("localhost:8000",r)
	if err != nil{
		log.Error(err)
	}

}