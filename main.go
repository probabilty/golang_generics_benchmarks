package generics
import (
	"reflect"
	"errors"
)
func addGeneric[T int|int64|string](a,b T)T  {
	return a+b
}
func addInt64(a,b int64)int64 {
	return a+b
}
func addInterface(a, b interface{}) (interface{}, error) {
	typeA := reflect.TypeOf(a).Kind()
	typeB := reflect.TypeOf(b).Kind()
	if typeA != typeB {
		return nil, errors.New("Both parameters should be of the same type")
	}
	switch typeA {
	case reflect.Int:
		return a.(int) + b.(int), nil
	case reflect.Int8:
		return a.(int8) + b.(int8), nil
	case reflect.Int16:
		return a.(int16) + b.(int16), nil
	case reflect.Int32:
		return a.(int32) + b.(int32), nil
	case reflect.Int64:
		return a.(int64) + b.(int64), nil
	case reflect.Float32:
		return a.(float32) + b.(float32), nil
	case reflect.Float64:
		return a.(float64) + b.(float64), nil
	case reflect.String:
		return a.(string) + b.(string), nil
	default:
		return nil, errors.New("Parameters type is not supported")
	}
}
func FibonacciGeneric[T int|int64](a T)T {
	if a <= 1 {
        return a
    }
    return FibonacciGeneric(a-1) + FibonacciGeneric(a-2)
}
func FibonacciInt64(a int64)int64 {
	if a <= 1 {
        return a
    }
    return FibonacciInt64(a-1) + FibonacciInt64(a-2)
}
func FibonacciInterface(a interface{})interface{}{
	typeA := reflect.TypeOf(a).Kind()
	if !(typeA == reflect.Int ||  typeA == reflect.Int64){
		return errors.New("Parameter type is not supported")
	}
	switch typeA {
	case reflect.Int:
		if a.(int) <= 1 {
			return a
		}
		return FibonacciInterface(a.(int)-1).(int) + FibonacciInterface(a.(int)-2).(int)
	case reflect.Int64:
		if a.(int64) <= 1 {
			return a
		}
		return FibonacciInterface(a.(int64)-1).(int64) + FibonacciInterface(a.(int64)-2).(int64)
	}
	return nil
}