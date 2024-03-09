Intuition
=========

It's a tree problem coupled with the following criteria which made me consider DFS.  
`The path from c to b and the path from c to a do not share any edges`

Approach
========

1.  Consider a root node, and find all nodes from each connected edge that have a reachable distance (`distance % signalSpeed == 0`).
    
    Eg. For instance, if edges `A`,`B`,`C` are connected to the root node, and through edge `A`, `B` and `C` there are `a`,`b` and `c` reachable nodes respectively. Then the total number of pairs possible are `a*b +b*c +c*a`.
    
    Pairwise sum can be obtained from a derived variation of the _whole squared_ formula,
    
    (a+b+c+..)2=a2+b2+..+2∗(a∗b+b∗c+c∗a+...)(a+b+c+..)^2 = a^2+b^2+.. + 2\*(a\*b+b\*c+c\*a+...)(a+b+c+..)2\=a2+b2+..+2∗(a∗b+b∗c+c∗a+...)
    
2.  Iteratively consider every node as a root node, and repeat the process above and store the results in an array.
    

Complexity
==========

*   Time complexity:  
    O(n2)O(n^2)O(n2)
    
*   Space complexity:  
    O(n)O(n)O(n)
    
