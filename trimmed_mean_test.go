package Go_Trimmed_Mean

import (
	"fmt"
	"reflect"
	"testing"
)

// test convertToFloat64
func TestConvertToFloat64(t *testing.T) {
	testCases := []struct {
		scenario       string
		input          []interface{}
		expected       []float64
		expectedResult bool
	}{
		{
			scenario:       "Integer",
			input:          []interface{}{10, 20, 30},
			expected:       []float64{10.0, 20.0, 30.0},
			expectedResult: false,
		},
		{
			scenario:       "Float64",
			input:          []interface{}{10.0, 20.0, 30.},
			expected:       []float64{10, 20, 30},
			expectedResult: false,
		},
		{
			scenario:       "Mixed",
			input:          []interface{}{10, 20.0, 30},
			expected:       []float64{10, 20, 30},
			expectedResult: false,
		},
		{
			scenario:       "strings",
			input:          []interface{}{"1", "2.2", "3"},
			expected:       nil,
			expectedResult: true,
		},
		{
			scenario:       "nil",
			input:          []interface{}{nil, nil, nil},
			expected:       nil,
			expectedResult: true,
		},
	}

	// Loop through the test cases
	for _, test := range testCases {
		t.Run(test.scenario, func(t *testing.T) {
			actual, err := convertToFloat64(test.input)
			if (err != nil) != test.expectedResult {
				t.Errorf("error = %v, wantErr %v", err, test.expectedResult)
				return
			}
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("got = %v, want %v", actual, test.expected)
			}
		})
	}
}

// test trim
func TestTrim(t *testing.T) {
	testCases := []struct {
		scenario         string
		nums             []float64
		lower_percentile float64
		upper_percentile float64
		expected         []float64
		expectedError    bool
	}{
		{
			scenario:         "10% Trime",
			nums:             []float64{0, 2, 4, 6, 8, 10, 50, 70, 90, 110, 130, 150},
			lower_percentile: 0.1,
			upper_percentile: 0.1,
			expected:         []float64{2, 4, 6, 8, 10, 50, 70, 90, 110, 130},
			expectedError:    false,
		},
		{
			scenario:         "20% Trime",
			nums:             []float64{0, 2, 4, 6, 8, 10, 50, 70, 90, 110, 130, 150},
			lower_percentile: 0.2,
			upper_percentile: 0.2,
			expected:         []float64{4, 6, 8, 10, 50, 70, 90, 110},
			expectedError:    false,
		},
		{
			scenario:         "30% Trime",
			nums:             []float64{0, 2, 4, 6, 8, 10, 50, 70, 90, 110, 130, 150},
			lower_percentile: 0.3,
			upper_percentile: 0.3,
			expected:         []float64{6, 8, 10, 50, 70, 90},
			expectedError:    false,
		},
		{
			scenario:         "40% Trime",
			nums:             []float64{0, 2, 4, 6, 8, 10, 50, 70, 90, 110, 130, 150},
			lower_percentile: 0.4,
			upper_percentile: 0.4,
			expected:         []float64{8, 10, 50, 70},
			expectedError:    false,
		},
		{
			scenario:         "45% Trime",
			nums:             []float64{0, 2, 4, 6, 8, 10, 50, 70, 90, 110, 130, 150},
			lower_percentile: 0.45,
			upper_percentile: 0.45,
			expected:         []float64{10, 50},
			expectedError:    false,
		},
		{
			scenario:         "Assymetric 10/30",
			nums:             []float64{0, 2, 4, 6, 8, 10, 50, 70, 90, 110, 130, 150},
			lower_percentile: 0.10,
			upper_percentile: 0.30,
			expected:         []float64{2, 4, 6, 8, 10, 50, 70, 90},
			expectedError:    false,
		},
		{
			scenario:         "Assymetric 30/10",
			nums:             []float64{0, 2, 4, 6, 8, 10, 50, 70, 90, 110, 130, 150},
			lower_percentile: 0.30,
			upper_percentile: 0.10,
			expected:         []float64{6, 8, 10, 50, 70, 90, 110, 130},
			expectedError:    false,
		},
		{
			scenario:         "Fail 100% Trim",
			nums:             []float64{0, 2, 4, 6, 8, 10, 50, 70, 90, 110, 130, 150},
			lower_percentile: 0.50,
			upper_percentile: 0.50,
			expected:         nil,
			expectedError:    true,
		},
	}

	// Loop through the test cases
	for _, test := range testCases {
		t.Run(test.scenario, func(t *testing.T) {
			actual, err := trim(test.nums, test.lower_percentile, test.upper_percentile)

			// Handle the expected error scenario
			if test.expectedError {
				if err == nil {
					t.Errorf("expected an error but got none")
				}
			} else {
				// In case no error is expected, check for unexpected errors
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				} else if !reflect.DeepEqual(actual, test.expected) {
					t.Errorf("got %v, want %v", actual, test.expected)
				}
			}
		})
	}
}

// testing sum
func TestAvg(t *testing.T) {
	testCases := []struct {
		scenario      string
		nums          []float64
		expectedAvg   float64
		expectedError error
	}{
		{
			scenario:      "average of positive numbers",
			nums:          []float64{10.0, 20.0, 30.0},
			expectedAvg:   20.0,
			expectedError: nil,
		},
		{
			scenario:      "emplty slice throws an error",
			nums:          []float64{},
			expectedAvg:   0.0,
			expectedError: fmt.Errorf("cannot average an empty slice"),
		},
		{
			scenario:      "mean of negative numbers",
			nums:          []float64{-10.0, -20.0, -30.0},
			expectedAvg:   -20.0,
			expectedError: nil,
		},
		{
			scenario:      "single number mean",
			nums:          []float64{10},
			expectedAvg:   10.0,
			expectedError: nil,
		},
		{
			scenario:      "mixed positive and engative numbers mean",
			nums:          []float64{-10.0, 0.0, 10.0},
			expectedAvg:   0.0,
			expectedError: nil,
		},
	}

	for _, test := range testCases {
		t.Run(test.scenario, func(t *testing.T) {
			actual, err := avg(test.nums)
			if err != nil && test.expectedError == nil {
				t.Errorf("test failed as expected")
			} else if err == nil && test.expectedError != nil {
				t.Errorf("failure expected but got none")
				// } else if err != nil && test.expectedError != nil && err.Error() != test.expectedError.Error() {
				// 	t.Errorf("Test '%s' failed: expected error '%v', got '%v'", tc.scenario, tc.expectedError, err)
			} else if actual != test.expectedAvg {
				t.Errorf("average does not match expected")
			}
		})
	}
}

// benchmarking
// calculates trimmed mean
func BenchmarkTrimmedMean(b *testing.B) {

	// setting up testing slice
	trimPercentile := 0.1
	sliceLength := 100
	nums := make([]interface{}, 0, sliceLength)
	for i := 0; i < sliceLength; i++ {
		nums = append(nums, i)
	}

	for i := 0; i < b.N; i++ {
		numsFloat, err := convertToFloat64(nums) //convert to float64
		if err != nil {
			b.Fatalf("error in type conversion %T\n", nums)
		}

		trimmedNums, err := trim(numsFloat, trimPercentile, trimPercentile) //trim the slice
		if err != nil {
			b.Fatalf("error in trimming %T\n", nums)
		}

		_, err = avg(trimmedNums) //calculate the mean
		if err != nil {
			b.Fatalf("error in averaging %T\n", nums)
		}
	}
}
