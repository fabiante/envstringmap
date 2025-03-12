package envstringmap

import (
	"os"
	"strings"
)

// GetMap returns a map with keys that have the given prefix.
//
// The returned map will have keys without this prefix. As a special case, a leading underscore
// will also be removed from each key.
//
// For example, if the environment contains:
// APP_VALUE_ABC_1=1
// APP_VALUE_ABC_2=2
// APP_ELSE=unrelated
//
// Calling GetMap("APP_VALUE") will return:
//
//	map[string]string{
//	    "ABC_1": "1",
//	    "ABC_2": "2",
//	}
//
// If you want the keys to be non-case-sensitive, you can use strings.ToLower() on the keys like so:
//
//	values := make(map[string]string)
//	for k, v := range envstringmap.GetMap("APP_VALUE") {
//	    k := strings.ToLower(k)
//	    // do something with k and v - you could also map it into a new map
//	    values[k] = v
//	}
func GetMap(prefix string) map[string]string {
	result := make(map[string]string)
	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		key, value := pair[0], pair[1]
		if strings.HasPrefix(key, prefix) {
			newKey := strings.TrimPrefix(key, prefix)
			newKey = strings.TrimPrefix(newKey, "_")
			result[newKey] = value
		}
	}
	return result
}
