class Solution:
    def getRow(self, rowIndex: int) -> List[int]:
        if rowIndex < 1:
            return [1]
        
        pascal = {}
        def get_pascal(i,j):
            if j > i:
                raise ValueError("Incorrect input")
            if (i,j) in pascal:
                return pascal[i,j]
            if i == 0 or j == 0 or j == i:
                pascal[i,j] = 1
            else:
                pascal[i,j] = get_pascal(i-1,j-1) + get_pascal(i-1, j)
                
            return pascal[i,j]
        
        return [get_pascal(rowIndex, j) for j in range(rowIndex+1)]
        