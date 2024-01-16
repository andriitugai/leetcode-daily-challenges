type RandomizedSet struct {
    set map[int]bool
    vals []int
}


func Constructor() RandomizedSet {
    set := make(map[int]bool)
    vals := []int{}
    return RandomizedSet{set, vals}
}


func (this *RandomizedSet) Insert(val int) bool {
    if _, ok := this.set[val]; ok {
        return false
    }
    this.set[val] = true
    this.vals = append(this.vals, val)
    return true
}


func (this *RandomizedSet) Remove(val int) bool {
    if _, ok := this.set[val]; !ok {
        return false
    }
    delete(this.set, val)
    valIdx := -1
    for idx, num := range this.vals {
        if num == val {
            valIdx = idx
            break
        }
    }
    this.vals[valIdx] = this.vals[len(this.vals) - 1]
    this.vals = this.vals[:len(this.vals) - 1]
    return true
}


func (this *RandomizedSet) GetRandom() int {
    rndIdx := rand.Intn(len(this.vals))
    return this.vals[rndIdx]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */