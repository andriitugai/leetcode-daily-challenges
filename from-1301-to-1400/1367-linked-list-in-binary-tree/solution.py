# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def isSubPath(self, head: Optional[ListNode], root: Optional[TreeNode]) -> bool:
        sub_path = ""
        while head:
            sub_path += str(head.val) + ":"
            head = head.next

        def dfs(node, path):
            if sub_path in path:
                return True
            if not node:
                return False
            new_path = path + str(node.val) + ":"
            return dfs(node.left, new_path) or dfs(node.right, new_path)
        
        return dfs(root, "")