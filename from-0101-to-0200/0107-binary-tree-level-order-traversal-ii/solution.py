# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution(object):
    def levelOrderBottom(self, root):
        """
        :type root: TreeNode
        :rtype: List[List[int]]
        """
        if not root:
            return root
        
        level = [root]
        result = []
        
        while True:
            childs = []
            for elem in level:
                if elem.left:
                    childs.append(elem.left)
                if elem.right:
                    childs.append(elem.right)
            result.append([e.val for e in level])
            if not childs:
                break
            else:
                level = childs
                
        return result[::-1]
                    
        