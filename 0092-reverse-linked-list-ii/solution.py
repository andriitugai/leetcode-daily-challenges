# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reverseBetween(self, head: Optional[ListNode], left: int, right: int) -> Optional[ListNode]:
        p_slow, p_fast = head, head
        pos = 1
        
        while pos < left:
            p_slow = p_slow.next
            p_fast = p_fast.next
            pos += 1
            
        stack = []
        while pos <= right:
            stack.append(p_fast.val)
            p_fast = p_fast.next
            pos += 1
            
        while stack:
            p_slow.val = stack.pop()
            p_slow = p_slow.next
            
        return head