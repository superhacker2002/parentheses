## Parentheses
Project that contains web service for generating string with random parentheses of the given length and client for reporting percentage of balanced and unbalanced generated strings.

## Usage
**1.** Download repository from GitHub  
**2.** Set up environment variables:
- Create .env file
- Set `PORT` variable to port on which the server will listen for incoming connections.  
- Set `SERVER_URL` to the URL or address where the web server is located.  

**3.** Run web service:
```go
go run cmd/generator/generator.go
```
**4.** Run client and get a report:
```go
go run cmd/reporter/reporter.go
```

## Description
#### 1 - Algorithm

The function of package `parentheses` verifies if the given string is a balanced sequence of brackets, i.e. each of the open brackets must be closed by the same type of bracket in the right order:
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
