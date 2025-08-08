## Toolkit Document

Find my toolkit document at [toolkit.md](toolkit.md)

---

## Running Instructions

1. **Clone the project**

2. **Change into the project directory**

   ```bash
   cd random-facts-api
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

### Endpoints

- To see a random fact: [http://localhost:8080](http://localhost:8080)
- To see all random facts: [http://localhost:8080/all](http://localhost:8080/all)

---
