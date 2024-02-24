type DSU struct {
    parent []int
}

func (d DSU) find(x int) int {
    if d.parent[x] != x {
        d.parent[x] = d.find(d.parent[x])
    }
    return d.parent[x]
}

func (d DSU) union(x, y int) {
    d.parent[d.find(x)] = d.find(y)
}

func newDSU(n int) DSU {
    p := make([]int, n)
    for i := range p {
        p[i] = i
    }
    return DSU{parent: p}
}

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
    // sort meetings by time
    sort.Slice(meetings, func(i, j int) bool {return meetings[i][2] < meetings[j][2]})
    // init a disjoint set union with `n` elements and make 0 refer to firstPerson
    dsu := newDSU(n)
    dsu.union(0, firstPerson)
    
    // `cur` record the current time of the meeting 
    cur := meetings[0][2]
    
    // `cache` record all person that have meeting in the same time
    cache := make(map[int]int)
    
    for i:=0; i<len(meetings); i++ {
        // reset `cur` and `cache` when meeting time changed
        if meetings[i][2] != cur {
            cur = meetings[i][2]
            
            // reset `dsu` elements; people attend the meeting in the same time may not be shared the secret
            for k, _ := range cache {
                if dsu.find(k) != firstPerson {
                    dsu.parent[k] = k
                }
            }
            cache = make(map[int]int)
        }
        cache[meetings[i][0]] = 0
        cache[meetings[i][1]] = 0
        // if a person in the meeting know the secret then union another person with this one
        if dsu.find(meetings[i][0]) == firstPerson {
            dsu.union(meetings[i][1], meetings[i][0])
        } else {
            dsu.union(meetings[i][0], meetings[i][1])
        }

    }
    var res []int
    for i:=0; i<n; i++ {
        if dsu.find(i) == firstPerson{
            res = append(res, i)
        }
    }
    return res
}