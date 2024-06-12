func sortColors(nums []int)  {
    count := make(map[int]int)
    for _, num := range nums {
        count[num] += 1
    }
    i := 0
    for j := 0; j < 3; j ++ {
        for k := 0; k < count[j]; k ++ {
            nums[i] = j
            i += 1
        }
    }
}