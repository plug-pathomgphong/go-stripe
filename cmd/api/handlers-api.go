***REMOVED***

***REMOVED***
	"encoding/json"
	"errors"
***REMOVED***
***REMOVED***
	"strconv"
	"strings"
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

// CreateAuthToken
func (app *application***REMOVED*** CreateAuthToken(w http.ResponseWriter, r *http.Request***REMOVED*** {
	var userInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
***REMOVED***
	err := app.readJSON(w, r, &userInput***REMOVED***
***REMOVED***
		app.badRequest(w, r, err***REMOVED***
		return
***REMOVED***

	// get the user from the database by email; send error if invalid email
	user, err := app.DB.GetUserByEmail(userInput.Email***REMOVED***
***REMOVED***
		app.invalidCredentials(w***REMOVED***
		return
***REMOVED***

	// validate the password;send error if invalid password
	validPassword, err := app.passwordMatches(user.Password, userInput.Password***REMOVED***
***REMOVED***
		app.invalidCredentials(w***REMOVED***
		return
***REMOVED***

	if !validPassword {
		app.invalidCredentials(w***REMOVED***
		return
***REMOVED***

	// generate the token
	token, err := models.GenerateToken(user.ID, 24*time.Hour, models.ScopeAuthentication***REMOVED***
***REMOVED***
		app.badRequest(w, r, err***REMOVED***
		return
***REMOVED***

	// save to database
	err = app.DB.InsertToken(token, user***REMOVED***
***REMOVED***
		app.badRequest(w, r, err***REMOVED***
		return
***REMOVED***

	// send respone
	var payload struct {
		Error   bool          `json:"error"`
		Message string        `json:"message"`
		Token   *models.Token `json:"authentication_token"`
***REMOVED***

	payload.Error = false
	payload.Message = fmt.Sprintf("token for %s created", userInput.Email***REMOVED***
	payload.Token = token

	_ = app.writeJSON(w, http.StatusOK, payload***REMOVED***
***REMOVED***

func (app *application***REMOVED*** authenticateToken(r *http.Request***REMOVED*** (*models.User, error***REMOVED*** {
	var u models.User
	// fmt.Println("Authorization"***REMOVED***
	authenticationHeader := r.Header.Get("Authorization"***REMOVED***
	if authenticationHeader == "" {
		return nil, errors.New("no authentication header received"***REMOVED***
***REMOVED***
	// fmt.Println("headerParts", authenticationHeader***REMOVED***
	headerParts := strings.Split(authenticationHeader, " "***REMOVED***
	fmt.Println("headerParts", headerParts, headerParts[0]***REMOVED***
	if len(headerParts***REMOVED*** != 2 || headerParts[0] != "Bearer" {
		return nil, errors.New("no authentication header received"***REMOVED***
***REMOVED***
	// fmt.Println("token"***REMOVED***
	token := headerParts[1]
	if len(token***REMOVED*** != 26 {
		return nil, errors.New("authentication token wrong size"***REMOVED***
***REMOVED***

	return &u, nil
***REMOVED***

func (app *application***REMOVED*** CheckAuthentication(w http.ResponseWriter, r *http.Request***REMOVED*** {
	// validate the token, and get associated user
	// fmt.Println("authenticateToken ----------"***REMOVED***
	user, err := app.authenticateToken(r***REMOVED***
***REMOVED***
		app.invalidCredentials(w***REMOVED***
		return
***REMOVED***

	// valid user
	var payload struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
***REMOVED***
	payload.Error = false
	payload.Message = fmt.Sprintf("authenticated user %s", user.Email***REMOVED***
	app.writeJSON(w, http.StatusOK, payload***REMOVED***
***REMOVED***

func (app *application***REMOVED*** VirtualTerminalPaymentSucceeded(w http.ResponseWriter, r *http.Request***REMOVED*** {
	var txnData struct {
		PaymentAmount   int    `json:"amount"`
		PaymentCurrency string `json:"currency"`
		FirstName       string `json:"first_name"`
		LastName        string `json:"last_name"`
		Email           string `json:"email"`
		PaymentIntent   string `json:"payment_intent"`
		PaymentMethod   string `json:"payment_method"`
		BankReturnCode  string `json:"bank_return_code"`
		ExpiryMonth     int    `json:"expiry_month"`
		ExpiryYear      int    `json:"expiry_year"`
		LastFour        string `json:"last_four"`
***REMOVED***
	// fmt.Println("readJSON process"***REMOVED***
	err := app.readJSON(w, r, &txnData***REMOVED***
***REMOVED***
		app.badRequest(w, r, err***REMOVED***
		return
***REMOVED***

	card := cards.Card{
		Secret: app.config.stripe.secret,
		Key:    app.config.stripe.key,
***REMOVED***
	// fmt.Println("pi process"***REMOVED***
	pi, err := card.RetrievePaymentIntent(txnData.PaymentIntent***REMOVED***
***REMOVED***
		app.badRequest(w, r, err***REMOVED***
		return
***REMOVED***

	// fmt.Println("pm process"***REMOVED***
	pm, err := card.GetPaymentMethod(txnData.PaymentMethod***REMOVED***
***REMOVED***
		app.badRequest(w, r, err***REMOVED***
		return
***REMOVED***

	txnData.LastFour = pm.Card.Last4
	txnData.ExpiryMonth = int(pm.Card.ExpMonth***REMOVED***
	txnData.ExpiryYear = int(pm.Card.ExpYear***REMOVED***

	txn := models.Transaction{
		Amount:              txnData.PaymentAmount,
		Currency:            txnData.PaymentCurrency,
		LastFour:            txnData.LastFour,
		ExpiryMonth:         txnData.ExpiryMonth,
		ExpiryYear:          txnData.ExpiryYear,
		PaymentIndent:       txnData.PaymentIntent,
		PaymentMethod:       txnData.PaymentMethod,
		BankReturnCode:      pi.Charges.Data[0].ID,
		TransactionStatusID: 2,
***REMOVED***

	// fmt.Println("SaveTransaction process"***REMOVED***
	_, err = app.SaveTransaction(txn***REMOVED***
***REMOVED***
		app.badRequest(w, r, err***REMOVED***
		return
***REMOVED***

	app.writeJSON(w, http.StatusOK, txn***REMOVED***
***REMOVED***
