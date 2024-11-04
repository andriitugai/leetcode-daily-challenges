class Solution:
    def compressedString(self, word: str) -> str:
        result = ""
        curC = word[0]
        cnt = 1
        for c in word[1:]:
            if c == curC and cnt < 9:
                cnt += 1
            else:
                result += f"{cnt}{curC}"
                curC = c
                cnt = 1
        result += f"{cnt}{curC}"
        return result