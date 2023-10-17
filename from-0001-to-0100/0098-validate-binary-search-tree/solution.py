# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def isValidBST(self, root: Optional[TreeNode]) -> bool:
        
        def is_valid(node, left_limit, right_limit):
            
            if not node: 
                return True
            
            if node.val <= left_limit or node.val >= right_limit:
                return False
            
            return (
                is_valid(node.left, left_limit, node.val) and 
                is_valid(node.right, node.val, right_limit)
            )
        
        return is_valid(root, float("-inf"), float("inf"))