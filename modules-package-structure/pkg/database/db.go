package database

var m = map[string]string{}

func Get(key string) (string, bool) {
	v, ok := m[key]
	return v, ok
}

func Put(key, val string) {
	m[key] = val
}
