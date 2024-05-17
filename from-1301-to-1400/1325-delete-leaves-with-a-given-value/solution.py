# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def removeLeafNodes(self, root: Optional[TreeNode], target: int) -> Optional[TreeNode]:
        def dfs(node: TreeNode, tgt: int)-> bool:
            if not node:
                return True
    
            if dfs(node.left, tgt):
                node.left = None
            if dfs(node.right, tgt):
                node.right = None
            if node.val == tgt and not node.left and not node.right:
                return True

            return False

        dfs(root, target)
        if not root.left and not root.right and root.val == target:
            return None

        return root