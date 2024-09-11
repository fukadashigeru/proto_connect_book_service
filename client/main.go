package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"connectrpc.com/connect"

	bookv1 "proto_connect_book_service/proto/gen/book/v1"
	"proto_connect_book_service/proto/gen/book/v1/bookv1connect"
)

func main() {
	client := bookv1connect.NewBookServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
	)

	// Add a book
	addResp, err := client.AddBook(
		context.Background(),
		connect.NewRequest(&bookv1.AddBookRequest{
			Book: &bookv1.Book{
				Title:  "The Go Programming Language",
				Author: "Alan A. A. Donovan & Brian W. Kernighan",
			},
		}),
	)
	if err != nil {
		log.Fatalf("AddBook failed: %v", err)
	}
	fmt.Printf("Added book with ID: %s\n", addResp.Msg.Id)

	// Get the book
	getResp, err := client.GetBook(
		context.Background(),
		connect.NewRequest(&bookv1.GetBookRequest{
			Id: addResp.Msg.Id,
		}),
	)
	if err != nil {
		log.Fatalf("GetBook failed: %v", err)
	}
	fmt.Printf("Retrieved book: %v\n", getResp.Msg.Book)

	// List all books
	listResp, err := client.ListBooks(
		context.Background(),
		connect.NewRequest(&bookv1.ListBooksRequest{}),
	)
	if err != nil {
		log.Fatalf("ListBooks failed: %v", err)
	}
	fmt.Println("All books:")
	for _, book := range listResp.Msg.Books {
		fmt.Printf("- %s by %s (ID: %s)\n", book.Title, book.Author, book.Id)
	}
}
