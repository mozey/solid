# Single Responsibility Principle (SRP)

## User example

Key points
- **Clear Separation of Concerns**: Each struct (`User`, `UserValidator`, `UserSerializer`, `UserRepository`) has a single, well-defined responsibility. This makes the code easier to understand, maintain, and modify
- **Reduced Coupling**: Changes to one struct are less likely to affect other structs. For example, you could *change the serialization format without affecting the user validation* or storage logic
- **Improved Testability**: Each struct can be tested independently. You can easily write unit tests for the `UserValidator` without needing to involve the database or serialization logic
- **Increased Reusability**: The structs can be reused in different parts of the application or even in other projects. For example, the `UserValidator` could be used in multiple contexts where user data needs to be validated
- **Easy to Extend**: If you need to add new functionality (e.g. a new validation rule, or a new serialization format), do so by creating a new struct or modifying an existing one without affecting other parts of the code

Run the code 
```bash
go run ./a-single-responsibility/user/main.go
```