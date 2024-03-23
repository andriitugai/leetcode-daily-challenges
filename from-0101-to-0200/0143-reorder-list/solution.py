# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reorderList(self, head: Optional[ListNode]) -> None:
        """
        Do not return anything, modify head in-place instead.
        """
        nodes = []
        p = head.next
        while p:
            nodes.append(p)
            p = p.next
            
        left, right = 0, len(nodes)-1
        p = head
        
        while left <= right:
            p.next = nodes[right]
            p = p.next
            right -= 1
            p.next = nodes[left]
            p = p.next
            left += 1
            
        p.next = None
        