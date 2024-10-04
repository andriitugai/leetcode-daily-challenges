class Solution:
    def dividePlayers(self, skill: List[int]) -> int:
        skill.sort()
        teamSkill = skill[0] + skill[-1]
        l, r = 1, len(skill) - 2
        result = skill[0] * skill[-1]
        while l < r:
            if skill[l] + skill[r] != teamSkill:
                return -1
            result += skill[l] * skill[r]
            l += 1
            r -= 1
        return result