# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def addOneRow(self, root: Optional[TreeNode], val: int, depth: int) -> Optional[TreeNode]:
        if depth == 1:
            return TreeNode(val=val, left=root)
            
        def dfs(node: TreeNode, cur_depth: int) -> None:
            nonlocal depth, val
            if not node or cur_depth > depth:
                return
            
            if cur_depth == depth - 1:
                node.left = TreeNode(val=val, left=node.left)
                node.right = TreeNode(val=val, right=node.right)
                return
                
            dfs(node.left, cur_depth + 1)
            dfs(node.right, cur_depth + 1)
            
        dfs(root, 1)
        
        return root
    
#         if depth == 1:
#                 return TreeNode(val=val, left=root)

#         from queue import SimpleQueue
#         q = SimpleQueue()
#         level = 1
#         q.put(root)
#         while level < depth-1:
#             for _ in range(q.qsize()):
#                 node = q.get()
#                 if node.left:
#                     q.put(node.left)
#                 if node.right:
#                     q.put(node.right)
#             level += 1

#         while not q.empty():
#             node = q.get()
#             node.left = TreeNode(val=val, left=node.left)
#             node.right = TreeNode(val=val, right=node.right)

#         return root