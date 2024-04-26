package main

import "fmt"

// we can store change record instead of intact array
type SnapshotArray struct {
	cache     map[int][][]int
	curSnapId int
	lenght    int
}

func main() {
	// ["SnapshotArray","set","snap","set","get"]
	// [[3],[0,5],[],[0,6],[0,0]]
	cur := Constructor(3)
	cur.Set(0, 5)
	cur.Snap()
	cur.Set(0, 6)
	fmt.Printf("cur.Get(0, 0): %v\n", cur.Get(0, 0))
}

func Constructor(length int) SnapshotArray {
	cache := make(map[int][][]int)
	return SnapshotArray{
		cache:     cache,
		curSnapId: 0,
		lenght:    length,
	}
}

func (this *SnapshotArray) Set(index int, val int) {
	curList := this.cache[index]
	if len(curList) != 0 && curList[len(curList)-1][0] == this.curSnapId {
		curList[len(curList)-1] = []int{this.curSnapId, val}
	} else {
		curList = append(curList, []int{this.curSnapId, val})
	}
	this.cache[index] = curList
}

func (this *SnapshotArray) Snap() int {
	this.curSnapId++
	return this.curSnapId - 1
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	left, right := 0, len(this.cache[index])
	for left < right {
		mid := left + (right-left)/2
		if this.cache[index][mid][0] > snap_id {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if left == 0 {
		return 0
	}
	return this.cache[index][left-1][1]
}

// type SnapshotArray struct {
// 	snapIndex  int
// 	curArray   []int
// 	snapArrays [][]int
// }

// func Constructor(length int) SnapshotArray {
// 	return SnapshotArray{
// 		snapIndex:  0,
// 		curArray:   make([]int, length),
// 		snapArrays: make([][]int, 0),
// 	}
// }

// func (this *SnapshotArray) Set(index int, val int) {
// 	this.curArray[index] = val
// }

// func (this *SnapshotArray) Snap() int {
// 	newArray := make([]int, len(this.curArray))
// 	copy(newArray, this.curArray)
// 	this.snapArrays = append(this.snapArrays, newArray)
// 	this.snapIndex++
// 	return this.snapIndex - 1
// }

// func (this *SnapshotArray) Get(index int, snap_id int) int {
// 	return this.snapArrays[snap_id][index]
// }
