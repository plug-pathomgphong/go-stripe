***REMOVED***

import "net/http"

func (app *application***REMOVED*** VirtualTerminal(w http.ResponseWriter, r *http.Request***REMOVED*** {
	stringMap := make(map[string]string***REMOVED***
	stringMap["publishable_key"] = app.config.stripe.key
	if err := app.renderTemplate(w, r, "terminal", &templateData{StringMap: stringMap***REMOVED******REMOVED***; err != nil {
		app.errorLog.Println(err***REMOVED***
***REMOVED***
***REMOVED***

func (app *application***REMOVED*** PaymentSucceeded(w http.ResponseWriter, r *http.Request***REMOVED*** {
	err := r.ParseForm(***REMOVED***
***REMOVED***
		app.errorLog.Println(err***REMOVED***
		return
***REMOVED***

	// read posted data
	cardHolder := r.Form.Get("cardholder_name"***REMOVED***
	email := r.Form.Get("email"***REMOVED***
	paymentIndent := r.Form.Get("payment_intent"***REMOVED***
	paymentMethod := r.Form.Get("payment_method"***REMOVED***
	paymentAmount := r.Form.Get("payment_amount"***REMOVED***
	paymentCurrency := r.Form.Get("payment_currency"***REMOVED***

	data := make(map[string]interface{***REMOVED******REMOVED***
	data["cardholder"] = cardHolder
	data["email"] = email
	data["pi"] = paymentIndent
	data["pm"] = paymentMethod
	data["pa"] = paymentAmount
	data["pc"] = paymentCurrency

	if err := app.renderTemplate(w, r, "succeeded", &templateData{Data: data***REMOVED******REMOVED***; err != nil {
		app.errorLog.Println(err***REMOVED***
***REMOVED***
***REMOVED***
