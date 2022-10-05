package lgproblems

import "os"

func simpleArraySum(ar []int32) int32 {
	// Write your code here
	if !constrain1(int32(len(ar))) {
		os.Exit(1)
	}

	for i := 0; i < len(ar); i++ {
		if !constrain1(ar[i]) {
			os.Exit(1)
		}
	}

	var acc int32 = 0
	for i := 0; i < len(ar); i++ {
		acc = acc + ar[i]
	}

	return acc
}
