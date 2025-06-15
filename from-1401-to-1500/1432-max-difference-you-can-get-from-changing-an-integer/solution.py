class Solution:
    def maxDiff(self, num: int) -> int:
        s = str(num)

        # Generate the maximum number by replacing first non-9 digit with '9'
        for digit in s:
            if digit != '9':
                max_num = s.replace(digit, '9')
                break
        else:
            max_num = s  # All digits are '9'

        # Generate the minimum number
        if s[0] != '1':
            min_num = s.replace(s[0], '1')
        else:
            for digit in s[1:]:
                if digit != '0' and digit != s[0]:
                    min_num = s.replace(digit, '0')
                    break
            else:
                min_num = s  # Already minimal

        return int(max_num) - int(min_num)
