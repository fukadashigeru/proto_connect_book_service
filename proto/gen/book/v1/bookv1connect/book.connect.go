// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: book/v1/book.proto

package bookv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	http "net/http"
	v1 "proto_connect_book_service/proto/gen/book/v1"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// BookServiceName is the fully-qualified name of the BookService service.
	BookServiceName = "book.v1.BookService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// BookServiceAddBookProcedure is the fully-qualified name of the BookService's AddBook RPC.
	BookServiceAddBookProcedure = "/book.v1.BookService/AddBook"
	// BookServiceGetBookProcedure is the fully-qualified name of the BookService's GetBook RPC.
	BookServiceGetBookProcedure = "/book.v1.BookService/GetBook"
	// BookServiceListBooksProcedure is the fully-qualified name of the BookService's ListBooks RPC.
	BookServiceListBooksProcedure = "/book.v1.BookService/ListBooks"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	bookServiceServiceDescriptor         = v1.File_book_v1_book_proto.Services().ByName("BookService")
	bookServiceAddBookMethodDescriptor   = bookServiceServiceDescriptor.Methods().ByName("AddBook")
	bookServiceGetBookMethodDescriptor   = bookServiceServiceDescriptor.Methods().ByName("GetBook")
	bookServiceListBooksMethodDescriptor = bookServiceServiceDescriptor.Methods().ByName("ListBooks")
)

// BookServiceClient is a client for the book.v1.BookService service.
type BookServiceClient interface {
	AddBook(context.Context, *connect.Request[v1.AddBookRequest]) (*connect.Response[v1.AddBookResponse], error)
	GetBook(context.Context, *connect.Request[v1.GetBookRequest]) (*connect.Response[v1.GetBookResponse], error)
	ListBooks(context.Context, *connect.Request[v1.ListBooksRequest]) (*connect.Response[v1.ListBooksResponse], error)
}

// NewBookServiceClient constructs a client for the book.v1.BookService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewBookServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) BookServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &bookServiceClient{
		addBook: connect.NewClient[v1.AddBookRequest, v1.AddBookResponse](
			httpClient,
			baseURL+BookServiceAddBookProcedure,
			connect.WithSchema(bookServiceAddBookMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getBook: connect.NewClient[v1.GetBookRequest, v1.GetBookResponse](
			httpClient,
			baseURL+BookServiceGetBookProcedure,
			connect.WithSchema(bookServiceGetBookMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listBooks: connect.NewClient[v1.ListBooksRequest, v1.ListBooksResponse](
			httpClient,
			baseURL+BookServiceListBooksProcedure,
			connect.WithSchema(bookServiceListBooksMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// bookServiceClient implements BookServiceClient.
type bookServiceClient struct {
	addBook   *connect.Client[v1.AddBookRequest, v1.AddBookResponse]
	getBook   *connect.Client[v1.GetBookRequest, v1.GetBookResponse]
	listBooks *connect.Client[v1.ListBooksRequest, v1.ListBooksResponse]
}

// AddBook calls book.v1.BookService.AddBook.
func (c *bookServiceClient) AddBook(ctx context.Context, req *connect.Request[v1.AddBookRequest]) (*connect.Response[v1.AddBookResponse], error) {
	return c.addBook.CallUnary(ctx, req)
}

// GetBook calls book.v1.BookService.GetBook.
func (c *bookServiceClient) GetBook(ctx context.Context, req *connect.Request[v1.GetBookRequest]) (*connect.Response[v1.GetBookResponse], error) {
	return c.getBook.CallUnary(ctx, req)
}

// ListBooks calls book.v1.BookService.ListBooks.
func (c *bookServiceClient) ListBooks(ctx context.Context, req *connect.Request[v1.ListBooksRequest]) (*connect.Response[v1.ListBooksResponse], error) {
	return c.listBooks.CallUnary(ctx, req)
}

// BookServiceHandler is an implementation of the book.v1.BookService service.
type BookServiceHandler interface {
	AddBook(context.Context, *connect.Request[v1.AddBookRequest]) (*connect.Response[v1.AddBookResponse], error)
	GetBook(context.Context, *connect.Request[v1.GetBookRequest]) (*connect.Response[v1.GetBookResponse], error)
	ListBooks(context.Context, *connect.Request[v1.ListBooksRequest]) (*connect.Response[v1.ListBooksResponse], error)
}

// NewBookServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewBookServiceHandler(svc BookServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	bookServiceAddBookHandler := connect.NewUnaryHandler(
		BookServiceAddBookProcedure,
		svc.AddBook,
		connect.WithSchema(bookServiceAddBookMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	bookServiceGetBookHandler := connect.NewUnaryHandler(
		BookServiceGetBookProcedure,
		svc.GetBook,
		connect.WithSchema(bookServiceGetBookMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	bookServiceListBooksHandler := connect.NewUnaryHandler(
		BookServiceListBooksProcedure,
		svc.ListBooks,
		connect.WithSchema(bookServiceListBooksMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/book.v1.BookService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case BookServiceAddBookProcedure:
			bookServiceAddBookHandler.ServeHTTP(w, r)
		case BookServiceGetBookProcedure:
			bookServiceGetBookHandler.ServeHTTP(w, r)
		case BookServiceListBooksProcedure:
			bookServiceListBooksHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedBookServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedBookServiceHandler struct{}

func (UnimplementedBookServiceHandler) AddBook(context.Context, *connect.Request[v1.AddBookRequest]) (*connect.Response[v1.AddBookResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("book.v1.BookService.AddBook is not implemented"))
}

func (UnimplementedBookServiceHandler) GetBook(context.Context, *connect.Request[v1.GetBookRequest]) (*connect.Response[v1.GetBookResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("book.v1.BookService.GetBook is not implemented"))
}

func (UnimplementedBookServiceHandler) ListBooks(context.Context, *connect.Request[v1.ListBooksRequest]) (*connect.Response[v1.ListBooksResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("book.v1.BookService.ListBooks is not implemented"))
}
