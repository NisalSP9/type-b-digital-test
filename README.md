# type-b-digital-test

### How to run the application.

#### 1. Clone the repository
```bash
git clone https://github.com/NisalSP9/type-b-digital-test.git
cd type-b-digital-test
```
#### 2. Initialize the Go module (if not already done)
```bash
go mod tidy
```
#### 3. Run the application
```bash
go run .\cmd\server\ 
```
#### The server will start on port 8080:
```bash
Server running on port 8080
```
#### 4. Check server health
```bash
curl "http://localhost:8080/health"     
```
Response:
```json
{
  "go version": "go1.22.4",
  "message": "TypeBDigital test server is up and running",
  "number of goroutines": 3,
  "status": "OK",
  "time": "2025-11-12T00:39:43+05:30"
}
```
#### 5. Test the API using curl or a browser
##### Valid request
```bash
curl "http://localhost:8080/hello-world?name=Alice  
```
##### Response:
```json
{
    "message": "Hello Alice"
}
```
##### Invalid request
```bash
curl "http://localhost:8080/hello-world?name=Zara
```
##### Response:
```json
{
    "error": "Invalid Input"
}
```
##### Missing name
```bash
curl "http://localhost:8080/hello-world?name=
```
##### Response:
```json
{
    "error": "Invalid Input"
}
```
### How to run the application.

```bash
go test ./internal/handler -v
```
#### Example output
```bash
--- PASS: TestCheckName (0.00s)
    --- PASS: TestCheckName/ValidFirstHalf_#1 (0.00s)
    --- PASS: TestCheckName/ValidFirstHalf_#2 (0.00s)
    --- PASS: TestCheckName/InvalidSecondHalf_#1 (0.00s)
    --- PASS: TestCheckName/InvalidSecondHalf_#2 (0.00s)
    --- PASS: TestCheckName/InvalidSymbol_#1 (0.00s)
    --- PASS: TestCheckName/InvalidSymbol_#2 (0.00s)
    --- PASS: TestCheckName/EmptyName (0.00s)
PASS
```
### Assumptions Made

1. Input is provided via a name query parameter (e.g. /hello-world?name=Alice).

2. Only English alphabet characters (A–Z or a–z) are considered valid for comparison.

3. The first letter determines validity:

    'a' to 'm' → accepted.

    'n' to 'z' → rejected.

    Empty or non-alphabetic names return a 400 Bad Request.
