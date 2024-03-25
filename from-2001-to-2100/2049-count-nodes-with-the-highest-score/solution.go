func countHighestScoreNodes(parents []int) int {
    children := make(map[int][]int)
    for i, p := range parents[1:] {
        children[p] = append(children[p], i + 1)
    }
    weights := make(map[int]int)
    var dfs func(parent int) int
    dfs = func(parent int) int {
        w := 1
        for _, c := range children[parent] {
            w += dfs(c)
        }
        weights[parent] = w
        return w
    }
    dfs(0)
    tot := weights[0]
    maxScore := -1
    numMaxScore := 0
    var score int

    for i := 0; i < len(parents); i ++ {
        score = tot - weights[i]
        if score == 0 {
            score = 1
        }
        for _, child := range children[i] {
            score *= weights[child]
        }
        if score > maxScore {
            maxScore = score
            numMaxScore = 1
        } else if score == maxScore {
            numMaxScore += 1
        }
    }
    return numMaxScore
}