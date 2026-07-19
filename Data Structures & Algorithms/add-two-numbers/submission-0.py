# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        cur = l1
        num1 = ""
        while cur: 
            num1 = str(cur.val) + num1 
            cur = cur.next
        
        cur = l2
        num2 = ""
        while cur: 
            num2 = str(cur.val) + num2
            cur = cur.next
            
        sum = int(num1) + int(num2)
        dummy = ListNode()   
        tail = dummy         

        for v in reversed(str(sum)):
            tail.next = ListNode(int(v))   
            tail = tail.next               

        return dummy.next    