package main
//题目39:数组中出现次数超过一半的数
//摩尔投票,正负抵消
func majorityElement(nums []int) int {
	if nums == nil {
		return -1
	}


    x, votes := 0, 1
    for i := 1; i < len(nums); i++ {
    	if votes == 0 {
    		x = nums[i]
		}

    	if nums[i] == x {
    		votes++
		} else {
			votes--
		}
	}

    return x
}
