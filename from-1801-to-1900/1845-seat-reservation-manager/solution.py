class SeatManager:

    def __init__(self, n: int):
        self.h = [i+1 for i in range(n)]
        heapq.heapify(self.h)

    def reserve(self) -> int:
        return heapq.heappop(self.h)

    def unreserve(self, seatNumber: int) -> None:
        heapq.heappush(self.h, seatNumber)


# Your SeatManager object will be instantiated and called as such:
# obj = SeatManager(n)
# param_1 = obj.reserve()
# obj.unreserve(seatNumber)