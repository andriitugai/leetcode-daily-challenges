# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def hasCycle(self, head: Optional[ListNode]) -> bool:
        if not head or not head.next:
            return False
        
        p_slow = head
        p_fast = head
        
        while p_fast.next:
            p_fast = p_fast.next
            if p_fast == p_slow:
                return True
            if p_fast.next:
                p_fast = p_fast.next
            else:
                return False            
            if p_fast == p_slow:
                return True
            p_slow = p_slow.next
            
        return False