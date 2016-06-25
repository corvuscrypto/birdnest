package sessions

//Session is the interface through which most sessions can be validated or checked for info
type Session interface {
	IsGuest() bool
	IsValid() bool
}
