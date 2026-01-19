package user

type PasswordHasher interface {
	Hash(rawPassword string) (string, error)
	Compare(hashedPassword, rawPassword string) bool
}
