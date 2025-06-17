## GoGinInitializer

### About
GoGinInitializer is a small program that helps quickly set up a new Go project using Gin framework.
It generates a basic project structure with essential files.

### Features
- Initializes a new Go project with Gin framework
- Creates a basic directory structure
- Generates essential files like `main.go` and `go.mod`
- Generates a sample route handler
- Generates a simple Error handler with custom errors
- Has additional feature that generates full api with --full-api flag, it generates additional directories like models/contracts, models/entities and repositories
- Has additional feature that initializes git repository with `--git` flag
- Has GUI for easy project setup with `--gui`/ `-g` flag or double click on executable

### Requirements
- Go 1.24.4

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/ErVaaL/GoGinInitializer.git
   ```
2. Navigate to the project directory:
   ```bash
    cd GoGinInitializer
   ```
3. Install dependencies:
   ```bash
    go mod download
   ```
4. Build the project:
   ```bash
    go build 
   ```
   or build version with GUI
   ```bash
    go build -tags gui
   ```

5. Run the executable:
   ```bash
    ./GoGinInitializer
   ```
   or double click on the executable if you are using GUI.

### Usage
To initialize a new Go project with Gin framework, run the following command:
```bash
./GoGinInitializer <project_name> [--full-api] [--git]
```
or
```bash
./GoGinInitializer [--gui/-g]
```
or
Launch the GUI by double clicking on the executable and fill in the input fields, then click on the "Generate" button.


