# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    def reorderList(self, head: Optional[ListNode]) -> None:
        stack = []
        curr = head
        
        while curr: 
            stack.append(curr)
            curr = curr.next 
        
        curr = head
        n = len(stack)
            
        for _ in range(n // 2):
            nxt = curr.next 
            back = stack.pop()
            
            curr.next = back
            back.next = nxt

            curr = nxt 
                
        curr.next = None 
        