package web

import (
	"fmt"
	"html/template"
	"net/http"
)

type service interface {
	GetOrderFromCache(orderUUID string) (rawJson string, err error)
}

type server struct {
	tmpHome     *template.Template
	tmpOrder    *template.Template
	tmpNotFound *template.Template
	srv         service
}

func New(srv service) *server {
	return &server{
		tmpHome:     template.Must(template.ParseFiles("../../html/home.html")),
		tmpOrder:    template.Must(template.ParseFiles("../../html/order.html")),
		tmpNotFound: template.Must(template.ParseFiles("../../html/notfound.html")),
		srv:         srv,
	}
}

func (s *server) Run() error {
	http.HandleFunc("/", s.handleHome)
	http.HandleFunc("/order", s.handleOrder)

	fmt.Println("serve")
	return http.ListenAndServe(":80", nil)
}
