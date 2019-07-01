package week3

type ComparatorFunc func(int, int, []interface{}) int

// Sorts data in descending order by user defined comparator function
// If data[i] > data[j] positive value must be returned
// If data[i] = data[j] zero must be returned
// If data[i] < data[j] negative value must be returned
func Sort(data []interface{}, comparatorFunc ComparatorFunc) {
	dataLen := len(data)
	if dataLen < 2 {
		return
	}
	for i0 := 0; i0 < dataLen; i0++ {
		for i := 0; i < dataLen-1; i++ {
			res := comparatorFunc(i, i+1, data)
			if res < 0 {
				tmp := data[i]
				data[i] = data[i+1]
				data[i+1] = tmp
			}
		}
	}
}
