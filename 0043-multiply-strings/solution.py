lass Solution:
    def multiply(self, num1: str, num2: str) -> str:
        if num1 == "0" or num2 == "0":
            return "0"
        
        if len(num1) < len(num2):
            num1, num2 = num2, num1        # num2 not longer than num1
            
        n1 = list(map(int, num1))
        n2 = list(map(int, num2))
        
        # print(n1, n2)
        
        def mul_arr(arr: list, d: int) -> list:
            if d == 0:
                return [0]
            elif d == 1:
                # print(f"{arr} x {d} = {arr}")
                return arr.copy()
            
            carry = 0
            result = []
            for i in range(len(arr)-1, -1, -1):
                prod = arr[i] * d + carry
                result.append(prod % 10)
                carry = prod //10
            if carry > 0:
                result.append(carry)
            
            # print(f"{arr} x {d} = {result[::-1]}")
            return result[::-1]
        
        results = []
        lg = 0
        maxlen = 0
        for d in n2[::-1]:
            result = mul_arr(n1, d)
            result.extend([0]*lg)
            lg += 1
            maxlen = max(maxlen, len(result))
            results.append(result)
          
        # ensure all results have same length:
        for i in range(len(results)):
            delta = maxlen - len(results[i])
            r = [0] * delta
            r.extend(results[i])
            results[i] = r
            
        # summ all results:
        ans = []
        carry = 0
        for col in range(len(results[0])-1, -1, -1):
            s = carry
            for row in range(len(results)):
                s = s + results[row][col]
            ans.append(str(s % 10))
            carry = s // 10
            
        if carry > 0:
            ans.append(str(carry))
            
        return ''.join(ans[::-1])
    