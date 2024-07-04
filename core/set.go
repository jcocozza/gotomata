package core

// Not quite a set
//
// will only add an element if already does not exist
type CellSet[T comparable] map[uint32]*Cell[T]

func (s CellSet[T]) Add(k uint32, v *Cell[T]) {
    if _, exists := s[k]; !exists {
       s[k] = v
    }
}
