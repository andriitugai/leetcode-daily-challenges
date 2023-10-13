class Solution:
    def solveNQueens(self, n: int) -> List[List[str]]:
        results = []
        self.solve(0, n, [], results)
        return results
    
    def solve(self, row, n, col_positions, results):
        if row == n:
            results.append(self.create_board(col_positions, n))
            
            return
    
        for col in range(n):
            col_positions.append(col)
            
            if self.is_valid(col_positions, n):
                self.solve(row+1, n, col_positions, results)
                
            col_positions.pop()
            
    def is_valid(self, col_positions, n):
        last_row = len(col_positions) - 1
        for row in range(last_row):
            col_distance = abs(col_positions[row] - col_positions[last_row])
            row_distance = last_row - row
            
            if col_distance == 0 or col_distance == row_distance:
                return False
        
        return True
    
    def create_board(self, col_positions, n):
        board = []
        for row in range(len(col_positions)):
            s = ""
            
            for col in range(n):
                if col == col_positions[row]:
                    s += 'Q'
                else: s += '.'
                    
            board.append(s)
            
        return board