class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        sortedArr = self.merge(nums1, nums2) 
        mid = len(sortedArr) // 2 
        median = 0 
        if len(sortedArr) % 2 == 0: 
            median = (sortedArr[mid] + sortedArr[mid - 1]) / 2
        else: 
            median = sortedArr[mid]
        
        return median
        
    def merge(self, arr1, arr2): 
        i = j = 0 
        merged = []
        
        while i < len(arr1) and j < len(arr2): 
            if arr1[i] <= arr2[j]: 
                merged.append(arr1[i])
                i += 1
            else: 
                merged.append(arr2[j]) 
                j += 1
                
        merged.extend(arr1[i:])
        merged.extend(arr2[j:])
            
        return merged 
        