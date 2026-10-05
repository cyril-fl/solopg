package primitive

import "errors"

type Fallible struct {
	errs []error
}

func (f *Fallible) SetErr(err error) {
	if err != nil {
		f.errs = append(f.errs, err)
	}
}
func (f *Fallible) ClearErr() {
	f.errs = nil
}

func (f *Fallible) GetErr() error {
	return errors.Join(f.errs...)
}

func (f *Fallible) HasErr() bool {
	return len(f.errs) > 0
}
