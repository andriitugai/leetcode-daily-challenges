# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def mergeTwoLists(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        if not l1:
            return l2
        if not l2:
            return l1
        
        lm = ListNode()
        p = lm
        while True:
            if l1.val < l2.val:
                p.val = l1.val
                l1 = l1.next
            else:
                p.val = l2.val
                l2 = l2.next
                
            if not l1 or not l2:
                break
                
            p.next = ListNode()
            p = p.next
            
        if not l1:
            p.next = l2
        if not l2:
            p.next = l1
            
        return lm