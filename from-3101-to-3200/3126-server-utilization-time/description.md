### 3126\. Server Utilization Time

Table: `Servers`

+----------------+----------+
| Column Name    | Type     |
+----------------+----------+
| server\_id      | int      |
| status\_time    | datetime |
| session\_status | enum     |
+----------------+----------+
(server\_id, status\_time, session\_status) is the primary key (combination of columns with unique values) for this table.
session\_status is an ENUM (category) type of ('start', 'stop').
Each row of this table contains server\_id, status\_time, and session\_status.

Write a solution to find the **total time** when servers were **running**. The output should be rounded down to the nearest number of **full days**.

Return _the result table in **any**_ _order._

The result format is in the following example.

**Example:**

**Input:**

Servers table:

+-----------+---------------------+----------------+
| server\_id | status\_time         | session\_status |
+-----------+---------------------+----------------+
| 3         | 2023-11-04 16:29:47 | start          |
| 3         | 2023-11-05 01:49:47 | stop           |
| 3         | 2023-11-25 01:37:08 | start          |
| 3         | 2023-11-25 03:50:08 | stop           |
| 1         | 2023-11-13 03:05:31 | start          |
| 1         | 2023-11-13 11:10:31 | stop           |
| 4         | 2023-11-29 15:11:17 | start          |
| 4         | 2023-11-29 15:42:17 | stop           |
| 4         | 2023-11-20 00:31:44 | start          |
| 4         | 2023-11-20 07:03:44 | stop           |
| 1         | 2023-11-20 00:27:11 | start          |
| 1         | 2023-11-20 01:41:11 | stop           |
| 3         | 2023-11-04 23:16:48 | start          |
| 3         | 2023-11-05 01:15:48 | stop           |
| 4         | 2023-11-30 15:09:18 | start          |
| 4         | 2023-11-30 20:48:18 | stop           |
| 4         | 2023-11-25 21:09:06 | start          |
| 4         | 2023-11-26 04:58:06 | stop           |
| 5         | 2023-11-16 19:42:22 | start          |
| 5         | 2023-11-16 21:08:22 | stop           |
+-----------+---------------------+----------------+

**Output:**

+-------------------+
| total\_uptime\_days |
+-------------------+
| 1                 |
+-------------------+

**Explanation:**

*   For server ID 3:
    *   From 2023-11-04 16:29:47 to 2023-11-05 01:49:47: ~9.3 hours
    *   From 2023-11-25 01:37:08 to 2023-11-25 03:50:08: ~2.2 hours
    *   From 2023-11-04 23:16:48 to 2023-11-05 01:15:48: ~1.98 hoursTotal for server 3: ~13.48 hours
*   For server ID 1:
    *   From 2023-11-13 03:05:31 to 2023-11-13 11:10:31: ~8 hours
    *   From 2023-11-20 00:27:11 to 2023-11-20 01:41:11: ~1.23 hoursTotal for server 1: ~9.23 hours
*   For server ID 4:
    *   From 2023-11-29 15:11:17 to 2023-11-29 15:42:17: ~0.52 hours
    *   From 2023-11-20 00:31:44 to 2023-11-20 07:03:44: ~6.53 hours
    *   From 2023-11-30 15:09:18 to 2023-11-30 20:48:18: ~5.65 hours
    *   From 2023-11-25 21:09:06 to 2023-11-26 04:58:06: ~7.82 hoursTotal for server 4: ~20.52 hours
*   For server ID 5:
    *   From 2023-11-16 19:42:22 to 2023-11-16 21:08:22: ~1.43 hoursTotal for server 5: ~1.43 hours

The accumulated runtime for all servers totals approximately 44.46 hours, equivalent to one full day plus some additional hours. However, since we consider only full days, the final output is rounded to 1 full day.
