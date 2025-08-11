## Toolkit Document

Find my toolkit document at [toolkit.md](toolkit.md)

---

## Running Instructions

1. **Clone the project**

2. **Change into the project directory**

   ```bash
   cd FACT-api-GO
   ```

3. **Run the server directly (without building):**

   ```bash
   go run main.go
   ```

4. **Build the server:**

   ```bash
   go build -o random-facts-api main.go
   ```
   On successful build, you’ll see a new file named `random-facts-api` created in your current directory.

5. **Run the built executable:**

   ```bash
   ./random-facts-api
   ```

The server starts at port **8080**.
The server starts at port **8080**.

### Endpoints

- To see a random fact: [http://localhost:8080](http://localhost:8080)
- To see all random facts: [http://localhost:8080/all](http://localhost:8080/all)

---

## Testing Instructions

1. **Make sure you are in the project directory.**

2. **Run all tests with:**
   ```bash
   go test
   ```
   This will automatically discover and run all tests in files ending with `_test.go`.

3. **What to expect:**
   - If all tests pass, you’ll see output like:
     ```
      ok  	random-facts-api	0.004s
     ```
   - If a test fails, you’ll see details about the failure.
5. **How to run a specific test:**
   - Use the `-run` flag with `go test` to run a specific test function by name. For example:
     ```bash
     go test -run TestNoImmediateRepetition
     ```
   - This will only run the test function named `TestNoImmediateRepetition` in the test file.
