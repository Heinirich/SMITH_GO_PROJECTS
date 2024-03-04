package handlers

import (
	"net/http"

	"github.com/Heinirich/basic-server/pkg/config"
	"github.com/Heinirich/basic-server/pkg/models"
	"github.com/Heinirich/basic-server/pkg/render"
)

// TemplateData holds data sent from handlers to templates

var Repo *Repository

// Repository is the repository used by the handlers
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new Repository with the given AppConfig.
// It takes a pointer to an AppConfig as a parameter and returns a pointer to a Repository.
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers initializes the Repository with the given parameter.
// r *Repository
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr

	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	
	//
	// Perfome
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
