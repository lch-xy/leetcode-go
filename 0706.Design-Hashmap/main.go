package main

type MyHashMap struct {
	list []int
}

func Constructor() MyHashMap {
	list := make([]int, 1000001)
	for i := range list {
		list[i] = -1
	}
	return MyHashMap{
		list: list,
	}
}

func (this *MyHashMap) Put(key int, value int) {
	this.list[key] = value
}

func (this *MyHashMap) Get(key int) int {
	return this.list[key]
}

func (this *MyHashMap) Remove(key int) {
	this.list[key] = -1
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
