class Solution:
    def minEatingSpeed(self, piles: List[int], h: int) -> int:
        lo, hi = 1, max(piles)
        
        while lo < hi: 
            k = (lo + hi) // 2
            if self.canFinish(piles, k, h): 
                hi = k
            else:
                lo = k + 1 
    
        return lo
        
    def canFinish(self, piles: List[int], k: int, h: int) -> bool: 
        hours = sum((p + k - 1) // k for p in piles)
        return hours <= h 