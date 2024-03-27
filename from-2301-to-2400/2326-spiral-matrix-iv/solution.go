/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func spiralMatrix(m int, n int, head *ListNode) [][]int {
    result := make([][]int, m)
    for i := 0; i < m; i ++ {
        result[i] = make([]int, n)
        for j := 0; j < n; j ++ {
            result[i][j] = -1
        }
    }

    top, bottom := 0, m
    left, right := 0, n
    p := head

    for left < right || top < bottom {
        for col := left; col < right; col ++ {
            result[top][col] = p.Val
            p = p.Next
            if p == nil {
                return result
            }
        }
        top += 1
        for row := top; row < bottom; row ++ {
            result[row][right - 1] = p.Val
            p = p.Next
            if p == nil {
                return result
            }
        }
        right -= 1
        for col := right - 1; col >= left; col -- {
            result[bottom - 1][col] = p.Val
            p = p.Next
            if p == nil {
                return result
            }
        }
        bottom -= 1
        for row := bottom - 1; row >= top; row -- {
            result[row][left] = p.Val
            p = p.Next
            if p == nil {
                return result
            }
        }
        left += 1
    }
    return result
}