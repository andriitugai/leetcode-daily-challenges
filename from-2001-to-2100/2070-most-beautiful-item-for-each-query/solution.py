class Solution:
    def maximumBeauty(self, items: List[List[int]], queries: List[int]) -> List[int]:
        items.sort(key=lambda x: x[0])
        n = len(items)
        for i in range(1, n):
            items[i][1] = max(items[i][1], items[i - 1][1])

        result = []
        for q in queries:
            r1 = bisect.bisect_right(items, q, key=itemgetter(0))
            
            # l, r = 0, n - 1
            # while l <= r:
            #     m = (l + r) >> 1
            #     p, b = items[m]
            #     if q < p:
            #         r = m - 1
            #     else:
            #         l = m + 1

            # print(r1, r)

            result.append(items[r1-1][1] if r1 > 0 else 0)

        return result