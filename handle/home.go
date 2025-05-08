package handle

import (
	"net/http"

	"github.com/nanvenomous/htmx_templ_daisy_starter/models"
	"github.com/nanvenomous/htmx_templ_daisy_starter/ui"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	var (
		err     error
		stts    models.HttpStatus
		partial = r.URL.Query().Get("partial")
	)

	comp := ui.PageHome()
	if partial != "" {
		comp = ui.Home()
	}

	stts, err = render(w, r,
		comp,
	)
	if err != nil {
		errorOut(w, stts, err)
	}
}

func GetTable(w http.ResponseWriter, r *http.Request) {
	var (
		err     error
		stts    models.HttpStatus
		partial = r.URL.Query().Get("partial")
	)

	comp := ui.PageTable()
	if partial != "" {
		comp = ui.Table()
	}

	stts, err = render(w, r,
		comp,
	)
	if err != nil {
		errorOut(w, stts, err)
	}
}
