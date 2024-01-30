class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        stack = []
        for token in tokens: 
            if token in "+-*/":            
                op1 = stack.pop()
                op2 = stack.pop()
                
                if token == '+':
                    op1 = op1 + op2
                elif token == '-':
                    op1 = op2 - op1
                elif token == '*':
                    op1 = op1 * op2
                else:
                    op1 = int(op2/op1)
                    
                stack.append(op1)
                
            else:
                stack.append(int(token))
                
        return stack.pop()