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
> I want to refactor the code so that no fact is repeated until all the facts in our facts array have been exhausted. How would I achieve that in Go? I think this project is doing too little — do you have any extra feature changes we can add to give it some depth?  

**Link to Curriculum:**  Using AI to Refactor

**AI’s Response Summary:**  It showed me how to refactor the code so each fact is shown only once before any are repeated by shuffling the array and tracking the index.  It also suggested useful feature expansions like:  
- `/facts` – list all facts  
- `/add-fact` – POST route to add new facts  
- Switching to JSON responses  

**Helpfulness:** 5/5

---

### 🔹 Prompt 4  
**Prompt Used:**  
Now that we want to add new features, before I get started: Does Go have an inbuilt linting tool like ESLint to help me catch syntax errors while doing so?  

**Link to Curriculum:**  *Using AI to Refactor*

**AI’s Response Summary:**  
Outlined various Go linters and how to use them.  
I chose `gofmt`, the top recommended one.  
Also recommended using the Go VS Code extension to aid with formatting and linting.  

**Helpfulness:** 5/5  
The linter worked and helped me catch a couple of syntax errors I had.

---

### 🔹 Prompt 5  
**Prompt Used:**  I want to test no repetition behaviour until fact list is exhausted this project, Can you  guide me throigh it? 

**Link to Curriculum:** Using AI for Testing and Iteration*

**AI’s Response Summary:**  
The AI provided a step-by-step guide to create a `main_test.go` file using Go's built-in `testing` package.  
It explained how the test works and how to run it using `go test`.  
**Helpfulness:** 5/5  
The test ran successfully

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