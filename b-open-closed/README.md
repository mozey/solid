# Open/Closed Principle (OCP)


## Greet Example

Demonstrates how types are open for extension by embedding.

Run the code 
```bash
go run ./b-open-closed/greet/main.go
```


## OctoCat Example

The method set of a type cannot be altered by embedding it into other types, and is therefore closed to modification.

Run the code 
```bash
go run ./b-open-closed/octocat/main.go
```


## Shape Example

Use common interfaces to add functionality without changing existing code.

Key points
- **Abstraction (Interface)**: The `Shape` interface defines a contract for all shape types
- **Open for Extension**: New shape types (like `Triangle`) can be added by simply implementing the `Shape` interface. No changes are required in the `AreaCalculator`. This makes the system flexible and maintainable
- **Closed for Modification**: The `AreaCalculator` class is closed for modification. You don't need to change its code to handle new shape types. It works with any type that satisfies the `Shape` interface

Run the code 
```bash
go run ./b-open-closed/shape/main.go
```

