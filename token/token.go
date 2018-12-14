package token

type Token interface {
	Generage() (string, string, error)
	Balance(string) (amount float64, err error)
}
