# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def pathSum(self, root: Optional[TreeNode], targetSum: int) -> int:
        # Hashmap to store prefix sums and their counts
        prefix_sum_count = {0: 1}  # Base case: one way to have a sum of 0 (no nodes)

        def dfs(node, current_sum) -> int:
            if node is None:
                return 0
            
            current_sum += node.val
            # Count the number of valid paths ending at the current node
            # Our found path = targetSum = current_sum - prefix_sum
            num_paths = prefix_sum_count.get(current_sum - targetSum, 0)

            # Update the hashmap with the current cumulative sum
            prefix_sum_count[current_sum] = prefix_sum_count.get(current_sum, 0) + 1

            # Recurse down to the left and right children
            num_paths += dfs(node.left, current_sum)
            num_paths += dfs(node.right, current_sum)
            
            # Backtrack - remove the current cumulative sum from the hashmap
            prefix_sum_count[current_sum] -= 1
            
            return num_paths
        
        return dfs(root, 0)