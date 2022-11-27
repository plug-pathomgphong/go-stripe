***REMOVED***

import "net/http"

func (app *application***REMOVED*** Auth(next http.Handler***REMOVED*** http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request***REMOVED*** {
		_, err := app.authenticateToken(r***REMOVED***
	***REMOVED***
			app.invalidCredentials(w***REMOVED***
			return
	***REMOVED***
		next.ServeHTTP(w, r***REMOVED***
***REMOVED******REMOVED***
***REMOVED***
