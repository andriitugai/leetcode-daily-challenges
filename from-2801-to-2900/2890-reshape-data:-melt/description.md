### 2890\. Reshape Data: Melt

DataFrame `report`
+-------------+--------+
| Column Name | Type   |
+-------------+--------+
| product     | object |
| quarter\_1   | int    |
| quarter\_2   | int    |
| quarter\_3   | int    |
| quarter\_4   | int    |
+-------------+--------+

Write a solution to **reshape** the data so that each row represents sales data for a product in a specific quarter.

The result format is in the following example.

**Example 1:**

**Input:**
+-------------+-----------+-----------+-----------+-----------+
| product     | quarter\_1 | quarter\_2 | quarter\_3 | quarter\_4 |
+-------------+-----------+-----------+-----------+-----------+
| Umbrella    | 417       | 224       | 379       | 611       |
| SleepingBag | 800       | 936       | 93        | 875       |
+-------------+-----------+-----------+-----------+-----------+
**Output:**
+-------------+-----------+-------+
| product     | quarter   | sales |
+-------------+-----------+-------+
| Umbrella    | quarter\_1 | 417   |
| SleepingBag | quarter\_1 | 800   |
| Umbrella    | quarter\_2 | 224   |
| SleepingBag | quarter\_2 | 936   |
| Umbrella    | quarter\_3 | 379   |
| SleepingBag | quarter\_3 | 93    |
| Umbrella    | quarter\_4 | 611   |
| SleepingBag | quarter\_4 | 875   |
+-------------+-----------+-------+
**Explanation:**
The DataFrame is reshaped from wide to long format. Each row represents the sales of a product in a quarter.
