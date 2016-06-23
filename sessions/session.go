package sessions

//Session is the interface through which most sessions can be validated or checked for info
type Session interface {
	IsGuest() bool
	IsValid() bool
}

//SessionProvider is the interface through which sessions are created or removed
type SessionProvider interface {
	NewSessionToken() []byte
	RemoveSession()
	RefreshSession()
}
