package domain

type Exercise struct {
	ID        int
	SessionID int
	Name      string
	Sets      []Set
}
