package models

***REMOVED***
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
***REMOVED***
***REMOVED***

const (
	ScopeAuthentication = "authentication"
***REMOVED***

// Token is the type for authentication tokens
type Token struct {
	PlainText string    `json:"token"`
	UserID    int64     `json:"-"`
	Hash      []byte    `json:"-"`
	Expiry    time.Time `json:"expiry"`
	Scope     string    `json:"-"`
***REMOVED***

// GenerateToken generates a token that lasts for ttl, and returns it
func GenerateToken(userID int, ttl time.Duration, scope string***REMOVED*** (*Token, error***REMOVED*** {
	token := &Token{
		UserID: int64(userID***REMOVED***,
		Expiry: time.Now(***REMOVED***.Add(ttl***REMOVED***,
		Scope:  scope,
***REMOVED***

	randomBytes := make([]byte, 16***REMOVED***
	_, err := rand.Read(randomBytes***REMOVED***
***REMOVED***
		return nil, err
***REMOVED***

	token.PlainText = base32.StdEncoding.WithPadding(base32.NoPadding***REMOVED***.EncodeToString(randomBytes***REMOVED***
	hash := sha256.Sum256(([]byte(token.PlainText***REMOVED******REMOVED******REMOVED***
	token.Hash = hash[:]
	return token, nil
***REMOVED***

func (m *DBModel***REMOVED*** InsertToken(t *Token, u User***REMOVED*** error {
	ctx, cancel := context.WithTimeout(context.Background(***REMOVED***, 3*time.Second***REMOVED***
	defer cancel(***REMOVED***

	stmt := `insert into tokens 
		(user_id, name, email, token_hash, expiry, created_at, updated_at***REMOVED*** 
		values (?, ?, ?, ?, ?, ?, ?***REMOVED***`

	_, err := m.DB.ExecContext(ctx, stmt,
		u.ID,
		u.LastName,
		u.Email,
		t.Hash,
		t.Expiry,
		time.Now(***REMOVED***,
		time.Now(***REMOVED***,
	***REMOVED***
***REMOVED***
		return err
***REMOVED***

	return nil
***REMOVED***

func (m *DBModel***REMOVED*** GetUserForToken(token string***REMOVED*** (*User, error***REMOVED*** {
	ctx, cancel := context.WithTimeout(context.Background(***REMOVED***, 3*time.Second***REMOVED***
	defer cancel(***REMOVED***

	tokenHash := sha256.Sum256([]byte(token***REMOVED******REMOVED***
	var user User

	query := `
		select 
			u.id, u.first_name, u.last_name, u.email
		from
			user u
			inner join tokens t on (u.id = t.user_id***REMOVED***
		where 
			t.token_hash = ?
			and t.expiry > ?
		`

	err := m.DB.QueryRowContext(ctx, query, tokenHash[:], time.Now(***REMOVED******REMOVED***.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
	***REMOVED***
***REMOVED***
		return nil, err
***REMOVED***
	return &user, nil
***REMOVED***
