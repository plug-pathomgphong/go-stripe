***REMOVED***

***REMOVED***
***REMOVED***

	"github.com/go-chi/chi/v5"
***REMOVED***

func (app *application***REMOVED*** routes(***REMOVED*** http.Handler {
	mux := chi.NewRouter(***REMOVED***
	mux.Use(SessionLoad***REMOVED***

	mux.Get("/", app.Home***REMOVED***

	mux.Get("/virture-terminal", app.VirtualTerminal***REMOVED***
	mux.Post("/payment-succeeded", app.PaymentSucceeded***REMOVED***

	mux.Get("/widget/{id***REMOVED***", app.ChangeOnce***REMOVED***

	fileServer := http.FileServer(http.Dir("./static"***REMOVED******REMOVED***
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer***REMOVED******REMOVED***

	return mux
***REMOVED***
