package logs

var (
	cache_registrable Registrable
	caches_author     = "System"
)

type Registrable interface {
	Register(log *Log) error
}

func Init(r Registrable) {
	cache_registrable = r
}

func SetAuthor(a string) {
	caches_author = a
}
