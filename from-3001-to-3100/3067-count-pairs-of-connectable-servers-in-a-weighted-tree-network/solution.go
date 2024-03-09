var g [][][]int
var sp int

func countPairsOfConnectableServers(edges [][]int, signalSpeed int) []int {
    n:=len(edges)+1
    sp=signalSpeed
    g=make([][][]int, n)
    for i:=range g{
        g[i]=[][]int{}
    }
    for _,e:=range edges{
        g[e[0]]=append(g[e[0]], []int{e[1],e[2]})
        g[e[1]]=append(g[e[1]], []int{e[0],e[2]})
    }
    ans:=make([]int, 0, n)
    for i:=range g{
        sqSum,sum:=0,0
        for _,e:=range g[i]{
            count:=dfs(i,e[0],e[1])
            sqSum+=count*count
            sum+=count
        }
        ans=append(ans, (sum*sum-sqSum)/2)
    }
    return ans
}

func dfs(p, i, d int) int {
    ans:=0
    if d%sp == 0{
        ans++
    }
    for _,e:=range g[i]{
        if e[0]==p {continue}
        ans+=dfs(i,e[0],d+e[1])
    }
    return ans
}