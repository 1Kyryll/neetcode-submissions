# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        # stack - 0(N), O(N)
        
        dummy = ListNode(0, head)
        
        stack = []
        curr = dummy
        
        while curr:
            stack.append(curr)
            curr = curr.next
            
        for _ in range(n): 
            stack.pop() 
        prev = stack.pop()
        prev.next = prev.next.next
        
        return dummy.next