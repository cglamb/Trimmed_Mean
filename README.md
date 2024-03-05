# Trimmed Mean Library
**Author:** Charles Lamb  
**Contact Info:** charlamb@gmail.com  
**Git Clone Command:** `git clone https://github.com/cglamb/Trimmed_Mean.git`  
**Go Get Command:** `go get github.com/cglamb/Trimmed_Mean`

## Introduction
This go package calculates trimmed means. The package can be added to a go project using `go get github.com/cglamb/Trimmed_Mean`  
The trimmed mean function can be accessed using `trimmed_mean.Tmean([]interface{}, lower_trim_percentile, upper_trim_percentile)`.  
The upper percentile input is optional and if left blank will default to the same value as the lower percentile.

## Testing
A number of unit tests are included in `trimmed_mean_test.go`. Amongst the unit tests is a `TestTMean` function which tests the trimmed mean calculation on two sequences (each input twice – once as a sequence of integers and once as a sequence of float64). The expected result was determined by using `mean(x,trim)` in R, per the following code:

```r
x <- seq(1, by = 5, length.out = 100)
print(x)
mean(x,trim=.10)
x <- seq(-50, by = 5, length.out = 100)
print(x)
mean(x,trim=.10)
```

## Example

An example of the library being used is provided here:

```go
package main

import (
    "fmt"
    trimmed_mean "github.com/cglamb/Trimmed_Mean"
)

func main() {
    // setup a test sequence
    trim_percentile := 0.1
    slice_length := 100
    nums := make([]interface{}, 0, slice_length)
    for i := 0; i < slice_length; i++ {
        nums = append(nums, i)
    }

    // function passing one percentile parameter
    mean, err := trimmed_mean.TMean(nums, trim_percentile)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("Trimmed mean is ", mean)

    // testing two percentiles passed
    mean1, err := trimmed_mean.TMean(nums, trim_percentile, trim_percentile)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("Trimmed mean is ", mean1)
}
```
## Testing Results

```plaintext
PS C:\Users\CGLam\go\src\trimmed_mean> go test -v
=== RUN   TestConvertToFloat64
=== RUN   TestConvertToFloat64/Integer
=== RUN   TestConvertToFloat64/Float64
=== RUN   TestConvertToFloat64/Mixed
=== RUN   TestConvertToFloat64/strings
=== RUN   TestConvertToFloat64/nil
--- PASS: TestConvertToFloat64 (0.00s)
    --- PASS: TestConvertToFloat64/Integer (0.00s)
    --- PASS: TestConvertToFloat64/Float64 (0.00s)
    --- PASS: TestConvertToFloat64/Mixed (0.00s)
    --- PASS: TestConvertToFloat64/strings (0.00s)
    --- PASS: TestConvertToFloat64/nil (0.00s)
=== RUN   TestTrim
=== RUN   TestTrim/10%_Trime
=== RUN   TestTrim/20%_Trime
=== RUN   TestTrim/30%_Trime
=== RUN   TestTrim/40%_Trime
=== RUN   TestTrim/45%_Trime
=== RUN   TestTrim/Assymetric_10/30
=== RUN   TestTrim/Assymetric_30/10
=== RUN   TestTrim/Fail_100%_Trim
--- PASS: TestTrim (0.00s)
    --- PASS: TestTrim/10%_Trime (0.00s)
    --- PASS: TestTrim/20%_Trime (0.00s)
    --- PASS: TestTrim/30%_Trime (0.00s)
    --- PASS: TestTrim/40%_Trime (0.00s)
    --- PASS: TestTrim/45%_Trime (0.00s)
    --- PASS: TestTrim/Assymetric_10/30 (0.00s)
    --- PASS: TestTrim/Assymetric_30/10 (0.00s)
    --- PASS: TestTrim/Fail_100%_Trim (0.00s)
=== RUN   TestAvg
=== RUN   TestAvg/average_of_positive_numbers
=== RUN   TestAvg/emplty_slice_throws_an_error
=== RUN   TestAvg/mean_of_negative_numbers
=== RUN   TestAvg/single_number_mean
=== RUN   TestAvg/mixed_positive_and_negative_numbers_mean
--- PASS: TestAvg (0.00s)
    --- PASS: Test
```
