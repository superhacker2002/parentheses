## Parentheses
Project that contains web service for generating string with random parentheses of the given length and client for reporting percentage of balanced and unbalanced generated strings.

## Usage
**1.** Get module from GitHub:
```go
go get "github.com/superhacker2002/parentheses"
```
**2.** Run web service:
```
go run cmd/generator/generator.go
```
**3.** Run client and get a report:
```
go run cmd/reporter/reporter.go
```

## Description
#### 1 - Algorithm

The function of package `parentheses` verifies if the given string is a balanceds sequence of brackets, i.e. each of the open brackets must be closed by the same type of bracket in the right order:
```
[(]) - unbalanced

{[{}]()} - balanced

((1 + 2) * 3) - 4)/5 - balanced
```
#### 2 - Service

Parentheses web service generates a random sequence of parentheses of the length n. Used the standard router for http requests.

GET /generate?n={length}

#### 3 - Report

Reporter calls the service N = 1000 times for each of the strings of the length 2,4,8, calculates the percent of balanced strings for each length and prints the results to stdout.
