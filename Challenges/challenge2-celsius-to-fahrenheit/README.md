# Challenge 2: Celsius to Fahrenheit

## Problem

Write a Go program that:

1. Reads a temperature in Celsius from the user.
2. Converts it to Fahrenheit using the formula:  
   `F = (C × 9/5) + 32`
3. Prints the converted temperature.

## Approach

- **Two methods** are demonstrated in the same program:
  1. **Using `bufio.NewReader`**
     - Reads input as a string.
     - Uses `strings.TrimSpace` to remove newline/spaces.
     - Converts the string to `float64` using `strconv.ParseFloat`.
     - Applies the conversion formula.
  2. **Using `fmt.Scanf`**
     - Directly scans a `float64` from user input without needing string conversion.
     - Applies the same formula.
- Printed results with two decimal places using `%.2f` in `fmt.Printf`.

## Example Output

=== Celsius to Fahrenheit (bufio.NewReader) ===
Enter temperature in Celsius: 25
25.00°C is 77.00°F

=== Celsius to Fahrenheit (fmt.Scanf) ===
Enter temperature in Celsius: 30
30.00°C is 86.00°F
