# Event Scheduler

## Overview
This is a **Go-based console application** for managing events without overlaps. It allows users to **schedule events** by specifying start and end times. The application ensures **no overlapping events** are added, and events are sorted by their start time for easy viewing.

---

## Features

1. **Add Event:**  
   - Users can add events with start and end times.  
   - Ensures events **do not overlap** with existing ones.

2. **View Events:**  
   - Displays all scheduled events in a **readable format** with 24-hour notation.  
   - Events are **sorted by start time**.

3. **Input Validation:**  
   - Ensures valid time ranges (start < end).  
   - Accepts **hours in the range of 0 to 24**.

4. **No Overlap Logic:**  
   - Events are **only added** if they do not overlap with any existing event.

---

## Example I/O
```1. Add Event
2. View Events
3. Exit
Enter your choice: 1
Enter start time (0-23): 9
Enter end time (1-24): 12
Event added successfully!

1. Add Event
2. View Events
3. Exit
Enter your choice: 2
Scheduled Events:
Event: 09:00 - 12:00
```

---

## How to Run

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd koach-event-scheduler

