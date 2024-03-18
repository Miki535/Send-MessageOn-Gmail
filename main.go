package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/smtp"
)

// templates:
var tpl = template.Must(template.ParseFiles("templates/index.html"))
var tplp = template.Must(template.ParseFiles("templates/about.html"))

func main() {
	http.HandleFunc("/about", about)
	http.HandleFunc("/", HomeFunc)
	//start server on localhost :8080
	http.ListenAndServe(":8080", nil)
}

func HomeFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		gmail := r.FormValue("gmail")
		message := r.FormValue("message")
		// dont forget change ENTERGMAIL@gmail.com on real gmail!
		// authentication in SMTP // Don`t forget change secret code!!!
		auth := smtp.PlainAuth("", "ENTERGMAIL@gmail.com", "YOU-SECRET-CODE", "smtp.gmail.com")

		to := []string{gmail}
		msg := message
		// send message on gmail

		// don`t forget change ENTERGMAIL@gmail.com on real gmail!
		err := smtp.SendMail("smtp.gmail.com:587", auth, "ENTERGMAIL@gmail.com", to, []byte(fmt.Sprint(msg)))

		if err != nil {
			http.Error(w, "ERROR! \n StatusBadRequest", http.StatusBadRequest)
			return
		}

	}
	tpl.Execute(w, nil)
}

func about(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

	}
	tplp.Execute(w, nil)

}
