class Solution:
    def countSeniors(self, details: List[str]) -> int:
        result = 0
        for d in details:
            if int(d[11:13]) > 60:
                result += 1

        return result