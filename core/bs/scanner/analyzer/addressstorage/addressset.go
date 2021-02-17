package addressstorage

type AddressSet struct {
	addresses map[string]struct{}
}

func NewSet() *AddressSet {

	return &AddressSet{make(map[string]struct{})}
}

func (l *AddressSet) HasValue(key string) bool {

	_, ok := l.addresses[key]
	return ok
}

func (l *AddressSet) SetValue(key string) {

	l.addresses[key] = struct{}{}
}
