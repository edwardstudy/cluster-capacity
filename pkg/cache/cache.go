package cache

// You only need **one** of these per package!
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . Cache
type Cache interface {
	Get(key string) ([]byte, error)
	Set(key string, data []byte) error
	Clean(key string) error
}
