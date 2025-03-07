class Solution:
    def closestPrimes(self, left: int, right: int) -> List[int]:
        is_prime = [True] * (right + 1)
        is_prime[0] = is_prime[1] = False
        step = 2
        while step * step <= right:
            if is_prime[step]:
                for num in range(step * step, right + 1, step):
                    is_prime[num] = False
            step += 1

        primes = [idx for idx in range(left, right + 1) if is_prime[idx]]

        if len(primes) < 2:
            return [-1, -1]

        result = [0, 0]
        min_diff = float('inf')
        for i in range(1, len(primes)):
            diff = primes[i] - primes[i- 1]
            if diff < min_diff:
                min_diff = diff
                result = [primes[i-1], primes[i]]

        return result
