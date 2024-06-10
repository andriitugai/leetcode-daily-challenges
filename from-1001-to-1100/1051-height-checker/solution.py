class Solution(object):
    def heightChecker(self, heights):
        """
        :type heights: List[int]
        :rtype: int
        """
        if not heights:
            return 0
        
        hc = heights[:]
        hc.sort()
        
        pi = 0
        pc = 0
        result = 0
        
        while pi < len(heights):
            if heights[pi] != hc[pi]:
                result += 1
            pi += 1
            
        return result
        