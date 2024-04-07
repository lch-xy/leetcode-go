package main

type ThroneInheritance struct {
	cache    map[string][]string
	delKey   map[string]struct{}
	kingName string
}

// use adjacency list to save children
// use map to save delete name
func Constructor(kingName string) ThroneInheritance {
	t := ThroneInheritance{
		cache:    make(map[string][]string),
		delKey:   make(map[string]struct{}),
		kingName: kingName,
	}
	return t
}

func (this *ThroneInheritance) Birth(parentName string, childName string) {
	this.cache[parentName] = append(this.cache[parentName], childName)
}

func (this *ThroneInheritance) Death(name string) {
	this.delKey[name] = struct{}{}
}

func (this *ThroneInheritance) GetInheritanceOrder() []string {
	res := make([]string, 0)
	this.dfs(this.kingName, &res)
	return res
}

// use dfs to search all children
func (this *ThroneInheritance) dfs(s string, res *[]string) {
	if _, ok := this.delKey[s]; !ok {
		*res = append(*res, s)
	}

	if _, ok := this.cache[s]; ok {
		for _, v := range this.cache[s] {
			this.dfs(v, res)
		}
	}
}
