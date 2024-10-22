# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def kthLargestLevelSum(self, root: Optional[TreeNode], k: int) -> int:
        q = deque([root])
        sums = []
        while q:
            n = len(q)
            levelSum = 0
            for i in range(n):
                node = q.popleft()
                levelSum += node.val
                if node.left:
                    q.append(node.left)
                if node.right:
                    q.append(node.right)
            sums.append(levelSum)
        if len(sums) < k:
            return -1

        sums.sort(reverse=True)
        return sums[k-1]
            