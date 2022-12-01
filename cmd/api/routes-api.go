package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		// ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	mux.Get("/api/", app.GetStartPage)
	mux.Post("/api/payment-intent", app.GetPaymentIntent)
	mux.Get("/api/widget/{id}", app.GetWidgetByID)

	mux.Post("/api/create-customer-and-subscribe-to-plan", app.CreateCustomerAndSubscribeToPlan)

	mux.Post("/api/authenticate", app.CreateAuthToken)
	mux.Post("/api/is-authenticated", app.CheckAuthentication)
	mux.Post("/api/forgot-password", app.SendPasswordResetEmail)
	mux.Post("/api/reset-password", app.ResetPassword)

	mux.Route("/api/admin", func(r chi.Router) {
		r.Use(app.Auth)

		r.Post("/virtual-terminal-succeeded", app.VirtualTerminalPaymentSucceeded)
		r.Post("/all-sales", app.AllSales)
		r.Post("/all-subscriptions", app.AllSubscriptions)

		r.Post("/get-sale/{id}", app.GetSale)

		r.Post("/refund", app.RefundCharge)
		r.Post("/cancel-subscription", app.CancelSubscription)

		r.Post("/all-users", app.AllUsers)
		r.Get("/all-users/{id}", app.OneUser)
		r.Post("/all-users/edit/{id}", app.EditUser)
		r.Get("/all-users/delete/{id}", app.DeleteUser)
		r.Get("/all-users/remove/{id}", app.RemoveUser)
	})

	return mux
}
