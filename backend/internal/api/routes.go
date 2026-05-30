package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.HandleMethodNotAllowed = true
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFoundResponse(w, r)
	})
	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.methodNotAllowedResponse(w, r)
	})

	router.HandlerFunc(http.MethodGet, "/api/v1/healthz", app.healthcheckHandler)

	router.HandlerFunc(http.MethodPost, "/api/v1/auth/register", app.registerUserHandler)
	router.HandlerFunc(http.MethodPost, "/api/v1/auth/login", app.loginUserHandler)

	router.HandlerFunc(http.MethodPost, "/api/v1/missions", app.createMissionOrganizationHandler)
	router.HandlerFunc(http.MethodGet, "/api/v1/missions", app.listMissionOrganizationsHandler)
	router.HandlerFunc(http.MethodGet, "/api/v1/missions/:id", app.getMissionOrganizationHandler)
	router.HandlerFunc(http.MethodPut, "/api/v1/missions/:id", app.updateMissionOrganizationHandler)
	router.HandlerFunc(http.MethodDelete, "/api/v1/missions/:id", app.deleteMissionOrganizationHandler)
	router.HandlerFunc(http.MethodPost, "/api/v1/missions/:id/members", app.addMissionMemberHandler)
	router.HandlerFunc(http.MethodGet, "/api/v1/missions/:id/members", app.listMissionMembersHandler)
	router.HandlerFunc(http.MethodDelete, "/api/v1/missions/:id/members/:user_id", app.removeMissionMemberHandler)
	router.HandlerFunc(http.MethodPost, "/api/v1/missions/:id/ventures", app.addMissionVentureHandler)
	router.HandlerFunc(http.MethodGet, "/api/v1/missions/:id/ventures", app.listMissionVenturesHandler)
	router.HandlerFunc(http.MethodDelete, "/api/v1/missions/:id/ventures/:venture_id", app.removeMissionVentureHandler)

	router.HandlerFunc(http.MethodPost, "/api/v1/churches", app.createChurchHandler)
	router.HandlerFunc(http.MethodGet, "/api/v1/churches", app.listChurchesHandler)
	router.HandlerFunc(http.MethodGet, "/api/v1/churches/:id", app.getChurchHandler)
	router.HandlerFunc(http.MethodPut, "/api/v1/churches/:id", app.updateChurchHandler)
	router.HandlerFunc(http.MethodDelete, "/api/v1/churches/:id", app.deleteChurchHandler)
	router.HandlerFunc(http.MethodPost, "/api/v1/churches/:id/members", app.addChurchMemberHandler)
	router.HandlerFunc(http.MethodGet, "/api/v1/churches/:id/members", app.listChurchMembersHandler)
	router.HandlerFunc(http.MethodPut, "/api/v1/churches/:id/members/:user_id", app.updateChurchMemberRoleHandler)
	router.HandlerFunc(http.MethodDelete, "/api/v1/churches/:id/members/:user_id", app.removeChurchMemberHandler)

	return app.recoverPanic(app.securityHeaders(router))
}

func (app *application) readParam(r *http.Request, name string) string {
	params := httprouter.ParamsFromContext(r.Context())
	return params.ByName(name)
}
