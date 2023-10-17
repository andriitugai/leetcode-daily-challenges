# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        
        from heapq import heappop, heappush
        heap = []
        for lst in lists:
            while lst:
                heappush(heap, lst.val)
                lst = lst.next
           
        head = ListNode(None)  # Dummy node
        p = head
        while heap:
            p.next = ListNode(heappop(heap))
            p = p.next
            
        return head.next