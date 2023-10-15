class Solution:
    def simplifyPath(self, path: str) -> str:
        stack = ['']*3000
        top = 0
        
        path_lst = path.split('/')
        for dir_ in path_lst:
            if not dir_:
                continue
                
            if dir_ == '.':
                continue
            elif dir_ == '..':
                if top > 0:
                    top -= 1
            else:
                stack[top] = dir_
                top += 1
                
        return '/' + '/'.join(stack[:top])