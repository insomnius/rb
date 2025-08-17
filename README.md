# rb: Ruby-Like Elegance in Go

[![Go Report Card](https://goreportcard.com/badge/github.com/insomnius/rb)](https://goreportcard.com/report/github.com/insomnius/rb)
[![Go Coverage](https://codecov.io/gh/insomnius/rb/branch/main/graph/badge.svg)](https://codecov.io/gh/insomnius/rb)
[![Go Lint](https://golangci.com/badges/github.com/insomnius/rb.svg)](https://golangci.com/r/github.com/insomnius/rb)

Welcome to `rb`, where the simplicity and elegance of Ruby meet the performance and robustness of Go. If you're a Rubyist curious about Go, or a Gopher longing for Ruby's expressive methods, this library bridges the gap by bringing Ruby-inspired utility methods to Go's strong typing and concurrency capabilities.

## Why rb?
- **Ruby Simplicity:** Experience Ruby-style methods like `downcase`, `upcase`, `chars`, and `count` in your Go projects.
- **Go Power:** Maintain the speed and safety of Go while leveraging the expressive power of Ruby-like syntax.
- **Familiarity:** Ideal for Ruby developers transitioning to Go or Go developers seeking more expressive tools.

## Features
- Ruby-inspired `String`, `Integer`, `Float`, `Boolean`, `Array`, `Hash`, and `Range` types.
- Comprehensive method coverage including `map`, `select`, `reject`, `find`, `any`, `all`, `none`, and more.
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
	fmt.Println("Capitalized:", str.Capitalize()) // Capitalized: Hello, world!
	fmt.Println("Reversed:", str.Reverse())   // Reversed: !dlroW ,olleH

	str.EnforceDowncase()
	fmt.Println("Enforce Downcase:", str) // Enforce Downcase: hello, world!

	// Ruby-style Array methods
	arr := rb.Array[rb.String]{"ruby", "go", "ruby", "python"}
	fmt.Println("Total Elements:", arr.Count())     // Total Elements: 4
	fmt.Println("Count 'ruby':", arr.Count("ruby")) // Count 'ruby': 2
	fmt.Println("Count with Predicate:", arr.Count(func(s rb.String) bool {
		return s.Length() > 2
	})) // Count with Predicate: 4

	// Array transformations
	mapped := arr.Map(func(s rb.String) rb.String {
		return s.Upcase()
	})
	fmt.Println("Mapped:", mapped) // Mapped: [RUBY GO RUBY PYTHON]

	selected := arr.Select(func(s rb.String) bool {
		return s.Length() > 3
	})
	fmt.Println("Selected:", selected) // Selected: [ruby ruby python]

	// Working with rb.Integer
	num := rb.Integer(5)
	fmt.Println("Is Prime:", num.IsPrime())        // Is Prime: true
	fmt.Println("Factorial:", num.Factorial())     // Factorial: 120
	fmt.Println("Divisors:", num.Divisors())       // Divisors: [1 5]

	// Integer iteration
	num.Times(func(i rb.Integer) {
		fmt.Printf("Count: %d\n", i)
	})

	// Working with rb.Float
	pi := rb.Float(3.14159)
	fmt.Println("Ceiled:", pi.Ceil())    // Ceiled: 4
	fmt.Println("Floored:", pi.Floor())  // Floored: 3
	fmt.Println("Rounded:", pi.Round())  // Rounded: 3

	// Working with rb.Boolean
	flag := rb.Boolean(true)
	flag.IfTrue(func() { fmt.Println("It's true!") })
	result := flag.Ternary("yes", "no")
	fmt.Println("Ternary result:", result) // Ternary result: yes

	// Working with rb.Hash
	hash := rb.Hash[rb.String, rb.Integer]{"a": 1, "b": 2, "c": 3}
	fmt.Println("Keys:", hash.Keys())           // Keys: [a b c]
	fmt.Println("Values:", hash.Values())       // Values: [1 2 3]
	fmt.Println("Has key 'a':", hash.HasKey("a")) // Has key 'a': true

	// Working with rb.Range
	rng := rb.NewRange(rb.Integer(1), rb.Integer(5))
	fmt.Println("Range size:", rng.Size())      // Range size: 5
	fmt.Println("Includes 3:", rng.Include(3)) // Includes 3: true
}
```

## Available Types

The library provides the following Ruby-inspired types:

- **`rb.String`** - String manipulation and query methods
- **`rb.Integer`** - Mathematical operations and iteration
- **`rb.Float`** - Mathematical functions and rounding
- **`rb.Boolean`** - Logical operations and conditional execution
- **`rb.Array[T]`** - Collection methods and transformations
- **`rb.Hash[K, V]`** - Key-value operations and iteration
- **`rb.Range[T]`** - Range iteration and query methods

Each type provides a comprehensive set of methods that mirror Ruby's functionality while maintaining Go's type safety and performance characteristics.

## Documentation
Explore the complete documentation and examples [here](https://pkg.go.dev/github.com/insomnius/rb).

## Contributing
We welcome contributions from the community! Feel free to submit issues, fork the repository, and open pull requests.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

