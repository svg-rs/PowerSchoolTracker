# 📌 PowerSchool Tracker

## 📖 Overview
PowerSchool Tracker is a Go-based automation tool that logs into PowerSchool, navigates to the missing assignments page, extracts the data, and saves it to a file. This helps students or parents easily track overdue work without manually checking the portal.

## 🚀 Features
- 🔐 **Secure Login** – Automates PowerSchool authentication
- 🔎 **Scrapes Missing Assignments** – Extracts missing assignments for a specific term (e.g., Q3)
- 💾 **Saves to File** – Stores missing assignments in a structured `missing.txt` file
- ⏳ **Automated Navigation** – Clicks through PowerSchool pages to retrieve data
- ⚡ **Fast & Efficient** – Runs Selenium automation for quick data extraction

## 🛠️ Setup

### 1️⃣ Prerequisites
- Install [Go](https://go.dev/doc/install)
- Install Google Chrome
- Download [ChromeDriver](https://sites.google.com/chromium.org/driver/)
  - Ensure that the ChromeDriver version matches your Chrome browser
  - Place the ChromeDriver file in `../chromedriver/`

### 2️⃣ Install Dependencies
Run the following command to install the required Selenium package:
```sh
go get github.com/tebeka/selenium
```

### 3️⃣ Configure and Run
Navigate to the project directory and execute:
```sh
go run .
```

## 📂 Output
After execution, the script generates a `missing.txt` file with missing assignments in this format:
```
Course: Math | Due Date: 2025-03-01 | Assignment: Algebra Homework
Course: Science | Due Date: 2025-03-03 | Assignment: Lab Report
```

## ⚠️ Important Notes
- Ensure **ChromeDriver** is properly installed and accessible.
- Your **PowerSchool login credentials** are required for the script to function.
- If PowerSchool updates its UI, element IDs or XPaths may need adjustments.

## 🎯 Future Improvements
- 🖥️ **Graphical User Interface (GUI)** – A more user-friendly experience
- 📧 **Email Notifications** – Get alerts for new missing assignments
- 📊 **Data Visualization** – Display trends of missing assignments
- 🔄 **Automated Scheduled Runs** – Set up daily or weekly tracking

---

Made with ❤️ in Go! 🚀
