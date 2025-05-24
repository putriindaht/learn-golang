package variadicSubtract

import "fmt"

func KurangVariadic(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) > 5 {
		fmt.Println("KurangVariadic only processes the first 5 parameters")
		nums = nums[:5]
	}

	result := nums[0]

	for _, n := range nums[1:] {
		result -= n
	}
	return result

}
