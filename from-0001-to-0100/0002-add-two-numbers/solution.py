# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        p1 = l1
        p2 = l2
        ls = ListNode()
        ps = ls
        tr = 0
        
        while True:
            summ = p1.val + p2.val + tr
            tr = 0 if summ < 10 else 1
            ps.val = summ - tr * 10
            p1 = p1.next
            p2 = p2.next
            if p1 and p2:
                ps.next = ListNode()
                ps = ps.next
            else:
                break
                
        while p1:
            ps.next = ListNode()
            ps = ps.next
            summ = p1.val + tr
            tr = 0 if summ < 10 else 1
            ps.val = summ - tr * 10
            p1 = p1.next
            
        while p2:
            ps.next = ListNode()
            ps = ps.next
            summ = p2.val + tr
            tr = 0 if summ < 10 else 1
            ps.val = summ - tr * 10
            p2 = p2.next
                
                
                
        if tr:
            ps.next = ListNode()
            ps.next.val = tr
                    
        return ls
            