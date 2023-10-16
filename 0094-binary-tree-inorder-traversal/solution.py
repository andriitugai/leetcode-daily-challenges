# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def inorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
#         res = []
#         curr = root
#         stack = []
        
#         while curr or stack:
#             while curr:
#                 stack.append(curr)
#                 curr = curr.left
                
#             curr = stack.pop()
#             res.append(curr.val)
#             curr = curr.right
            
#         return res

        result = []
        def inorder(node):
            nonlocal result
            if not node:
                return
            
            inorder(node.left)
            result.append(node.val)
            inorder(node.right)
            
        inorder(root)
        
        return result