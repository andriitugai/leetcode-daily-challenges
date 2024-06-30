class Solution:
    def maxHeightOfTriangle(self, red: int, blue: int) -> int:
        return max(self._helper(red, blue), self._helper(blue, red))

    def _helper(self, red: int, blue: int) -> int:
        height = 0
        rowToComplete = 1

        while True:
            if rowToComplete % 2 == 1:
                if red < rowToComplete:
                    break
                red -= rowToComplete
            else:
                if blue < rowToComplete:
                    break
                blue -= rowToComplete

            height += 1
            rowToComplete += 1
        return height
        