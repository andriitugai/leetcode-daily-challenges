# """
# This is the interface that allows for creating nested lists.
# You should not implement it, or speculate about its implementation
# """
#class NestedInteger:
#    def isInteger(self) -> bool:
#        """
#        @return True if this NestedInteger holds a single integer, rather than a nested list.
#        """
#
#    def getInteger(self) -> int:
#        """
#        @return the single integer that this NestedInteger holds, if it holds a single integer
#        Return None if this NestedInteger holds a nested list
#        """
#
#    def getList(self) -> [NestedInteger]:
#        """
#        @return the nested list that this NestedInteger holds, if it holds a nested list
#        Return None if this NestedInteger holds a single integer
#        """

class NestedIterator:
    def __init__(self, nestedList: [NestedInteger]):
        self.store = []
        self._flatten(nestedList)
        self.pointer = 0
        
    def _flatten(self, lst):
        for elem in lst:
            lst = elem.getList()
            if lst:
                self._flatten(lst)
            elif elem.isInteger():
                self.store.append(elem.getInteger())
    
    def next(self) -> int:
        if self.hasNext():
            rv = self.store[self.pointer]
            self.pointer += 1
            return rv
        else:
            raise StopIteration()        
    
    def hasNext(self) -> bool:
        return self.pointer < len(self.store)
         

# Your NestedIterator object will be instantiated and called as such:
# i, v = NestedIterator(nestedList), []
# while i.hasNext(): v.append(i.next())