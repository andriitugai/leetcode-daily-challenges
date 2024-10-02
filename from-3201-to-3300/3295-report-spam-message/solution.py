class Solution:
    def reportSpam(self, message: List[str], bannedWords: List[str]) -> bool:
        banned = {word : True for word in bannedWords}
        # banned = set(bannedWords)
        banCount = 0
        for word in message:
            if banned.get(word, False):
            # if word in banned:
                banCount += 1
            if banCount == 2:
                return True
        return False