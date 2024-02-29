# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def isEvenOddTree(self, root: Optional[TreeNode]) -> bool:
        def strictly_inc(vals: list[int]) -> bool:
            for i in range(1, len(vals)):
                if vals[i] <= vals[i-1]:
                    return False
            return True
        
        level = 0
        q = deque()
        q.append(root)
        
        while q:
            vals = []
            for _ in range(len(q)):
                node = q.popleft()
                vals.append(node.val)
                if node.left:
                    q.append(node.left)
                if node.right:
                    q.append(node.right)
                    
            if not strictly_inc(vals if level & 1 == 0 else vals[::-1]):
                return False
            
            cond = (lambda x: x % 2 == 1) if level & 1 else (lambda x: x % 2 == 0)
            if len(list(filter(cond, vals))) != 0:
                return False
            
            level += 1
            
        return True