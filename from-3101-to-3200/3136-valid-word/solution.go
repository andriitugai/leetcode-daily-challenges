func isValid(word string) bool {
    if len(word) < 3 {
        return false
    }
    if matched, _ := regexp.MatchString("^[a-zA-Z0-9]*$", word); !matched {
        return false
    }
    if matched, _ := regexp.MatchString("[aeiouAEIOU]+", word); !matched {
        return false
    }
    if matched, _ := regexp.MatchString("[^aeiouAEIOU0-9]+", word); !matched {
        return false
    }
    return true
}