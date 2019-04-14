package httpserver

import (
	"context"
	"net/http"

	"github.com/yakaa/cuter/lib/system"
)

func StartServer(srv *http.Server) error {
	system.AddWrapUpListener(func() {
		srv.Shutdown(context.Background())
	})

	return srv.ListenAndServe()
}
