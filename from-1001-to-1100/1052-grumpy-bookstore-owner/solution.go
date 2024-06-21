func maxSatisfied(customers []int, grumpy []int, minutes int) int {
    happy, maxHappy := 0, 0
    left := 0
    satisfied := 0
    for i := 0; i < len(customers); i ++ {
        if grumpy[i] == 1 {
            happy += customers[i]
        } else {
            satisfied += customers[i]
        }

        if i - left + 1 > minutes {
            if grumpy[left] == 1 {
                happy -= customers[left]
            }
            left += 1
        }

        if happy > maxHappy {
            maxHappy = happy
        }
    }
    return satisfied + maxHappy
}