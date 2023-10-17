class Solution(object):
    def removeKdigits(self, num, k):
        """
        :type num: str
        :type k: int
        :rtype: str
        """
        if k == 0:
            return num

        if k == len(num):
            return "0"

        nums = [int(c) for c in num]

        j = 0
        while j < k:
            idx = 1
            is_deleted = False
            while True:
                prev = nums[idx-1]
    
                if nums[idx] < prev:
                    del nums[idx - 1]
                    is_deleted = True

                idx += 1
                if is_deleted or idx == len(nums):
                    break

            if not is_deleted:
                del nums[-1]

            j += 1

        return str(int(''.join([str(n) for n in nums])))