class Solution:
    def lemonadeChange(self, bills: List[int]) -> bool:
        bank = {5: 0, 10: 0, 20: 10}
        for bill in bills:
            bank[bill] += 1
            change = bill - 5
            while change > 0:
                while change >= 20 and bank[20] > 0:
                    bank[20] -= 1
                    change -= 20

                while change >= 10 and bank[10] > 0:
                    bank[10] -= 1
                    change -= 10
                    
                while change >= 5 and bank[5] > 0:
                    bank[5] -= 1
                    change -= 5

                if change > 0:
                    return False

        return True