//165. Compare Version Numbers

package twopointer

func compareVersion(version1 string, version2 string) int {
	arr1, arr2 := []int{}, []int{}
	// temp := ""
	temp := 0
	for i := 0; i < len(version1); i++ {
		if version1[i] == '.' || i == len(version1)-1 {
			if i == len(version1)-1 {
				temp = temp*10 + int(version1[i]-48)
			}
			arr1 = append(arr1, temp)
			temp = 0
			continue
		}
		temp = temp*10 + int(version1[i]-48)
	}
	for i := 0; i < len(version2); i++ {
		if version2[i] == '.' || i == len(version2)-1 {
			if i == len(version2)-1 {
				temp = temp*10 + int(version2[i]-48)
			}
			arr2 = append(arr2, temp)
			temp = 0
			continue
		}
		temp = temp*10 + int(version2[i]-48)
	}
	if len(arr1) != len(arr2) {
		apd(&arr1, &arr2)
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] > arr2[i] {
			return 1
		}
		if arr1[i] < arr2[i] {
			return -1
		}
	}
	return 0
}

func apd(arr1, arr2 *[]int) {
	if len(*arr1) < len(*arr2) {
		for len(*arr1) < len(*arr2) {
			*arr1 = append(*arr1, 0)
		}
	} else if len(*arr1) > len(*arr2) {
		for len(*arr1) > len(*arr2) {
			*arr2 = append(*arr2, 0)
		}
	}
}
