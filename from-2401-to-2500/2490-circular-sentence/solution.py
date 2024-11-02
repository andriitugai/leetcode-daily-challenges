class Solution:
    def isCircularSentence(self, sentence: str) -> bool:
        if sentence[0] != sentence[len(sentence) - 1]:
            return False
        words = sentence.split(' ')
        prevWord = words[0]
        for currWord in words[1:]:
            if prevWord[len(prevWord) - 1] != currWord[0]:
                return False
            prevWord = currWord
        return True