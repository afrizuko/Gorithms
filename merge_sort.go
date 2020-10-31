package gorithms

//Merge merges the sorted individual entries
func Merge(l, r []int) []int {
	ret := make([]int, 0, (len(l) + len(r)))
	for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
			return append(ret, r...)
		}
		if len(r) == 0 {
			return append(ret, l...)
		}
		if l[0] <= r[0] {
			ret = append(ret, l[0])
			l = l[1:]
		} else {
			ret = append(ret, r[0])
			r = r[1:]
		}
	}
	return ret
}

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	n := len(arr) / 2
	l := MergeSort(arr[:n])
	r := MergeSort(arr[n:])

	return Merge(l, r)
}
