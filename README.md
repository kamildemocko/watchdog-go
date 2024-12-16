# Watchdog-Go

## Description

Watchdog-Go is a process monitoring tool written in Go. It tracks the start and end times of specified processes on your system and logs these events to a CSV file with option to add more engines to support other formats. This tool is useful for monitoring the activity of specific applications, ensuring they are running as expected, and keeping a historical record of their execution times.

## Installation

### Prerequisites

- Go
- Git

### Steps

1. Clone the repository:
    ```sh
    git clone https://github.com/kamildemocko/watchdog-go.git
    cd watchdog-go
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Build the project:
    ```sh
    make build-prod
    ```

4. Use executable:
    ```sh
    bin/app
    ```

## Usage

1. Configure the processes to monitor by editing the `settings.toml` file:
    ```toml
    processes = ["list.exe", "of.exe", "processes.exe"]
    log_file = "./log/watchdog.log"
    refresh_seconds = 1
    ```

2. Run the application:
    ```sh
    make run
    ```

3. To stop the application:
    ```sh
    make stop
    ```

### Example

Here is an example of the log file output (`watchdog.log`):
```csv
Event,Timestamp,Pid,Name,Exe,Cmd,CreateTime,Seconds
start,2024-12-16T15:44:14.4682262+01:00,34888,python.exe,C:\Users\kamil\AppData\Local\Programs\Python\Python312\python.exe,"C:\Users\kamil\AppData\Local\Programs\Python\Python312\python.exe ""F:\Development\Python\Mood.py"" ",2024-12-16T15:43:36.329+01:00,38
end,2024-12-16T15:44:20.4814404+01:00,34888,python.exe,C:\Users\kamil\AppData\Local\Programs\Python\Python312\python.exe,"C:\Users\kamil\AppData\Local\Programs\Python\Python312\python.exe ""F:\Development\Python\Mood.py"" ",2024-12-16T15:43:36.329+01:00,44
```

## License

This project is licensed under the MIT License. 
