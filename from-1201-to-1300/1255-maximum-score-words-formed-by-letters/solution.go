func maxScoreWords(words []string, letters []byte, score []int) int {
    var cnt [26]int
	for _, letter := range letters {
		cnt[letter - 'a']++
	}

	d := map[[26]int]int{};
    d[cnt] = 0
	res := 0

	for _, word := range words {
		new_d := map[[26]int]int{};

		for rem, cur := range d {
			new_rem := rem
			new_cur := cur
			valid := true

			for _, c := range word {
                tmp := c - 'a'
				new_rem[tmp]--
				new_cur += score[tmp]
				if new_rem[tmp] < 0 {
					valid = false
					break
				}
			}

			if !valid {
				continue
			}

			if new_cur > res {
				res = new_cur
			}

            new_d[new_rem] = max(new_d[new_rem], new_cur);
		}

		for k, val := range new_d {
            d[k] = max(d[k], val)
		}
	}

	return res
}