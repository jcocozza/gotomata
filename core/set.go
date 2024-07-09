package core

// Not quite a set
//
// will only add an element if already does not exist
type CellSet[T comparable] map[uint64]*Cell[T]

func (s CellSet[T]) Add(k uint64, v *Cell[T]) {
    if _, exists := s[k]; !exists {
       s[k] = v
    }
}
