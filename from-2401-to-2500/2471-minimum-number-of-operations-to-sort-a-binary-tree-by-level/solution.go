/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func minimumOperations(root *TreeNode) int {
    var lvl []int
    var node *TreeNode
    result := 0
    q := []*TreeNode{root}

    for len(q) > 0 {
        n := len(q)
        lvl = []int{}
        for i := 0; i < n; i ++ {
            node = q[0]
            q = q[1:]
            lvl = append(lvl, node.Val)

            if node.Left != nil {
                q = append(q, node.Left)
            }
            if node.Right != nil {
                q = append(q, node.Right)
            }
        }
        fmt.Print(lvl)
        ms := minSwapsToSort(lvl)
        fmt.Println(ms)
        result += ms
    }
    return result
}

func minSwapsToSort(arr []int) int {
    srt := make([]int, len(arr))
    copy(srt, arr)
    sort.Ints(srt)

    idx := make(map[int]int)
    for i, el := range arr {
        idx[el] = i
    }
    result := 0
    for i := 0; i < len(srt); i ++ {
        if srt[i] != arr[i] {
            result += 1
            newPos := idx[srt[i]]
            idx[arr[i]] = newPos
            idx[arr[newPos]] = i
            arr[i], arr[newPos] = arr[newPos], arr[i]
        }
    }
    return result
}