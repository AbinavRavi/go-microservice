package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"os"
)

var MySigninKey = []byte(os.Getenv("SECRET_KEY"))

func homepage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Secret Information, Uncle SAM is watching you")
}

func isAuthorized(w http.ResponseWriter, r *http.Request){
	return http.HandleFunc(func( w http.ResponseWriter, r *http.Request){
		if r.Header["Token"] != nil{
			
		}
	})
}

func handleRequests(){
	http.Handle("/",isAuthorized(homepage))
	log.Fatal(http.ListenAndServe("9001",nil))

}

func main(){
	fmt.Printf("Authentication Server")
	handleRequests()
}