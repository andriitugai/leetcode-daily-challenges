class Solution:
    def convert(self, s: str, numRows: int) -> str:
        if numRows == 1 or len(s) <= numRows:
            return s

        zz = [[''] * len(s) for _ in range(numRows)]
        row, col = 0, 0
        d_row = 1
        idx = 0

        while idx < len(s):
            zz[row][col] = s[idx]
            idx += 1
            row = row + d_row
            if row == 0 or row == numRows-1:
                d_row *= -1
                col += 1

        return ''.join([''.join([c for c in r if c]) for r in zz])