package walk

import (
	"reflect"
)

func xwalk(x interface{}, fn func(s string)) {
	val := getValue(x)
	numberOfValues := 0
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Slice, reflect.Array:
		numberOfValues = val.Len()
		getField = val.Index
	case reflect.Struct:
		numberOfValues = val.NumField()
		getField = val.Field
	case reflect.Map:
		for _, key := range val.MapKeys() {
			xwalk(val.MapIndex(key).Interface(), fn)
		}
	}
	for i := 0; i < numberOfValues; i++ {
		xwalk(getField(i).Interface(), fn)
	}
}

func walk(x interface{}, fn func(s string)) {
	val := getValue(x)
	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	}

}

func getValue(v interface{}) reflect.Value {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}
