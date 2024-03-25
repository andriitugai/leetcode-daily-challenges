/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 type Codec struct {
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    if root == nil {
        return "$,"
    }
    return strconv.Itoa(root.Val) + "," + this.serialize(root.Left) + this.serialize(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
    i, a := 0, strings.Split(data[:len(data) - 1], ",")
    return preorder(&i, a)
}

func preorder(i *int, a []string) *TreeNode {
    data := a[(*i)]
    (*i)++
    if data == "$" { return nil }
    val, _ := strconv.Atoi(data)
    return &TreeNode{Val: val, Left: preorder(i, a), Right: preorder(i, a)}
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */