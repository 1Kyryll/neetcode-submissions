class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        for j in range(len(matrix)): 
            resR = self.search(matrix[j], target) 
            if resR == -1: 
                continue
            else: 
                return True
                
        return False
        
    def search(self, nums: List[int], target: int) -> int:
        left = 0 
        right = len(nums) - 1
        
        while left <= right: 
            mid = (left + right) // 2
            
            if nums[mid] == target:
                return mid
                
            if nums[mid] < target: 
                left = mid + 1 
            else: 
                right = mid - 1 
                
        return -1 
            