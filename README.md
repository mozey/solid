# SOLID

SOLID principles applied to Go with examples.

[Origins of SOLID, DDD, and Software Ethics](https://web.archive.org/web/20180824134615/https://www.infoq.com/podcasts/uncle-bob-solid-ddd/)
- Michael Feathers wrote to Bob (Robert C. Martin) and said if you rearrange the order of the design principles, it spells SOLID<sup>[[InfoQ Podcast](https://www.youtube.com/watch?v=hFtVBAxjeK0&t=130s)]</sup>
- Clean Architecture is a way to develop software with low coupling
- You find the bounded context of [DDD](https://en.wikipedia.org/wiki/Domain-driven_design) at the innermost circles of a clean architecture<sup>[[InfoQ Podcast](https://www.youtube.com/watch?v=hFtVBAxjeK0&t=503s)]</sup>
- Services do not form an architecture. They form a deployment pattern<sup>[[InfoQ Podcast](https://youtu.be/hFtVBAxjeK0?si=Zs28Kl0YIRgLnyzK&t=667)]</sup>
- Advantage of clean architecture is that, the highest level decisions (those that make or save the money), are unaffected by GUI, schema, and framework changes. Dependencies must point inwards<sup>[[InfoQ Podcast](https://www.youtube.com/watch?v=hFtVBAxjeK0&t=831s)]</sup>
- Protect the business rules from changes
- Sidecar (service talks to localhost) and service mesh (allows the sidecar to discover other services)<sup>[InfoQ Podcast](https://www.youtube.com/watch?v=hFtVBAxjeK0&t=987s)</sup>
- Where you want to live is between convenient code, and "clean" code<sup>[InfoQ Podcast](https://www.youtube.com/watch?v=hFtVBAxjeK0&t=1064s)</sup>
- Create the boundaries anyway, and then remove them if performance matters. Worse may be better, example of function calls taking nanoseconds
- Clean architecture requires some kind of indirection to cross boundaries, polymorphism (in an Object-Oriented language) or pointers to functions<sup>[InfoQ Podcast](https://www.youtube.com/watch?v=hFtVBAxjeK0&t=1625s)</sup>


## Go is not Object-Oriented

Go is not an [Object-Oriented](https://en.wikipedia.org/wiki/Object-oriented_programming) language, but it has some **similar features**
- **Methods**: Define methods on structs, similar to methods on objects
- **Interfaces**: Define a set of methods that a type must implicitly implement ([subtype polymorphism](https://en.wikipedia.org/wiki/Polymorphism_(computer_science)))
- **Encapsulation**: Control the visibility of fields and methods using capitalization. Use packages to organise code

Go **doesn't have**
- **Classes**: It has structs
- **Inheritance**: Promotes composition over inheritance

**Composition** in Go
- [Embedding](https://go.dev/doc/effective_go#embedding): Embedding types within a struct or interface. It's like saying *"this struct has-a"* instead of *"this struct is-a"*
- [Interfaces](https://go.dev/doc/effective_go#interfaces_and_types): Specify the behavior of an *"object"*: if something can do this, then it can be used here. Define a set of methods that a type must implement. Allows you to write code that works with any type that implements the interface, regardless of its underlying type. E.g. *io.Writer* interface has the *Write()* method

**Benefits of Composition**
- **Code Reusability**: Reuse existing types to build new ones, avoiding code duplication
- **Flexibility**: Easily swap out implementations by using interfaces. *Small interfaces lead to simple implementations*<sup>[[Go UK 2016](https://www.youtube.com/watch?v=zzAdEt3xZ1M&t=652s)]</sup>
- **Loose Coupling**: Components are less dependent on each other, making it easier to modify or replace them
- **Testability**: Smaller components are easier to test in isolation


## SOLID Principles

**SOLID** is an acronym that represents five key principles, that may be applied when taking a software engineering approach to *Object-Oriented Design*. These principles are aimed at making software more maintainable, scalable, and flexible. It helps developers create code that is easier to understand, modify, and extend over time.<sup>[[Wikipedia: SOLID](https://en.wikipedia.org/wiki/SOLID)]</sup>

**TL;DR** Applied to Go, we can summarise this approach as follows:<sup>[[Go UK 2016](https://www.youtube.com/watch?v=zzAdEt3xZ1M&t=1361s)] [[Go Time #16: SOLID Go Design](https://changelog.com/gotime/16)]</sup>

*Interfaces let you apply the SOLID principles to Go programs*

Here's a breakdown of each principle:


### S - Single Responsibility Principle (SRP)

**Concept**: 

*A class should have only one reason to change.* -- [Robert C. Martin](https://en.wikipedia.org/wiki/Robert_C._Martin), "Agile Software Development, Principles, Patterns, and Practices", 2003

This means that each class should have a single responsibility or job within the software system. If different actors (stakeholder or user) require different changes, those changes should be separated into different classes.

**Benefits**:
- Improved code organization and readability.
- Easier to maintain and modify individual classes.
- Reduced risk of unintended side effects when changes are made.

**Applied to Go**:
- Coupling: *two things that change together* & Cohesion: *pieces of code that naturally fit together*<sup>[[Go UK 2016](https://www.youtube.com/watch?v=zzAdEt3xZ1M&t=316s)]</sup>
- *Avoid miscellaneous packages*, e.g. package server, utils, common
- Structure functions, types, and methods into *packages that exhibit natural cohesion*, e.g. net/http, os/exec, and encoding/json<sup>[[Go UK 2016](https://www.youtube.com/watch?v=zzAdEt3xZ1M&t=1278s)]</sup>
- The *idea of a bounded context*, to *group things up by responsibility*, folders for cmd, internal, and domain models
- Go *stops you from importing from other projects’ internal folders*<sup>[[Go Time #273: Domain-driven design with Go](https://changelog.com/gotime/273)]</sup>
- *DDD can inform SRP*: Understanding the domain helps you identify the appropriate responsibilities for your classes. By focusing on domain concepts, you can avoid classes that are overloaded with unrelated tasks
- Use composition instead of inheritance
- Methods can be defined on structs
- Encapsulate private struct fields, use packages to organise code
- Types may be embedded

[SRP Examples](https://github.com/mozey/solid/tree/main/a-single-responsibility)
- Common Go Mistakes: *"type embedding can help avoid boilerplate code; ensure that doing so doesn’t lead to visibility issues where some fields should have remained hidden"*<sup>[[Possible problems with type embedding](https://100go.co/#not-being-aware-of-the-possible-problems-with-type-embedding-10)]</sup>


### O - Open/Closed Principle (OCP)

**Concept**: 

*Software modules should be open for extension, but closed for modification.* -- Bertrand Meyer, "Object-Oriented Software Construction", 1988

Software entities (classes, modules, functions, etc.) should be open for extension, but closed for modification. This means that you should be able to add new functionality to a class without altering its existing code.   

**Benefits**:
- Enables extending functionality without introducing bugs in existing code.
- Promotes code reusability and maintainability.
- Facilitates easier adaptation to changing requirements.

**Applied to Go**:
- Compose simple types with embedding rather than extending them through inheritance<sup>[[Go UK 2016](https://youtu.be/zzAdEt3xZ1M?si=mqvhWHYkr7Guybnd&t=1294)]</sup>
- Types are *open for extension* by embedding<sup>[[Go UK 2016](https://youtu.be/zzAdEt3xZ1M?si=mqvhWHYkr7Guybnd&t=484)]</sup>
- The method set of a type cannot be altered by embedding it into other types, and is therefore *closed to modification*
- Methods in Go are syntactic sugar around functions with a pre-declared formal parameter, that being it's receiver.
- Use common interfaces to add functionality without changing existing code

[OCP Examples](https://github.com/mozey/solid/tree/main/b-open-closed)
- Common Go Mistakes: *"Abstractions should be discovered, not created. To prevent unnecessary complexity, create an interface when you need it and not when you foresee needing it"*<sup>[[Interface pollution](https://100go.co/#interface-pollution-5)]</sup>


### L - Liskov Substitution Principle (LSP)

**Concept**: 

*Substitutability: if S is a subtype of T, then objects of type T can be replaced with objects of type S without altering the desirable properties of the program* -- [Barbara Liskov](https://en.wikipedia.org/wiki/Barbara_Liskov), "Data abstraction and hierarchy", 1987

Objects of a derived class should be substitutable for objects of their base class without altering the correctness of the program.

**Benefits**:
- Ensures that inheritance is used correctly and consistently.
- Prevents unexpected behavior when using subclasses.
- Supports polymorphism and code flexibility.

**Applied to Go**:
- Two types are substitutable if they exhibit behaviour such that the caller is unable to tell the difference<sup>[[Go UK 2016](https://youtu.be/zzAdEt3xZ1M?si=mqvhWHYkr7Guybnd&t=618s)]</sup>
- Types that implement same interface (implicitly) are interchangeable
- Express dependencies between packages in terms of interfaces, not concrete types<sup>[[Go UK 2016](https://youtu.be/zzAdEt3xZ1M?si=mqvhWHYkr7Guybnd&t=1304s)]</sup>
- Use interfaces for: *common behaviour*, *decoupling*, and *restricting behaviour*

[LSP Examples](https://github.com/mozey/solid/tree/main/c-liskov-substitution)


### I - Interface Segregation Principle (ISP)

**Concept**: 

*Clients should not be forced to depend on methods that they don't use.* -- Robert C. Martin

Instead, many specific interfaces are better than one general interface. This means that interfaces should be small and focused on specific sets of related methods.

**Benefits**:
- Reduces dependencies and coupling between classes.
- Improves code flexibility and maintainability.
- Prevents unnecessary implementations of methods.

**Applied to Go**:
- Interfaces are implemented implicitly in Go
- Small interfaces lead to simple implementations
- Accept interfaces, return structs<sup>[[Go UK 2016](https://youtu.be/zzAdEt3xZ1M?si=mqvhWHYkr7Guybnd&t=772s)]</sup>
- Define functions and methods that depend only on the behaviour that they need<sup>[[Go UK 2016](https://youtu.be/zzAdEt3xZ1M?si=mqvhWHYkr7Guybnd&t=1322s)]</sup>

[ISP Examples](https://github.com/mozey/solid/tree/main/d-interface-segregation)
- Common Go Mistakes: *"The bigger the interface, the weaker the abstraction"*<sup>[[Interface pollution](https://100go.co/5-interface-pollution/)]</sup>
- Common Go Mistakes: *"Keeping interfaces on the client side avoids unnecessary abstractions"*<sup>[[Interface on the producer side](https://100go.co/#interface-on-the-producer-side-6)]</sup>


### D - Dependency Inversion Principle (DIP)

**Concept**: 

*High-level modules should not depend on low-level modules. Both should depend on abstractions. Abstractions should not depend on details. Details should depend on abstractions.* -- Robert C. Martin

This means that classes should depend on interfaces or abstract classes rather than concrete classes.   

**Benefits**:
- Decouples classes and reduces dependencies.
- Increases code reusability and testability.
- Makes it easier to change implementations without affecting other parts of the system.

**Note**:
- The *DIP* guides how you structure the dependencies between different parts of the software
- *Dependency Injection (DI)*: is how you create and provision dependencies

**Applied to Go**:
- Avoid injecting implementations (structs), inject abstractions (interfaces)
- Types of explicit DI: *Constructor* or *Setter* (Method)
- Structure of the import graph must be acyclic, otherwise compilation error
- Push specifics as high as possible up the import graph<sup>[[Go UK 2016](https://youtu.be/zzAdEt3xZ1M?si=mqvhWHYkr7Guybnd&t=1142s)]</sup>
- Refactor dependencies from compile time to run time<sup>[[Go UK 2016](https://youtu.be/zzAdEt3xZ1M?si=mqvhWHYkr7Guybnd&t=1342s)]</sup>

[DIP Examples](https://github.com/mozey/solid/tree/main/e-dependency-inversion)
- Common Go Mistakes: *"Not using the functional options pattern"*. Use the functional options pattern to configure options in an API-friendly manner. Works well for dependency injection. It makes your constructors more flexible, readable, and maintainable, especially when dealing with multiple dependencies and configuration options, some of which may be optional<sup>[[Functional options pattern](https://100go.co/#not-using-the-functional-options-pattern-11)]</sup>

**DI container libs**:
- For complex projects with many dependencies, use a DI or Inversion of Control (IoC) container library. A framework that automates the process of dependency injection, instead of your code explicitly creating and injecting its dependencies
- [google/wire](https://github.com/google/wire): Operates without runtime state or reflection
- [uber-go/dig](https://github.com/uber-go/dig): Reflection based dependency injection toolkit, resolves the object graph during process startup
- [aziis98/go-sl](https://github.com/aziis98/go-sl): Uses generics to provide a type safe interface without using reflection, implements the [Service Locator](https://martinfowler.com/articles/injection.html#UsingAServiceLocator) pattern


### Benefits of SOLID Principles

Clean Architecture is a way to develop software with low coupling. The architecture is divided into layers, with the core business logic in the center and external concerns (like UI, database) in the outer layers. Dependencies flow inwards, towards the core business logic. Inner layers don't depend on outer layers.

SOLID principles provide guidelines for creating well-designed components that fit within the layers of Clean Architecture.

By applying both SOLID principles and Clean Architecture, you can create software that is:

**Maintainable**: 

Is easier to understand, modify, and maintain over time.

**Reusable**: 

Encourages the creation of reusable code components that can be used in different parts of the system.

**Flexible**: 

Enables code that is adaptable to changing requirements, and independent of the underlying technologies.

**Testable**: 

Easier to write test for specific functionality. Write unit tests for individual components. Integration and end-to-end tests for groups of components, swapping out dependencies for stubs (test doubles with pre-programmed responses) and mocks (test doubles with expectations) as required 

**Scalable**: 

Clean Architecture is easier to optimise, new features and requirements can be implemented without negatively impacting the performance of existing code.


## Error handling in Go

Consider the discussion in [Go Time #16](https://changelog.com/gotime/16):

All the developers on the team are encouraged to think about errors while writing the code. **Exceptions** allow you to delay thinking about how the error will be handled.

Error values may be compared to predefined **Sentinel errors**, e.g. `io.EOF` or `sql.ErrNoRows`. Do this sparingly, avoid comparing errors everywhere in the code. Usually it's enough to return the error all the way up the caller stack. Wrap and unwrap error value as required.

**Who** is the error message (value) for, the *end-user or developers*? Programs shouldn't say something unless there is an error. Stack traces are useful to developers, and may be included with the error.

**Structured logs** are not useful to the end-user, but are useful in development. RequestID (*tracing*) and UserID (*audit logs*) are examples where structured logging is especially useful.

Common Go Mistakes: [Handling an error twice](https://100go.co/#handling-an-error-twice-52). In most situations, an error should be handled only once. Logging an error is handling an error. Therefore, you have to choose between logging or returning an error. In many cases, error wrapping is the solution as it allows you to provide additional context to an error and return the source error.

Don't put metrics in the logs, keep these in a separate system.

Common Go Mistakes: [Unnecessary nested code](https://100go.co/#unnecessary-nested-code-2). Avoid many nested levels, it causes [bad line of sight](https://www.youtube.com/watch?v=yeetIgNeIkc&t=330s), try to keep the happy path aligned on the left. If possible, make the happy return the last statement. This makes it easier to build a mental mode of the code

[Don't just check errors, handle them gracefully](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully)
- Never inspect the output of `error.Error`
- **Sentinel errors** become part of your public API. They create a dependency between two packages
- An **error type** is a type that you create that implements the error interface, i.e. custom errors. Error types must be made public
- **Opaque errors**: is when you know an error occurred, but you don't have the ability to see inside the error
- Interactions with the world outside your process, like network activity, require that the caller investigate the nature of the error to decide if it is reasonable to retry the operation. *Assert errors for behaviour, not type*. Errors may implement the Temporary interface, then call `IsTemporary` to determine if the error could be retried
- Add context to errors, use [pkg/errors](github.com/pkg/errors). It has methods like `Wrap`, `Cause`, and `Is`. Can be used to add a stack trace to errors

**Summary**: 
- For maximum flexibility try to treat all errors as opaque. 
- In situations where you cannot do that, try to assert errors for behaviour, instead of type or value. 
- Minimise the number of sentinel error values

