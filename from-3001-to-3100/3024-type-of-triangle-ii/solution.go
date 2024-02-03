func triangleType(nums []int) string {
    a, b, c := nums[0], nums[1], nums[2]
    
    if a + b <= c || b + c <= a || a + c <= b {
        return "none"
    } else if a == b && b == c {
        return "equilateral"
    } else if a == b || a == c || b == c {
        return "isosceles"
    }
    return "scalene"
}