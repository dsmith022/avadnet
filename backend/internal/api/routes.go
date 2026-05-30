package api

import (
	"bytes"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/healthz", app.healthcheckHandler)

	mux.HandleFunc("POST /api/v1/auth/register", app.registerUserHandler)
	mux.HandleFunc("POST /api/v1/auth/login", app.loginUserHandler)

	mux.HandleFunc("POST /api/v1/missions", app.createMissionOrganizationHandler)
	mux.HandleFunc("GET /api/v1/missions", app.listMissionOrganizationsHandler)
	mux.HandleFunc("GET /api/v1/missions/{id}", app.getMissionOrganizationHandler)
	mux.HandleFunc("PUT /api/v1/missions/{id}", app.updateMissionOrganizationHandler)
	mux.HandleFunc("DELETE /api/v1/missions/{id}", app.deleteMissionOrganizationHandler)
	mux.HandleFunc("POST /api/v1/missions/{id}/members", app.addMissionMemberHandler)
	mux.HandleFunc("GET /api/v1/missions/{id}/members", app.listMissionMembersHandler)
	mux.HandleFunc("DELETE /api/v1/missions/{id}/members/{user_id}", app.removeMissionMemberHandler)
	mux.HandleFunc("POST /api/v1/missions/{id}/ventures", app.addMissionVentureHandler)
	mux.HandleFunc("GET /api/v1/missions/{id}/ventures", app.listMissionVenturesHandler)
	mux.HandleFunc("DELETE /api/v1/missions/{id}/ventures/{venture_id}", app.removeMissionVentureHandler)

	mux.HandleFunc("POST /api/v1/churches", app.createChurchHandler)
	mux.HandleFunc("GET /api/v1/churches", app.listChurchesHandler)
	mux.HandleFunc("GET /api/v1/churches/{id}", app.getChurchHandler)
	mux.HandleFunc("PUT /api/v1/churches/{id}", app.updateChurchHandler)
	mux.HandleFunc("DELETE /api/v1/churches/{id}", app.deleteChurchHandler)
	mux.HandleFunc("POST /api/v1/churches/{id}/members", app.addChurchMemberHandler)
	mux.HandleFunc("GET /api/v1/churches/{id}/members", app.listChurchMembersHandler)
	mux.HandleFunc("PUT /api/v1/churches/{id}/members/{user_id}", app.updateChurchMemberRoleHandler)
	mux.HandleFunc("DELETE /api/v1/churches/{id}/members/{user_id}", app.removeChurchMemberHandler)

	return app.recoverPanic(app.securityHeaders(app.routeErrors(mux)))
}

func (app *application) routeErrors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		recorder := newResponseRecorder()
		next.ServeHTTP(recorder, r)

		switch recorder.statusCode {
		case http.StatusNotFound:
			app.notFoundResponse(w, r)
		case http.StatusMethodNotAllowed:
			for key, values := range recorder.header {
				if key != "Content-Type" {
					for _, value := range values {
						w.Header().Add(key, value)
					}
				}
			}
			app.methodNotAllowedResponse(w, r)
		default:
			recorder.writeTo(w)
		}
	})
}

type responseRecorder struct {
	header     http.Header
	body       bytes.Buffer
	statusCode int
}

func newResponseRecorder() *responseRecorder {
	return &responseRecorder{header: make(http.Header)}
}

func (rw *responseRecorder) Header() http.Header {
	return rw.header
}

func (rw *responseRecorder) Write(data []byte) (int, error) {
	if rw.statusCode == 0 {
		rw.statusCode = http.StatusOK
	}

	return rw.body.Write(data)
}

func (rw *responseRecorder) WriteHeader(statusCode int) {
	if rw.statusCode != 0 {
		return
	}

	rw.statusCode = statusCode
}

func (rw *responseRecorder) writeTo(w http.ResponseWriter) {
	statusCode := rw.statusCode
	if statusCode == 0 {
		statusCode = http.StatusOK
	}

	for key, values := range rw.header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.WriteHeader(statusCode)
	_, _ = w.Write(rw.body.Bytes())
}
