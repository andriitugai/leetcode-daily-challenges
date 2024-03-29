func countSubarrays(nums []int, k int) int64 {
    mx := 0
    mxCnt := 0
    for _, num := range nums {
        if num > mx {
            mx = num
            mxCnt = 1
        } else if num == mx {
            mxCnt += 1
        }
    }
    if mxCnt < k {
        return 0
    }
    var result int64 = 0
    cnt, left := 0, 0
    for right := 0; right < len(nums); right ++ {
        if nums[right] == mx {
            cnt += 1
        }
        for cnt == k {
            if nums[left] == mx {
                cnt -= 1
            }
            left += 1
        }
        result += int64(left)
    }
    return result
}