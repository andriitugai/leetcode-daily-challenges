class Solution:
    def minHeightShelves(self, books: List[List[int]], shelfWidth: int) -> int:
        n = len(books)
        
        @lru_cache(None)
        def dp(i, cur_height, rem_width):
            nonlocal n, books, shelfWidth
            if i == n:
                return cur_height
            
            w, h = books[i]
            
            # put on the next level
            ans = cur_height + dp(i+1, h, shelfWidth - w)
            
            if rem_width >= w:
                ans = min(ans, dp(i+1, max(h, cur_height), rem_width - w))
                
            return ans
        
        return dp(0, 0, shelfWidth)