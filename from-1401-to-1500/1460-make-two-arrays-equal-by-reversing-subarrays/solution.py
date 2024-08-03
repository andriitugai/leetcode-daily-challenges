class Solution(object):
    def canBeEqual(self, target, arr):
        """
        :type target: List[int]
        :type arr: List[int]
        :rtype: bool
        """
        from collections import defaultdict
        d_arr = defaultdict(int)
        d_tar = defaultdict(int)
        
        for item in arr:
            d_arr[item] += 1
            
        for item in target:
            d_tar[item] += 1
            
#         for key in d_arr.keys():
#             if key not in d_tar.keys() or d_tar[key] != d_arr[key]:
#                 return False
            
        return d_tar == d_arr
    
    
        