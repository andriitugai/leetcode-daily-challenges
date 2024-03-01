/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */

 func getImportance(employees []*Employee, id int) int {
    empInfo := make(map[int]*Employee)
    for _, emp := range employees {
        empInfo[emp.Id] = emp
    }
    var importance func(id int) int
    importance = func(id int) int {
        value := empInfo[id].Importance
        for _, subId := range empInfo[id].Subordinates {
            value += importance(subId)
        }
        return value
    }
    return importance(id)
}