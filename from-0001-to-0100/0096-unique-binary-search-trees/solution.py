class Solution:
    def numTrees(self, n: int) -> int:
        num_tree = [1] * (n + 1)

        for nodes in range(2, n+1):
            total = 0
            for root in range(1, nodes + 1):
                left = root - 1
                right = nodes - root
                total += num_tree[left] * num_tree[right]

            num_tree[nodes] = total

        return num_tree[-1]