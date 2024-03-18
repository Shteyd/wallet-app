package contract

type SecretHasher interface {
	Hash(raw string) (string, error)
}
