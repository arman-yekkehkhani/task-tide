package security

type HashAlgorithm string

const (
	BCRYPT HashAlgorithm = "bcrypt"
)

type HashService interface {
	Hash(algorithm HashAlgorithm, raw string) (string, error)
}
