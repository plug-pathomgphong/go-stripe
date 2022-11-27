CREATE TABLE sessions (
	token CHAR(43***REMOVED*** PRIMARY KEY,
	data BLOB NOT NULL,
	expiry TIMESTAMP(6***REMOVED*** NOT NULL
***REMOVED***;

CREATE INDEX sessions_expiry_idx ON sessions (expiry***REMOVED***;