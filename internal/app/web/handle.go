package web

import (
	"fmt"
	"net/http"
	"strings"
)

func (s *server) handleHome(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Println(r.FormValue("orderUUID"))
		w.Write([]byte(r.FormValue("orderUUID")))
		return
	}
	s.tmpHome.Execute(w, nil)
}

func (s *server) handleOrder(w http.ResponseWriter, r *http.Request) {
	uuid := r.FormValue("orderUUID")

	if uuid == "" {
		w.Write([]byte("You need fill order id"))
		return
	}
	data, err := s.srv.GetOrderFromCache(strings.Trim(uuid, " "))

	if err != nil {
		s.tmpNotFound.Execute(w, nil)
		return
	}
	s.tmpOrder.Execute(w, data)

}
