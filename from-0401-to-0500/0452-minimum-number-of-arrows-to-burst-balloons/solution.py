class Solution:
    def findMinArrowShots(self, points: List[List[int]]) -> int:
        points.sort(key=lambda x: x[1])
        arrow_x = points[0][1]
        count = 1
        for start, end in points:
            if start > arrow_x:
                count += 1
                arrow_x = end
                
        return count