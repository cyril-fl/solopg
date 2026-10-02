package logs

var cache Registrable

type Registrable interface {
	Register(log Log) error
}

func Init(r Registrable) {
	cache = r
}

// ---
func TestLog(log Log) error {
	return cache.Register(log)
}
