package repository

//Simple key value storage
type Interface interface {
	Read(key string) (string, error)
	Write(key, value string) error
}
