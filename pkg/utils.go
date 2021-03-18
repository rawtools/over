package pkg

import (
	"reflect"
	"sort"
)

// func SortKeys(m map[string]interface{}) []string {
// 	keys := make([]string, len(m))
// 	i := 0
// 	for k := range m {
// 		keys[i] = k
// 		i++
// 	}
// 	sort.Strings(keys)
// 	return keys
// }

func SortedStringKeys(i interface{}) []string {

	v := reflect.ValueOf(i)

	if v.Kind() != reflect.Map {
		panic("parameter is not a map")
	}
	if v.Type().Key().Kind() != reflect.String {
		panic("parameter is not a map with string keys")
	}

	keys := v.MapKeys()
	strKeys := []string{}

	for _, key := range keys {
		strKeys = append(strKeys, key.Interface().(string))
	}

	sort.Strings(strKeys)
	return strKeys
}
