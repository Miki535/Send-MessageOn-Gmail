package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/smtp"
)

var tpl = template.Must(template.ParseFiles("templates/index.html"))

func main() {
	http.HandleFunc("/", HomeFunc)
	http.ListenAndServe(":8080", nil)
}

func HomeFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		ownGmail := r.FormValue("ownGmail")
		secretKEY := r.FormValue("KEY")
		gmail := r.FormValue("gmail")
		message := r.FormValue("message")
		auth := smtp.PlainAuth("", ownGmail, secretKEY, "smtp.gmail.com")

		to := []string{gmail}
		msg := message

		err := smtp.SendMail("smtp.gmail.com:587", auth, ownGmail, to, []byte(fmt.Sprint(msg)))

		if err != nil {
			http.Error(w, "ERROR! \n StatusBadRequest", http.StatusBadRequest)
			return
		}

	}
	tpl.Execute(w, nil)
}
