func totalMoney(n int) int {
    weeks := n / 7
    sw := weeks * 28 + weeks * (weeks - 1) * 7 / 2
    days := n % 7
    sd := (weeks + 1) * days + days * (days - 1) / 2
    return sd + sw
}