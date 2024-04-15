# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def sumNumbers(self, root: Optional[TreeNode]) -> int:
        result = 0

        def dfs(node: TreeNode, curr: int) -> None:
            if not node:
                return

            nonlocal result
            curr = curr * 10 + node.val
            if not node.left and not node.right:
                result += curr
                return

            dfs(node.left, curr)
            dfs(node.right, curr)
            return

        dfs(root, 0)
        return result
        