***REMOVED***

***REMOVED***
	"encoding/json"
	"errors"
	"io"
***REMOVED***

	"golang.org/x/crypto/bcrypt"
***REMOVED***

// writeJSON wrties aribtarary data out as JSON
func (app *application***REMOVED*** writeJSON(w http.ResponseWriter, status int, data interface{***REMOVED***, headers ...http.Header***REMOVED*** error {
	out, err := json.MarshalIndent(data, "", "\t"***REMOVED***
***REMOVED***
		return err
***REMOVED***

	if len(headers***REMOVED*** > 0 {
		for k, v := range headers[0] {
			w.Header(***REMOVED***[k] = v
	***REMOVED***
***REMOVED***

	w.Header(***REMOVED***.Set("Content-Type", "application/json"***REMOVED***
	w.WriteHeader(status***REMOVED***
	w.Write(out***REMOVED***
	return nil
***REMOVED***

func (app *application***REMOVED*** readJSON(w http.ResponseWriter, r *http.Request, data interface{***REMOVED******REMOVED*** error {
	maxBytes := 1048576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes***REMOVED******REMOVED***

	dec := json.NewDecoder(r.Body***REMOVED***
	err := dec.Decode(data***REMOVED***
***REMOVED***
		return err
***REMOVED***

	err = dec.Decode(&struct{***REMOVED***{***REMOVED******REMOVED***
	if err != io.EOF {
		return errors.New("body must only have a single JSON value"***REMOVED***
***REMOVED***

	return nil
***REMOVED***

func (app *application***REMOVED*** badRequest(w http.ResponseWriter, r *http.Request, err error***REMOVED*** error {
	var payload struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
***REMOVED***

	payload.Error = true
	payload.Message = err.Error(***REMOVED***

	out, err := json.MarshalIndent(payload, "", "   "***REMOVED***
***REMOVED***
		return err
***REMOVED***

	w.Header(***REMOVED***.Set("Content-Type", "application/json"***REMOVED***
	w.Write(out***REMOVED***
	return nil
***REMOVED***

func (app *application***REMOVED*** invalidCredentials(w http.ResponseWriter***REMOVED*** error {
	var payload struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
***REMOVED***

	payload.Error = true
	payload.Message = "invalid authentication credentials"

	err := app.writeJSON(w, http.StatusUnauthorized, payload***REMOVED***
***REMOVED***
		return err
***REMOVED***
	return nil
***REMOVED***

func (app *application***REMOVED*** passwordMatches(hash, password string***REMOVED*** (bool, error***REMOVED*** {
	err := bcrypt.CompareHashAndPassword([]byte(hash***REMOVED***, []byte(password***REMOVED******REMOVED***
***REMOVED***
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword***REMOVED***:
			return false, nil
		default:
			return false, err
	***REMOVED***
***REMOVED***
	return true, nil
***REMOVED***
