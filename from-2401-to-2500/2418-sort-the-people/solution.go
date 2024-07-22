type info struct {
    name string
    height int
}

func sortPeople(names []string, heights []int) []string {
    n := len(names)
    people := make([]info, n)
    for i := 0; i < n; i ++ {
        people[i].name = names[i]
        people[i].height = heights[i]
    }
    sort.Slice(people, func(i, j int) bool {
        return people[i].height > people[j].height
    })
    result := make([]string, n)
    for i := 0; i < n; i ++ {
        result[i] = people[i].name
    }
    return result
}