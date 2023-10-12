# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def removeNthFromEnd(self, head, n):
        """
        :type head: ListNode
        :type n: int
        :rtype: ListNode
        """
        if head.val is None or head.next is None:
            return None
        
        to_remove = head
        pointer = head
        for i in range(n):
            if pointer.next is None and i == n-1:
                head = head.next
                return head
            pointer = pointer.next
            
        while pointer.next:
            pointer = pointer.next
            to_remove = to_remove.next
            
        # remove the node under to_remove
        # if to_remove.next is not None:
        to_remove.next = to_remove.next.next
            
        # to_remove = None
            
        return head