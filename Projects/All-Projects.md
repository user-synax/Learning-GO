# 1. CLI Todo App

Concepts
slices
structs
functions
loops
conditionals
range
Features
Add task
List tasks
Mark completed
Delete task
Example Struct
type Task struct {
Title string
Done bool
}
Practice Focus
Updating slice elements
Looping through structs
Function separation

# 2. Student Grade Manager

Concepts
maps
structs
loops
functions
Features
Store student marks
Calculate average
Find topper
Show pass/fail
Example
map[string]int
Practice Focus
Real understanding of maps
Key-value operations
Conditional logic

# 3. Expense Tracker

Concepts
slices
structs
functions
variadic functions
Features
Add expenses
Show total
Category-wise spending
Monthly summary
Example Struct
type Expense struct {
Title string
Amount float64
}
Bonus

Use variadic function:

func total(nums ...float64) float64

# 4. Simple Banking System

Concepts
pointers
structs
methods mindset
conditionals
Features
Deposit
Withdraw
Balance check
Important Practice

Use pointers to update balance:

func Deposit(balance \*float64, amount float64)

This project teaches mutation through pointers clearly.

# 5. Contact Book

Concepts
maps
structs
slices
loops
Features
Save contacts
Search contact
Delete contact
Show all contacts
Example
map[string]Contact

Very common production pattern.

# 6. Quiz Game

Concepts
loops
switch
functions
arrays/slices
Features
Ask MCQs
Store score
Final result
Example
type Question struct {
Question string
Options []string
Answer string
}
Practice Focus
Nested loops
Struct handling
User input flow

# 7. Inventory Management System

Concepts
structs
maps
slices
pointers
Features
Add products
Update stock
Remove items
Search products
Example
type Product struct {
Name string
Stock int
}
Real-World Value

Very close to backend CRUD logic.

# 8. Library Management System

Concepts
structs
slices
functions
loops
pointers
Features
Add books
Borrow books
Return books
List available books
Example
type Book struct {
Title string
Borrowed bool
}
Good For

Understanding state updates.

# 9. CLI Calculator

Concepts
switch-case
functions
variadic functions
Features
Add
Subtract
Multiply
Divide
Bonus

Add:

multiple number addition
percentage
modulus
Practice Focus
Function organization
Input handling
Switch mastery

# 10. Mini URL Shortener Logic (Without Backend)

Concepts
maps
functions
random generation
structs
Features
Generate short code
Store original URL
Retrieve original URL
Example
map[string]string
Why This Matters

This is your first “backend-style” project logic.

Recommended Build Order

Build them in this order:

Calculator
Quiz Game
Todo App
Student Grade Manager
Expense Tracker
Banking System
Contact Book
Library System
Inventory System
URL Shortener Logic
