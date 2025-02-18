# Open/Closed Principle (OCP)

## Shape Example

Key points
- **Abstraction (Interface)**: The `Shape` interface defines a contract for all shape types. This is a key aspect of OCP
- **Extension through Implementation**: New shape types (like `Triangle`) can be added by simply implementing the `Shape` interface. No changes are required in the `AreaCalculator`
- **Closed for Modification**: The `AreaCalculator` class is closed for modification. You don't need to change its code to handle new shape types. It works with any type that satisfies the `Shape` interface
- **Open for Extension**: The system is open for extension. You can easily add new shape types without altering existing code. This makes the system flexible and maintainable

Run the code 
```bash
go run ./b-open-closed/shape/main.go
```

