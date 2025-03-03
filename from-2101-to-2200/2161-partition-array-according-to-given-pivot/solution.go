func pivotArray(nums []int, pivot int) []int {
    lt, eq, gt := make([]int, 0), make([]int, 0), make([]int, 0)
    for _, num := range(nums) {
        if num < pivot {
            lt = append(lt, num)
        } else if num > pivot {
            gt = append(gt, num)
        } else {
            eq = append(eq, num)
        }
    }
    return append(lt, append(eq, gt...)...)
}