class Solution:
    def findMin(self, nums: List[int]) -> int:
        l = 0
        r = len(nums) - 1
        res = nums[0] 
        
        while l < r: 
            mid = (l + r) // 2
            
            if nums[mid] > nums[r]:
                l = mid + 1
                res = nums[l]
            else: 
                r = mid
            
        return res
    