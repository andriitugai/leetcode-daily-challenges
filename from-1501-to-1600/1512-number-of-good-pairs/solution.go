func numIdenticalPairs(nums []int) int {
    cnt := make(map[int]int)
    for _, num := range nums {
        cnt[num] += 1
    }

    ans := 0
    for _, v := range cnt {
        ans += (v * (v - 1) / 2)
    }
    return ans
}