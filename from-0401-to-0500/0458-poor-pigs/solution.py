import math

class Solution:
    def poorPigs(self, buckets: int, minutesToDie: int, minutesToTest: int) -> int:
        trials = minutesToTest // minutesToDie + 1
        return math.ceil(math.log(buckets)/math.log(trials))