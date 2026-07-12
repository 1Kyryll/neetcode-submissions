
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
        
        for t, value in self.storage[key]: 
            if t <= timestamp: 
                answer = value
            else: 
                break 
        
        return answer
        