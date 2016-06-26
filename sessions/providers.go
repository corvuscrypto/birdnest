package sessions

import (
	"database/sql"
	"fmt"

	"github.com/corvuscrypto/birdnest/security"
)

//SessionProvider is the interface through which sessions are created or removed
type SessionProvider interface {
	NewSession() *Session
	RemoveSession()
	RefreshSession()
}

//SQLSessionProvider is the built-in session provider for SQL data stores.
type SQLSessionProvider struct {
	db         *sql.DB
	TableName  string
	OwnerField string
}

//NewSession creates a session and inserts it into the database table specified by the provider.
func (p *SQLSessionProvider) NewSession(owner interface{}) *Session {
	csrfToken := security.GenerateCSRFToken()
	session := NewSession(owner)
	//insert token
	insertStmt := fmt.Sprintf("INSERT into %s (%s, SessionToken, CSRFToken) values (?,?,?)",
		p.TableName, p.OwnerField)
	p.db.Exec(insertStmt, owner, session.rawToken, csrfToken)
	return session
}
