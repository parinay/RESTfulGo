package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"

	// gorm driver for posgreSQL
	"github.com/jinzhu/gorm"
	// gorm driver for posgres
	_ "github.com/jinzhu/gorm/dialects/postgres"

	//jwt
	jwt "github.com/dgrijalva/jwt-go"
)

var signingKey = os.Getenv("MY_JWT")

// isAuthorized() validates the token in the header and if so calls the endpoint
func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {

			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				byteSigningKey := []byte(signingKey)
				return byteSigningKey, nil
			})
			if err != nil {
				fmt.Fprintf(w, err.Error())
			}
			if token.Valid {
				endpoint(w, r)
			}
		} else {
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}

// homePage() to handle all request to root URL
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Homepage!")
	log.Println("Endpoint Hit: homepage")
}

// returnArticles() returns all articles
func returnArticles(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: returnArticles")

	// Connecto DB and create a record
	db, err := gorm.Open("postgres", "host=localhost port=5432 dbname=test_db user=postgres password=Par1nay! sslmode=disable")
	if err != nil {
		panic("failed to connect to Database")
	}
	defer db.Close()

	var articles []Article
	db.Find(&articles)

	json.NewEncoder(w).Encode(Articles)
}

// returnArticle() returns particualr article matching {id}
func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: returnSingleArticle")

	// Connecto DB and create a record
	db, err := gorm.Open("postgres", "host=localhost port=5432 dbname=test_db user=postgres password=Par1nay! sslmode=disable")
	if err != nil {
		panic("failed to connect to Database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	key := vars["id"]
	key1, _ := strconv.Atoi(key)

	var article1 Article
	db.Where("Ide = ?", key1).Find(&article1)
	db.Find(&article1)

	/*for _, article := range Articles {
		if article.Ide == key1 {
		}
	}*/
	json.NewEncoder(w).Encode(article1)
}

// createNewArticle() creates a new article
func createNewArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createNewArticle")
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))

	var article Article
	err = json.Unmarshal(reqBody, &article)
	if err != nil {
		log.Fatalf("Failed to Unmarshal request body %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Create new Article Failed due to some error"))
	}
	// Connecto DB and create a record
	db, err := gorm.Open("postgres", "host=localhost port=5432 dbname=test_db user=postgres password=Par1nay! sslmode=disable")
	if err != nil {
		panic("failed to connect to Database")
	}
	defer db.Close()
	db.Create(&article)

	Articles = append(Articles, article)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200  Create article successful"))
	json.NewEncoder(w).Encode(article)
}

// createNewArticleV2() creates a new article
func createNewArticleV2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createNewArticleV2")
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))

	var article Article
	err := json.Unmarshal(reqBody, &article)
	if err != nil {
		log.Fatalf("Failed to Unmarshal request body %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Create new Article Failed due to some error"))
	}
	// Connecto DB and create a record
	db, err := gorm.Open("postgres", "host=localhost port=5432 dbname=test_db user=postgres password=Par1nay! sslmode=disable")
	if err != nil {
		panic("failed to connect to Database")
	}
	defer db.Close()
	db.Create(&article)
	Articles = append(Articles, article)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200 - Create new Article Successful"))
	json.NewEncoder(w).Encode(article)
}

// deleteArticle() deletes an article
func deleteArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deleteArticle")

	vars := mux.Vars(r)
	id := vars["id"]
	id1, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Connecto DB and create a record
	db, err := gorm.Open("postgres", "host=localhost port=5432 dbname=test_db user=postgres password=Par1nay! sslmode=disable")
	if err != nil {
		panic("failed to connect to Database")
	}
	defer db.Close()

	var article1 Article
	// db.Debug().Model(&Article{}).Where("id = ?", id1).Take(&Article{}).Delete(&Article{})
	db.Where("Ide = ?", id1).Find(&article1)
	db.Delete(&article1)

	/*for index, article := range Articles {
		if article.Ide == id1 {
			Articles = append(Articles[:index], Articles[index+1:]...)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("200 - Delete article successful"))
		} else {
			log.Println("Record not found")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Delete article failed due to some error"))
		}
	}*/
}

// updateArticle() update existing article id
func updateArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: updateArticle")

	// Connecto DB and create a record
	db, err := gorm.Open("postgres", "host=localhost port=5432 dbname=test_db user=postgres password=Par1nay! sslmode=disable")
	if err != nil {
		panic("failed to connect to Database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]
	id1, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var article1 Article
	db.Where("Ide = ?", id1).Find(&article1)

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &article1)
	db.Save(&article1)

	json.NewEncoder(w).Encode(article1)
	/*for index, article := range Articles {
		if article.Ide == id1 {
			var article Article
			json.Unmarshal(reqBody, &article)
			Articles[index].Ide = article.Ide
			Articles[index].Title = article.Title
			Articles[index].Desc = article.Desc
			Articles[index].Content = article.Content
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("200  Update article successful"))
			json.NewEncoder(w).Encode(article)

		} else {
			log.Println("Record not found")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Delete article failed due to some error"))
		}
	}*/

}

// HandleRequests  matches URL paths to defined function
func HandleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.Handle("/api/v1/article", isAuthorized(returnArticles)).Methods("GET")
	myRouter.Handle("/api/v1/article/{id}", isAuthorized(returnSingleArticle)).Methods("GET")
	myRouter.Handle("/api/v1/article/{id}", isAuthorized(deleteArticle)).Methods("DELETE")
	myRouter.Handle("/api/v1/article/{id}", isAuthorized(updateArticle)).Methods("PUT")
	myRouter.Handle("/api/v1/article", isAuthorized(createNewArticle)).Methods("POST")
	myRouter.Handle("/api/v2/article", isAuthorized(createNewArticleV2)).Methods("POST")

	myRouter.Handle("/api/v1", isAuthorized(homePage))

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

// InitialMigration Schema migration
func InitialMigration() {
	// db, err = gorm.Open("postgres", "postgres:Par1nay!@/test.db?charset=utf8&parseTime=True&loc=Local")
	db, err = gorm.Open("postgres", "host=localhost port=5432 dbname=test_db user=postgres password=Par1nay! sslmode=disable")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect to database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Article{})
}
