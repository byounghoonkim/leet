/*
 * @lc app=leetcode id=80 lang=golang
 *
 * [80] Remove Duplicates from Sorted Array II
 *
 * https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/description/
 *
 * algorithms
 * Medium (40.86%)
 * Likes:    714
 * Dislikes: 554
 * Total Accepted:    210.5K
 * Total Submissions: 515.3K
 * Testcase Example:  '[1,1,1,2,2,3]'
 *
 * Given a sorted array nums, remove the duplicates in-place such that
 * duplicates appeared at most twice and return the new length.
 * 
 * Do not allocate extra space for another array, you must do this by modifying
 * the input array in-place with O(1) extra memory.
 * 
 * Example 1:
 * 
 * 
 * Given nums = [1,1,1,2,2,3],
 * 
 * Your function should return length = 5, with the first five elements of nums
 * being 1, 1, 2, 2 and 3 respectively.
 * 
 * It doesn't matter what you leave beyond the returned length.
 * 
 * Example 2:
 * 
 * 
 * Given nums = [0,0,1,1,1,1,2,3,3],
 * 
 * Your function should return length = 7, with the first seven elements of
 * nums being modified to 0, 0, 1, 1, 2, 3 and 3 respectively.
 * 
 * It doesn't matter what values are set beyond the returned length.
 * 
 * 
 * Clarification:
 * 
 * Confused why the returned value is an integer but your answer is an array?
 * 
 * Note that the input array is passed in by reference, which means
 * modification to the input array will be known to the caller as well.
 * 
 * Internally you can think of this:
 * 
 * 
 * // nums is passed in by reference. (i.e., without making a copy)
 * int len = removeDuplicates(nums);
 * 
 * // any modification to nums in your function would be known by the caller.
 * // using the length returned by your function, it prints the first len
 * elements.
 * for (int i = 0; i < len; i++) {
 * print(nums[i]);
 * }
 * 
 * 
 */
func removeDuplicates(nums []int) int {
	l,r := 0,0 // left, reignt index
	c := 0 // 동일 숫자가 몇개나 진행 되었는지 카운트

	// r 인덱스가 끝까지 루프 종료
	for ;r < len(nums); {
		// 카운트가 하나도 진행되지 않았다면 현재 r 인덱스를 l인덱스로 복사 하고
		// r 진행 
		if c == 0 {
			nums[l] = nums[r]
			r++
			c++
			continue
		}


		// 두 인덱스 아래 숫자가 다르면 l을 하나 증가 시키고 카운트로 초기화 한 다음
		// 다음 루프에서 복사할 수 있도록 진행
		if nums[l] != nums[r] {
			l++
			c = 0
		} else {
			// 두 번호가 같은데 카운트가 2 미만이면 l 인덱스를 하나 진행 시키고
			// 숫자를 복사 한 다음 r 도 진행 시킴
			if c < 2 {
				l++
				nums[l] = nums[r]
				r++
				c++
			} else {
				// 카운트가 2 이상이면 r 인덱스만 진행 시킴
				r++
				c++
			}
		}
	}

	// 루프를 탈출 했을 때 l 이 결과 슬라이스의 마지막을 가리키므로
	// l + 1 만큼의 길이를 가진 답임.
	return l+1

	// 이 알고리즘은 O(n)
    
}

