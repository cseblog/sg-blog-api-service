package main

import (
    "encoding/json"
    "log"
    "math/rand"
    "strconv"
    "net/http"
    "github.com/gorilla/mux"
)


// Model 
type Article struct {
    ID    string   `json:"id"`
    Title string   `json:"title"`
    Content  string   `json:"content"`
    Author   string `json:"author"`
}

type RequestError struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}


type ResponseResult struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type IdData struct {
	ID string `json:"id"`
}

var articles []Article

// Http handling functions
func GetArticle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(req)
    for _, item := range articles {
        if item.ID == params["id"] {
        	log.Printf("id: %s", item.ID)
        	var responseResult ResponseResult
			responseResult.Status = http.StatusOK
			responseResult.Message = "Sucess"
			responseResult.Data = item
    		json.NewEncoder(w).Encode(responseResult)
            return
        }
    }
    //
    var error404 RequestError
    error404.Status = http.StatusNotFound
    error404.Message = "Article is not found!"
    error404.Data = nil
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode(error404)
}

func GetAllArticles(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var responseResult ResponseResult
	responseResult.Status = http.StatusOK
	responseResult.Message = "Sucess"
	responseResult.Data = articles
    json.NewEncoder(w).Encode(responseResult)
}

func CreateArticle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    var article Article
    err := json.NewDecoder(req.Body).Decode(&article)
    if err != nil {
    	var invalidReq RequestError
    	invalidReq.Status = http.StatusBadRequest
    	invalidReq.Message = "Bad request!"
    	invalidReq.Data = nil
    	w.WriteHeader(http.StatusBadRequest)
    	json.NewEncoder(w).Encode(invalidReq)
    	return
    }

    
    article.ID = strconv.Itoa(rand.Intn(100000000))
    articles = append(articles, article)
    w.WriteHeader(http.StatusCreated)

    var responseResult ResponseResult
    var id IdData
    id.ID = article.ID
	responseResult.Status = http.StatusCreated
	responseResult.Message = "Sucess"
	responseResult.Data = id
    json.NewEncoder(w).Encode(responseResult)
}


func main() {
    router := mux.NewRouter()
    router.HandleFunc("/articles", GetAllArticles).Methods("GET")
    router.HandleFunc("/articles/{id}", GetArticle).Methods("GET")
    router.HandleFunc("/articles", CreateArticle).Methods("POST")
    log.Fatal(http.ListenAndServe(":8080", router))
}