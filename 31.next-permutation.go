/*
 * @lc app=leetcode id=31 lang=golang
 *
 * [31] Next Permutation
 */

func reverse(nums []int) {
	l := len(nums)
	for i, _ := range nums {
		j := l - 1 - i
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func nextPermutation(nums []int) {

	// 예를 들면 1243 다음 순열을 찾는다고 하면
	//  index  0123
	// 결과는 1324 이어야 한다.

	h := -1 // h는 순열 뒤에서 검색했을 때 값이 적어기 직전 인덱스이다.

	// h를 찾는다.
	// 위 예시에서는 h 값은 인덱스 2이다.
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			h = i
			break
		}
	}

	// h 값을 찾지 못하는 경우는 첫번째 순열로 정렬되어 있는 상태이다.
	// 이때는 단순히 마지막 순열을 구하기 위해 순열을 뒤집는다.
	// 1234 -> 4321 경우
	if h == -1 {
		reverse(nums)
		return
	}

	// h 인덱스가 있는 경우는 h 를 포함해 뒷부분만 순서를 역방향으로 뒤집는다.
	// 예제는 h 값이 2이므로 1243 -> 1234 가 된다.
	reverse(nums[h:])

	// h 이후의 값이 정렬되었으므로 h바로 앞의 값 nums[h-1]보다 큰 수를 찾아서
	// 위치를 스왑해 준다.
	// 위 예에서 h는 2 이고 현재 1234 인 상태
	// nums[h-1] 은 2 이고 이보다 큰 수를 34 에서 검색하는데 3이므로
	// 1234 -> 1324 가 된다.
	for j := h; j < len(nums); j++ {
		if nums[h-1] < nums[j] {
			nums[h-1], nums[j] = nums[j], nums[h-1]
			break
		}
	}

	// 이 알고리즘은 nums의 최악의 경우 nums의 요소 개수 만큼 루프를 2회
	// 수행하기 되므로 2n 의 복잡도를 보이며 O(N) 으로 표기 할 수 있다.
}

