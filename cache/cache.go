package cache

type Cache interface {
	All() (map[string]string, error)
	Add(string, string) error
	Get(string) (string, error)
}
