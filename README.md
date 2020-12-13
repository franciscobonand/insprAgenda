
# insprAgenda

  

**Structure of the tasks:**

- ID (int)

- Title (string)

- Description (string)

- Priority (int)

  - 1(min) to 10(max)

- Deadline (date)

- Time estimate (date)

  - HH:mm

- Dependency ([int])

- Status (int)

  - To Do (1)

  - Working (2)

  - Closed (3)

  - Done (4)

- Work start date (date)

- Work end date (date)

  

**User Methods/Functionalities:**

- CreateTask

  - Params: (Title, Description, Priority, Deadline, Time Estimate, Dependency)

- RemoveTask

  - Params: (ID)

- MoveOnBoard

  - Params: (ID, Status)

- GetTimeSpent

  - Params: (ID)

- ListTasksBy

  - Params: (Filter) - Filter options: Priority, Deadline or Added time

- ShowTaskDetails

  - Params: (ID)

  

**Menu divisions:**

- Main menu

  - Options: See board, Manage board, Show calendar, Exit

- Board visualization menu

  - Options: By priority, By deadline, By Added time

- Board management menu

  - Options: Create task, Remove task, Update task, Show task details