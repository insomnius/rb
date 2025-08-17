// Package main demonstrates the usage of rb library with various examples.
package main

import (
	"fmt"

	"github.com/insomnius/rb"
)

func main() {
	fmt.Println("=== Ruby-like Methods in Go ===")

	// String methods
	fmt.Println("--- String Methods ---")
	str := rb.String("  Hello, World!  ")
	fmt.Printf("Original: '%s'\n", str)
	fmt.Printf("Strip: '%s'\n", str.Strip())
	fmt.Printf("Capitalize: '%s'\n", str.Strip().Capitalize())
	fmt.Printf("Reverse: '%s'\n", str.Strip().Reverse())
	fmt.Printf("Upcase: '%s'\n", str.Strip().Upcase())
	fmt.Printf("Downcase: '%s'\n", str.Strip().Downcase())
	fmt.Printf("Title: '%s'\n", str.Strip().Title())
	fmt.Printf("Swapcase: '%s'\n", str.Strip().Swapcase())
	fmt.Printf("IsEmpty: %t\n", str.IsEmpty())
	fmt.Printf("IsBlank: %t\n", str.IsBlank())
	fmt.Printf("StartWith 'Hello': %t\n", str.Strip().StartWith("Hello"))
	fmt.Printf("EndWith 'World!': %t\n", str.Strip().EndWith("World!"))
	fmt.Printf("Include 'World': %t\n", str.Strip().Include("World"))
	fmt.Printf("Length: %d\n", str.Strip().Length())

	// String splitting
	words := str.Strip().Words()
	fmt.Printf("Words: %v\n", words)
	chars := str.Strip().Chars()
	fmt.Printf("Chars: %v\n", chars)

	// String substitution
	text := rb.String("hello world hello")
	fmt.Printf("Original: '%s'\n", text)
	fmt.Printf("Gsub 'o' -> '0': '%s'\n", text.Gsub("o", "0"))
	fmt.Printf("Sub 'o' -> '0': '%s'\n", text.Sub("o", "0"))

	fmt.Println()

	// Integer methods
	fmt.Println("--- Integer Methods ---")
	num := rb.Integer(12)
	fmt.Printf("Number: %d\n", num)
	fmt.Printf("IsOdd: %t\n", num.IsOdd())
	fmt.Printf("IsEven: %t\n", num.IsEven())
	fmt.Printf("IsPositive: %t\n", num.IsPositive())
	fmt.Printf("IsZero: %t\n", num.IsZero())
	fmt.Printf("Abs: %d\n", num.Abs())
	fmt.Printf("IsPrime: %t\n", num.IsPrime())
	fmt.Printf("Factorial: %d\n", rb.Integer(5).Factorial())
	fmt.Printf("Divisors: %v\n", num.Divisors())
	fmt.Printf("IsDivisibleBy 3: %t\n", num.IsDivisibleBy(3))
	fmt.Printf("Next: %d\n", num.Next())
	fmt.Printf("Pred: %d\n", num.Pred())

	// Integer iteration
	fmt.Println("Counting 3 times:")
	num.Times(func(i rb.Integer) {
		fmt.Printf("  Count: %d\n", i)
	})

	fmt.Println("Counting from 1 to 3:")
	rb.Integer(1).Upto(3, func(i rb.Integer) {
		fmt.Printf("  %d\n", i)
	})

	fmt.Println()

	// Float methods
	fmt.Println("--- Float Methods ---")
	pi := rb.Float(3.14159)
	fmt.Printf("Pi: %f\n", pi)
	fmt.Printf("Ceil: %f\n", pi.Ceil())
	fmt.Printf("Floor: %f\n", pi.Floor())
	fmt.Printf("Round: %f\n", pi.Round())
	fmt.Printf("Abs: %f\n", rb.Float(-3.14).Abs())
	fmt.Printf("Power 2: %f\n", pi.Power(2))
	fmt.Printf("Sqrt: %f\n", rb.Float(16).Sqrt())
	fmt.Printf("IsPositive: %t\n", pi.IsPositive())
	fmt.Printf("IsFinite: %t\n", pi.IsFinite())
	fmt.Printf("IsInteger: %t\n", pi.IsInteger())

	fmt.Println()

	// Boolean methods
	fmt.Println("--- Boolean Methods ---")
	flag := rb.Boolean(true)
	fmt.Printf("Flag: %t\n", flag)
	fmt.Printf("And true: %t\n", flag.And(true))
	fmt.Printf("Or false: %t\n", flag.Or(false))
	fmt.Printf("Not: %t\n", flag.Not())
	fmt.Printf("Xor true: %t\n", flag.Xor(true))
	fmt.Printf("IsTrue: %t\n", flag.IsTrue())
	fmt.Printf("Ternary: %s\n", flag.Ternary("yes", "no"))

	flag.IfTrue(func() {
		fmt.Println("  It's true!")
	})

	fmt.Println()

	// Array methods
	fmt.Println("--- Array Methods ---")
	arr := rb.Array[rb.String]{"ruby", "go", "ruby", "python", "javascript"}
	fmt.Printf("Original: %v\n", arr)
	fmt.Printf("Count: %d\n", arr.Count(nil))
	fmt.Printf("Count 'ruby': %d\n", arr.Count("ruby"))
	fmt.Printf("Count length > 3: %d\n", arr.Count(func(s rb.String) bool {
		return s.Length() > 3
	}))

	// Array transformations
	mapped := arr.Map(func(s rb.String) rb.String {
		return s.Upcase()
	})
	fmt.Printf("Mapped: %v\n", mapped)

	selected := arr.Select(func(s rb.String) bool {
		return s.Length() > 3
	})
	fmt.Printf("Selected: %v\n", selected)

	rejected := arr.Reject(func(s rb.String) bool {
		return s.Length() <= 3
	})
	fmt.Printf("Rejected: %v\n", rejected)

	reversed := arr.Reverse()
	fmt.Printf("Reversed: %v\n", reversed)

	unique := arr.Uniq()
	fmt.Printf("Unique: %v\n", unique)

	// Array query methods
	fmt.Printf("First: %v\n", arr.First())
	fmt.Printf("Last: %v\n", arr.Last())
	fmt.Printf("IsEmpty: %t\n", arr.IsEmpty())
	fmt.Printf("Any length > 5: %t\n", arr.Any(func(s rb.String) bool {
		return s.Length() > 5
	}))
	fmt.Printf("All length > 1: %t\n", arr.All(func(s rb.String) bool {
		return s.Length() > 1
	}))
	fmt.Printf("None length > 10: %t\n", arr.None(func(s rb.String) bool {
		return s.Length() > 10
	}))

	// Array utility methods
	taken := arr.Take(3)
	fmt.Printf("Take 3: %v\n", taken)

	dropped := arr.Drop(2)
	fmt.Printf("Drop 2: %v\n", dropped)

	fmt.Println()

	// Hash methods
	fmt.Println("--- Hash Methods ---")
	hash := rb.Hash[rb.String, rb.Integer]{"a": 1, "b": 2, "c": 3, "d": 4}
	fmt.Printf("Original: %v\n", hash)
	fmt.Printf("Keys: %v\n", hash.Keys())
	fmt.Printf("Values: %v\n", hash.Values())
	fmt.Printf("Size: %d\n", hash.Size())
	fmt.Printf("HasKey 'a': %t\n", hash.HasKey("a"))
	fmt.Printf("HasKey 'z': %t\n", hash.HasKey("z"))

	// Hash iteration
	fmt.Println("Hash contents:")
	hash.Each(func(k rb.String, v rb.Integer) {
		fmt.Printf("  %s: %d\n", k, v)
	})

	// Hash transformation
	selectedHash := hash.Select(func(_ rb.String, v rb.Integer) bool {
		return v > 2
	})
	fmt.Printf("Selected (value > 2): %v\n", selectedHash)

	fmt.Println()

	// Range methods
	fmt.Println("--- Range Methods ---")
	rng := rb.NewRange(rb.Integer(1), rb.Integer(5))
	fmt.Printf("Range: %s\n", rng.ToS())
	fmt.Printf("Size: %d\n", rng.Size())
	fmt.Printf("IsEmpty: %t\n", rng.IsEmpty())
	fmt.Printf("Min: %d\n", rng.Min())
	fmt.Printf("Max: %d\n", rng.Max())
	fmt.Printf("First: %d\n", rng.First())
	fmt.Printf("Last: %d\n", rng.Last())
	fmt.Printf("Include 3: %t\n", rng.Include(3))
	fmt.Printf("Include 7: %t\n", rng.Include(7))

	// Range iteration
	fmt.Println("Range values:")
	rng.Each(func(i rb.Integer) {
		fmt.Printf("  %d\n", i)
	})

	// Range with step
	fmt.Println("Range with step 2:")
	rng.Step(2, func(i rb.Integer) {
		fmt.Printf("  %d\n", i)
	})

	fmt.Println()

	fmt.Println("\n=== All methods demonstrated successfully! ===")
}
