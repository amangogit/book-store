package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/amangogit/book-store/pkg/models"
	"github.com/amangogit/book-store/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(resp http.ResponseWriter, req *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	resp.Header().Set("Content-Type", "pkglication/json")
	resp.WriteHeader(http.StatusOK)
	resp.Write(res)
}

func CreateBook(resp http.ResponseWriter, req *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(req, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	resp.WriteHeader(http.StatusOK)
	resp.Write(res)

}

func GetBookById(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	resp.Header().Set("Content-Type", "pkhlication/json")
	resp.WriteHeader(http.StatusOK)
	resp.Write(res)
}

func DeleteBookById(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error in parsing")
	}

	book := models.DeleteBookID(ID)
	res, _ := json.Marshal(book)

	resp.Header().Set("Content-Type", "pkglication/json")
	resp.WriteHeader(http.StatusOK)
	resp.Write(res)
}

func UpdateBookById(resp http.ResponseWriter, req *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(req, updateBook)

	vars := mux.Vars(req)

	bookId := vars["bookId"]
	
	ID, err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("error in parsing")
	}
	bookDetails , db := models.GetBookById(ID)
	if updateBook.Name != ""{
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != ""{
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != ""{
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	resp.Header().Set("Content-Type", "pkglication/json")
	resp.WriteHeader(http.StatusOK)
	resp.Write(res)
}