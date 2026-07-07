package main
import (
	"fmt"
	"net/http"
)
func helloHanler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "hello world")
}

func main(){
	router := http.NewServeMux()
	router.HandleFunc("GET /hello", helloHanler)

	fmt.Println("server starting on localhost:8080")
	if err := http.ListenAndServe(":8080", router);err != nil{
		panic(err)
	}
}