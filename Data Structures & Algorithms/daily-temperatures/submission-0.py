from typing import List

class Stack:
    def __init__(self):
        self.stack = []

    def push(self, val):
        self.stack.append(val)

    def pop(self):
        if not self.isEmpty():
            return self.stack.pop()
        return "Stack is empty"

    def peek(self):
        if not self.isEmpty():
            return self.stack[-1]
        return "Stack is empty"

    def isEmpty(self):
        return len(self.stack) == 0

    def size(self):
        return len(self.stack)


class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        stack = Stack()
        res = [0] * len(temperatures)
        
        for num in range(len(temperatures)):
            while not stack.isEmpty() and temperatures[stack.peek()] < temperatures[num]:
                i = stack.pop()
                res[i] = num - i 
            
            stack.push(num)
            
        return res
