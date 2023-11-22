func findDiagonalOrder(nums [][]int) []int {
    rankToNum := make(map[int]int)
    ranks := []int{}
    var curRank int

    for r := 0; r < len(nums); r ++ {
        for c := 0; c < len(nums[r]); c++ {
            curRank = (r + c) * 200000 + c
            ranks = append(ranks, curRank)
            rankToNum[curRank] = nums[r][c]
        }
    }
    sort.Ints(ranks)
    
    result := make([]int, 0, len(ranks))
    for i := 0; i < len(ranks); i ++ {
        result = append(result, rankToNum[ranks[i]])
    }
    return result
}