# gitl

`gitl` is an environment manager designed for developers. It helps you manage different development environments efficiently with support for jason and go configurations. 

> ⚠️**Note:** `gitl` is currently in the early stages of development. As such, there may be some ``bugs`` or incomplete features. Contributions are welcome to help improve the project!

## Features

- Manage and configure development environments.
- Supports json configurations for easy setup.
- Designed with security in mind, preventing unauthorized access and harmful activities.

## gitl comment's 
The `gitl init` command is used to initialize a new project environment. When run, it creates a `.gitl` directory in the root of your project. This directory is used to store the project description, configurations, and other relevant files that are part of the `gitl` environment setup.
``` bash
gitl init
```
The `gitl fsetup` command is used to configure the user's information, such as the username and email, for the `gitl` environment. This information is saved for future use and can be helpful when managing environments or project configurations within `gitl`.
```bash
gitl fsetup
```
___
### Prerequisites

- Go (version 1.24.1 or higher)
- json parser

### Steps

1. Clone the repository:

   ```bash
   git clone https://github.com/neel-xyt/gitl.git
   cd gitl && go run cmd/gitl/main.go
   ```
   - This will execute the main gitl command.
### Contributing
As `gitl` is in its early development stage, contributions are highly appreciated! Feel free to fork the repository, improve the code, or suggest new features via pull requests.
   
