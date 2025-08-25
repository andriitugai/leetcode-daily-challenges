class Solution:
    def findDiagonalOrder(self, mat: List[List[int]]) -> List[int]:
        m, n  = len(mat), len(mat[0])
        result = [0] * m * n
        row, col = 0, 0

        for i in range(m * n):
            result[i] = mat[row][col]
            # print(i, row, col, result)

            if (row + col) & 1:
                if row == m - 1:
                    col += 1
                elif col == 0:
                    row += 1
                else:
                    row += 1
                    col -= 1
            else:
                if col == n - 1:
                    row += 1
                elif row == 0:
                    col += 1
                else:
                    row -= 1
                    col += 1

        return result