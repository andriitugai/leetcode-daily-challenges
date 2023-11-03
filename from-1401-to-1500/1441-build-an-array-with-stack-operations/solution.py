class Solution(object):
    def buildArray(self, target, n):
        """
        :type target: List[int]
        :type n: int
        :rtype: List[str]
        """
        result = []
        
        for nxt in range(1, n+1):  
            
            result.append("Push")
            
            if nxt not in target:
                result.append("Pop")
                
            if nxt == target[-1]:
                break
                
        return result
                