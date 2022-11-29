***REMOVED***

***REMOVED***
***REMOVED***

	"github.com/go-chi/chi/v5"
***REMOVED***

func (app *application***REMOVED*** routes(***REMOVED*** http.Handler {
	mux := chi.NewRouter(***REMOVED***
	mux.Use(SessionLoad***REMOVED***

	mux.Get("/", app.Home***REMOVED***

	mux.Route("/admin", func(r chi.Router***REMOVED*** {
		r.Use(app.Auth***REMOVED***
		r.Get("/virtual-terminal", app.VirtualTerminal***REMOVED***
***REMOVED******REMOVED***

	// mux.Post("/virtual-terminal-payment-succeeded", app.VirtualTerminalPaymentSucceeded***REMOVED***
	// mux.Get("/virtual-terminal-receipt", app.VirtualTerminalReceipt***REMOVED***

	mux.Get("/widget/{id***REMOVED***", app.ChangeOnce***REMOVED***
	mux.Post("/payment-succeeded", app.PaymentSucceeded***REMOVED***
	mux.Get("/receipt", app.Receipt***REMOVED***

	mux.Get("/plan/bronze", app.BronzePlan***REMOVED***
	mux.Get("/receipt/bronze", app.BronzePlanReceipt***REMOVED***

	// auth route
	mux.Get("/login", app.LoginPage***REMOVED***
	mux.Post("/login", app.PostLoginPage***REMOVED***
	mux.Get("/logout", app.Logout***REMOVED***
	mux.Get("/forgot-password", app.ForgotPassword***REMOVED***
	mux.Get("/reset-password", app.ShowResetPassword***REMOVED***

	fileServer := http.FileServer(http.Dir("./static"***REMOVED******REMOVED***
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer***REMOVED******REMOVED***

	return mux
***REMOVED***
