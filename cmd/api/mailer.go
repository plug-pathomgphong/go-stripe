***REMOVED***

***REMOVED***
	"bytes"
	"embed"
***REMOVED***
***REMOVED***
***REMOVED***

	mail "github.com/xhit/go-simple-mail/v2"
***REMOVED***

//go:embed templates
var emailTemplateFS embed.FS

func (app *application***REMOVED*** SendMail(from, to, subject, tmpl string, data interface{***REMOVED******REMOVED*** error {
	templateToRender := fmt.Sprintf("templates/%s.html.tmpl", tmpl***REMOVED***

	t, err := template.New("email-html"***REMOVED***.ParseFS(emailTemplateFS, templateToRender***REMOVED***
***REMOVED***
		app.errorLog.Println(err***REMOVED***
		return err
***REMOVED***

	var tpl bytes.Buffer
	if err = t.ExecuteTemplate(&tpl, "body", data***REMOVED***; err != nil {
		app.errorLog.Println(err***REMOVED***
		return err
***REMOVED***

	formattedMessage := tpl.String(***REMOVED***

	templateToRender = fmt.Sprintf("templates/%s.plain.tmpl", tmpl***REMOVED***
	t, err = template.New("email-plain"***REMOVED***.ParseFS(emailTemplateFS, templateToRender***REMOVED***
***REMOVED***
		app.errorLog.Println(err***REMOVED***
		return err
***REMOVED***

	if err = t.ExecuteTemplate(&tpl, "body", data***REMOVED***; err != nil {
		app.errorLog.Println(err***REMOVED***
		return err
***REMOVED***

	plainMessage := tpl.String(***REMOVED***
	fmt.Println("formattedMessage"***REMOVED***
	app.infoLog.Println(formattedMessage, plainMessage***REMOVED***

	// send the mail
	server := mail.NewSMTPClient(***REMOVED***
	server.Host = app.config.smtp.host
	server.Port = app.config.smtp.port
	server.Username = app.config.smtp.username
	server.Password = app.config.smtp.password
	server.Encryption = mail.EncryptionTLS
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	smtpClient, err := server.Connect(***REMOVED***
***REMOVED***
		return err
***REMOVED***

	email := mail.NewMSG(***REMOVED***
	email.SetFrom(from***REMOVED***.AddTo(to***REMOVED***.SetSubject(subject***REMOVED***

	email.SetBody(mail.TextHTML, formattedMessage***REMOVED***
	email.AddAlternative(mail.TextPlain, plainMessage***REMOVED***

	err = email.Send(smtpClient***REMOVED***
***REMOVED***
		app.errorLog.Println(err***REMOVED***
		return err
***REMOVED***

	app.infoLog.Println("send mail"***REMOVED***

	return nil
***REMOVED***
