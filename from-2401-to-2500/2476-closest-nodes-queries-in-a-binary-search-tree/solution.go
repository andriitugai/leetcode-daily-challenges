/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func closestNodes(root *TreeNode, queries []int) [][]int {
    result := [][]int{}
    nums := []int{}
    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode) {
        if node != nil {
            dfs(node.Left)
            nums = append(nums, node.Val)
            dfs(node.Right)
        }
    }
    dfs(root)
    n := len(nums)

    var binarySearch func(num int) []int
    binarySearch = func(num int) []int {
        left, right := 0, n - 1
        var mid int
        s, e := -1, -1
        for left <= right {
            mid = (left + right) / 2
            if nums[mid] == num {
                return []int{num, num}
            }

            if nums[mid] < num {
                s = nums[mid]
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
        left, right = 0, n - 1
        for left <= right {
            mid = (left + right) / 2

            if nums[mid] > num {
                e = nums[mid]
                right = mid -1
            } else {
                left = mid + 1
            }
        }
        return []int{s, e}
    }

    for _, num := range queries {
        result = append(result, binarySearch(num))
    }
    return result
}

