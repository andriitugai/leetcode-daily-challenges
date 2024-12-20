# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def reverseOddLevels(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        q = [root]
        level = 0
        while len(q) > 0:
            n = len(q)
            if level % 2:
                left, right = 0, n - 1
                while left < right:
                    q[left].val, q[right].val = q[right].val, q[left].val
                    left += 1
                    right -= 1

            for i in range(n):
                if q[i].left:
                    q.append(q[i].left)
                if q[i].right:
                    q.append(q[i].right)
            q = q[n:]
            level += 1
        return root
            
