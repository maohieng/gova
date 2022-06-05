package collect

import "reflect"

type Collection interface {
	[]any | ~string | map[any]any /*| chan string*/
}

func IndexOf[T Collection](list T, obj any) int {
	for index, item := range list {
		if reflect.DeepEqual(item, obj) {
			return index
		}
	}

	return -1
}

func Contains[T Collection](list T, obj any) bool {
	return IndexOf(list, obj) > -1
}
