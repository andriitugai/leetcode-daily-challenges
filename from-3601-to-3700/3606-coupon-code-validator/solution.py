class Solution:
    def validateCoupons(self, code: List[str], businessLine: List[str], isActive: List[bool]) -> List[str]:
        valid_categories = {"electronics", "grocery", "pharmacy", "restaurant"}
        valid_pattern = r"^[0-9a-zA-Z\_]+$"
        n = len(code)
        valid_codes = defaultdict(list)
        for i in range(n):
            if isActive[i] and businessLine[i] in valid_categories and re.match(valid_pattern, code[i]):
                valid_codes[businessLine[i]].append(code[i])

        result = []
        for cat in sorted(valid_codes.keys()):
            result.extend(sorted(valid_codes[cat]))
        return result