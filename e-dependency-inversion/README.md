#  Dependency Inversion Principle (DIP)

Types of explicit DI:


## Constructor Example

Most common and generally recommended approach. You pass the dependencies as arguments to the constructor of your struct
- **Advantages:** Clear dependencies, easy to test (you can pass mock loggers), enforces dependencies at creation time.
- **Disadvantages:** Can become verbose if you have many dependencies.

Run the code 
```bash
go run ./e-dependency-inversion/constructor/main.go
```


## Setter Example

Provide setter methods to inject dependencies after the object is created
- **Advantages:** Allows for optional dependencies, more flexible in some scenarios.
- **Disadvantages:** Dependencies might not be set, requires nil checks, can make dependencies less obvious. Generally less preferred than constructor injection.

Run the code 
```bash
go run ./e-dependency-inversion/setter/main.go
```
