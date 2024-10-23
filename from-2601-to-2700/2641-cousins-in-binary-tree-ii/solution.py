# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def replaceValueInTree(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        levelSums = defaultdict(int)
        lvl = 0
        q = deque([root])

        while q:
            n = len(q)
            
            for i in range(n):
                node = q.popleft()
                levelSums[lvl] += node.val
                if node.left:
                    q.append(node.left)
                if node.right:
                    q.append(node.right)

            lvl += 1

        def dfs(node: TreeNode, level: int, siblingsSum: int) -> None:
            if not node:
                return
            node.val = levelSums[level] - siblingsSum
            ss = 0
            if node.left:
                ss += node.left.val
            if node.right:
                ss += node.right.val

            dfs(node.left, level + 1, ss)
            dfs(node.right, level + 1, ss)

        dfs(root, 0, root.val)
        return root
        