class Solution:
    def intToRoman(self, num: int) -> str:
        num2roman = {
            1: "I", 4: "IV", 5: "V", 9:"IX", 10: "X", 40: "XL", 50: "L", 
            90: "XC", 100: "C", 400: "CD", 500: "D", 900: "CM", 1000: "M"
        }
        ans  = []
        for key in sorted(num2roman.keys(), reverse=True):
            while num >= key:
                num -= key
                ans.append(num2roman[key])

        return "".join(ans)
        