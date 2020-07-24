package common


import (
	"sync"
)

const (
	maxWorlds   = 500
	worldsCount = 10000
)

var worldPool = &sync.Pool{
	New: func() interface{} {
		return new(World)
	},
}

var worldsPool = &sync.Pool{
	New: func() interface{} {
		return &Worlds{
			W: make([]World, 0, maxWorlds),
		}
	},
}

func AcquireWorld() *World {
	return worldPool.Get().(*World)
}

func ReleaseWorld(w *World) {
	w.ID = 0
	w.RandomNumber = 0
	worldPool.Put(w)
}

func AcquireWorlds() *Worlds {
	return worldsPool.Get().(*Worlds)
}

func ReleaseWorlds(w *Worlds) {
	w.W = w.W[:0]
	worldsPool.Put(w)
}
