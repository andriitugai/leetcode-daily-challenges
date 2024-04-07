type FrequencyTracker struct {
    counter map[int]int
    frequencies map[int]int
}


func Constructor() FrequencyTracker {
    cnt := make(map[int]int)
    frq := make(map[int]int)
    return FrequencyTracker{counter: cnt, frequencies: frq}
}


func (this *FrequencyTracker) Add(number int)  {
    if this.frequencies[this.counter[number]] > 0 {
        this.frequencies[this.counter[number]] -= 1
    }
    this.counter[number] += 1
    this.frequencies[this.counter[number]] += 1
}


func (this *FrequencyTracker) DeleteOne(number int)  {
    if v, ok := this.counter[number]; ok {
        if v > 0 {
            this.frequencies[this.counter[number]] -= 1
            this.counter[number] = v - 1
            this.frequencies[this.counter[number]] += 1
        }
    }
}


func (this *FrequencyTracker) HasFrequency(frequency int) bool {
    val, ok := this.frequencies[frequency]
    return ok && val > 0
}


/**
 * Your FrequencyTracker object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * obj.DeleteOne(number);
 * param_3 := obj.HasFrequency(frequency);
 */