# Building a Simple REST API in Go – A Beginner’s Guide

## 1. Title & Objective
**Technology Chosen:** Go (Golang)

**Why I Chose It:**
My background is in Node.js and its backend frameworks like Express and Fastify. Within this frameworks 
for a long time ive been working with microservices so I wanted to try a new language where i can apply concepts i already understand.And the language has to also be very miscroservice friendly.

While i also love my stack my main issue is always generics hell  expecially when modifying internal behaviours to optimise some performance aspects. So i was also specifically looking also for a language with simpler type system i wanted one priortising explict type behaivour ,prefarably minimalistic syntax, less boilerplate. And go came up as a top recommended option.

Go’s strong type system, built-in concurrency, simplicity, and performance and its popularity in miscroservice architecture just made it appealing.

**End Goal:**
Use Go’s built-in HTTP capabilities to create a simple REST API showing random facts. The facts will be defined in the project. No repeats until all facts used up.

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
- **OS:** Linux
- **Tools/Editors:**  
  - VS Code recommended with Go extension https://marketplace.visualstudio.com/items?itemName=golang.go
- **Packages:**  
  - No external packages required for this example (we’ll use Go’s built-in `net/http`).

---

## 4. Installation & Setup Instructions
For other OS GO TO  https://go.dev/dl/   Download the installer for your OS and follow the setup instructions.
AI prompt journal 
### Linux

**1. Download Go**  
Navigate to [https://go.dev/dl/](https://go.dev/dl/) and copy the URL for the latest binary release tarball.  
For this example, we’ll use:

```bash
https://go.dev/dl/go1.24.6.linux-amd64.tar.gz
```

Retrieve the tarball using curl
```bash
curl -OL https://go.dev/dl/go1.24.6.linux-amd64.tar.gz
```
Extract the tarball

```bash
sudo tar -C /usr/local -xvf go1.24.6.linux-amd64.tar.gz
```

Set Go env paths. We do this by editing .profile file usually stored in users home firectory . Open the file with nano

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

## 5. Compile & Run the Project

**Run the server directly (without building):**

```bash
go run main.go
```

**Build the server:**

```bash
go build -o random-facts-api main.go
```
On successful build, you’ll see a new file named random-facts-api created in your current directory.

**Run the built executable**
```bash
./random-facts-api
```
The server starts at port 8080

To see a random fact http://localhost:8080
To see all random facts  http://localhost:8080/all