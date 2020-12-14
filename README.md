
  

# insprAgenda

This is an application that implements a simple Kanban board  using CLI as it's frontend.  

**Details on how it works:**  
- A task can only move foward on the board (it's not possible to return a Working task back to To Do)  
- A task can't be unremoved
- All task visualization is based on the 4-status board. Therefore, if you choose to display tasks by a filter like priority, they will be ordered from max to min priority in between other tasks with the same status
- Priority filter orders from max to min priority (10 to 1)
- Deadline filter orders from closer to farther delivery date
- Added time filter orders from most recent to most distant creation date
- Time spent working on a task can be seen using "Show task details" functionality, inside the "Manage tasks" menu
- ### The application is not yet user-proof, so, please, input information as it's asked, otherwise it will most likely break

**Structure of the tasks:**

  

- ID (int)

  

- Title (string)

  

- Description (string)

  

- Priority (int)

  

  - 1(min) to 10(max)

  

- Deadline (date)

  

- Time estimate (int)

  

- Dependency (string)

  

- Status (int)

  

  - To Do (1)

  

  - Working (2)

  

  - Closed (3)

  

  - Done (4)

  

- Work start date (date)

  

- Work end date (date)

  
- Creation date (date)
  

**User Methods/Functionalities:**

  

- CreateTask


  

- RemoveTask


  

- UpdateTask (move task foward on the board)

  

- ShowTaskDetails


  

- ListTasksBy

  

  - Filter options: Priority, Deadline or Added time


- ShowCalendar  

  

**Menu divisions:**

  

- Main menu

  

  - Options: Visualize board, Manage board, Show calendar, Exit

  

- Board visualization menu

  

  - Options: By priority, By deadline, By Added time

  

- Board management menu

  

  - Options: Create task, Remove task, Update task status, Show task details