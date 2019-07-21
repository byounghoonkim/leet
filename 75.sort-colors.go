/*
 * @lc app=leetcode id=75 lang=golang
 *
 * [75] Sort Colors
 *
 * https://leetcode.com/problems/sort-colors/description/
 *
 * algorithms
 * Medium (42.67%)
 * Likes:    1761
 * Dislikes: 161
 * Total Accepted:    334.4K
 * Total Submissions: 783.7K
 * Testcase Example:  '[2,0,2,1,1,0]'
 *
 * Given an array with n objects colored red, white or blue, sort them in-place
 * so that objects of the same color are adjacent, with the colors in the order
 * red, white and blue.
 * 
 * Here, we will use the integers 0, 1, and 2 to represent the color red,
 * white, and blue respectively.
 * 
 * Note: You are not suppose to use the library's sort function for this
 * problem.
 * 
 * Example:
 * 
 * 
 * Input: [2,0,2,1,1,0]
 * Output: [0,0,1,1,2,2]
 * 
 * Follow up:
 * 
 * 
 * A rather straight forward solution is a two-pass algorithm using counting
 * sort.
 * First, iterate the array counting number of 0's, 1's, and 2's, then
 * overwrite array with total number of 0's, then 1's and followed by 2's.
 * Could you come up with a one-pass algorithm using only constant space?
 * 
 * 
 */
func sortColors(nums []int)  {
	w, b, r := 0, 0, 0 // 각 색의 갯수
	last := len(nums)-1
	for r + w + b < len(nums){ // 각 색의 갯수가 전체와 같아 지면 탈출
		// r 을 인덱스로 보고 앞에서 부터 0이면 진행
		if nums[r]== 0 {
			r++
			continue
		}

		// r 위치에 1이 있으면 뒤에서 w+b 갯수 만큼 떨어진 위치와 스왑
		// w 갯수를 증가
		if nums[r] == 1 {
			wp := last - w - b
			nums[r], nums[wp] = nums[wp], nums[r] 
			w++
			continue
		}

		// r 위치에 2이면 뒤에서 b 만큼 개수 앞에서 스왑
		// b 갯수를 증가 시키고
		if nums[r] == 2 {
			bp := last - b
			nums[r], nums[bp] = nums[bp], nums[r] 
			b++

			// 이 알고리즘의 핵심.
			// w 갯수가 1 이상이면 스왑한 부분이 w를 침범한 것이므로
			// w 개수를 하나 뺌 
			if w > 0 {
				w--
			}

			continue
		}

	}

	// 이 알고리즘의 특징은 
	// 정렬할 대상이 한정적일 때 원패스 만에 정렬을 할 수 있다는 점이다.
}

