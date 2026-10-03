package logs

var cache Registrable

type Registrable interface {
	Register(log *Entry) error
}

func Init(r Registrable) {
	cache = r
}

// ---
//
//	TODO DELETE THIS FUNCTION WHEN THE LOGS ARE TESTED AND WORKING
func TestLog(log *Entry) error {
	return cache.Register(log)
}
