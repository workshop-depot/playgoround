package labsort_test

import (
	"fmt"
	"sort"
	"testing"
	"time"

	labsort "github.com/dc0d/playgoround/lab-sort"

	"github.com/stretchr/testify/assert"
)

func ExampleSorter_bytes() {
	collection := []byte{100, 90, 80}

	sorter := &labsort.Sorter{
		LenFunc:  func() int { return len(collection) },
		LessFunc: func(i, j int) bool { return collection[i] < collection[j] },
		SwapFunc: func(i, j int) { collection[i], collection[j] = collection[j], collection[i] },
	}
	sort.Sort(sorter)

	fmt.Println(collection)

	// Output:
	// [80 90 100]
}

func ExampleSorter_strings() {
	collection := []string{"ABC", "ZBC", "CBC"}

	sorter := &labsort.Sorter{
		LenFunc:  func() int { return len(collection) },
		LessFunc: func(i, j int) bool { return collection[i] < collection[j] },
		SwapFunc: func(i, j int) { collection[i], collection[j] = collection[j], collection[i] },
	}
	sort.Sort(sorter)

	fmt.Println(collection)

	// Output:
	// [ABC CBC ZBC]
}

func ExampleSorter_timeDurations() {
	collection := []time.Duration{time.Second * 10, time.Second * 120, time.Millisecond, time.Nanosecond, time.Microsecond}

	sorter := &labsort.Sorter{
		LenFunc:  func() int { return len(collection) },
		LessFunc: func(i, j int) bool { return collection[i] < collection[j] },
		SwapFunc: func(i, j int) { collection[i], collection[j] = collection[j], collection[i] },
	}
	sort.Sort(sorter)

	fmt.Println(collection)

	// Output:
	// [1ns 1µs 1ms 10s 2m0s]
}

func TestSorter(t *testing.T) {
	collection := []float64{4, 6.9, 19, 113.3, 0.3, 0.7}

	sorter := &labsort.Sorter{
		LenFunc:  func() int { return len(collection) },
		LessFunc: func(i, j int) bool { return collection[i] < collection[j] },
		SwapFunc: func(i, j int) { collection[i], collection[j] = collection[j], collection[i] },
	}
	sort.Sort(sorter)

	assert.EqualValues(t, []float64{0.3, 0.7, 4, 6.9, 19, 113.3}, collection)
}

var _ sort.Interface = &labsort.Sorter{}
