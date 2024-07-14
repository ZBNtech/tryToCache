package tryToCache

type cache struct {
	key   int
	value interface{}
}

var cacheM map[string]interface{}

func Set(key string, val interface{}) {
	cacheM[key] = val
}

func Get(key string) interface{} {
	return cacheM[key].(interface{})
}

func Delete(key string) {
	delete(cacheM, key)
}
