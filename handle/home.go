package handle

import (
	"net/http"

	"github.com/nanvenomous/htmx_templ_daisy_starter/ui"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	stts, err := render(w, r, ui.Home())
	if err != nil {
		errorOut(w, stts, err)
		return
	}
}
