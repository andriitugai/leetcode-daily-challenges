class Solution:
    def solveSudoku(self, board: List[List[str]]) -> None:
        """
        Do not return anything, modify board in-place instead.
        """
        self.helper(board)

    def helper(self, board) -> bool:
        digits = "123456789"
        def possible(r, c, n):
            nonlocal board
            for i in range(9):
                if board[r][i] == n or board[i][c] == n:
                    return False
            r0 = (r // 3) * 3
            c0 = (c // 3) * 3
            for i in range(3):
                for j in range(3):
                    if board[r0+i][c0+j] == n:
                        return False

            return True

        for row in range(9):
            for col in range(9):
                if board[row][col] == ".":
                    for n in digits:
                        if possible(row, col, n):
                            board[row][col] = n
                            if self.helper(board):
                                return True
                            else:
                                board[row][col] = "."

                    return False
        return True