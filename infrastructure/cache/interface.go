package cache

type Redis interface {
	Get(key string, data interface{}) error
	Set(key string, value interface{}, expiresOn int) (bool, error)
	Delete(key string) error
	Ping() (string, error)
}
