quickServer

# Introduction

This software add a http static file server, using go and net/http, to windows right click menu.

# Usage
1. Download and install.
2.  
    - Right click a folder and choose "Open Folder With quickServer".
	- Open a folder, right click the background and choose "Open Folder With quickServer".
3. Input new host and port such as 0.0.0.0:80 or press enter to use default.

# Build from source
1. Build an exe.  
	`go build -o .\build\quickServer[64|32].exe .`

2. Using Visual Studio to add right click function.  
	Open "Setup32.sln" or "Setup64.sln" with Visual Studio. Change file path if you need it. Then build new msi/exe file.
# Others
Welcome to use, test and modify it.