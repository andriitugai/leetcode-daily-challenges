# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def swapPairs(self, head: Optional[ListNode]) -> Optional[ListNode]:
        def swap(node):
            if not node or not node.next:
                return node
            
            new_head = node.next
            
            tail = new_head.next
            new_head.next = node
            node.next = swap(tail)
            
            return new_head
        
        return swap(head)