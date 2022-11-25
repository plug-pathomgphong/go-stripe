***REMOVED***

***REMOVED***
***REMOVED***
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/plug-pathomgphong/dotnet-webapi/internal/cards"
***REMOVED***

func (app *application***REMOVED*** Home(w http.ResponseWriter, r *http.Request***REMOVED*** {
	// stringMap := make(map[string]string***REMOVED***
	// stringMap["publishable_key"] = app.config.stripe.key
	if err := app.renderTemplate(w, r, "home", &templateData{***REMOVED******REMOVED***; err != nil {
		app.errorLog.Println(err***REMOVED***
***REMOVED***
***REMOVED***

func (app *application***REMOVED*** VirtualTerminal(w http.ResponseWriter, r *http.Request***REMOVED*** {
	// stringMap := make(map[string]string***REMOVED***
	// stringMap["publishable_key"] = app.config.stripe.key
	if err := app.renderTemplate(w, r, "terminal", &templateData{***REMOVED***, "stripe-js"***REMOVED***; err != nil {
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

	card := cards.Card{
		Secret: app.config.stripe.secret,
		Key:    app.config.stripe.key,
***REMOVED***

	pi, err := card.RetrievePaymentIntent(paymentIndent***REMOVED***
***REMOVED***
		app.errorLog.Println(err***REMOVED***
		return
***REMOVED***

	pm, err := card.GetPaymentMethod(paymentMethod***REMOVED***
***REMOVED***
		app.errorLog.Println(err***REMOVED***
		return
***REMOVED***

	lastFour := pm.Card.Last4
	expiryMonth := pm.Card.ExpMonth
	expireYear := pm.Card.ExpYear

	data := make(map[string]interface{***REMOVED******REMOVED***
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

	if err := app.renderTemplate(w, r, "succeeded", &templateData{Data: data***REMOVED******REMOVED***; err != nil {
		app.errorLog.Println(err***REMOVED***
***REMOVED***
***REMOVED***

func (app *application***REMOVED*** ChangeOnce(w http.ResponseWriter, r *http.Request***REMOVED*** {
	id := chi.URLParam(r, "id"***REMOVED***
	widgetID, _ := strconv.Atoi(id***REMOVED***

	widget, err := app.DB.GetWidget(widgetID***REMOVED***
***REMOVED***
		app.errorLog.Panicln(err***REMOVED***
		return
***REMOVED***

	data := make(map[string]interface{***REMOVED******REMOVED***
	data["widget"] = widget
	if err := app.renderTemplate(w, r, "buy-once", &templateData{Data: data***REMOVED***, "stripe-js"***REMOVED***; err != nil {
		app.errorLog.Println(err***REMOVED***
***REMOVED***
***REMOVED***
