class Solution:
    def findCenter(self, edges: List[List[int]]) -> int:
        vrt = dict()
        for v1, v2 in edges:
            if vrt.get(v1, False):
                return v1
            vrt[v1] = True
            if vrt.get(v2, False):
                return v2
            vrt[v2] = True

        return -1