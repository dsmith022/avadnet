package api

import "net/http"

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	err := app.writeJSON(w, http.StatusOK, envelope{
		"status":      "available",
		"environment": app.config.AppEnv,
	}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
