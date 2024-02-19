func largestNumber(nums []int) string {
    n:=len(nums)
    arr:=make([]string,n)
    for i:=0; i<n; i++ {
        arr[i]=fmt.Sprintf("%d",nums[i])
    }
    sort.Slice(arr,func (i,j int)bool{
        if len(arr[i])==len(arr[j]){
            return arr[i]>arr[j]
        }
        return (arr[i]+arr[j]) > (arr[j]+arr[i])
    })
    
    if arr[0] == "0" {
        return "0"
    }  

    result:=""
    for i:=0; i<n; i++{
        result =fmt.Sprintf("%s%s",result,arr[i])
    }
    return result
}