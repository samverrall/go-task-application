package hasher

type Hasher interface {
	Generate(string) (string, error)
}
