# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def isSameTree(self, p: Optional[TreeNode], q: Optional[TreeNode]) -> bool:
#         def bfs(root):
#             result = []
#             stack = []
#             stack.append(root)
            
#             while stack:
#                 for _ in range(len(stack)):
#                     node = stack.pop()
#                     if not node:
#                         result.append(None)
#                     else:
#                         result.append(node.val)
#                         stack.append(node.right)
#                         stack.append(node.left)
                        
#             return result
        
#         return bfs(p) == bfs(q)

        if not p and not q:
            return True
        if not p or not q or p.val != q.val:
            return False
        
        return self.isSameTree(p.left, q.left) and self.isSameTree(p.right, q.right)
    