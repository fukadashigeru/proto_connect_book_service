// bookapi/book_service.go
package bookapi

import (
	"context"
	"fmt"
	"net/http"

	bookv1 "proto_connect_book_service/proto/gen/book/v1"
	"proto_connect_book_service/proto/gen/book/v1/bookv1connect"

	"connectrpc.com/connect"
)

type BookService struct {
	books map[string]*bookv1.Book
}

func NewBookService() *BookService {
	return &BookService{
		books: make(map[string]*bookv1.Book),
	}
}

func HandleBook(mux *http.ServeMux) {
	bookService := NewBookService()
	path, handler := bookv1connect.NewBookServiceHandler(bookService)
	mux.Handle(path, handler)
}

func (s *BookService) AddBook(
	ctx context.Context,
	req *connect.Request[bookv1.AddBookRequest],
) (*connect.Response[bookv1.AddBookResponse], error) {
	book := req.Msg.Book
	book.Id = fmt.Sprintf("%d", len(s.books)+1)
	s.books[book.Id] = book
	return connect.NewResponse(&bookv1.AddBookResponse{Id: book.Id}), nil
}

func (s *BookService) GetBook(
	ctx context.Context,
	req *connect.Request[bookv1.GetBookRequest],
) (*connect.Response[bookv1.GetBookResponse], error) {
	book, ok := s.books[req.Msg.Id]
	if !ok {
		return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("book not found"))
	}
	return connect.NewResponse(&bookv1.GetBookResponse{Book: book}), nil
}

func (s *BookService) ListBooks(
	ctx context.Context,
	req *connect.Request[bookv1.ListBooksRequest],
) (*connect.Response[bookv1.ListBooksResponse], error) {
	books := make([]*bookv1.Book, 0, len(s.books))
	for _, book := range s.books {
		books = append(books, book)
	}
	return connect.NewResponse(&bookv1.ListBooksResponse{Books: books}), nil
}
