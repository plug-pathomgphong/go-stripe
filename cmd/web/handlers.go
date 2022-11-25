package main

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/plug-pathomgphong/dotnet-webapi/internal/cards"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	// stringMap := make(map[string]string)
	// stringMap["publishable_key"] = app.config.stripe.key
	if err := app.renderTemplate(w, r, "home", &templateData{}); err != nil {
		app.errorLog.Println(err)
	}
}

func (app *application) VirtualTerminal(w http.ResponseWriter, r *http.Request) {
	// stringMap := make(map[string]string)
	// stringMap["publishable_key"] = app.config.stripe.key
	if err := app.renderTemplate(w, r, "terminal", &templateData{}, "stripe-js"); err != nil {
		app.errorLog.Println(err)
	}
}

func (app *application) PaymentSucceeded(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	// read posted data
	cardHolder := r.Form.Get("cardholder_name")
	email := r.Form.Get("email")
	paymentIndent := r.Form.Get("payment_intent")
	paymentMethod := r.Form.Get("payment_method")
	paymentAmount := r.Form.Get("payment_amount")
	paymentCurrency := r.Form.Get("payment_currency")

	card := cards.Card{
		Secret: app.config.stripe.secret,
		Key:    app.config.stripe.key,
	}

	pi, err := card.RetrievePaymentIntent(paymentIndent)
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	pm, err := card.GetPaymentMethod(paymentMethod)
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	lastFour := pm.Card.Last4
	expiryMonth := pm.Card.ExpMonth
	expireYear := pm.Card.ExpYear

	data := make(map[string]interface{})
	data["cardholder"] = cardHolder
	data["email"] = email
	data["pi"] = paymentIndent
	data["pm"] = paymentMethod
	data["pa"] = paymentAmount
	data["pc"] = paymentCurrency
	data["last_four"] = lastFour
	data["expire_month"] = expiryMonth
	data["expire_year"] = expireYear
	data["bank_return_code"] = pi.LatestCharge.ID

	if err := app.renderTemplate(w, r, "succeeded", &templateData{Data: data}); err != nil {
		app.errorLog.Println(err)
	}
}

func (app *application) ChangeOnce(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	widgetID, _ := strconv.Atoi(id)

	widget, err := app.DB.GetWidget(widgetID)
	if err != nil {
		app.errorLog.Panicln(err)
		return
	}

	data := make(map[string]interface{})
	data["widget"] = widget
	if err := app.renderTemplate(w, r, "buy-once", &templateData{Data: data}, "stripe-js"); err != nil {
		app.errorLog.Println(err)
	}
}
