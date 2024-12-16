class Solution:
    def buttonWithLongestTime(self, events: List[List[int]]) -> int:
        max_button, prev_time = events[0]
        max_time = prev_time

        for button, time in events[1:]:
            elapsed_time = time - prev_time
            if elapsed_time > max_time:
                max_time = elapsed_time
                max_button = button
            elif elapsed_time == max_time and max_button > button:
                max_button = button

            prev_time = time

        return max_button