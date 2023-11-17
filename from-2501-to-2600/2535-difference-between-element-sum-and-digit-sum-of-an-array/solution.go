unc differenceOfSum(nums []int) int {
    getDigitSum := func(num int) int {
        ds := 0
        for num > 0 {
            ds += num % 10
            num /= 10
        }
        return ds
    }
    elemSum, digitSum := 0, 0
    for _, num := range nums {
        elemSum += num
        digitSum += getDigitSum(num)
    }

    if elemSum > digitSum {
        return elemSum - digitSum
    }
    return digitSum - elemSum
}