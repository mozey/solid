# Liskov Substitution Principle (LSP)

## Decoupling Example

Key Points
- Clear interface: `CustomerStorer` requires the `StoreCustomer` method
- LSP compliant type: `QueryLogger`
- Interface decouples the implementation: use the concrete implementation for integration tests, use a mock for unit tests, create query logger, etc

Run the code 
```bash
go run ./c-liskov-substitution/decoupling/main.go
```
