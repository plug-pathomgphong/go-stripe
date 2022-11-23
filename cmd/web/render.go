***REMOVED***

***REMOVED***
	"embed"
***REMOVED***
***REMOVED***
	"strings"
***REMOVED***
***REMOVED***

type templateData struct {
	StringMap       map[string]string
	IntMap          map[string]int
	FloatMap        map[string]float32
	Data            map[string]interface{***REMOVED***
	CSRFToke        string
	Flash           string
	Warning         string
	Error           string
	IsAuthenticated int
	API             string
	CSSVersion      string
***REMOVED***

var functions = template.FuncMap{***REMOVED***

//go:embed templates
var templateFS embed.FS

func (app *application***REMOVED*** addDefaultData(td *templateData, r *http.Request***REMOVED*** *templateData {
	td.API = app.config.api
	return td
***REMOVED***

func (app *application***REMOVED*** renderTemplate(w http.ResponseWriter, r *http.Request, page string, td *templateData, partials ...string***REMOVED*** error {
	var t *template.Template
	var err error
	templateToRender := fmt.Sprintf("templates/%s.page.gohtml", page***REMOVED***

	_, templateInMap := app.templateCache[templateToRender]
	if app.config.env == "production" && templateInMap {
		t = app.templateCache[templateToRender]
***REMOVED*** else {
		t, err = app.parseTemplate(partials, page, templateToRender***REMOVED***
	***REMOVED***
			app.errorLog.Println(err***REMOVED***
			return err
	***REMOVED***
***REMOVED***

	if td == nil {
		td = &templateData{***REMOVED***
***REMOVED***

	td = app.addDefaultData(td, r***REMOVED***

	err = t.Execute(w, td***REMOVED***
***REMOVED***
		app.errorLog.Println(err***REMOVED***
		return err
***REMOVED***
	return nil
***REMOVED***

func (app *application***REMOVED*** parseTemplate(partials []string, page, templateToRender string***REMOVED*** (*template.Template, error***REMOVED*** {
	var t *template.Template
	var err error

	// build partials
	if len(partials***REMOVED*** > 0 {
		for i, x := range partials {
			partials[i] = fmt.Sprintf("templates/%s.partials.gohtml", x***REMOVED***
	***REMOVED***
***REMOVED***

	if len(partials***REMOVED*** > 0 {
		t, err = template.New(fmt.Sprintf("%s.page.gohtml", page***REMOVED******REMOVED***.Funcs(functions***REMOVED***.ParseFS(templateFS, "templates/base.layout.gohtml", strings.Join(partials, ","***REMOVED***, templateToRender***REMOVED***
***REMOVED*** else {
		t, err = template.New(fmt.Sprintf("%s.page.gohtml", page***REMOVED******REMOVED***.Funcs(functions***REMOVED***.ParseFS(templateFS, "templates/base.layout.gohtml", templateToRender***REMOVED***
***REMOVED***

***REMOVED***
		app.errorLog.Println(err***REMOVED***
		return nil, err
***REMOVED***

	app.templateCache[templateToRender] = t
	return t, nil
***REMOVED***
