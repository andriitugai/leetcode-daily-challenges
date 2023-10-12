class Solution(object):
    def longestCommonPrefix(self, strs):
        """
        :type strs: List[str]
        :rtype: str
        """
        if not strs:
            return ''
        
        srt_strs = sorted(strs)
        if not srt_strs[0]:
            return ''
        
        result = []
        
        for tp in zip(srt_strs[0],srt_strs[-1]):
            if tp[0] == tp[1]:
                result.append(tp[0])
            else:
                break
                
        return ''.join(result)
            
        