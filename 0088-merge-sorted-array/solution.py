lass Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        """
        Do not return anything, modify nums1 in-place instead.
        """
        nums3 = nums1[:m]
        p1, p2, p3 = 0, 0, 0
        
        while p2 < n and p3 < m:
            if nums2[p2] < nums3[p3]:
                nums1[p1] = nums2[p2]
                p2 += 1
            else:
                nums1[p1] = nums3[p3]
                p3 += 1
            p1 += 1
            
        while p2 < len(nums2):
            nums1[p1] = nums2[p2]
            p2 += 1
            p1 += 1
            
        while p3 < len(nums3):
            nums1[p1] = nums3[p3]
            p3 += 1
            p1 += 1
        