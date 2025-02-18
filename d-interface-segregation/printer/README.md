# Interface Segregation Principle (ISP)

## Printer Example

Key Points
- **Thick Interface**: `ThickPrinter` interface is too broad, and violates ISP
- **Segregated Interfaces**: Breaks the interface down into `SPrinter`, `SFax`, and `SScanner`
- **Flexibility**: Possible to implement all three interfaces, or implement only what is needed
- **Alternate Solution**: The `GPrinter` interface embeds `io.Writer`, and `GScanner` interface embeds `io.Reader`, from the standard library. 
- **Composable Pipelines**: Combine different implementations to create I/O pipelines. For example, you can chain `gzip.Writer | bufio.Writer | os.File` to compress data before writing it to a file, all while working with a single simple interface
- **Adaptor Pattern**: Possible to implement the alternative solution on top of the segregated interfaces

Run the code 
```bash
go run ./d-interface-segregation/printer/main.go
```
