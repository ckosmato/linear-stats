
# Linear Regression and Pearson Correlation Coefficient Calculator

This Go program calculates the Linear Regression Line and Pearson Correlation Coefficient (PCC) for a given set of data. The program takes input from either user input or a file and outputs the results for the linear regression line and PCC.

## Features
- Calculates the **Linear Regression Line** in the form `y = mx + b` where `m` is the slope and `b` is the intercept.
- Computes the **Pearson Correlation Coefficient** (PCC) to measure the strength and direction of the relationship between two variables.
- Accepts input from the user or from a text file (`data.txt`).

## Usage

### Option 1: User Input

1. Compile and run the program:
   ```bash
   go run main.go
   ```
2. Enter values one by one in the terminal. After each value, press Enter. The program will automatically generate the Linear Regression Line and the PCC.

### Option 2: Input from `data.txt`

You can provide data through a file named `data.txt`, where each line contains a `y` value and the `x` values are automatically generated as the line numbers (indices).

1. Ensure that the `data.txt` file is in the same directory as the `main.go` file.
2. Run the program with the following command:
   ```bash
   go run main.go < data.txt
   ```

The program will read the contents of the `data.txt` file, process the data, and output the results for the Linear Regression Line and PCC.

## Example Input (data.txt)

```
5
10
15
20
25
```

## Output

The program will print something similar to the following:

```
Linear Regression Line: y = 5.000000x + 0.000000
Pearson Correlation Coefficient: 1.0000000000
```

## Author

Author: Chris Kosmatos 
