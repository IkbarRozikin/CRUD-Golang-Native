package homecontroller

import (
	"crud-go-native/config"
	"html/template"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {

	session, _ := config.Strore.Get(r, config.SESSION_ID)
	if len(session.Values) == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {

			data := map[string]interface{}{
				"username": session.Values["username"],
			}
			temp, _ := template.ParseFiles("views/home/index.html")
			temp.Execute(w, data)
		}
	}
}
