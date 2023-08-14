package authcontroller

import (
	"crud-go-native/config"
	"crud-go-native/entities"
	"crud-go-native/models/usermodel"
	"errors"
	"html/template"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/auth/login.html")

		if err != nil {
			panic(err)
		}
		temp.Execute(w, nil)
	} else if r.Method == "POST" {
		var message error

		r.ParseForm()
		userinput := entities.UserInput{
			Username: r.Form.Get("username"),
			Password: r.Form.Get("password"),
		}
		log.Println(userinput)

		result, isValid := usermodel.GetUsername(userinput.Username)
		log.Println(result)

		if !isValid {
			message = errors.New("email salah")
		} else {
			match := CheckPasswordHash(userinput.Password, result.Password)
			if !match {
				message = errors.New("password salah")
				log.Println(match)
			}
		}

		if message != nil {
			data := map[string]interface{}{
				"error": message,
			}

			temp, _ := template.ParseFiles("views/auth/login.html")
			temp.Execute(w, data)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			session, _ := config.Strore.Get(r, config.SESSION_ID)

			session.Values["loggedIn"] = true
			session.Values["username"] = result.Username
			session.Values["email"] = result.Email
			session.Values["password"] = result.Password

			session.Save(r, w)

			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

	}
}

func Register(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/auth/register.html")

		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var user entities.User

		user.Username = r.FormValue("username")
		user.Email = r.FormValue("email")
		user.Password = r.FormValue("password")

		errMessage := make(map[string]interface{})

		if user.Username == "" {
			errMessage["Username"] = "Username harus di isi"
		}
		if user.Email == "" {
			errMessage["Email"] = "Email harus di isi"
		}
		if user.Password == "" {
			errMessage["Password"] = "Password harus di isi"
		}

		if len(errMessage) > 0 {
			data := map[string]interface{}{
				"validation": errMessage,
			}

			temp, err := template.ParseFiles("views/auth/register.html")

			if err != nil {
				panic(err)
			}

			temp.Execute(w, data)
		} else {
			hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			user.Password = string(hashPassword)

			_, err := usermodel.CreatUser(user)

			var message string

			if err != nil {
				message = "Proses registrasi gagal: " + message
			} else {
				message = "Registrasi berhasil, silahkan login"
			}

			data := map[string]interface{}{
				"pesan": message,
			}

			temp, err := template.ParseFiles("views/auth/register.html")

			if err != nil {
				panic(err)
			}

			temp.Execute(w, data)
		}
	}

}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Strore.Get(r, config.SESSION_ID)

	session.Options.MaxAge = -1
	session.Save(r, w)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
