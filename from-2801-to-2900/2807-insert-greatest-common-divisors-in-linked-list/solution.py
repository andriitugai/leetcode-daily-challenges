# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def insertGreatestCommonDivisors(self, head: Optional[ListNode]) -> Optional[ListNode]:
        def gcd(a: int, b: int) -> int:
            while a != b:
                if a > b:
                    a -= b
                else:
                    b -= a
            return a

        p = head
        while p.next:
            newNode = ListNode(val=gcd(p.val, p.next.val), next=p.next)
            p.next = newNode
            p = p.next.next
        return head