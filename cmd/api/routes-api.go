***REMOVED***

***REMOVED***
***REMOVED***

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
***REMOVED***

func (app *application***REMOVED*** routes(***REMOVED*** http.Handler {
	mux := chi.NewRouter(***REMOVED***

	mux.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"***REMOVED***, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"***REMOVED***,
		// AllowOriginFunc:  func(r *http.Request, origin string***REMOVED*** bool { return true ***REMOVED***,
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"***REMOVED***,
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"***REMOVED***,
		// ExposedHeaders:   []string{"Link"***REMOVED***,
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
***REMOVED******REMOVED******REMOVED***

	mux.Get("/api/", app.GetStartPage***REMOVED***
	mux.Post("/api/payment-intent", app.GetPaymentIntent***REMOVED***
	mux.Get("/api/widget/{id***REMOVED***", app.GetWidgetByID***REMOVED***

	mux.Post("/api/create-customer-and-subscribe-to-plan", app.CreateCustomerAndSubscribeToPlan***REMOVED***

	mux.Post("/api/authenticate", app.CreateAuthToken***REMOVED***
	mux.Post("/api/is-authenticated", app.CheckAuthentication***REMOVED***

	mux.Route("/api/admin", func(r chi.Router***REMOVED*** {
		r.Use(app.Auth***REMOVED***

		r.Post("/virtual-terminal-succeeded", app.VirtualTerminalPaymentSucceeded***REMOVED***
***REMOVED******REMOVED***

	return mux
***REMOVED***
