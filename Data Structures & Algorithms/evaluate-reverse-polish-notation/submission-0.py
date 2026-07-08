from functools import reduce
from typing import List
import operator

class Stack: 
    def __init__(self): 
        self.stack = []
        
    def push(self, element): 
        self.stack.append(element)
    
    def pop(self): 
        if self.isEmpty(): 
            return "Stack is empty"
        
        return self.stack.pop() 
    
    def peek(): 
        if self.isEmpty(): 
            return "Stack is empty"
        
        return self.stack[-1]
        
    def isEmpty(self): 
        return len(self.stack) == 0 
    
    def size(self):
        return len(self.stack)
    
operations = {
    "+": operator.add,
    "-": operator.sub,
    "*": operator.mul,
    "/": lambda a, b: int(a / b) 
}

class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        st = Stack() 
        
        for i in range(len(tokens)): 
            if tokens[i] not in operations:
                st.push(int(tokens[i]))
                continue
            
            num1 = st.pop()
            num2 = st.pop()
            res = operations[tokens[i]](num2, num1)
            st.push(res)
            
        return st.pop() 