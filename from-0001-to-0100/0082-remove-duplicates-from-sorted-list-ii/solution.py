# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def deleteDuplicates(self, head: Optional[ListNode]) -> Optional[ListNode]:
        prev = None
        curr = head
        
        while curr and curr.next:
            if curr.val != curr.next.val:
                prev = curr
                curr = curr.next
            else:
                while curr.next and curr.val == curr.next.val:
                    curr = curr.next
                curr = curr.next
                if prev:
                    prev.next = curr
                else:
                    head = curr
                    
        return head