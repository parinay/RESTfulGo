package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

var signingKey = os.Getenv("MY_JWT")

// var signingKey = []byte("restfuleAPIinGO ")

func GenerateJWT() (string, error) {
	fmt.Println("In GenerateJWT")

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "RESTfulAPIinGO"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	fmt.Printf("My Signing Key is %s\n", signingKey)
	byteSigningKey := []byte(signingKey)
	tokenString, err := token.SignedString(byteSigningKey)
	// fmt.Printf("TokenString : %s", tokenString)

	if err != nil {
		//		log.Errorf("Something went wrong %v", err)
		fmt.Printf("Something went wrong %v\n", err)
		return "", err
	}

	return tokenString, nil
}
func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Println("In homePage")
	validToken, err := GenerateJWT()
	fmt.Printf("Valid Token - %s\n", validToken)

	if err != nil {
		log.Printf("Failed to generate JWT with error %v", err)
		fmt.Printf("Failed to generate JWT with error %v\n", err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:10000/api/v1/article", nil)
	if err != nil {
		fmt.Printf("New request failed with error: %v\n", err)
	}
	req.Header.Set("Token", validToken)
	res, err := client.Do(req)

	if err != nil {
		fmt.Fprintf(w, "Error:%s", err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Errorf("Error reading the http response %v", err)
		fmt.Printf("Error reading the http response %v\n", err)
	}
	fmt.Fprintf(w, string(body))
}
func handleRequests() {
	fmt.Println("In handleRequest")
	http.HandleFunc("/", homePage)

	log.Fatal(http.ListenAndServe(":10001", nil))
}
func main() {
	fmt.Println("In main")
	handleRequests()
}
