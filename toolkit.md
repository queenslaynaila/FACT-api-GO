# Building a Simple REST API in Go

## 1. Title & Objective
**Technology Chosen:** Go (Golang)

**Why I Chose It:**
My background is in Node.js and its backend frameworks like Express and Fastify, primarily working with microservices. I wanted to try a new language where I could apply concepts I’m familiar with, and that is well-suited for microservices.

While I like my current stack, I often run into generics hell in TypeScript where complex generic types make the code harder to read and maintain, especially when optimizing internal behaviors that require stricter type safety.Because of this, I was looking for a language with a simpler, more explicit type system, minimal syntax, and less boilerplate. Go emerged as a strong candidate.

Go’s strong, static type system, built-in concurrency support, simplicity, and performance, along with its popularity in microservice architectures, made it an appealing choice for me.

**End Goal:**
Use Go’s built-in HTTP capabilities to create a simple REST API showing random facts. The facts will be defined in the project.

---

## 2. Quick Summary of the Technology
**What is it?**  
Go is an open-source programming language developed by Google. It’s statically typed, compiled, and designed for simplicity, concurrency, and speed.

**Where is it used?**  
Go is widely used for building backend APIs, microservices, CLI tools, and infrastructure tools.

**Real-World Examples:**  
Docker, Kubernetes, Terraform

---

## 3. System Requirements
- **OS:** Linux/Windows
- **Tools/Editors:**  
  - VS Code recommended with Go extension https://marketplace.visualstudio.com/items?itemName=golang.go
- **Packages:**  
  - No external packages required for this example (we’ll use Go’s built-in `net/http`).

---

## 4. Installation & Setup Instructions
For other OS GO TO  https://go.dev/dl/   Download the installer for your OS and follow the setup instructions.
### Linux

**1. Download Go**  
Navigate to [https://go.dev/dl/](https://go.dev/dl/) and copy the URL for the latest binary release tarball.  

Retrieve the tarball using curl. We'll use the latest stable release 1.24.6
```bash
curl -OL https://go.dev/dl/go1.24.6.linux-amd64.tar.gz
```
Extract the tarball

```bash
sudo tar -C /usr/local -xvf go1.24.6.linux-amd64.tar.gz
```

Set Go env paths, so the system knows where to look for it. We do this by editing .profile file usually stored in home directory . Open the file with nano

```bash
sudo nano ~/.profile
```

Add the following lines to the end of the file:

```bash
export PATH=$PATH:/usr/local/go/bin
```
Close the file. Save Changes. Refresh file with 

```bash
source ~/.profile
```

Confirm installation by running 

```bash
go version
```

expected output 

```bash
go version go1.24.6 linux/amd64
```

---

## 5. Minimal Working Example

Create a new file called `main.go` and paste the following code:

```go
package main

import "fmt"

func main() {
  fmt.Println("Welcome to Go!")
}
```

Run it with:

```bash
go run main.go
```

It will print Welcome to Go on terminal


## 6. AI Prompt Journal

### 🔹 Prompt 1  
**Prompt Used:**  
> Give me a step-by-step guide to install Go on Ubuntu 22. Version installed must be latest stable release.  

**Link to Curriculum:**  Learning with AI

** AI’s Response Summary:**  
The AI gave me terminal commands to install Go version `go1.24.6`.  

**Helpfulness:** 5/5  
Installation proceeded seamlessly.

---

### 🔹 Prompt 2  
**Prompt Used:**  
> I'm ready to make my first Go project. A simple HTTP server that shows random facts (the facts will be defined in the project as an array — no third-party API or database connection needed). Could you guide me through it? Please explain the syntax. Also, coming from Express, are there any differences I should note?  

**Link to Curriculum:**  Learning New Language with AI-Guided Implementation

**AI’s Response Summary:**  
It outlined server and route setup differences in both stacks — for example:  
- Express: `app.get('/path', handler)`  
- Go: `http.HandleFunc("/path", handler)`  

It also highlighted differences in error handling and middleware behavior.  
The AI wrote a `main.go` file using Go’s built-in `net/http` package, defined an array of strings, used `math/rand` and `time` packages to return a random fact on each request, and showed how to run the file.  

**Helpfulness:** 5/5  
Code was explained line by line.

---

### 🔹 Prompt 3  
**Prompt Used:**  
> I want to refactor the code so that no fact is repeated until all the facts in our facts array have been exhausted. How would I achieve that in Go?  

**Link to Curriculum:**  Using AI to Refactor

**AI’s Response Summary:**  It showed me how to refactor the code so each fact is shown only once before any are repeated by shuffling the array and tracking the index.  It also suggested useful feature expansions like:  
- `/facts` – list all facts  
- `/add-fact` – POST route to add new facts  
- Switching to JSON responses  

**Helpfulness:** 5/5

---

### 🔹 Prompt 4  
**Prompt Used:**  
I want to make this code more readable and maintable [code] Please help by idnetifying difficylt parts, pointing out better var names, identify complex sections to break down and point out any incosistent stype eg naming formating issues? Do you have any linting recomendations for go? 

**Link to Curriculum:** [Using AI to Refactor – Code Readability Improvement](https://ai.moringaschool.com/ai-software/ai-use-cases/usecases-refactor/)

**AI’s Response Summary:** 
Pointed out better variable names, and suggested breaking down some parts
Outlined various Go linters and how to use them.  
I chose `gofmt`, the top recommended one.  
Also recommended using the Go VS Code extension to aid with formatting and linting.  

**Helpfulness:** 3/5  
**Feedback:** **I found the recommended refactored code to have too much abstraction. I ended up using only about half of the abstraction suggested, but the naming was on point.**
---

### 🔹 Prompt 5  
**Prompt Used:**  Im learning how to test in Go and i want to understand what behaviours to test., Help me make a testing plan by asking me questions about the codes behavioir , For each behaviour ask me how i would test it and give me hints if i miss sth ? 

**Link to Curriculum:** [Using AI for Testing – What to Test](https://ai.moringaschool.com/ai-software/ai-use-cases/usecases-testing-simpler/)

**AI’s Response Summary:**  
The AI asked me questions about the general code behaviour eg whats the function to do, and what happens after a user has requested all facts, how does the function decide which facs to present. After i replied it sent me behaviours ive missed and gave me tests suggestions, it also had a basic guide to testing in go, how the test works and how to run it using `go test`.  
**Helpfulness:** 5/5  


## 7. Common Errors

### 1. `go: command not found`
**Cause:** Go is not installed correctly or user did not set `$PATH` is not set.

**Fix:**
- Ensure Go is installed in `/usr/local/go`.
- Check if you added this line to your `.profile` (or `.bashrc` / `.zshrc` for other os):

    ```bash
    export PATH=$PATH:/usr/local/go/bin
    ```

- Make sure to Apply the changes with:

    ```bash
    source ~/.profile
    ```

### 2. Syntax Errors
- Use the following command to auto-format your Go code and catch common syntax issues.Make sure uve imported the fmt package :
  ```go
  import ( "fmt")
  ```
- Then run 
  ```bash
  go fmt
  ```

### 3. `go: cannot find main module, but found....`
Even if the project was bult and ran without a module go tools like go test go get expects a project to be a module as when you run go test go is looking for go.mod under the hood so make the project a module by runing 
**Fix:**
 ```bash
 go mod init <module-name>
 ```

## Technology Comparison GO vs Express

| Feature            | Go (net/http)                                 | Express (Node.js)                        |
|--------------------|-----------------------------------------------|------------------------------------------|
| **Routing**        | `http.HandleFunc("/", handler)`               | `app.get("/", handler)`                  |
| **Middleware**     | Manual chaining, explicit                     | Built-in, easy chaining                  |
| **Ecosystem**      | Smaller, focused, Go standard library stdlib already has a lot of what we need  so we dont always need extra packaging and so theres less reliance on 3rd parties, package ecosystem exists but its more focused and doesnt have small libraries and their duplicates like npm,             | Huge, npm registry, many packages for almost anything but this also means most pkgs are overlaping. In general theres overrelianace on many 3rd party pkgs        |
| **Error Handling** | Explicit and local to each handler. If a handler does not handle an error, it is ignored—no global error handler by default. | Global error-handling middleware can catch errors not handled in a route. |
| **Deployment**     | Code is compiled to a Single binary, that runs directly with no runtime needed                | Needs Node.js runtime                    |
| **Function & Var Syntax** | `func` defines the function. Type comes after name: `var name string = "Queenslay"` | `function` defines the function. Type comes after colon: `let name: string = "Queenslay"` |
| **Type Declaration** | Struct tags: `type FactResponse struct { }` | Interfaces: `interface FactResponse { }` |
| **Arrays**         | `[]type` — e.g., `[]string{"a","b"}`           | `type[]` or `Array<type>` — e.g., `["a","b"]` |
| **Handler Signature** | `func handler(w http.ResponseWriter, r *http.Request)` | `(req: Request, res: Response) => { ... }` |
| **Entry Point**    | Must have `package main` and `func main()`, which the Go runtime calls automatically | No `main()` — Node.js executes the entry file (e.g., `index.js`) from top to bottom |
| **Hot Reloading**     | Needs tools like `air` or `reflex`. | Common with `nodemon` or `ts-node-dev`. |
| **Testing** | Built in `go test` auto discoveres `_test.go` files for testing and runs the testing func starting with `Test` | No built in testing must use a test framework eg `jest` `mocha`. The framework willl either disover the test files based on its file naming patters or config but most times you must explicetly point the file paths to it |
