# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def zigzagLevelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        result = []
        if not root:
            return result

        q = deque()
        q.append(root)
        dir = 1
        while q:
            level = []
            for _ in range(len(q)):
                node = q.popleft()            
                level.append(node.val)
                if node.right:
                    q.append(node.right)
                if node.left:
                    q.append(node.left)
                
            dir *= -1
            result.append(level[::dir])

        return result