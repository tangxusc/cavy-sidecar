package unit

import "sync"

var entryMap = sync.Map{}

type LockEntry struct {
	mu  sync.Mutex
	Key interface{}
}

func (entry *LockEntry) Unlock() {
	//删除lock
	entryMap.Delete(entry.Key)
	entry.mu.Unlock()
}

func Of(key interface{}) *LockEntry {
	entry := &LockEntry{}
	actual, loaded := entryMap.LoadOrStore(key, entry)
	if loaded {
		return actual.(*LockEntry)
	}
	return entry
}

func (entry *LockEntry) Lock() {
	entry.mu.Lock()
}
