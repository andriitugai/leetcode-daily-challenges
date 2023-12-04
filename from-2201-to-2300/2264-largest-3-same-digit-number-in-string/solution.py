class Solution:
    def largestGoodInteger(self, num: str) -> str:
        maxstring = ""
        for i in range(len(num)-2):
            if num[i] == num[i+1] == num[i+2] and num[i:i+3] > maxstring:
                maxstring = num[i:i+3]
                
        return maxstring