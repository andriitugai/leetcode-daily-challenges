class Solution:
    def isValid(self, s: str) -> bool:
        stack = []
        for c in s:
            if c == '{':
                stack.append('}')
            elif c == '(':
                stack.append(')')
            elif c == '[':
                stack.append(']')
            else:
                if stack:
                    if stack.pop() != c:
                        return False
                else:
                    return False
                
        return not stack