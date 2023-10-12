/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

 func findInMountainArray(target int, mountainArr *MountainArray) int {
    pIdx := findPeakIndex(mountainArr)
    maLen := mountainArr.length()

    lIdx := findOnLeftSide(target, 0, pIdx, mountainArr)
    if lIdx >= 0 {
        return lIdx
    }
    return findOnRightSide(target, pIdx, maLen - 1, mountainArr)
}

func findPeakIndex(ma *MountainArray) int {
    left, right := 0, ma.length() - 1
    for left < right {
        mid := left + (right - left) / 2
        if ma.get(mid) > ma.get(mid+1) {
            right = mid
        } else {
            left = mid + 1
        }
    }
    return left
}

func findOnLeftSide(target, left, right int, ma *MountainArray) int {
    for left <= right {
        mid := left + (right - left) / 2
        midItem := ma.get(mid)
        if midItem == target {
            return mid
        } else if midItem > target {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return -1
}

func findOnRightSide(target, left, right int, ma *MountainArray) int {
    for left <= right {
        mid := left + (right - left) / 2
        midItem := ma.get(mid)
        if midItem == target {
            return mid
        } else if midItem < target {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return -1
}