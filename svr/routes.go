package svr

import (
	"net/http"

	"github.com/pisgahi/xcerpted-admin/hndl"
)

func (s *Svr) MountHdlrs() {
	s.Router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	s.Router.Get("/", hndl.HomePageHdlr)
}
