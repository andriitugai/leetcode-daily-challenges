func median(list []int) float64 {
	mid := len(list) / 2
	if len(list)%2 == 0 {
		return float64(list[mid]+list[mid-1]) / 2.0
	}
	return float64(list[mid])
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	answer := []int{}
	for i, j := 0, 0; i < len(nums1) || j < len(nums2); {
		if i < len(nums1) && (j == len(nums2) || nums1[i] <= nums2[j]) {
			answer = append(answer, nums1[i])
			i += 1
		} else {
			answer = append(answer, nums2[j])
			j += 1
		}

	}
	return median(answer)
}