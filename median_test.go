package median

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

var array1, array2 []uint32
var expectedMedian uint32

func init() {
	rand.Seed(time.Now().UTC().UnixNano())

	l := 100000

	array := make([]uint32, l)
	for i := 0; i < l; i++ {
		array[i] = uint32(i)
	}
	expectedMedian = (array[l/2] + array[l/2-1]) / 2

	rand.Shuffle(len(array), func(i, j int) { array[i], array[j] = array[j], array[i] })

	array1 = array[:l/2]
	array2 = array[l/2:]

	sort.Slice(array1, func(i, j int) bool { return array1[i] < array1[j] })
	sort.Slice(array2, func(i, j int) bool { return array2[i] < array2[j] })
}

func BenchmarkMedian_ON(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if m := GetMedian_ON(array1, array2); m != expectedMedian {
			b.Errorf("median is %d, expected %d", m, expectedMedian)
		}
	}
}

func BenchmarkMedian_ON2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if m := GetMedian_ON2(array1, array2); m != expectedMedian {
			b.Errorf("median is %d, expected %d", m, expectedMedian)
		}
	}
}

func BenchmarkMedian_OLOGN(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if m := GetMedian_LOGN(array1, array2); m != expectedMedian {
			b.Errorf("median is %d, expected %d", m, expectedMedian)
		}
	}
}
