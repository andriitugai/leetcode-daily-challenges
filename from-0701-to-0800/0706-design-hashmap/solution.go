type HashNode struct {
    key int
    val int
    next *HashNode
}

func NewHashNode(key, val int) *HashNode {
    return &HashNode {
        key: key,
        val: val,
        next: nil,
    }
}

type MyHashMap struct {
    hash []*HashNode
}

var NUM_BUCKETS int = 10000

func Constructor() MyHashMap {
    hash := make([]*HashNode, NUM_BUCKETS)
    for i := 0; i < NUM_BUCKETS; i++ {
        hash[i] = NewHashNode(-1, -1)
    }
    return MyHashMap{
        hash: hash,
    }
}


func (this *MyHashMap) Put(key int, value int)  {
    h := hasher(key)
    p := this.hash[h]

    for p.next != nil {
        if p.next.key == key {
            p.next.val = value
            return
        }
        p = p.next
    }
    p.next = NewHashNode(key, value)
}


func (this *MyHashMap) Get(key int) int {
    h := hasher(key)
    p := this.hash[h]

    for p != nil && p.next != nil {
        if p.next.key == key {
            return p.next.val
        }
        p = p.next
    }
    return -1
}


func (this *MyHashMap) Remove(key int)  {
    h := hasher(key)
    p := this.hash[h]

    for p != nil && p.next != nil {
        if p.next.key == key {
            p.next = p.next.next
        }
        p = p.next
    }
}


func hasher(key int) int {
    return key % NUM_BUCKETS
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */