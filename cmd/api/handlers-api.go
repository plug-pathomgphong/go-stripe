***REMOVED***

***REMOVED***
	"encoding/json"
***REMOVED***
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/plug-pathomgphong/dotnet-webapi/internal/cards"
***REMOVED***

type stripePayload struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
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
