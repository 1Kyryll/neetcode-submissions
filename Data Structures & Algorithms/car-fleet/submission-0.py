class Solution:
    def carFleet(self, target: int, position: List[int], speed: List[int]) -> int:
        if len(position) != len(speed): 
            return 0
        
        cars = sorted(zip(position, speed), reverse=True)
        
        stack = []
        
        for pos, sp in cars: 
            time = (target - pos) / sp 
            
            if not stack or time > stack[-1]:
                stack.append(time)
            
        return len(stack)