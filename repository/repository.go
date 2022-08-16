package repository

type Repository interface {
	GetAll() (map[string]string, error)
	Add(id, source string) error
	Get(shortUrl string) (string, error)
}
