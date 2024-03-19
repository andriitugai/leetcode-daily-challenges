/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 type CBTInserter struct {
    Root *TreeNode
    prevRow []*TreeNode
}


func Constructor(root *TreeNode) CBTInserter {
    var prevRow, lastRow []*TreeNode
    q := []*TreeNode{root}
    prevRow = []*TreeNode{root}

    for len(q) > 0 {
        n := len(q)
        lastRow = []*TreeNode{}
        for i := 0; i < n; i ++ {
            node := q[0]
            q = q[1:]
            if node.Left != nil {
                // q = append(q, node.Left)
                lastRow = append(lastRow, node.Left)
            } else {
                // this is the last row
                break
            }
            if node.Right != nil {
                // q = append(q, node.Right)
                lastRow = append(lastRow, node.Right)
            } else {
                break
            }
        }
        q = append(q, lastRow...)
        // if we have the last row full
        if len(lastRow) == 2 * len(prevRow) {
            prevRow = lastRow
        } else {
            break
        }
    }
    return CBTInserter{Root: root, prevRow: prevRow}
}


func (this *CBTInserter) Insert(val int) int {
    lastRow := []*TreeNode{}
    // look for empty child in the prevRow
    for i := 0; i < len(this.prevRow); i ++ {
        if this.prevRow[i].Left == nil {
            this.prevRow[i].Left = &TreeNode{Val: val}
            return this.prevRow[i].Val
        } else {
            lastRow = append(lastRow, this.prevRow[i].Left)
        }
        if this.prevRow[i].Right == nil {
            this.prevRow[i].Right = &TreeNode{Val: val}
            return this.prevRow[i].Val
        } else {
            lastRow = append(lastRow, this.prevRow[i].Right)
        }
    }
    // the lastRow is full, so make it prevRow
    this.prevRow = lastRow
    this.prevRow[0].Left = &TreeNode{Val: val}
    return this.prevRow[0].Val
}


func (this *CBTInserter) Get_root() *TreeNode {
    return this.Root
}


/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(val);
 * param_2 := obj.Get_root();
 */