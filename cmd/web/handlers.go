***REMOVED***

***REMOVED***
***REMOVED***
	"strconv"
***REMOVED***

	"github.com/go-chi/chi/v5"
	"github.com/plug-pathomgphong/dotnet-webapi/internal/cards"
***REMOVED***
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

type TransactionData struct {
	FirstName       string
	LastName        string
	Email           string
	PaymentIndentID string
	PaymentMethodID string
	PaymentAmount   int
	PaymentCurrency string
	LastFour        string
	ExpiryMonth     int
	ExpiryYear      int
	BankReturnCode  string
***REMOVED***

// GetTransactionData gets txn data from post and stripe
func (app *application***REMOVED*** GetTransactionData(r *http.Request***REMOVED*** (TransactionData, error***REMOVED*** {
	var tnxData TransactionData
	err := r.ParseForm(***REMOVED***
***REMOVED***
		app.errorLog.Println(err***REMOVED***
		return tnxData, err
***REMOVED***

	// read posted data
	firstName := r.Form.Get("first_name"***REMOVED***
	lastName := r.Form.Get("last_name"***REMOVED***
	email := r.Form.Get("email"***REMOVED***
	paymentIndent := r.Form.Get("payment_intent"***REMOVED***
	paymentMethod := r.Form.Get("payment_method"***REMOVED***
	paymentAmount := r.Form.Get("payment_amount"***REMOVED***
	paymentCurrency := r.Form.Get("payment_currency"***REMOVED***
	amount, _ := strconv.Atoi(paymentAmount***REMOVED***

	card := cards.Card{
		Secret: app.config.stripe.secret,
		Key:    app.config.stripe.key,
***REMOVED***

	pi, err := card.RetrievePaymentIntent(paymentIndent***REMOVED***
***REMOVED***
		app.errorLog.Println(err***REMOVED***
		return tnxData, err
***REMOVED***

	pm, err := card.GetPaymentMethod(paymentMethod***REMOVED***
***REMOVED***
		app.errorLog.Println(err***REMOVED***
		return tnxData, err
***REMOVED***

	lastFour := pm.Card.Last4
	expiryMonth := pm.Card.ExpMonth
	expireYear := pm.Card.ExpYear

	txnData := TransactionData{
		FirstName:       firstName,
		LastName:        lastName,
		Email:           email,
		PaymentIndentID: paymentIndent,
		PaymentMethodID: paymentMethod,
		PaymentAmount:   amount,
		PaymentCurrency: paymentCurrency,
		LastFour:        lastFour,
		ExpiryMonth:     int(expiryMonth***REMOVED***,
		ExpiryYear:      int(expireYear***REMOVED***,
		BankReturnCode:  pi.Charges.Data[0].ID,
***REMOVED***

	return txnData, nil

***REMOVED***

// PaymentSucceeded displays the receipt page
func (app *application***REMOVED*** PaymentSucceeded(w http.ResponseWriter, r *http.Request***REMOVED*** {
	err := r.ParseForm(***REMOVED***
***REMOVED***
		app.errorLog.Println(err***REMOVED***
		return
***REMOVED***

	// read posted data
	widgetID, _ := strconv.Atoi(r.Form.Get("product_id"***REMOVED******REMOVED***
	txnData, err := app.GetTransactionData(r***REMOVED***
***REMOVED***
		app.errorLog.Println(err***REMOVED***
		return
***REMOVED***

	// create a new customer
	customerID, err := app.SaveCustomer(txnData.FirstName, txnData.LastName, txnData.Email***REMOVED***
***REMOVED***
		app.errorLog.Println(err***REMOVED***
		return
***REMOVED***

	// create a new transaction
	txn := models.Transaction{
		Amount:              txnData.PaymentAmount,
		Currency:            txnData.PaymentCurrency,
		LastFour:            txnData.LastFour,
		ExpiryMonth:         txnData.ExpiryMonth,
		ExpiryYear:          txnData.ExpiryYear,
		BankReturnCode:      txnData.BankReturnCode,
		PaymentIndent:       txnData.PaymentIndentID,
		PaymentMethod:       txnData.PaymentMethodID,
		TransactionStatusID: 2,
***REMOVED***

	txnID, err := app.SaveTransaction(txn***REMOVED***
***REMOVED***
		app.errorLog.Println(err***REMOVED***
		return
***REMOVED***

	// create a new order
	order := models.Order{
		WidgetID:      widgetID,
		TransactionID: txnID,
		CustomerID:    customerID,
		StatusID:      1,
		Quantity:      1,
		Amount:        txnData.PaymentAmount,
		CreatedAt:     time.Now(***REMOVED***,
		UpdatedAt:     time.Now(***REMOVED***,
***REMOVED***

	_, err = app.SaveOrder(order***REMOVED***
***REMOVED***
		app.errorLog.Println(err***REMOVED***
		return
***REMOVED***

	// data := make(map[string]interface{***REMOVED******REMOVED***
	// data["cardholder"] = cardHolder
	// data["email"] = txnData.Email
	// data["pi"] = txnData.PaymentIndentID
	// data["pm"] = txnData.PaymentMethodID
	// data["pa"] = txnData.PaymentAmount
	// data["pc"] = txnData.PaymentCurrency
	// data["last_four"] = txnData.LastFour
	// data["expire_month"] = txnData.ExpiryMonth
	// data["expire_year"] = txnData.ExpiryYear
	// data["bank_return_code"] = txnData.BankReturnCode
	// data["first_name"] = txnData.FirstName
	// data["last_name"] = txnData.LastName

	// should write this data to session, and then redirect user to new page
	app.Session.Put(r.Context(***REMOVED***, "receipt", txnData***REMOVED***
	http.Redirect(w, r, "/receipt", http.StatusSeeOther***REMOVED***

	// if err := app.renderTemplate(w, r, "succeeded", &templateData{Data: data***REMOVED******REMOVED***; err != nil {
	// 	app.errorLog.Println(err***REMOVED***
	// ***REMOVED***
***REMOVED***

// VirtualTerminalPaymentSucceeded displays the receipt page
func (app *application***REMOVED*** VirtualTerminalPaymentSucceeded(w http.ResponseWriter, r *http.Request***REMOVED*** {

	txnData, err := app.GetTransactionData(r***REMOVED***
***REMOVED***
		app.errorLog.Println(err***REMOVED***
		return
***REMOVED***

	// create a new transaction
	txn := models.Transaction{
		Amount:              txnData.PaymentAmount,
		Currency:            txnData.PaymentCurrency,
		LastFour:            txnData.LastFour,
		ExpiryMonth:         txnData.ExpiryMonth,
		ExpiryYear:          txnData.ExpiryYear,
		BankReturnCode:      txnData.BankReturnCode,
		PaymentIndent:       txnData.PaymentIndentID,
		PaymentMethod:       txnData.PaymentMethodID,
		TransactionStatusID: 2,
***REMOVED***

	_, err = app.SaveTransaction(txn***REMOVED***
***REMOVED***
		app.errorLog.Println(err***REMOVED***
		return
***REMOVED***

	app.Session.Put(r.Context(***REMOVED***, "receipt", txnData***REMOVED***
	http.Redirect(w, r, "/virtual-terminal-receipt", http.StatusSeeOther***REMOVED***

***REMOVED***

func (app *application***REMOVED*** Receipt(w http.ResponseWriter, r *http.Request***REMOVED*** {
	txn := app.Session.Get(r.Context(***REMOVED***, "receipt"***REMOVED***.(TransactionData***REMOVED***
	data := make(map[string]interface{***REMOVED******REMOVED***
	data["txn"] = txn

	app.Session.Remove(r.Context(***REMOVED***, "receipt"***REMOVED***
	if err := app.renderTemplate(w, r, "receipt", &templateData{
		Data: data,
***REMOVED******REMOVED***; err != nil {
		app.errorLog.Println(err***REMOVED***
***REMOVED***
***REMOVED***

func (app *application***REMOVED*** VirtualTerminalReceipt(w http.ResponseWriter, r *http.Request***REMOVED*** {
	txn := app.Session.Get(r.Context(***REMOVED***, "receipt"***REMOVED***.(TransactionData***REMOVED***
	data := make(map[string]interface{***REMOVED******REMOVED***
	data["txn"] = txn

	app.Session.Remove(r.Context(***REMOVED***, "receipt"***REMOVED***
	if err := app.renderTemplate(w, r, "virtual-terminal-receipt", &templateData{
		Data: data,
***REMOVED******REMOVED***; err != nil {
		app.errorLog.Println(err***REMOVED***
***REMOVED***
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

func (app *application***REMOVED*** BronzePlan(w http.ResponseWriter, r *http.Request***REMOVED*** {
	widget, err := app.DB.GetWidget(2***REMOVED***
***REMOVED***
		app.errorLog.Panicln(err***REMOVED***
		return
***REMOVED***
	data := make(map[string]interface{***REMOVED******REMOVED***
	data["widget"] = widget

	if err := app.renderTemplate(w, r, "bronze-plan", &templateData{
		Data: data,
***REMOVED******REMOVED***; err != nil {
		app.errorLog.Println(err***REMOVED***
***REMOVED***
***REMOVED***

func (app *application***REMOVED*** BronzePlanReceipt(w http.ResponseWriter, r *http.Request***REMOVED*** {

	if err := app.renderTemplate(w, r, "receipt-plan", &templateData{***REMOVED******REMOVED***; err != nil {
		app.errorLog.Println(err***REMOVED***
***REMOVED***
***REMOVED***
