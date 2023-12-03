func squareIsWhite(coordinates string) bool {
    r := int(coordinates[0] - 'a') % 2
    c := int(coordinates[1] - '1') % 2
    return r ^ c > 0
}