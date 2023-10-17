class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
              
        def check_nine(nine):
            nines = set()
            n_digs = 0
            for c in nine:
                if c == '.':
                    continue
                else:
                    try:
                        ci = int(c)
                        if ci > 0 and ci < 10:
                            n_digs +=1
                            nines.add(c)
                        else:
                            return False
                    except ValueError:
                        return False

            if len(nines) == n_digs:
                return True
            return False

        for row in board:
            if not check_nine(row):
                return False
        # rows = [ row for row in board ]
        # if not all(map(check_nine, rows)):
        #     return False

        cols = ([board[j][i] for j in range(9)] for i in range(9))
        for col in cols:
            if not check_nine(col):
                return False
        # if not all(map(check_nine, cols)):
        #     return False

        squares = (
            [
                board[3*i][3*j],  board[3*i+1][3*j],  board[3*i+2][3*j],
                board[3*i][3*j+1],board[3*i+1][3*j+1],board[3*i+2][3*j+1],
                board[3*i][3*j+2],board[3*i+1][3*j+2],board[3*i+2][3*j+2]
            ] for i in range(3) for j in range(3)
        )
        for square in squares:
            if not check_nine(square):
                return False
        # if not all(map(check_nine, squares)):
        #     return False

        return True