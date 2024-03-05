//calculates trimmed mean

package trimmed_mean

import (
	"fmt"
	"sort"
)

// converts input interface{} to float64
func convertToFloat64(nums []interface{}) ([]float64, error) {
	var nums_float []float64 // slice to hold the float64 numbers
	for _, v := range nums {
		switch value := v.(type) { // type assertion
		case int: // if int convert to float64
			nums_float = append(nums_float, float64(value))
		case float64: // if float64, append to x_float
			nums_float = append(nums_float, value)
		default:
			return nil, fmt.Errorf("unsupported type: %T", v)
		}
	}
	return nums_float, nil
}

// removes the upper and lower percentile
func trim(nums_float []float64, lower_percentile float64, upper_percentile float64) ([]float64, error) {

	// sort
	sort.Float64s(nums_float)

	// calculates the # of elements to remove
	lower := int(lower_percentile * float64(len(nums_float)))
	upper := int(upper_percentile * float64(len(nums_float)))

	//add a condition that if lower + upper > len(nums_float), return an error
	if lower+upper >= len(nums_float) {
		return nil, fmt.Errorf("percentile is too high, all elements will be removed")
	}

	// return lower to len(nums_float)-upper elements of the sorted list
	nums_float = nums_float[lower : len(nums_float)-upper]

	return nums_float, nil
}

// sums across a slice of float64
func sum(data []float64) float64 {
	sum := 0.0

	for _, v := range data {
		sum += v
	}

	return sum
}

// calculates the mean
func avg(trimmed_nums []float64) (float64, error) {

	//if empty slice pass an error message
	if len(trimmed_nums) == 0 {
		return 0.0, fmt.Errorf("cannot average an empty slice")
	}

	return sum(trimmed_nums) / float64(len(trimmed_nums)), nil
}

// base function for trimmed mean
// assumes that three inputs are passed with the last two being the upper and lower precentiles
func trimmed_mean_base(nums []interface{}, lower_percentile float64, upper_percentile float64) (float64, error) {

	// converting input string to float64
	nums_float, err := convertToFloat64(nums)
	if err != nil {
		return 0.0, fmt.Errorf("error in type conversion %T\n", nums)
	}

	// calculte the trimmed slice
	trimmed_nums, err := trim(nums_float, lower_percentile, upper_percentile)
	if err != nil {
		return 0.0, fmt.Errorf("error in trimming %T\n", nums)
	}

	mean, err := avg(trimmed_nums)
	if err != nil {
		return 0.0, fmt.Errorf("error in averaging %T\n", nums)
	}

	return mean, nil
}

// calculates the trimmed mean
// defaults two percentiles to the same value if only 1 parameter is passed
func TMean(nums []interface{}, lowerPercentile float64, upperPercentiles ...float64) (float64, error) {

	// default to symmetric trimming
	upperPercentile := lowerPercentile
	if len(upperPercentiles) > 0 {
		upperPercentile = upperPercentiles[0]
	}

	// call the base function
	return trimmed_mean_base(nums, lowerPercentile, upperPercentile)
}

// for testing
// func main() {

// 	// for prototyping, we can use a slice of interface{} to hold the numbers
// 	trim_percentile := 0.1
// 	slice_length := 100
// 	nums := make([]interface{}, 0, slice_length)
// 	for i := 0; i < slice_length; i++ {
// 		nums = append(nums, i)
// 	}

// 	//testing one percentile passed
// 	mean, err := trimmedMean(nums, trim_percentile)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("Trimmed mean is ", mean)

// 	//testing two percentiles passed
// 	mean1, err := trimmedMean(nums, trim_percentile, trim_percentile)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("Trimmed mean is ", mean1)

// }
