package models

***REMOVED***
	"context"
	"database/sql"
***REMOVED***
***REMOVED***

// DBModel is the type for database connection values
type DBModel struct {
	DB *sql.DB
***REMOVED***

// Modelis the wrapper for all models
type Models struct {
	DB DBModel
***REMOVED***

// NewModels returns a model type with database connect pool
func NewModels(db *sql.DB***REMOVED*** Models {
	return Models{
		DB: DBModel{DB: db***REMOVED***,
***REMOVED***
***REMOVED***

// Widget is the type for all widgets
type Widget struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	InventoryLevel int       `json:"inventory_level"`
	Price          int       `json:"price"`
	Image          string    `json:"image"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
***REMOVED***

// Order is the type for all order
type Order struct {
	ID            int       `json:"id"`
	WidgetID      int       `json:"widget_id"`
	TransactionID int       `json:"transaction_id"`
	CustomerID    int       `json:"customer_id"`
	StatusID      int       `json:"status_id"`
	Quantity      int       `json:"quantity"`
	Amount        int       `json:"amount"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`
***REMOVED***

// Status is the type for order statuses
type Status struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
***REMOVED***

// TransactionStatus is the type for Transaction statuses
type TransactionStatus struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
***REMOVED***

// Transaction is the type for Transactions
type Transaction struct {
	ID                  int       `json:"id"`
	Amount              int       `json:"amount"`
	Currency            string    `json:"currency"`
	LastFour            string    `json:"last_four"`
	ExpiryMonth         int       `json:"expiry_month"`
	ExpiryYear          int       `json:"expiry_year"`
	PaymentIndent       string    `json:"payment_indent"`
	PaymentMethod       string    `json:"payment_method"`
	BankReturnCode      string    `json:"bank_return_code"`
	TransactionStatusID int       `json:"transaction_status_id"`
	CreatedAt           time.Time `json:"-"`
	UpdatedAt           time.Time `json:"-"`
***REMOVED***

// User is the type for users
type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
***REMOVED***

// Customer is the type for customers
type Customer struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
***REMOVED***

func (m *DBModel***REMOVED*** GetWidget(id int***REMOVED*** (Widget, error***REMOVED*** {
	ctx, cancel := context.WithTimeout(context.Background(***REMOVED***, 3*time.Second***REMOVED***
	defer cancel(***REMOVED***

	var widget Widget

	row := m.DB.QueryRowContext(ctx, `
		select 
			id, name, description, inventory_level, price, coalesce(image, ''***REMOVED***,
			created_at, updated_at
		from
			widgets
		where id = ?`, id***REMOVED***
	err := row.Scan(
		&widget.ID,
		&widget.Name,
		&widget.Description,
		&widget.InventoryLevel,
		&widget.Price,
		&widget.Image,
		&widget.CreatedAt,
		&widget.UpdatedAt,
	***REMOVED***
***REMOVED***
		return widget, err
***REMOVED***
	return widget, nil
***REMOVED***

// InsertTransaction inserts a new txn, and return its id
func (m *DBModel***REMOVED*** InsertTransaction(txn Transaction***REMOVED*** (int, error***REMOVED*** {
	ctx, close := context.WithTimeout(context.Background(***REMOVED***, 3*time.Second***REMOVED***
	defer close(***REMOVED***

	stmt := `insert into transactions 
		(amount, currency, last_four, bank_return_code, expiry_month, expiry_year,
		payment_indent, payment_method,
		transaction_status_id, created_at, updated_at***REMOVED*** 
		values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?***REMOVED***`

	result, err := m.DB.ExecContext(ctx, stmt,
		txn.Amount,
		txn.Currency,
		txn.LastFour,
		txn.BankReturnCode,
		txn.ExpiryMonth,
		txn.ExpiryYear,
		txn.PaymentIndent,
		txn.PaymentMethod,
		txn.TransactionStatusID,
		time.Now(***REMOVED***,
		time.Now(***REMOVED***,
	***REMOVED***
***REMOVED***
		return 0, err
***REMOVED***

	id, err := result.LastInsertId(***REMOVED***
***REMOVED***
		return 0, err
***REMOVED***

	return int(id***REMOVED***, nil
***REMOVED***

// InsertCustomer inserts a new customer, and return its id
func (m *DBModel***REMOVED*** InsertCustomer(c Customer***REMOVED*** (int, error***REMOVED*** {
	ctx, close := context.WithTimeout(context.Background(***REMOVED***, 3*time.Second***REMOVED***
	defer close(***REMOVED***

	stmt := `insert into customers 
		(first_name, last_name, email, created_at, updated_at***REMOVED*** 
		values (?, ?, ?, ?, ?***REMOVED***`

	result, err := m.DB.ExecContext(ctx, stmt,
		c.FirstName,
		c.LastName,
		c.Email,
		time.Now(***REMOVED***,
		time.Now(***REMOVED***,
	***REMOVED***
***REMOVED***
		return 0, err
***REMOVED***

	id, err := result.LastInsertId(***REMOVED***
***REMOVED***
		return 0, err
***REMOVED***

	return int(id***REMOVED***, nil
***REMOVED***

// InsertOrder inserts a new order, and return its id
func (m *DBModel***REMOVED*** InsertOrder(order Order***REMOVED*** (int, error***REMOVED*** {
	ctx, close := context.WithTimeout(context.Background(***REMOVED***, 3*time.Second***REMOVED***
	defer close(***REMOVED***

	stmt := `insert into orders 
		(widget_id, transaction_id, status_id, quantity, customer_id,
		amount, created_at, updated_at***REMOVED*** 
		values (?, ?, ?, ?, ?, ?, ?, ?***REMOVED***`

	result, err := m.DB.ExecContext(ctx, stmt,
		order.WidgetID,
		order.TransactionID,
		order.StatusID,
		order.Quantity,
		order.CustomerID,
		order.Amount,
		time.Now(***REMOVED***,
		time.Now(***REMOVED***,
	***REMOVED***
***REMOVED***
		return 0, err
***REMOVED***

	id, err := result.LastInsertId(***REMOVED***
***REMOVED***
		return 0, err
***REMOVED***

	return int(id***REMOVED***, nil
***REMOVED***
