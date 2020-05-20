package main

import "fmt"

//题目21:调整数组顺序使奇数位于偶数前面
/*
 输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。
 输入：nums = [1,2,3,4]
 输出：[1,3,2,4]
 注：[3,1,2,4] 也是正确的答案之一。
 */

 //思路:使用前后指针
 //p1从前往后找,p2从后往前找,p1找到第一个偶数,p2找到第一个奇数,就停下来交换
func exchange(nums []int) []int {
	if len(nums) < 1 {
		return nil
	}
	if len(nums) == 1 {
		return nums
	}

	p1, p2 := 0, len(nums) - 1
	for {
		if p1 >= p2 {
			break
		}

		//p1找到第一个偶数,p2找到第一个奇数,交换
		if (nums[p1] & 1 == 0) && (nums[p2] & 1 == 1) {
			nums[p1], nums[p2] = nums[p2], nums[p1]
			p1++
			p2--
		}

		if nums[p1] & 1 == 1 {
			p1++
		}

		if nums[p2] & 1 == 0 {
			p2--
		}
	}

	return nums
}

func main() {
	nums := []int{3, 10, 6, 9, 8, 11, 12, 14}
	nums = exchange(nums)
	fmt.Printf("整理后exchange为:%v\n", nums)
}
