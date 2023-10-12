class Solution(object):
    def myAtoi(self, str):
        """
        :type str: str
        :rtype: int
        """
        if not str:
            return 0

        int_min = -2 ** 31
        int_max = 2 ** 31 - 1

        result = 0
        sign = 1

        idx = 0
        while idx < len(str) and str[idx] == ' ':
            idx += 1

        if idx < len(str):
            if str[idx] == '-':
                sign = -1
                idx += 1
            elif str[idx] == '+':
                idx += 1

        while idx < len(str):
            if '0' <= str[idx] <= '9':
                result = result * 10 + int(str[idx])
            else:
                break
            idx += 1

        result *= sign

        if result > int_max:
            return int_max
        if result < int_min:
            return int_min

        return result