class Solution:
    def subsetsWithDup(self, nums: List[int]) -> List[List[int]]:
        result = []
        nums.sort()
        
        def backtrack(nums,index=0,arr=[]):
            result.append( arr.copy() )
            for i in range( index, len(nums)):
                if i != index and nums[i] == nums[i-1]: #skip the duplicates, except for the first time
                    continue
                arr.append(nums[i]) #include the element
                backtrack(nums,i+1,arr ) #explore
                arr.pop() #remove the element
            
        backtrack(nums)
        return result