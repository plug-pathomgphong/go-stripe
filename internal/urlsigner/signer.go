package urlsigner

***REMOVED***
***REMOVED***
	"strings"
***REMOVED***

	goalone "github.com/bwmarrin/go-alone"
***REMOVED***

type Signer struct {
	Secret []byte
***REMOVED***

func (s *Signer***REMOVED*** GenerateTokenFromString(data string***REMOVED*** string {
	var urlToSign string

	crypt := goalone.New(s.Secret, goalone.Timestamp***REMOVED***
	if strings.Contains(data, "?"***REMOVED*** {
		urlToSign = fmt.Sprintf("%s&hash=", data***REMOVED***
***REMOVED*** else {
		urlToSign = fmt.Sprintf("%s?hash=", data***REMOVED***
***REMOVED***

	tokenBypes := crypt.Sign([]byte(urlToSign***REMOVED******REMOVED***
	token := string(tokenBypes***REMOVED***
	return token
***REMOVED***

func (s *Signer***REMOVED*** VerifyToken(token string***REMOVED*** bool {
	crypt := goalone.New(s.Secret, goalone.Timestamp***REMOVED***
	_, err := crypt.Unsign([]byte(token***REMOVED******REMOVED***

***REMOVED***
		fmt.Println(err***REMOVED***
		return false
***REMOVED***

	return true
***REMOVED***

func (s *Signer***REMOVED*** Expired(token string, minutesUntilExpire int***REMOVED*** bool {
	crypt := goalone.New(s.Secret, goalone.Timestamp***REMOVED***
	ts := crypt.Parse([]byte(token***REMOVED******REMOVED***

	return time.Since(ts.Timestamp***REMOVED*** > time.Duration(minutesUntilExpire***REMOVED****time.Minute
***REMOVED***
