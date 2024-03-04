package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Heinirich/basic-server/pkg/config"
	"github.com/Heinirich/basic-server/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig;

// NewTemplate description of the Go function.
// Receives a pointer to a AppConfig and does not return anything.
func NewTemplate(a *config.AppConfig) { 
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renders template using html
func RenderTemplate(w http.ResponseWriter, tmpl string,td *models.TemplateData) {

	var tc map[string]*template.Template

	if app.UseCache {
	
		tc = app.TemplateCache 
	}else{
		tc, _ = CreateTemplateCache()
		
	}

	// tc := app.TemplateCache

	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)

	if err != nil {
		log.Fatal(err)
	}

	//
	//psTemplate, err := template.ParseFiles("./templates/" + tmpl + ".page.html")
	//if err != nil {
	//	fmt.Println("Error parsing template:", err)
	//	http.Error(w, "Internal server error", http.StatusInternalServerError)
	//	return
	//}
	//
	//err = psTemplate.Execute(w, nil)
	//if err != nil {
	//	fmt.Println("Error executing template:", err)
	//	http.Error(w, "Internal server error", http.StatusInternalServerError)
	//}
}

// CreateTemplateCache creates template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")

	if err != nil {
		fmt.Println(err)
	}

	for _, page := range pages {

		fmt.Println(page)
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			fmt.Println("Error parsing template:", err)
			// http.Error(w, "Internal server error", http.StatusInternalServerError)
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")

		if err != nil {
			fmt.Println(err)
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				fmt.Println("Error parsing template:", err)
				// http.Error(w, "Internal server error", http.StatusInternalServerError)
				return myCache, err
			}

			myCache[name] = ts
		}
	}

	return myCache, nil
}
