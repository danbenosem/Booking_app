
Booking Application

A lightweight console-based event ticket booking application built in Go. This project explores fundamental Go design concepts, focusing on package management, structural visibility, data input handling, and separation of concerns.

🚀 Core Concepts Learned

• Package Management: Learned how Go organizes code by directory names rather than file names, avoiding circular dependencies.
• Visibility Rules: Used uppercase field names (FirstName, LastName) to export fields publicly across packages, replacing Java-style getters and setters with direct field access.
• Console Scanning: Utilized pointer-based memory allocation via fmt.Scanln(&user.Field) to accept user terminal inputs directly into structural fields.
• Separation of Concerns: Deconstructed the application into custom sub-modules (model) to isolate application state from terminal workflows (main).

📁 Project Structure

text
Booking_app/
├── go.mod                # Tracks dependencies and names the project module
├── collectDetails.go     # Part of package main; prompts user for entry details
├── main.go               # Entry point of the application
└── model/
    └── models.go         # Declares the public User struct and its data fields


🛠️ How To Run

1. Initialize your Go module tracking workspace (if you haven't already):bash
go mod init Booking_app
Use code with caution.
2. Build and run the multi-file application from your project root:bash
go run .


