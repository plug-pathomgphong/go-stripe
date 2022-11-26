***REMOVED***

***REMOVED***
	"encoding/json"
***REMOVED***
	"strconv"
***REMOVED***

	"github.com/go-chi/chi/v5"
	"github.com/plug-pathomgphong/dotnet-webapi/internal/cards"
***REMOVED***
	"github.com/stripe/stripe-go/v72"
***REMOVED***

type stripePayload struct {
	Currency      string `json:"currency"`
	Amount        string `json:"amount"`
	PaymentMethod string `json:"payment_method"`
	Email         string `json:"email"`
	CardBrand     string `json:"card_brand"`
	ExpiryMonth   int    `json:"exp_month"`
	ExpiryYear    int    `json:"exp_year"`
	LastFour      string `json:"last_four"`
	Plan          string `json:"plan"`
	ProductID     string `json:"product_id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
***REMOVED***

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message,omitempty"`
	Content string `json:"content,omitempty"`
	ID      int    `json:"id,omitempty"`
***REMOVED***

func (app *application***REMOVED*** GetStartPage(w http.ResponseWriter, r *http.Request***REMOVED*** {
	myJsonString := `{"name":"hello world"***REMOVED***`

	out, err := json.MarshalIndent(myJsonString, "", "   "***REMOVED***
***REMOVED***
		app.errorLog.Println(err***REMOVED***
		return
***REMOVED***
	w.Header(***REMOVED***.Set("Content-Type", "application/json"***REMOVED***
	w.Write(out***REMOVED***
***REMOVED***

func (app *application***REMOVED*** GetPaymentIntent(w http.ResponseWriter, r *http.Request***REMOVED*** {
	var payload stripePayload

	err := json.NewDecoder(r.Body***REMOVED***.Decode(&payload***REMOVED***
***REMOVED***
		app.errorLog.Println(r.Method***REMOVED***
		app.errorLog.Println(err***REMOVED***
		return
***REMOVED***

	amount, err := strconv.Atoi(payload.Amount***REMOVED***
***REMOVED***
		app.errorLog.Println(err***REMOVED***
		return
***REMOVED***

	card := cards.Card{
		Secret:   app.config.stripe.secret,
		Key:      app.config.stripe.key,
		Currency: payload.Currency,
***REMOVED***

	okay := true

	pi, msg, err := card.Charge(payload.Currency, amount***REMOVED***
***REMOVED***
		okay = false
***REMOVED***

	if okay {
		out, err := json.MarshalIndent(pi, "", "   "***REMOVED***
	***REMOVED***
			app.errorLog.Println(err***REMOVED***
			return
	***REMOVED***

		w.Header(***REMOVED***.Set("Content-Type", "application/json"***REMOVED***
		w.Write(out***REMOVED***
***REMOVED*** else {
		j := jsonResponse{
			OK:      false,
			Message: msg,
			Content: "",
	***REMOVED***

		out, err := json.MarshalIndent(j, "", "   "***REMOVED***
	***REMOVED***
			app.errorLog.Panicln(err***REMOVED***
	***REMOVED***

		w.Header(***REMOVED***.Set("Content-Type", "application/json"***REMOVED***
		w.Write(out***REMOVED***
***REMOVED***

***REMOVED***

func (app *application***REMOVED*** GetWidgetByID(w http.ResponseWriter, r *http.Request***REMOVED*** {
	id := chi.URLParam(r, "id"***REMOVED***
	widgetID, _ := strconv.Atoi(id***REMOVED***

	widget, err := app.DB.GetWidget(widgetID***REMOVED***
***REMOVED***
		app.errorLog.Panicln(err***REMOVED***
		return
***REMOVED***

	out, err := json.MarshalIndent(widget, "", "   "***REMOVED***
***REMOVED***
		app.errorLog.Panicln(err***REMOVED***
		return
***REMOVED***

	w.Header(***REMOVED***.Set("Content-Type", "application/json"***REMOVED***
	w.Write(out***REMOVED***
***REMOVED***

func (app *application***REMOVED*** CreateCustomerAndSubscribeToPlan(w http.ResponseWriter, r *http.Request***REMOVED*** {
	var data stripePayload
	err := json.NewDecoder(r.Body***REMOVED***.Decode(&data***REMOVED***
***REMOVED***
		app.errorLog.Panicln(err***REMOVED***
		return
***REMOVED***

	app.infoLog.Println(data.Email, data.LastFour, data.PaymentMethod, data.Plan***REMOVED***

	card := cards.Card{
		Secret:   app.config.stripe.secret,
		Key:      app.config.stripe.key,
		Currency: data.Currency,
***REMOVED***

	okey := true
	var subscription *stripe.Subscription
	txnMsg := "Transaction successful"

	stripeCustomer, msg, err := card.CreateCustomer(data.PaymentMethod, data.Email***REMOVED***
***REMOVED***
		app.errorLog.Panicln(err***REMOVED***
		okey = false
		txnMsg = msg
***REMOVED***

	if okey {
		subscription, err = card.SubscribeToPlan(stripeCustomer, data.Plan, data.Email, data.LastFour, ""***REMOVED***
	***REMOVED***
			app.errorLog.Panicln(err***REMOVED***
			okey = false
			txnMsg = "Error subscribing customer"
	***REMOVED***
		app.infoLog.Println("subscription id is", subscription.ID***REMOVED***
***REMOVED***

	if okey {
		productID, _ := strconv.Atoi(data.ProductID***REMOVED***
		customerID, err := app.SaveCustomer(data.FirstName, data.LastName, data.Email***REMOVED***
	***REMOVED***
			app.errorLog.Panicln(err***REMOVED***
			return
	***REMOVED***

		// create a new txn
		amount, _ := strconv.Atoi(data.Amount***REMOVED***
		// expiryMonth, _ := strconv.Atoi(data.ExpiryMonth***REMOVED***
		// expiryYear, _ := strconv.Atoi(data.ExpiryYear***REMOVED***
		txn := models.Transaction{
			Amount:              amount,
			Currency:            "cad",
			LastFour:            data.LastFour,
			ExpiryMonth:         data.ExpiryMonth,
			ExpiryYear:          data.ExpiryYear,
			TransactionStatusID: 2,
	***REMOVED***

		txnID, err := app.SaveTransaction(txn***REMOVED***
	***REMOVED***
			app.errorLog.Panicln(err***REMOVED***
			return
	***REMOVED***

		// create order
		order := models.Order{
			WidgetID:      productID,
			TransactionID: txnID,
			CustomerID:    customerID,
			StatusID:      1,
			Quantity:      1,
			Amount:        amount,
			CreatedAt:     time.Now(***REMOVED***,
			UpdatedAt:     time.Now(***REMOVED***,
	***REMOVED***

		_, err = app.SaveOrder(order***REMOVED***
	***REMOVED***
			app.errorLog.Panicln(err***REMOVED***
			return
	***REMOVED***
***REMOVED***
	// msg := ""

	resp := jsonResponse{
		OK:      okey,
		Message: txnMsg,
***REMOVED***

	out, err := json.MarshalIndent(resp, "", "   "***REMOVED***
***REMOVED***
		app.errorLog.Panicln(err***REMOVED***
		return
***REMOVED***

	w.Header(***REMOVED***.Set("Content-Type", "application/json"***REMOVED***
	w.Write(out***REMOVED***

***REMOVED***

// SaveCustomer saves a customer and retruns id
func (app *application***REMOVED*** SaveCustomer(firstName, lastName, email string***REMOVED*** (int, error***REMOVED*** {
	customer := models.Customer{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
***REMOVED***

	id, err := app.DB.InsertCustomer(customer***REMOVED***
***REMOVED***
		return 0, err
***REMOVED***
	return id, nil
***REMOVED***

// SaveTransaction saves a transaction and retruns id
func (app *application***REMOVED*** SaveTransaction(txn models.Transaction***REMOVED*** (int, error***REMOVED*** {
	id, err := app.DB.InsertTransaction(txn***REMOVED***
***REMOVED***
		return 0, err
***REMOVED***
	return id, nil
***REMOVED***

// SaveOrder saves a order and retruns id
func (app *application***REMOVED*** SaveOrder(order models.Order***REMOVED*** (int, error***REMOVED*** {
	id, err := app.DB.InsertOrder(order***REMOVED***
***REMOVED***
		return 0, err
***REMOVED***
	return id, nil
***REMOVED***
