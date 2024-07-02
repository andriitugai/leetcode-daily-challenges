class Solution:
    def intersect(self, nums1: List[int], nums2: List[int]) -> List[int]:
        c1 = Counter(nums1)
        c2 = Counter(nums2)
        
        ans = []
        for key in c1.keys():
            if key in c2:
                ans.extend([key] * min(c1[key], c2[key]))
                
        return ans