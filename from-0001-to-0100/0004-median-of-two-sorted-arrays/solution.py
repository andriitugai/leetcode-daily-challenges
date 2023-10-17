class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        total = len(nums1) + len(nums2)
        half = total // 2

        if len(nums1) > len(nums2):
            nums1, nums2 = nums2, nums1

        left, right = 0, len(nums1) - 1
        while True:
            i = (left + right) // 2
            j = half - i - 2

            aleft = nums1[i] if i >= 0 else float('-inf')
            aright = nums1[i+1] if i+1 < len(nums1) else float('inf')
            bleft = nums2[j] if j >= 0 else float('-inf')
            bright = nums2[j+1] if j+1 < len(nums2) else float('inf')

            if aleft <= bright and bleft <= aright:
                if total % 2:
                    return min(aright, bright)
                return (max(aleft, bleft) + min(aright, bright)) / 2
            elif aleft > bright:
                right = i - 1
            else:
                left = i + 1
