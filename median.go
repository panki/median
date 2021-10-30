package median

import "math"

// GetMedian_ON finds median in O(N)
func GetMedian_ON(array1 []uint32, array2 []uint32) uint32 {
	l := len(array1)
	if l == 0 {
		return 0
	}

	// Quick win, both array consists of only one element
	if l == 1 {
		return (array1[0] + array2[0]) / 2
	}

	// Quick win, if arrays do not intersect
	if array1[l-1] <= array2[0] {
		return (array1[l-1] + array2[0]) / 2
	} else if array2[l-1] <= array1[0] {
		return (array2[l-1] + array1[0]) / 2
	}

	a := make([]uint32, l*2)

	// Merge sort
	for i, i1, i2 := 0, 0, 0; i < l*2; i++ {
		if i1 < l && i2 < l {
			if array1[i1] <= array2[i2] {
				a[i] = array1[i1]
				i1++
			} else {
				a[i] = array2[i2]
				i2++
			}
		} else if i1 < l {
			a[i] = array1[i1]
			i1++
		} else {
			a[i] = array2[i2]
			i2++
		}
	}

	return (a[l] + a[l-1]) / 2
}

// GetMedianON finds median in O(N/2)
func GetMedian_ON2(array1 []uint32, array2 []uint32) uint32 {
	l := len(array1)
	if l == 0 {
		return 0
	}

	// Quick win, both array consists of only one element
	if l == 1 {
		return (array1[0] + array2[0]) / 2
	}

	// Quick win, if arrays do not intersect
	if array1[l-1] <= array2[0] {
		return (array1[l-1] + array2[0]) / 2
	} else if array2[l-1] <= array1[0] {
		return (array2[l-1] + array1[0]) / 2
	}

	// This array will hold 2 elements to calculate median
	a := make([]uint32, 2)

	for i, i1, i2 := 0, 0, 0; i <= l; i++ {
		if array1[i1] <= array2[i2] {
			a[i%2] = array1[i1]
			i1++
		} else {
			a[i%2] = array2[i2]
			i2++
		}
	}

	return (a[0] + a[1]) / 2
}

// GetMedianON finds median in O(logN), see https://www.youtube.com/watch?v=LPFhl65R7ww
func GetMedian_LOGN(arrayL []uint32, arrayR []uint32) uint32 {
	l := len(arrayL)

	if l == 0 {
		return 0
	}

	// Quick win, both array consists of only one element
	if l == 1 {
		return (arrayL[0] + arrayR[0]) / 2
	}

	// Quick win, if arrays do not intersect
	if arrayL[l-1] <= arrayR[0] {
		return (arrayL[l-1] + arrayR[0]) / 2
	} else if arrayR[l-1] <= arrayL[0] {
		return (arrayR[l-1] + arrayL[0]) / 2
	}

	left, right := 0, l<<1

	for left <= right {
		cutR := (left + right) >> 1
		cutL := l<<1 - cutR

		L1 := math.MinInt32
		if cutL != 0 {
			L1 = int(arrayL[(cutL-1)/2])
		}

		L2 := math.MinInt32
		if cutR != 0 {
			L2 = int(arrayR[(cutR-1)/2])
		}

		R1 := math.MaxInt32
		if cutL != l<<1 {
			R1 = int(arrayL[cutL/2])
		}

		R2 := math.MaxInt32
		if cutR != l<<1 {
			R2 = int(arrayR[cutR/2])
		}

		if L1 > R2 {
			left = cutR + 1
		} else if L2 > R1 {
			right = cutR - 1
		} else {
			return uint32((max(L1, L2) + min(R1, R2)) / 2)
		}
	}

	return 0
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
