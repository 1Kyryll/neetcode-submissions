class Solution:
    def findDuplicate(self, nums: List[int]) -> int:
        seen = defaultdict(int)
        
        for i in range(len(nums)): 
            seen[nums[i]] += 1
            
        for k, v in seen.items(): 
            v -= 1
            if v != 0: 
                return k