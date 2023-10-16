class Solution:
    def restoreIpAddresses(self, s: str) -> List[str]:
        n, result = len(s), []
        if n > 12:
            return result

        def backtrack(i, cur_ip, dots):
            if i == n and dots == 4:
                result.append(cur_ip[:-1])
                return
            if dots > 4:
                return

            for j in range(i, min(i+3, n)):
                cand = s[i:j+1]
                if (len(cand)==1 or cand[0] != '0') and int(cand) < 256:
                    backtrack(j+1, cur_ip + cand + ".", dots+1)

        backtrack(0, "", 0)

        return result