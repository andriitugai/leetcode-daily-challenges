func combinationSum3(k int, n int) [][]int {
    result := [][]int{}
    var helper func(target int, startNum int, path []int)
    helper = func(target int, startNum int, path []int){
        if len(path) == k && target == 0 {
            result = append(result, path)
            return
        }
        for i := startNum; i < 10 && i <= target ; i ++ {
            path = append(path, i)
            helper(target - i, i + 1, append([]int{}, path...))
            path = path[:len(path)-1]
        }
    }
    helper(n, 1, []int{})
    return result
}