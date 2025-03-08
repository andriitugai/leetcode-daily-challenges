class Solution:
    def minimumRecolors(self, blocks: str, k: int) -> int:
        cur_white = 0
        for i in range(k):
            if blocks[i] == 'W':
                cur_white += 1

        min_white = cur_white
        for i in range(k, len(blocks)):
            if blocks[i] == 'W':
                cur_white += 1
            if blocks[i - k] == 'W':
                cur_white -= 1
            min_white = min(min_white, cur_white)
        
        return min_white