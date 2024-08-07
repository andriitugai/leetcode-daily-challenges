class Solution:
    def numberToWords(self, num: int) -> str:
        if num == 0:
            return "Zero"
    
        zillions = {
           # 10**33: "Decillion",  10**30: "Nonillion",   10**27: "Octillion",    10**24: "Septillion",
           # 10**21: "Sextillion", 10**18: "Quintillion", 10**15: "Quadrillion",  10**12: "Trillion",
           10**9:  "Billion",    10**6:  "Million",     10**3:  "Thousand", 
        }    
        words = {
            100:  "Hundred",     90: "Ninety",    80: "Eighty",   70: "Seventy",
            60:   "Sixty",       50: "Fifty",    40: "Forty",   30: "Thirty",
            20:   "Twenty",      19: "Nineteen", 18: "Eighteen", 17: "Seventeen",
            16:   "Sixteen",     15: "Fifteen",  14: "Fourteen", 13: "Thirteen",
            12:   "Twelve",      11: "Eleven",   10: "Ten",       9: "Nine",
             8:   "Eight",        7: "Seven",     6: "Six",       5: "Five",
             4:   "Four",         3: "Three",     2: "Two",       1: "One"
        }
        ans = []
        
        def thou_by_words (num):
            nonlocal words
            by_words = []
            for key in words:
                num_keys = 0
                while num >= key:
                    num -= key
                    num_keys += 1

                if num_keys > 0:
                    if num_keys == 1:
                        if key > 90:
                            by_words.append("One")
                    else:
                        by_words.append(words[num_keys])

                    by_words.append(words[key])

                if num == 0:
                    break

            return " ".join(by_words)
        
        for zill in zillions:
            if num >= zill:
                ans.append(thou_by_words(num // zill))
                ans.append(zillions[zill])
                
                num %= zill
            
        # The rest (all less than Thousand):
        ans.append(thou_by_words(num))
        
        return " ".join(ans).strip()
        