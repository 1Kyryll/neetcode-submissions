# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode], carry: int = 0) -> Optional[ListNode]:
        if not l1 and not l2 and carry == 0:
            return None

        v1 = l1.val if l1 else 0
        v2 = l2.val if l2 else 0

        total = v1 + v2 + carry
        digit = total % 10        
        new_carry = total // 10   

        node = ListNode(digit)
        node.next = self.addTwoNumbers(
            l1.next if l1 else None,
            l2.next if l2 else None,
            new_carry
        )
        return node