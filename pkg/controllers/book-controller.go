package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/bxbdev/go-bookstore/pkg/models"
	"github.com/bxbdev/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	// get data from db
	newBooks := models.GetAllBooks()
	// convert data to json
	res, _ := json.Marshal(newBooks)
	// set header content type
	w.Header().Set("Content-Type", "pkglication/json")
	// check http status if 200
	w.WriteHeader(http.StatusOK)
	// write data to response
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	// from url values
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	// convert string to int
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	// get data from db
	bookDetails, _ := models.GetBookById(ID)
	// convert data to json
	res, _ := json.Marshal(bookDetails)
	// set header content type
	w.Header().Set("Content-Type", "pkglication/json")
	// check http status if 200
	w.WriteHeader(http.StatusOK)
	// write data to response
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	newBook := &models.Book{}
	utils.ParseBody(r, newBook)
	// because (b *Book) CreateBook() is a pointer in models
	b := newBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// make a new book object
	var updateBook = &models.Book{}
	// use utils to parse body
	utils.ParseBody(r, updateBook)
	// set value from http request
	vars := mux.Vars(r)
	// set bookId from vars["bookId"]
	bookId := vars["bookId"]
	// convert string to int
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	// found the book from database by id and stored into bookDetails
	bookDetails, db := models.GetBookById(ID)
	// check the condition if value is not empty then update the book details
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	// final save to database
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
