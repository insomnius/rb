# rb: Ruby-Like Elegance in Go

Welcome to `rb`, where the simplicity and elegance of Ruby meet the performance and robustness of Go. If you're a Rubyist curious about Go, or a Gopher longing for Ruby's expressive methods, this library bridges the gap by bringing Ruby-inspired utility methods to Go's strong typing and concurrency capabilities.

## Why rb?
- **Ruby Simplicity:** Experience Ruby-style methods like `downcase`, `upcase`, `chars`, and `count` in your Go projects.
- **Go Power:** Maintain the speed and safety of Go while leveraging the expressive power of Ruby-like syntax.
- **Familiarity:** Ideal for Ruby developers transitioning to Go or Go developers seeking more expressive tools.

## Features
- Ruby-inspired `String` and `Integer` types.
- Ruby-style `Array` with enhanced methods like `count`.
- Seamless integration into your Go projects with idiomatic Go practices.

## Ruby-like Idiomatic Changes

In `rb-go`, the following Ruby-like method semantics have been adapted to Go's idiomatic patterns:

- `!` → **`Enforce`**: Methods with Ruby's "bang" `!` suffix are now implemented with `Enforce`.
  Example:
  ```go
  str.EnforceDowncase()
  ```

- `?` → **`Is`**: Methods with Ruby's "question" `?` suffix are now implemented with `Is`.
  Example:
  ```go
  5.IsOdd()
  ```

These changes ensure alignment with Go's naming conventions while maintaining Ruby's expressiveness.

## Installation

```bash
go get github.com/insomnius/rb
```

## Quick Start

```go
package main

import (
	"fmt"

	"github.com/insomnius/rb"
)

func main() {
	// Working with rb.String
	str := rb.String("Hello, World!")
	fmt.Println("Original:", str)             // Original: Hello, World!
	fmt.Println("Downcased:", str.Downcase()) // Downcased: hello, world!

	str.EnforceDowncase()
	fmt.Println("Enforce Downcase:", str) // Enforce Downcase: hello, world!

	// Ruby-style Array Count
	arr := rb.Array[rb.String]{"ruby", "go", "ruby"}
	fmt.Println("Total Elements:", arr.Count())     // Total Elements: 3
	fmt.Println("Count 'ruby':", arr.Count("ruby")) // Count 'ruby': 2
	fmt.Println("Count with Predicate:", arr.Count(func(s rb.String) bool {
		return s.Length() > 2
	})) // Count with Predicate: 2
}

```

## Documentation
Explore the complete documentation and examples [here](https://pkg.go.dev/github.com/insomnius/rb).

## Contributing
We welcome contributions from the community! Feel free to submit issues, fork the repository, and open pull requests.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

