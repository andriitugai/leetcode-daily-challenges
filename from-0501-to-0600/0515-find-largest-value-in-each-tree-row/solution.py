# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def largestValues(self, root: Optional[TreeNode]) -> List[int]:
        q = deque()
        result = []
        if root:
            q.append(root)

        while q:
            n = len(q)
            rowMax = float('-inf')
            for _ in range(n):
                node = q.popleft()
                rowMax = max(rowMax, node.val)

                if node.left:
                    q.append(node.left)
                if node.right:
                    q.append(node.right)

            result.append(rowMax)
        return result