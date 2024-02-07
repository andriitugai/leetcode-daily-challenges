class Solution(object):
    def frequencySort(self, s):
        """
        :type s: str
        :rtype: str
        """
        if not s:
            return s
        
        import collections
        freqs = collections.defaultdict(int)
        for c in s:
            freqs[c]+=1
        char_items = sorted(list(freqs.items()), key=lambda item:item[1], reverse=True)
        result = []
        for item in char_items:
            result.extend([item[0]] * item[1])
            
        return ''.join(result)