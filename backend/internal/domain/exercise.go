package domain

type Exercise struct {
	ID        int64
	SessionID int64
	Name      string
	Sets      []Set
}
