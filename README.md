# SOLID

SOLID principles applied to Go with examples

Origins of the acronym: *"Michael Feathers wrote to Bob (Robert C. Martin) and said if you rearrange the order of the design principles, it spells SOLID"*<sup>[[1](https://github.com/mozey/solid/tree/main?tab=readme-ov-file#1-origins)]</sup>


## Examples

Go examples to demonstrate the SOLID principles:<sup>[[3](https://github.com/mozey/solid/tree/main?tab=readme-ov-file#3-go-patterns--examples)]</sup>
- [Single responsibility](https://github.com/mozey/solid/tree/main/a-single-responsibility)
- [Open closed](https://github.com/mozey/solid/tree/main/b-open-closed)
- [Liskov substitution](https://github.com/mozey/solid/tree/main/c-liskov-substitution)
- [Interface segregation](https://github.com/mozey/solid/tree/main/d-interface-segregation)
- [Dependency inversion](https://github.com/mozey/solid/tree/main/e-dependency-inversion)


## Go is not Object-Oriented

Go is not an [Object-Oriented](https://en.wikipedia.org/wiki/Object-oriented_programming) language, but it has some **similar features**
- *Methods*: Define methods on structs, similar to methods on objects
- *Interfaces*: Define a set of methods that a type must implement ([polymorphism](https://en.wikipedia.org/wiki/Polymorphism_(computer_science)))
- *Encapsulation*: Control the visibility of fields and methods using capitalization. Use packages to organise code

Go **doesn't have**
- *Classes*: It has structs
- *Inheritance*: Promotes composition over inheritance

**Composition** in Go
- [Embedding](https://go.dev/doc/effective_go#embedding): Embedding types within a struct or interface. It's like saying *"this struct has-a"* instead of *"this struct is-a"*
- [Interfaces](https://go.dev/doc/effective_go#interfaces_and_types): Specify the behavior of an object: if something can do this, then it can be used here. Define a set of methods that a type must implement. Allows you to write code that works with any type that implements the interface, regardless of its underlying type. E.g. *io.Writer* has the *Write()* method

**Benefits of Composition**
- *Code Reusability*: Reuse existing types to build new ones, avoiding code duplication
- *Flexibility*: Easily swap out implementations by using interfaces
- *Loose Coupling*: Components are less dependent on each other, making it easier to modify or replace them
- *Testability*: Smaller components are easier to test in isolation


## SOLID Principles

**SOLID** is an acronym that represents five key principles of object-oriented design in software engineering. These principles are aimed at making software more maintainable, scalable, and flexible. By adhering to SOLID principles, developers can create code that is easier to understand, modify, and extend over time.<sup>[[2](https://github.com/mozey/solid/tree/main?tab=readme-ov-file#2-solid)]</sup>

Here's a breakdown of each principle:


### S - Single Responsibility Principle (SRP)

**Concept**: 
A class should have one, and only one, reason to change. This means that each class should have a single responsibility or job within the software system.

**Benefits**:
- Improved code organization and readability.
- Easier to maintain and modify individual classes.
- Reduced risk of unintended side effects when changes are made.

**Applied to Go**:
- Use composition instead of inheritance
- Methods can be defined on structs
- Encapsulate private struct fields
- Types may be embedded

[Common Go Mistakes: Possible problems with type embedding](https://100go.co/#not-being-aware-of-the-possible-problems-with-type-embedding-10)


### O - Open/Closed Principle (OCP)

**Concept**: 
Software entities (classes, modules, functions, etc.) should be open for extension, but closed for modification. This means that you should be able to add new functionality to a class without altering its existing code.   

**Benefits**:
- Enables extending functionality without introducing bugs in existing code.
- Promotes code reusability and maintainability.
- Facilitates easier adaptation to changing requirements.

**Applied to Go**:
- Common interface to add functionality without changing existing code


### L - Liskov Substitution Principle (LSP)

**Concept**: 
Objects of a derived class should be substitutable for objects of their base class without altering the correctness of the program. In simpler terms, if you have a class (e.g., "Animal") and a subclass (e.g., "Dog"), you should be able to use a "Dog" object wherever an "Animal" object is expected.

**Benefits**:
- Ensures that inheritance is used correctly and consistently.
- Prevents unexpected behavior when using subclasses.
- Supports polymorphism and code flexibility.


### I - Interface Segregation Principle (ISP)

**Concept**: 
Clients should not be forced to depend on interfaces that they don't use. Instead, many specific interfaces are better than one general interface. This means that interfaces should be small and focused on specific sets of related methods.

**Benefits**:
- Reduces dependencies and coupling between classes.
- Improves code flexibility and maintainability.
- Prevents unnecessary implementations of methods.

**Applied to Go**:

[Common Go Mistakes: Interface on the producer side](https://100go.co/#interface-on-the-producer-side-6)


### D - Dependency Inversion Principle (DIP)

**Concept**: 
High-level modules should not depend on low-level modules. Both should depend on abstractions. Abstractions should not depend on details. Details should depend on abstractions. This means that classes should depend on interfaces or abstract classes rather than concrete classes.   

**Benefits**:
- Decouples classes and reduces dependencies.
- Increases code reusability and testability.
- Makes it easier to change implementations without affecting other parts of the system.


### Benefits of SOLID Principles

**Maintainability**: 
Makes code easier to understand, modify, and maintain over time.

**Scalability**: 
Promotes code that can be easily extended and scaled to accommodate new features and requirements.

**Flexibility**: 
Enables code that is adaptable to changing needs and technologies.

**Testability**: 
Easier to write unit tests for individual classes and components.

**Reusability**: 
Encourages the creation of reusable code components that can be used in different parts of the system.


## References

### [1] Origins

[Origins of SOLID, DDD, & Software Ethics](https://web.archive.org/web/20180824134615/https://www.infoq.com/podcasts/uncle-bob-solid-ddd/)

### [2] SOLID

[Wikipedia: SOLID](https://en.wikipedia.org/wiki/SOLID)

### [3] Go Patterns & Examples

[Go Examples: Introduction to SOLID](https://hackernoon.com/go-design-patterns-an-introduction-to-solid)

[Go Time #16: SOLID Go Design](https://changelog.com/gotime/16)

