package core

import (
	"runtime"
	"sync"
)

var shardCount = runtime.NumCPU() * 8

// a sharded map to allow for sparse, parallel computation
//
// each shard supports several readers or one writer
type sparseCellGrid[T comparable] struct {
	shards []map[uint64]*Cell[T]
	mu     []sync.RWMutex
}

func NewSparseCellGrid[T comparable]() *sparseCellGrid[T] {
	sm := &sparseCellGrid[T]{
		shards: make([]map[uint64]*Cell[T], shardCount),
		mu:     make([]sync.RWMutex, shardCount),
	}
	for i := 0; i < shardCount; i++ {
		sm.shards[i] = make(map[uint64]*Cell[T])
	}
	return sm
}

// getShard returns the index of the shard for the given key
/*
func (sm *sparseCellGrid[T]) getShard(key uint32) int {
	return int(key % uint32(shardCount))
}
*/

func (sm *sparseCellGrid[T]) getShard(key uint64) int {
	    // Use the higher bits for better distribution
		    return int((key >> 56) % uint64(shardCount))
		}

// Set sets a value in the sharded map
func (sm *sparseCellGrid[T]) Set(key uint64, value *Cell[T]) {
	shard := sm.getShard(key)
	sm.mu[shard].Lock()
	defer sm.mu[shard].Unlock()
	sm.shards[shard][key] = value
}

// delete a key from the sharded map
func (sm *sparseCellGrid[T]) Delete(key uint64) {
	shard := sm.getShard(key)
	sm.mu[shard].Lock()
	defer sm.mu[shard].Unlock()
	delete(sm.shards[shard], key)
}

// Get gets a value from the sharded map
//
// if ok == false, then the vlaue is not there
func (sm *sparseCellGrid[T]) Get(key uint64) (*Cell[T], bool) {
	shard := sm.getShard(key)
	sm.mu[shard].RLock()
	defer sm.mu[shard].RUnlock()
	value, ok := sm.shards[shard][key]
	return value, ok
}

// GetAllKeys returns a slice of all keys in the sharded map
func (sm *sparseCellGrid[T]) GetAllKeys() []uint64 {
	var keys []uint64

	for i := 0; i < shardCount; i++ {
		sm.mu[i].RLock()
		shard := sm.shards[i]
		sm.mu[i].RUnlock()

		for key := range shard {
			keys = append(keys, key)
		}
	}

	return keys
}

// call a function in parallel on the shards
func (sm *sparseCellGrid[T]) ProcessShard(f func(shardNum int, shard map[uint64]*Cell[T])) {
	var wg sync.WaitGroup
	wg.Add(shardCount)

	for i := 0; i < shardCount; i++ {
		go func(shard int) {
			defer wg.Done()
			sm.mu[shard].RLock()
			defer sm.mu[shard].RUnlock()
			f(shard, sm.shards[shard])
		}(i)
	}
	wg.Wait()
}
