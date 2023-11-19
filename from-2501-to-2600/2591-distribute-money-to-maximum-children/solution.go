func distMoney(money int, children int) int {
    if money < children {
        return -1
    }
    result := 0
    for i := 1; i <= children; i ++ {
        restMoney := money - 8
        restChildren := children - i
        if restMoney >= restChildren {
            money = restMoney
            result += 1
        } else {
            restChildren = children - i + 1
            if restChildren == 1 && money == 4 {
                result -= 1
            }
            money = 0
            break
        }
    }
    if money > 0 {
        result -= 1
    }
    return result
}