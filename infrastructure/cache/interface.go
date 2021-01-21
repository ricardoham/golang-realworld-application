package cache

type Redis interface {
	Get(key string) (string, error)
	Ping() (string, error)
}
