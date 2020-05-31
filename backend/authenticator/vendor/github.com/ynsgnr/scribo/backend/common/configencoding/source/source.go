package source

type Source interface {
	GetValue(tag string) (string, bool)
}
