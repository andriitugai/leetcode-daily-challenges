class Solution:
    def maxWeight(self, pizzas: List[int]) -> int:
        pizzas.sort(reverse=True) 
        n = len(pizzas)
        m = n // 4  # Total number of days (each day we eat 4 pizzas)
        odd = (m + 1) // 2  # Number of odd days
        even = m - odd  # Number of even days
        
        total_weight = 0
        largest = 0
        
        # On odd days, we gain the weight of the heaviest pizza in the set of 4
        for _ in range(odd):
            total_weight += pizzas[largest]  # Pick the largest pizza
            largest += 1
        
        # On even days, we gain the weight of the second heaviest pizza in the set of 4
        for _ in range(even):
            largest += 1 # Skip third largest pizza 
            total_weight += pizzas[largest]  # Gain the weight of the second largest pizza
            largest += 1
        
        return total_weight 