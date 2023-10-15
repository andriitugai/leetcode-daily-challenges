# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def deleteDuplicates(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if not head:
            return None
        p = head
        pn = p.next
        
        while pn:
            if p.val == pn.val:
                p.next = pn.next
            else:
                p = p.next
            
            pn = pn.next
                
        return head