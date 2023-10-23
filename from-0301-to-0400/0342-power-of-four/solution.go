func isPowerOfFour(n int) bool {
    mask := 0x55555555
    isPowerOfTwo := (n > 0) && (n & (n - 1)) == 0

    return isPowerOfTwo && (n & mask) == n
}