type Solution struct {
    idxMap map[int][]int
}


func Constructor(nums []int) Solution {
    idxMap := make(map[int][]int)
    for idx, num := range nums {
        if _, ok := idxMap[num]; !ok {
            idxMap[num] = []int{}
        }
        idxMap[num] = append(idxMap[num], idx)
    }
    return Solution{idxMap: idxMap}
}


func (this *Solution) Pick(target int) int {
    indexes := this.idxMap[target]
    ri := rand.Intn(len(indexes))
    return indexes[ri]
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */