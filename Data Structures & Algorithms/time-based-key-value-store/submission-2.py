class TimeMap:
    def __init__(self):
        self.storage = {} 

    def set(self, key: str, value: str, timestamp: int) -> None:
        if key not in self.storage: 
            self.storage[key] = []
            
        self.storage[key].append((timestamp, value))

    def get(self, key: str, timestamp: int) -> str:
        if key not in self.storage:
            return ""
            
        answer = ""
        
        l = 0 
        r = len(self.storage[key]) - 1 
        
        while l <= r: 
            mid = (l + r) // 2 
            t, value = self.storage[key][mid] 
            
            if t == timestamp: 
                return value
            elif t < timestamp: 
                answer = value 
                l = mid + 1 
            else: 
                r = mid - 1 
        
        return answer