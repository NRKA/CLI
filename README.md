# **CLI**

## **Introduction**
Console command functionality has been integrated into the project with an architecture that demands minimal adjustments when introducing new commands.

## Installation
Clone this repository
  ```bash
    git clone https://github.com/NRKA/CLI.git
```
## How to Run
```bash
  go run cmd/main.go <argument> <argument>
```
  Second argument depends on type of command, for extra information type ```go run cmd/main.go help```
## **Command Types**
- help command
- spell command
- reformat command

## **Command Overview**
  ### **Help Command**
    Shows all available commands. Usage: <command> 
    Example: go run cmd/main.go help
  ### **Spell Command**
    Receives one or more arguments and converts them into letters, separated by spaces. Usage: <command> <argument>
    Example: go run cmd/main.go spell hello
  ### **Reformat Command**
    Receives one .txt format argument and formats the data. Usage: <command> <argument>
    Example: go run cmd/main.go reformat example.txt
## **After running the project with the help command:**
  ```
  go go run cmd/main.go help 
  Available commands:                                                                                                      
  help - shows all available commands. Usage: <command>        
  reformat - receives one .txt format argument and formats the data. Usage: <command> <argument>
  spell - receives one or more arguments and converts them into letters, separated by spaces. Usage: <command> <argument>
  ```

## **Run unit tests**
```
  go test ./...
```

## **Run unit tests with coverage**
```
go test ./... -cover
```
