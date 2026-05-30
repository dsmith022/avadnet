package api

import "net/http"

func (app *application) createChurchHandler(w http.ResponseWriter, r *http.Request) {
	app.notImplementedResponse(w, r, "create church not implemented")
}

func (app *application) listChurchesHandler(w http.ResponseWriter, r *http.Request) {
	app.notImplementedResponse(w, r, "list churches not implemented")
}

func (app *application) getChurchHandler(w http.ResponseWriter, r *http.Request) {
	app.notImplementedResponse(w, r, "get church not implemented")
}

func (app *application) updateChurchHandler(w http.ResponseWriter, r *http.Request) {
	app.notImplementedResponse(w, r, "update church not implemented")
}

func (app *application) deleteChurchHandler(w http.ResponseWriter, r *http.Request) {
	app.notImplementedResponse(w, r, "delete church not implemented")
}

func (app *application) addChurchMemberHandler(w http.ResponseWriter, r *http.Request) {
	app.notImplementedResponse(w, r, "add church member not implemented")
}

func (app *application) listChurchMembersHandler(w http.ResponseWriter, r *http.Request) {
	app.notImplementedResponse(w, r, "list church members not implemented")
}

func (app *application) updateChurchMemberRoleHandler(w http.ResponseWriter, r *http.Request) {
	app.notImplementedResponse(w, r, "update church member role not implemented")
}

func (app *application) removeChurchMemberHandler(w http.ResponseWriter, r *http.Request) {
	app.notImplementedResponse(w, r, "remove church member not implemented")
}
