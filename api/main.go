package main

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"os"
)

var MySigninKey = []byte(os.Getenv("SECRET_KEY"))

func homepage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Secret Information, Uncle SAM is watching you")
}

func isAuthorized(endpoint func(w http.ResponseWriter, r *http.Request)) http.Handler{
	return http.HandlerFunc(func ( w http.ResponseWriter, r *http.Request){
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func( token *jwt.Token) (interface{}, error){
				if _, ok := token.Method.(*jwt.SigningMethodHMAC);!ok{
					return nil, fmt.Errorf("Invalid signing method")
				}
				aud := "billing.jwtgo.io"
				checkAudience := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
				if !checkAudience{
					return nil, fmt.Errorf("invalid audience")
				}
				iss := "jwtgo.io"
				checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
				if !checkIss{
					return nil, fmt.Errorf("Invalid Issuer")
				}
				return MySigninKey, nil
			})

			if err != nil{
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid{
				endpoint(w,r)
			}
			
		}else {
			fmt.Fprintf(w, "No Authorization token provided")
		}
	})
}

func handleRequests(){
	http.Handle("/",isAuthorized(homepage))
	log.Fatal(http.ListenAndServe(":9001",nil))

}

func main(){
	fmt.Printf("Authentication Server")
	handleRequests()
}