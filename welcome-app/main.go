package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

//Create a struct that holds information to be displayed in our HTML file
type Welcome struct {
	Name string
	Time string
}

func main() {
	welcome := Welcome{"User", time.Now().Format(time.Stamp)}

	templates := template.Must(template.ParseFiles("templates/welcome-template.html"))

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if name := r.FormValue("name"); name != "" {
			welcome.Name = name
		}
		if err := templates.ExecuteTemplate(w, "welcome-template.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/login", login)

	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
func login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Error")
	}
	fmt.Fprintln(w, "Username : ", r.Form.Get("username"))
	fmt.Fprintln(w, "Password : ", r.Form.Get("password"))
	fmt.Println(r.Form.Get("username"))
	fmt.Println(r.Form.Get("password"))
}
