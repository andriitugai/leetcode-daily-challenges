# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reverseList(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        if not head or not head.next:
            return head
        
        new_head = head
        while head.next:
            neck = head.next
            head.next = neck.next
            neck.next = new_head
            new_head = neck
            
        return new_head