package cards

***REMOVED***
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/paymentintent"
	"github.com/stripe/stripe-go/v74/paymentmethod"
***REMOVED***

type Card struct {
	Secret   string
	Key      string
	Currency string
***REMOVED***

type Transaction struct {
	TransactionStatusID int
	Amount              int
	Currency            string
	LastFour            string // Last Four digit on card
	BankReturnCode      string
***REMOVED***

func (c *Card***REMOVED*** Charge(currency string, amount int***REMOVED*** (*stripe.PaymentIntent, string, error***REMOVED*** {
	return c.CreatePaymentIntent(currency, amount***REMOVED***
***REMOVED***

func (c *Card***REMOVED*** CreatePaymentIntent(currency string, amount int***REMOVED*** (*stripe.PaymentIntent, string, error***REMOVED*** {

	stripe.Key = c.Secret

	// create a payment intent
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(amount***REMOVED******REMOVED***,
		Currency: stripe.String(currency***REMOVED***,
***REMOVED***

	// params.AddMetaData("key", "value"***REMOVED***

	pi, err := paymentintent.New(params***REMOVED***
***REMOVED***
		msg := ""
		if stripeErr, ok := err.(*stripe.Error***REMOVED***; ok {
			msg = cardErrorMessage(stripeErr.Code***REMOVED***
	***REMOVED***
		return nil, msg, err
***REMOVED***
	return pi, "", nil
***REMOVED***

// GetPAymentMethod gets the payment method by payment intend id
func (c *Card***REMOVED*** GetPaymentMethod(s string***REMOVED*** (*stripe.PaymentMethod, error***REMOVED*** {
	stripe.Key = c.Secret
	pm, err := paymentmethod.Get(s, nil***REMOVED***
***REMOVED***
		return nil, err
***REMOVED***
	return pm, nil
***REMOVED***

// RetrievePaymentIntent gets an existing payment intend by id
func (c *Card***REMOVED*** RetrievePaymentIntent(id string***REMOVED*** (*stripe.PaymentIntent, error***REMOVED*** {
	stripe.Key = c.Secret
	pi, err := paymentintent.Get(id, nil***REMOVED***
***REMOVED***
		return nil, err
***REMOVED***
	return pi, nil
***REMOVED***

func cardErrorMessage(code stripe.ErrorCode***REMOVED*** string {
	var msg = ""
	switch code {
	case stripe.ErrorCodeCardDeclined:
		msg = "Your card was declined"
	case stripe.ErrorCodeExpiredCard:
		msg = "Your card is rxpired"
	case stripe.ErrorCodeIncorrectCVC:
		msg = "Incorrect CVC code"
	case stripe.ErrorCodeIncorrectZip:
		msg = "Incorrect zip/postal code"
	case stripe.ErrorCodeAmountTooLarge:
		msg = "The amount is too large to charge to your card"
	case stripe.ErrorCodeAmountTooSmall:
		msg = "The amount is too small to charge to your card"
	case stripe.ErrorCodeBalanceInsufficient:
		msg = "Insufficient balance"
	case stripe.ErrorCodePostalCodeInvalid:
		msg = "Your postal code is invalid"
	default:
		msg = "Your card was declined"
***REMOVED***
	return msg
***REMOVED***
