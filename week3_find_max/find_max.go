package week3_find_max

type ComparatorFunc func(int, int, []interface{}) int

// Sorts data in descending order by user defined comparator function
// If data[i] > data[j] positive value must be returned
// If data[i] = data[j] zero must be returned
// If data[i] < data[j] negative value must be returned
func FindMax(data []interface{}, comparatorFunc ComparatorFunc) interface{} {
	dataLen := len(data)
	if dataLen < 2 {
		return data[0]
	}

	maxIndex := 0
	for i := 1; i < dataLen; i++ {
		res := comparatorFunc(maxIndex, i, data)
		if res < 0 {
			maxIndex = i
		}
	}

	return data[maxIndex]

}
