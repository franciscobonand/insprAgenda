
  

# Inspr Task Manager (Agenda)

This is an application that implements a simple Kanban board  using CLI as it's frontend.  

## Install and execute
After cloning the git repository, git bash to **insprAgenda** folder and run the following command to pull the official PostgreSQL image from Docker Hub and configure the database used by Insper Task Manager. **Remember, you only have to run this ONCE**:
```
bash install.sh
```
Having installed and created the container, run the following command (still in the same insprAgenda repository) to execute Inspr Task Manager. **You should do this every time you want to run the application**:
```
bash exec.sh
```  
  
  
## Details on how it works 
- A task can only move foward on the board (i.e. it's not possible to return a Working task back to To Do)  
- A task can't be unremoved
- All task visualization is based on the 4-status board. Therefore, if you choose to display tasks by a filter (like priority) they will be ordered from max to min priority in between other tasks with the same status
  - Priority filter orders from max to min priority (10 to 1)
  - Deadline filter orders from closer to farther delivery date
  - Added time filter orders from most recent to most oldest date
- Time spent working on a task can be seen using "Show task details" functionality, inside the "Manage tasks" menu
- ***It should be noted that the application is not yet completely user-proof, so, please, input information as it's asked, otherwise it will most likely break***

## Structure of a task

  

- ID (int)*

  

- Title (string)

  

- Description (string)

  

- Priority (int)

  

  - 1(min) to 10(max)

  

- Deadline (date)

  

- Time estimate (int)

  

- Dependency (string)

  

- Status (int)*

  

  - To Do (1)

  

  - Working (2)

  

  - Closed (3)

  

  - Done (4)

  

- Work start date (date)*

  

- Work end date (date)*

  
- Creation date (date)*
  
\*These fields are not defined by the user directly  

## User Methods/Functionalities

  

- Create task


  

- Remove task


  

- Update task (move task foward on the board)

  

- Show task details (contains time spent working on task)


  

- List tasks by (lists all the tasks ordered by chosen option)

  

  - Order options: Priority, Deadline or Added time


- Filter tasks by (filter tasks by value defined for the filter)

  

  - Filter options: Priority, Deadline or Added time

  

**Menu divisions:**

  

- Main menu

  

  - Options: Visualize board, Manage board, Show calendar, Exit

  

- Board visualization menu

  

  - Options: By priority, By deadline, By Added time

  

- Board management menu

  

  - Options: Create task, Remove task, Update task status, Show task details

- Board filter menu

  

  - Options: By priority, By deadline, By Added time