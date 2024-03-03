func resultArray(nums []int) []int {
    la1, la2 := nums[0], nums[1]
    arr1 := []int{la1}
    arr2 := []int{la2}
    for i := 2; i < len(nums); i ++ {
        if la1 > la2 {
            la1 = nums[i]
            arr1 = append(arr1, la1)
        } else {
            la2 = nums[i]
            arr2 = append(arr2, la2)
        }
    }
    return append(arr1, arr2...)
}