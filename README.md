# junosOps
![Static Badge](https://img.shields.io/badge/Project-IN_PROGRESS:V1.0.0-orange) 
![Static Badge](https://img.shields.io/badge/Go-blue) 

| **⚠️ WARNING: This project is still in progress** ⚠️ |

junosOpst is a tool designed for network administrators and engineers to automate the auditing of network devices. This application leverages SSH for remote execution of network commands, enabling the collection of interface data directly from the network devices. 

## Current Capabilities:

Command Execution: Currently, the application supports the following SSH commands for data collection:
- ![Static Badge](https://img.shields.io/badge/COMPLETED-green) [JUNOS - show interfaces | display xml]

The tool connects to a list of network devices via SSH, executes specific commands to retrieve interface information, and processes the data to identify interfaces that meet certain criteria. The results are logged and summarised for further analysis.

### Key Features
- Concurrency:
Utilises a worker pool to handle multiple devices concurrently, significantly reducing the overall execution time
- Flexible Thresholds:
Allows the user to specify a time threshold (e.g., 2m for 2 minutes, 1h for 1 hour) to filter interfaces based on the last flap time.
- Logging:
Logs are written to a file for persistent storage and analysis
- Reporting:
Provides execution time, the number of devices processed, total interfaces processed, and the count of interfaces that meet the criteria.
Ensures that interfaces are only reported if they are in the 'down' state and have a non-blank description.

### Run the Tool:
- Execute the tool with required flags: `--u <username> --p <password> --t <threshold>`
Example: `junosOps --u admin --p password --t 1h`
- Input Devices:
The tool prompts the user to enter a list of device hostnames or IP addresses to monitor.
- Results:
Logs and results are written to `junosOps-application.log`
Detailed statistics are printed to the console.
