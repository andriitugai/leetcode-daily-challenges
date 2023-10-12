class Solution:
    def countAndSay(self, n: int) -> str:
        if n == 1:
            return "1"

        nstr = self.countAndSay(n-1)
        cur_char = nstr[0]
        count = 1
        idx = 1
        result = ""
        while idx < len(nstr):
            if cur_char == nstr[idx]:
                count += 1
            else:
                result += str(count*10 + int(cur_char))
                cur_char = nstr[idx]
                count = 1

            idx += 1

        result += str(count*10 + int(cur_char))
        return result
            