class Solution:
    def intersection(self, nums1: List[int], nums2: List[int]) -> List[int]:
        c1 = Counter(nums1)
        c2 = Counter(nums2)
        
        ans = []
        for k in c1.keys():
            if k in c2:
                ans.append(k)
                
        return ans
    