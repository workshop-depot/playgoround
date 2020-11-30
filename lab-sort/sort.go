package labsort

// Sorter provides a reusable implementation of sort.Interface interface.
// There is not much value in avoiding having a type-specialized implementation for sortable collections.
// Sorter just demonstrates that the closures can provide half of the functionality of
// the parametric polymorphism (generics) in some cases.
// Some examples are provided which shows how to use Sorter for sorting slices of bytes, strings and time.Durations.
type Sorter struct {
	LenFunc  func() int
	LessFunc func(i, j int) bool
	SwapFunc func(i, j int)
}

func (sorter *Sorter) Len() int           { return sorter.LenFunc() }
func (sorter *Sorter) Less(i, j int) bool { return sorter.LessFunc(i, j) }
func (sorter *Sorter) Swap(i, j int)      { sorter.SwapFunc(i, j) }
