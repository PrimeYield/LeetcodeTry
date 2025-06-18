//	303. Range Sum Query - Immutable

package array

// type NumArray struct {
// 	nums        []int
// 	left, right int
// }

// func Constructor(nums []int) NumArray {
// 	return NumArray{
// 		nums: nums,
// 	}
// }

// func (n *NumArray) SumRange(left int, right int) int {
// 	sum := 0
// 	for i := left; i <= right; i++ {
// 		sum += n.nums[i]
// 	}
// 	return sum
// }
// runtime 37ms
// Memory 9.31MB Beats 82.21%

//新的想法：[left,right]的總合 = [0,right] - [0,left-1] 這就是前綴和的概念

type NumArray struct {
	sumArr []int
}

func Constructor(nums []int) NumArray {
	obj := NumArray{sumArr: make([]int, len(nums)+1, len(nums)+1)}

	//nums從0開始遍歷 sumArr[0] = 0 既定的 從1開始sum
	origInd := 0
	arrInd := 1

	for arrInd < len(obj.sumArr) {
		obj.sumArr[arrInd] = nums[origInd] + obj.sumArr[origInd]
		origInd++
		arrInd++
	}

	return obj
}

func (n *NumArray) SumRange(left int, right int) int {
	return n.sumArr[right+1] - n.sumArr[left]
}

//runtime 1ms Beats 75.45%
//Memory 9.69 Beats 61.29%
