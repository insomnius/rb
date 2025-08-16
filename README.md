# rb: Ruby-Like Elegance in Go

Welcome to `rb`, where the simplicity and elegance of Ruby meet the performance and robustness of Go. If you're a Rubyist curious about Go, or a Gopher longing for Ruby's expressive methods, this library bridges the gap by bringing Ruby-inspired utility methods to Go's strong typing and concurrency capabilities.

## Why rb?
- **Ruby Simplicity:** Experience Ruby-style methods like `downcase`, `upcase`, `chars`, and `count` in your Go projects.
- **Go Power:** Maintain the speed and safety of Go while leveraging the expressive power of Ruby-like syntax.
- **Familiarity:** Ideal for Ruby developers transitioning to Go or Go developers seeking more expressive tools.

## Features
- Ruby-inspired `String`, `Integer`, `Float`, `Boolean`, `Symbol`, `Array`, `Hash`, and `Range` types.
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

	// Working with rb.Symbol
	sym := rb.NewSymbol("hello")
	fmt.Println("Symbol:", sym)                 // Symbol: hello
	fmt.Println("Is alpha:", sym.IsAlpha())     // Is alpha: true
	fmt.Println("Is numeric:", sym.IsNumeric()) // Is numeric: false
}
```

## Available Types and Methods

### String Methods
- **Case manipulation**: `Upcase()`, `Downcase()`, `Capitalize()`, `Title()`, `Swapcase()`
- **Whitespace handling**: `Strip()`, `Lstrip()`, `Rstrip()`
- **Transformation**: `Reverse()`, `Gsub()`, `Sub()`
- **Query methods**: `IsEmpty()`, `IsBlank()`, `StartWith()`, `EndWith()`, `Include()`
- **Splitting**: `Split()`, `Lines()`, `Words()`, `Chars()`
- **Conversion**: `ToI()`, `ToF()`
- **Enforce variants**: `EnforceUpcase()`, `EnforceDowncase()`, `EnforceCapitalize()`, etc.

### Integer Methods
- **Mathematical**: `Power()`, `Sqrt()`, `Abs()`, `Factorial()`, `GCD()`, `LCM()`
- **Query methods**: `IsOdd()`, `IsEven()`, `IsPositive()`, `IsNegative()`, `IsZero()`, `IsPrime()`
- **Iteration**: `Times()`, `Upto()`, `Downto()`, `Step()`
- **Utility**: `Next()`, `Pred()`, `Succ()`, `Divisors()`, `IsDivisibleBy()`
- **Conversion**: `ToF()`, `ToS()`

### Float Methods
- **Mathematical**: `Power()`, `Sqrt()`, `Sin()`, `Cos()`, `Tan()`, `Log()`, `Log10()`, `Exp()`
- **Rounding**: `Ceil()`, `Floor()`, `Round()`
- **Query methods**: `IsPositive()`, `IsNegative()`, `IsZero()`, `IsFinite()`, `IsInfinite()`, `IsNaN()`
- **Utility**: `Min()`, `Max()`, `Clamp()`, `Between()`, `IsInteger()`
- **Conversion**: `ToI()`, `ToS()`

### Boolean Methods
- **Logical operations**: `And()`, `Or()`, `Not()`, `Xor()`, `Nand()`, `Nor()`, `Xnor()`, `Implies()`
- **Query methods**: `IsTrue()`, `IsFalse()`
- **Conditional execution**: `IfTrue()`, `IfFalse()`, `If()`, `Ternary()`
- **Conversion**: `ToI()`, `ToF()`, `ToS()`

### Array Methods
- **Transformation**: `Map()`, `Select()`, `Reject()`, `Reverse()`, `Sort()`, `Uniq()`, `Compact()`
- **Query methods**: `Find()`, `Any()`, `All()`, `None()`, `First()`, `Last()`, `IsEmpty()`
- **Iteration**: `Each()`, `EachWithIndex()`
- **Utility**: `Take()`, `Drop()`, `Join()`, `Flatten()`
- **Counting**: `Count()`, `Length()`, `Size()`

### Hash Methods
- **Access**: `Keys()`, `Values()`, `HasKey()`, `HasValue()`, `Get()`, `Fetch()`
- **Modification**: `Set()`, `Delete()`, `Clear()`, `Merge()`, `Update()`, `Replace()`
- **Transformation**: `Select()`, `Reject()`, `Map()`, `Invert()`, `Clone()`
- **Iteration**: `Each()`, `EachKey()`, `EachValue()`
- **Utility**: `Size()`, `Length()`, `IsEmpty()`, `Default()`

### Range Methods
- **Iteration**: `Each()`, `EachWithIndex()`, `Step()`
- **Query methods**: `Include()`, `Cover()`, `IsEmpty()`, `Overlap()`, `Contains()`
- **Utility**: `Min()`, `Max()`, `First()`, `Last()`, `Size()`, `Length()`
- **Conversion**: `ToArray()`, `ToS()`

### Symbol Methods
- **Case manipulation**: `Upcase()`, `Downcase()`, `Capitalize()`, `Title()`, `Swapcase()`
- **Whitespace handling**: `Strip()`, `Lstrip()`, `Rstrip()`
- **Transformation**: `Reverse()`, `Gsub()`, `Sub()`
- **Query methods**: `IsEmpty()`, `IsBlank()`, `IsPresent()`, `StartWith()`, `EndWith()`, `Include()`
- **Character classification**: `IsAlpha()`, `IsAlphanumeric()`, `IsDigit()`, `IsSpace()`, `IsUpper()`, `IsLower()`
- **Splitting**: `Split()`, `Lines()`, `Words()`, `Chars()`
- **Conversion**: `ToI()`, `ToF()`, `ToS()`, `ToSym()`

## Documentation
Explore the complete documentation and examples [here](https://pkg.go.dev/github.com/insomnius/rb).

## Contributing
We welcome contributions from the community! Feel free to submit issues, fork the repository, and open pull requests.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

