# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def rotateRight(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        if not head:
            return None
        
        if k == 0:
            return head
        
        pointer = head
        n = 1
        while pointer.next:
            pointer = pointer.next
            n += 1
            
        last_elem = pointer
        length = n
            
        k = k % length
        if k == 0:
            return head
        
        n = 1
        pointer = head
                
        while n < length - k:            
            pointer = pointer.next
            n += 1
            
        prev = pointer
        pointer = pointer.next
            
        # pointer now is the head of the rotated list
        last_elem.next = head
        prev.next = None
        
        return pointer   