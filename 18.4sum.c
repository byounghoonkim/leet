/*
 * @lc app=leetcode id=18 lang=c
 *
 * [18] 4Sum
 *
 * https://leetcode.com/problems/4sum/description/
 *
 * algorithms
 * Medium (29.80%)
 * Total Accepted:    215.1K
 * Total Submissions: 721.6K
 * Testcase Example:  '[1,0,-1,0,-2,2]\n0'
 *
 * Given an array nums of n integers and an integer target, are there elements
 * a, b, c, and d in nums such that a + b + c + d = target? Find all unique
 * quadruplets in the array which gives the sum of target.
 * 
 * Note:
 * 
 * The solution set must not contain duplicate quadruplets.
 * 
 * Example:
 * 
 * 
 * Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.
 * 
 * A solution set is:
 * [
 * ⁠ [-1,  0, 0, 1],
 * ⁠ [-2, -1, 1, 2],
 * ⁠ [-2,  0, 0, 2]
 * ]
 * 
 * 
 */
/**
 * Return an array of arrays of size *returnSize.
 * Note: The returned array must be malloced, assume caller calls free().
 */

int* entry_alloc(int a, int b, int c, int d) {
	int * ret = malloc(sizeof(int) * 4);

	ret[0] = a;
	ret[1] = b;
	ret[2] = c;
	ret[3] = d;

	return ret;
}



int** fourSum(int* nums, int numsSize, int target, int* returnSize) {

	int** ret = malloc(100000 * sizeof(int*));
	*returnSize = 0;

	for(int i = 0; i < numsSize; i++) {
		for(int j = i+1; j < numsSize; j++) {
			if(nums[i] > nums[j])
			{
				int tmp = nums[i];
				nums[i] = nums[j];
				nums[j] = tmp;
			}
		}
	}

	for(int i = 0; i < numsSize; i++) {
		if(i>0 && nums[i-1] == nums[i]) continue;
		for(int j = i+1; j < numsSize; j++) {
			if(j>i+1 && nums[j-1] == nums[j]) continue;
			for(int k = j+1; k < numsSize; k++) {
				if(k>j+1 && nums[k-1] == nums[k]) continue;
				for(int l = k+1; l < numsSize; l++) {
					if(l>k+1 && nums[l-1] == nums[l]) continue;
					if(target == nums[i] + nums[j] + nums[k] + nums[l])
					{
						ret[*returnSize] = entry_alloc(nums[i],nums[j],nums[k],nums[l]);
						(*returnSize)++;
					}

				}
			}
		}
	}

	return ret;

}
