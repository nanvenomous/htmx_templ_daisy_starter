package handle

import (
	"bytes"
	"embed"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/a-h/templ"
	"github.com/amalfra/etag/v3"
	"github.com/nanvenomous/htmx_templ_daisy_starter/models"
	"github.com/nanvenomous/htmx_templ_daisy_starter/ui"
)

//go:embed build/*
var buildFS embed.FS

var (
	embeddedResources fs.FS
)

// render takes a single templ component and writes it's html to the handler output
// it returns an error if something fails as well as a net/http status code
func render(w http.ResponseWriter, r *http.Request, cmp templ.Component) (models.HttpStatus, error) {
	err := cmp.Render(r.Context(), w)
	if err != nil {
		return models.StatusHTMLTemplatingError, fmt.Errorf("render: %v", err)
	}
	return models.StatusOK, nil
}

// renderMulti takes multiple templ components and writes their html to the handler output
// it returns an error if something fails as well as a net/http status code
func renderMulti(w http.ResponseWriter, r *http.Request, cmp ...templ.Component) (models.HttpStatus, error) {
	err := ui.Multi(cmp...).Render(r.Context(), w)
	if err != nil {
		return models.StatusHTMLTemplatingError, fmt.Errorf("renderMulti: %v", err)
	}
	return models.StatusOK, nil
}

// errorOut handles the case when the error should not be reflected in the UI
// TODO probably change this to a redirect to a dedicated error page
func errorOut(w http.ResponseWriter, status models.HttpStatus, err error) {
	if err != nil {
		log.Println(err)
	}
	http.Error(w, err.Error(), int(status))
}

// Setup initializes application net/http handlers
func Setup(mux *http.ServeMux) error {
	var err error
	embeddedResources, err = fs.Sub(buildFS, "build")
	if err != nil {
		return err
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if slices.Contains([]string{"", "/"}, r.URL.Path) {
			GetHome(w, r)
			return
		}
		serveResourceCachedETag(w, r, getBundledFile)
	})

	mux.HandleFunc("/homeButton", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second * 2)
		stts, err := renderMulti(w, r,
			ui.Button(ui.PropsButton{ID: "seeToastButton", Label: "WasClicked", Type: "submit", CSS: "btn-primary"}),
			ui.Alert(ui.PropsAlert{Label: "hi there", Type: ui.AlertTypeInfo}),
		)
		if err != nil {
			errorOut(w, stts, err)
		}
	})

	mux.HandleFunc("/table", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			GetTable(w, r)
			break
		default:
			errorOut(w, models.StatusMethodNotAllowed, err)
		}
	})

	return nil
}

func getBundledFile(r *http.Request) ([]byte, error) {
	if len(r.URL.Path) > 500 {
		return []byte{}, fmt.Errorf(
			"The filepath you entered: %s was suspiciously long",
			r.URL.Path,
		)
	}

	var pth = r.URL.Path
	if strings.HasPrefix(pth, "/") {
		pth = strings.TrimPrefix(pth, "/")
	}

	fl, err := embeddedResources.Open(pth)
	if err != nil {
		return []byte{}, err
	}
	defer fl.Close()

	return io.ReadAll(fl)
}

func serveResourceCachedETag(w http.ResponseWriter, r *http.Request,
	fileCheck func(r *http.Request) ([]byte, error),
) {
	w.Header().Set("Cache-Control", "max-age=0")
	content, err := fileCheck(r)
	if err != nil {
		http.Error(w, "Could not locate the file to serve: "+err.Error(), http.StatusNotFound)
		return
	}

	etg := etag.Generate(string(content), false)
	if r.Header.Get("If-None-Match") == etg {
		w.WriteHeader(http.StatusNotModified)
		return
	}

	w.Header().Set("ETag", etg)
	http.ServeContent(w, r, r.URL.Path, time.Unix(0, 0), bytes.NewReader(content))
}
