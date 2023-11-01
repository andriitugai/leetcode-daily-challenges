# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def findMode(self, root: Optional[TreeNode]) -> List[int]:
        c = defaultdict(int)
        def dfs(node):
            if not node:
                return
            nonlocal c
            c[node.val] += 1
            dfs(node.left)
            dfs(node.right)
            
        dfs(root)
        
        counts = sorted(c.items(), key=lambda x: x[1], reverse=True)
        res = []
        mode = counts[0][1]
        p = 0
        while p < len(counts) and counts[p][1] == mode:
            res.append(counts[p][0])
            p += 1
            
        return res
            