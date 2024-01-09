# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def leafSimilar(self, root1: Optional[TreeNode], root2: Optional[TreeNode]) -> bool:
        def get_leaves(root) -> List[int]:
            leaves = []
            def dfs(node: TreeNode) -> None:
                nonlocal leaves
                if node.left is None and node.right is None:
                    leaves.append(node.val)
                    return
                if node.left:
                    dfs(node.left)
                if node.right:
                    dfs(node.right)

                return

            dfs(root)
            return leaves

        return get_leaves(root1) == get_leaves(root2)
        