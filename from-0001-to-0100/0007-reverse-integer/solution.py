class Solution(object):

    INT_MIN = -2 ** 31
    INT_MAX = 2 ** 31 - 1

    def reverse(self, x):
        """
        :type x: int
        :rtype: int
        """       

        if x < 0:
            result = -int(str(x)[1:][::-1])
        else:
            result = int(str(x)[::-1])
        
        if result > self.INT_MAX or result < self.INT_MIN:
            return 0
        
        return result