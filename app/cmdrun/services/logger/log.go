package logger


type Register interface {
	Register(log Log) error
}

var cache Register

func Init(r Register) {
	cache = r
}

func TestLog(log Log) error {
	return cache.Register(log)
}