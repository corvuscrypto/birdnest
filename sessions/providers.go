package sessions

import "database/sql"

//SessionProvider is the interface through which sessions are created or removed
type SessionProvider interface {
	NewSessionToken() []byte
	RemoveSession()
	RefreshSession()
}

//SQLSessionProvider is the built-in session provider for SQL data stores.
type SQLSessionProvider struct {
	*sql.DB
}

func (provider *SQLSessionProvider) NewSessionToken() []byte {

	return []byte{}
}
