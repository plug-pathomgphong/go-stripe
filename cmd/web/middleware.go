***REMOVED***

import "net/http"

func SessionLoad(next http.Handler***REMOVED*** http.Handler {
	return session.LoadAndSave(next***REMOVED***
***REMOVED***

func (app *application***REMOVED*** Auth(next http.Handler***REMOVED*** http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request***REMOVED*** {
		if !app.Session.Exists(r.Context(***REMOVED***, "userID"***REMOVED*** {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect***REMOVED***
			return
	***REMOVED***
		next.ServeHTTP(w, r***REMOVED***
***REMOVED******REMOVED***
***REMOVED***
