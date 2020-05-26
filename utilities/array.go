package utilities

func Insert(array []interface{}, item interface{}, index int) {
	array = append(array, 0 /* use the zero value of the element type */)
	copy(array[index+1:], array[index:])
	array[index] = item
}
