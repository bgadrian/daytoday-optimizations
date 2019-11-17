package batchread

// SetCache serve the same function as the global GetRandValue functionality
// but internally keeps a reference to a specific Category/Subcategory list
// improving consecutive calls performance (no more runtime.mapaccess1_faststr)
type SetCache interface {
	GetRandValue() string
}

// NewDataListCache creates a new reference to a Data category set
// for faster access
func NewDataListCache(f Intn, category, subcatgory string) (SetCache, error) {
	res := &simpleList{}
	res.randomizer = f
	res.db = Data[category][subcatgory]

	return res, nil
}

type simpleList struct {
	randomizer Intn
	db         []string
}

func (w *simpleList) GetRandValue() string {
	return w.db[w.randomizer.Intn(len(w.db))]
}
