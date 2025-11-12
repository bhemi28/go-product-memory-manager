package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
)

var (
	pagesDir    = "web/pages"
	templateDir = "web/templates"

	templateFiles *template.Template
)

type AuthData struct {
	IsLogin  bool
	IsSignup bool
}

func init() {
	var err error
	templateFiles, err = parseFiles(templateDir)
	if err != nil {
		log.Fatal("failed to parse template Files....!!!!!")
	}
}

func RegisterWebRoutes(r *chi.Mux) {
	r.Get("/signup", func(w http.ResponseWriter, r *http.Request) {
		RenderTemplate(w, "main", "auth.html", AuthData{
			IsLogin:  false,
			IsSignup: true,
		})
	})
}

func parseFiles(dir string) (*template.Template, error) {

	tmpl := template.New("")

	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && filepath.Ext(path) == ".html" {
			_, parseErr := tmpl.ParseFiles(path)
			if parseErr != nil {
				return fmt.Errorf("failed to parse %s: %w", path, parseErr)
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return tmpl, nil
}

// RenderTemplate renders a composite HTML template by combining a shared template set (templateFiles)
// with a specific page file. It looks up the named template (tmplName) and executes it with the given data.
//
// Parameters:
//   - w: http.ResponseWriter to write the rendered template output.
//   - tmplName: name of the template block to execute (e.g., "auth_main", "profile_main").
//   - pageFile: relative file path under pagesDir to parse (e.g., "auth.html", "profile.html").
//   - data: dynamic data passed to the template execution (e.g., structs, maps).
//
// Example usage:
//
//	RenderTemplate(w, "auth_main", "auth.html", AuthPageData{IsLogin: true})
//
// Errors:
//   - Returns a 500 Internal Server Error if parsing or execution fails.
//   - Returns a 404 Not Found if tmplName is not defined in the base template set.
func RenderTemplate(w http.ResponseWriter, tmplName string, pageFile string, data any) {
	if templateFiles == nil {
		http.Error(w, "Template files not initialized", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(templateFiles.Clone())

	if tmpl.Lookup(tmplName) == nil {
		http.Error(w, fmt.Sprintf("Template %s not found", tmplName), http.StatusNotFound)
		return
	}

	tmpl = template.Must(tmpl.ParseFiles(fmt.Sprintf("%s/%s", pagesDir, pageFile)))

	err := tmpl.ExecuteTemplate(w, tmplName, data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to render template: %v", err), http.StatusInternalServerError)
	}
}
