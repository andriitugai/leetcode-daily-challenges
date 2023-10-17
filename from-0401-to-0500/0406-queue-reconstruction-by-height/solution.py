class Solution(object):
    def reconstructQueue(self, people):
        """
        :type people: List[List[int]]
        :rtype: List[List[int]]
        """
        ppl = sorted(people, key=lambda p: (-p[0], p[1]))
        
        result = []
        for p in ppl:
            result.insert(p[1], p)
        return result