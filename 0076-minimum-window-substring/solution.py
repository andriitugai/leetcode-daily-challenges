class Solution:
    def minWindow(self, s, t):
        need = collections.Counter(t)            #hash table to store char frequency
        missing = len(t)                         #total number of chars we care
        start, end = 0, 0
        i = 0
        for j, char in enumerate(s, 1):          #index j from 1
            if need[char] > 0:
                missing -= 1
            need[char] -= 1
            if missing == 0:                     #match all chars
                while i < j and need[s[i]] < 0:  #remove chars to find the real start
                    need[s[i]] += 1
                    i += 1
                need[s[i]] += 1                  #make sure the first appearing char satisfies need[char]>0
                missing += 1                     #we missed this first char, so add missing by 1
                if end == 0 or j-i < end-start:  #update window
                    start, end = i, j
                i += 1                           #update i to start+1 for next window
        return s[start:end]
#     def minWindow(self, s: str, t: str) -> str:
#         if not t or not s or len(s) < len(t):
#             return ""

#         from collections import Counter
#         t_counter = Counter(t)
#         t_set = t_counter.keys()
        
#         min_win = ""
#         min_len = float('inf')
        
#         left, right = 0, len(t)
#         w_counter = Counter([c for c in s[left:right] if c in t_set])
        
#         while right < len(s) - 1:
#             print(right)
#             # extend window until the substring becomes 'desired'
#             while right < len(s)-1:
#                 right += 1
#                 new_char = s[right]
#                 if new_char in t_set:
#                     w_counter[new_char] = w_counter.get(new_char, 0) + 1
                    
#                     if t_counter == w_counter:
#                         if right - left < min_len:
#                             min_len = right - left
#                             min_win = s[left:right+1]
#                         break
                    
#             # contract the window until the substring is 'desired'
#             while right - left > len(t):
#                 old_char = s[left]
#                 left += 1
#                 if old_char in t_set:
#                     w_counter[old_char] -= 1 # w_counter becomes != t_counter
#                     break
#                 if right - left < min_len:
#                     min_len = right - left
#                     min_win = s[left:right+1]
                    
#         return min_win
    