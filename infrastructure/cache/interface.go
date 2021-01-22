package cache

type Redis interface {
	Get(key string) (string, error)
	Set(key string, value string, expiresOn int) (bool, error)
	Delete(key string) error
	Ping() (string, error)
}
