***REMOVED***

***REMOVED***
***REMOVED***

	"github.com/go-chi/chi/v5"
***REMOVED***

func (app *application***REMOVED*** routes(***REMOVED*** http.Handler {
	mux := chi.NewRouter(***REMOVED***

	mux.Get("/virture-terminal", app.VirtualTerminal***REMOVED***
	mux.Post("/payment-succeeded", app.PaymentSucceeded***REMOVED***
	return mux
***REMOVED***
