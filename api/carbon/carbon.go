package carbon

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/saigees/carbon/api"
	"github.com/saigees/carbon/internal"
)

func Main() {
	router := mux.NewRouter()

	api.GetRoutes(router)

	addr := fmt.Sprintf("0.0.0.0:%s", internal.Port)
	internal.Logger.Infof("Starting server - %s", addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		internal.Logger.Error("Failed to start http server.")
	}
}
