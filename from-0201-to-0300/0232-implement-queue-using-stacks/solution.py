class MyQueue:

    def __init__(self):
        self.stack = []
        self.backup = []

    def push(self, x: int) -> None:
        self.stack.append(x)

    def pop(self) -> int:
        while self.stack:
            self.backup.append(self.stack.pop())
            
        popval = self.backup.pop()
        
        while self.backup:
            self.stack.append(self.backup.pop())
            
        return popval

    def peek(self) -> int:
        while self.stack:
            self.backup.append(self.stack.pop())
            
        peekval = self.backup[-1]
        
        while self.backup:
            self.stack.append(self.backup.pop())
            
        return peekval

    def empty(self) -> bool:
        return len(self.stack) == 0


# Your MyQueue object will be instantiated and called as such:
# obj = MyQueue()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.peek()
# param_4 = obj.empty()