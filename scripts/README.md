# Real-Time Performance Monitoring in Go

This repository contains a Go script that monitors and logs system performance metrics in real-time. The script uses the **gopsutil** library to extract critical system information such as CPU usage, memory usage, disk usage, and network usage, making it suitable for observing the performance of servers, virtual machines, containers, and other systems.

## 🚀 Features

- **CPU Usage**: Tracks the percentage of CPU usage over time.
- **Memory Usage**: Monitors the system's RAM and reports on usage.
- **Disk Usage**: Measures storage space consumption and disk I/O.
- **Network Usage**: Monitors the network throughput, showing data sent/received.

## 📦 Use Case

Useful for real-time performance monitoring in:
- Cloud environments (AWS, Azure, GCP)
- Containers (Docker, Kubernetes)
- Bare-metal or Virtual Machines

## 🛠️ Getting Started

Follow these instructions to set up and run the Real-Time Performance Monitoring script on your local machine.

### ✅ Prerequisites

- Go version **1.23** or higher  
- Internet access to download dependencies

### 🔧 Installing Go

1. Download Go from the official website: [https://go.dev/dl/](https://go.dev/dl/)
2. Follow the instructions for your OS (Windows, macOS, Linux)

### 📁 Clone the Repository

git clone https://github.com/yourusername/golang-learning.git
### 📂 Navigate to the Scripts Folder
```bash
cd golang-learning/scripts
```

### 🧱 Initialize Go Module
```bash
go mod init
go mod tidy
```
### 📥 Install Dependencies
```bash
go get github.com/shirou/gopsutil/cpu
go get github.com/shirou/gopsutil/mem
go get github.com/shirou/gopsutil/disk
go get github.com/shirou/gopsutil/net
```
### ▶️ Run the Script
``` bash
go run Real-Time\ Performance\ Monitoring.go
```
### 💡 Sample Output
```yaml
CPU Usage: 15.3%
Memory Usage: 70%
Disk Usage: 25%
Network Usage: 5.6MB/s
```
### 🤝 Contributing
- Feel free to fork the repo, raise issues, or submit pull requests for improvements or new features.

#### 📄 License
- This project is licensed under the MIT License.
