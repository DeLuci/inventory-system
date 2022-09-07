package handlers

import (
	"encoding/json"
	"github.com/DeLuci/inventory-system/internal/config"
	"github.com/DeLuci/inventory-system/internal/driver"
	"github.com/DeLuci/inventory-system/internal/forms"
	"github.com/DeLuci/inventory-system/internal/helpers"
	"github.com/DeLuci/inventory-system/internal/models"
	"github.com/DeLuci/inventory-system/internal/render"
	"github.com/DeLuci/inventory-system/internal/repository"
	"github.com/DeLuci/inventory-system/internal/repository/dbrepo"
	"github.com/go-chi/chi/v5"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

// Repo the repository used by handlers
var Repo *Repository

// Repository is holds the app config and the db repository
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db.SQL, a),
	}
}

// NewHandlers set the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "login.page.gohtml", &models.TemplateData{
		Form: forms.New(nil),
	})
}

// PostLogin handles logging the user in
func (m *Repository) PostLogin(w http.ResponseWriter, r *http.Request) {
	_ = m.App.Session.RenewToken(r.Context())
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	pwd := r.Form.Get("password")
	email := r.Form.Get("email")

	form := forms.New(r.PostForm)
	form.Required("email", "password")
	form.IsEmail("email")
	if !form.Valid() {
		render.Template(w, r, "login.page.gohtml", &models.TemplateData{
			Form: form,
		})
		return
	}

	id, _, err := m.DB.Authenticate(email, pwd)
	if err != nil {
		log.Println(err)

		m.App.Session.Put(r.Context(), "error", "Invalid login credentials")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	m.App.Session.Put(r.Context(), "user_id", id)
	m.App.Session.Put(r.Context(), "flash", "Logged in successfully")

}

func (m *Repository) SignUp(w http.ResponseWriter, r *http.Request) {
	var emptySignUp models.Signup
	data := make(map[string]interface{})
	data["signup"] = emptySignUp
	render.Template(w, r, "signup.page.gohtml", &models.TemplateData{
		Form: forms.New(nil),
	})
}

// PostSignUp sign up handles the posting of a reservation form
func (m *Repository) PostSignUp(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	pwd := []byte(r.Form.Get("password"))
	hashedPwd, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	user := models.User{
		Email:             r.Form.Get("email"),
		EmailConfirmation: r.Form.Get("email_confirmation"),
		Password:          string(hashedPwd),
	}

	user2 := models.User{
		Email:             r.Form.Get("email"),
		EmailConfirmation: r.Form.Get("email_confirmation"),
	}

	form := forms.New(r.PostForm)
	// validation
	form.Required("email", "email_confirmation", "password")
	form.IsEmail("email")
	form.IsEqual("email", "email_confirmation")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["signup"] = user2

		render.Template(w, r, "signup.page.gohtml", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	err = m.DB.InsertUser(user)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
}

func (m *Repository) Logout(w http.ResponseWriter, r *http.Request) {
	_ = m.App.Session.Destroy(r.Context())
	_ = m.App.Session.RenewToken(r.Context())

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (m *Repository) AdminDashboard(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "admin-dashboard.page.gohtml", &models.TemplateData{})
}

func (m *Repository) Dashboard(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "dashboard.page.gohtml", &models.TemplateData{
		Form: forms.New(nil),
	})
}

// SearchBarData gets the url parameter and finds the data in the
func (m *Repository) SearchBarData(w http.ResponseWriter, r *http.Request) {
	productParam := chi.URLParam(r, "product")
	//decodedValue := url.QueryEscape(productParam)
	//splitProduct := strings.Split(productParam, " ")
	items, err := m.DB.SearchProduct(productParam)
	if err != nil {
		log.Println(err)
	}
	j, _ := json.MarshalIndent(items, "", "   ")
	log.Println(string(j))
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(j)
	if err != nil {
		log.Print(err)
	}
}

func (m *Repository) ScanProduct(w http.ResponseWriter, r *http.Request) {
	var code models.ScanCode

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&code)
	if err != nil {
		panic(err)
	}

	id := code.Code[12:21]
	size := code.Code[21:]

	products := models.ScanProduct{
		ID:   id,
		Size: size,
	}

	err = m.DB.InsertNewProduct(products)
	if err != nil {
		helpers.ServerError(w, err)
	}
}
