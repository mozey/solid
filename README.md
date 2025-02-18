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
- Create the boundaries anyway, and then remove then if performance matters. Worse may be better, example of function calls taking nanoseconds
- Clean architecture requires some kind of indirection to cross boundaries, polymorphism (in an Object-Oriented language) or pointers to functions<sup>[InfoQ Podcast](https://www.youtube.com/watch?v=hFtVBAxjeK0&t=1625s)</sup>


## Examples

Go examples to demonstrate the SOLID principles:
- [Dependency inversion](https://github.com/mozey/solid/tree/main/e-dependency-inversion)


## Go is not Object-Oriented

Go is not an [Object-Oriented](https://en.wikipedia.org/wiki/Object-oriented_programming) language, but it has some **similar features**
- *Methods*: Define methods on structs, similar to methods on objects
- *Interfaces*: Define a set of methods that a type must implicitly implement ([subtype polymorphism](https://en.wikipedia.org/wiki/Polymorphism_(computer_science)))
- *Encapsulation*: Control the visibility of fields and methods using capitalization. Use packages to organise code

Go **doesn't have**
- *Classes*: It has structs
- *Inheritance*: Promotes composition over inheritance

**Composition** in Go
- [Embedding](https://go.dev/doc/effective_go#embedding): Embedding types within a struct or interface. It's like saying *"this struct has-a"* instead of *"this struct is-a"*
- [Interfaces](https://go.dev/doc/effective_go#interfaces_and_types): Specify the behavior of an *"object"*: if something can do this, then it can be used here. Define a set of methods that a type must implement. Allows you to write code that works with any type that implements the interface, regardless of its underlying type. E.g. *io.Writer* interface has the *Write()* method

**Benefits of Composition**
- *Code Reusability*: Reuse existing types to build new ones, avoiding code duplication
- *Flexibility*: Easily swap out implementations by using interfaces. *Small interfaces lead to simple implementations*<sup>[[Go UK 2016](https://www.youtube.com/watch?v=zzAdEt3xZ1M&t=652s)]</sup>
- *Loose Coupling*: Components are less dependent on each other, making it easier to modify or replace them
- *Testability*: Smaller components are easier to test in isolation


## SOLID Principles

**SOLID** is an acronym that represents five key principles of *Object-Oriented Design* in software engineering. These principles are aimed at making software more maintainable, scalable, and flexible. By adhering to SOLID principles, developers can create code that is easier to understand, modify, and extend over time.<sup>[[Wikipedia: SOLID](https://en.wikipedia.org/wiki/SOLID)]</sup>

Applied to Go, we can summarise this approach as follows:<sup>[[Go UK 2016](https://www.youtube.com/watch?v=zzAdEt3xZ1M&t=1361s)] [[Go Time #16: SOLID Go Design](https://changelog.com/gotime/16)]</sup>

*Interfaces let you apply the SOLID principles to Go programs*

Here's a breakdown of each principle:


### S - Single Responsibility Principle (SRP)

**Concept**: 
A class should have one, and only one, reason to change. This means that each class should have a single responsibility or job within the software system.

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
Software entities (classes, modules, functions, etc.) should be open for extension, but closed for modification. This means that you should be able to add new functionality to a class without altering its existing code.   

**Benefits**:
- Enables extending functionality without introducing bugs in existing code.
- Promotes code reusability and maintainability.
- Facilitates easier adaptation to changing requirements.

**Applied to Go**:
- Common interface to add functionality without changing existing code

[OCP Examples](https://github.com/mozey/solid/tree/main/b-open-closed)
- Common Go Mistakes: *"Abstractions should be discovered, not created. To prevent unnecessary complexity, create an interface when you need it and not when you foresee needing it"*<sup>[[Interface pollution](https://100go.co/#interface-pollution-5)]</sup>


### L - Liskov Substitution Principle (LSP)

**Concept**: 
Objects of a derived class should be substitutable for objects of their base class without altering the correctness of the program. In simpler terms, if you have a class (e.g., "Bird") and a subclass (e.g., "Eagle"), you should be able to use an "Eagle" object wherever a "Bird" object is expected.

**Benefits**:
- Ensures that inheritance is used correctly and consistently.
- Prevents unexpected behavior when using subclasses.
- Supports polymorphism and code flexibility.

**Applied to Go**:
- Types that implement same interface (implicitly) are interchangeable
- Use interfaces for: *common behaviour*, *decoupling*, and *restricting behaviour*

[LSP Examples](https://github.com/mozey/solid/tree/main/c-liskov-substitution)
- Common Go Mistakes: *"The bigger the interface, the weaker the abstraction"*<sup>[[Interface pollution](https://100go.co/5-interface-pollution/)]</sup>


### I - Interface Segregation Principle (ISP)

**Concept**: 
Clients should not be forced to depend on interfaces that they don't use. Instead, many specific interfaces are better than one general interface. This means that interfaces should be small and focused on specific sets of related methods.

**Benefits**:
- Reduces dependencies and coupling between classes.
- Improves code flexibility and maintainability.
- Prevents unnecessary implementations of methods.

**Applied to Go**:
- Interfaces are implemented implicitly in Go
- Abstractions should be discovered, not created, see OCP
- Small interfaces lead to simple implementations
- Client (consumer) to decide whether it needs an abstraction and the level

[ISP Examples](https://github.com/mozey/solid/tree/main/d-interface-segregation)
- Common Go Mistakes: *"Keeping interfaces on the client side avoids unnecessary abstractions"*<sup>[[Interface on the producer side](https://100go.co/#interface-on-the-producer-side-6)]</sup>


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

