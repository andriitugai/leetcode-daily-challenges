# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def sortedArrayToBST(self, nums: List[int]) -> Optional[TreeNode]:
        if not nums:
            return None
        nums_len = len(nums)
        if nums_len == 1:
            return TreeNode(val=nums[0])
        
        mid = nums_len // 2
        result = TreeNode(val=nums[mid])
        result.left = self.sortedArrayToBST(nums[:mid])
        result.right = self.sortedArrayToBST(nums[mid+1:])
        
        return result