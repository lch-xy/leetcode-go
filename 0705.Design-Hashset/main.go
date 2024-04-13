package main

type MyHashSet struct {
	cnt   int
	cache map[int]struct{}
}

func Constructor() MyHashSet {
	return MyHashSet{
		cnt:   0,
		cache: make(map[int]struct{}),
	}
}

func (this *MyHashSet) Add(key int) {
	this.cache[key] = struct{}{}
}

func (this *MyHashSet) Remove(key int) {
	delete(this.cache, key)
}

func (this *MyHashSet) Contains(key int) bool {
	if _, ok := this.cache[key]; ok {
		return true
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
