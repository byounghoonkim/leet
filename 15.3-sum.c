/*
 * @lc app=leetcode id=15 lang=c
 *
 * [15] 3Sum
 *
 * https://leetcode.com/problems/3sum/description/
 *
 * algorithms
 * Medium (23.49%)
 * Total Accepted:    491.6K
 * Total Submissions: 2.1M
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * Given an array nums of n integers, are there elements a, b, c in nums such
 * that a + b + c = 0? Find all unique triplets in the array which gives the
 * sum of zero.
 * 
 * Note:
 * 
 * The solution set must not contain duplicate triplets.
 * 
 * Example:
 * 
 * 
 * Given array nums = [-1, 0, 1, 2, -1, -4],
 * 
 * A solution set is:
 * [
 * ⁠ [-1, 0, 1],
 * ⁠ [-1, -1, 2]
 * ]
 * 
 * 
 */
/**
 * Return an array of arrays of size *returnSize.
 * Note: The returned array must be malloced, assume caller calls free().
 */

int* entry_alloc(int a, int b, int c) {
    int * ret = malloc(sizeof(int) * 3);

    ret[0] = a;
    ret[1] = b;
    ret[2] = c;

    return ret;
}

int** threeSum(int* nums, int numsSize, int* returnSize) {

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

    for(int ai = 0; ai < numsSize; ai++) {
        if(ai > 0 && nums[ai] == nums[ai-1])
        {
            continue;
        }

        for(int bi = numsSize -1; bi > ai; bi--) {
            if(bi < numsSize - 1 && nums[bi] == nums[bi+1])
            {
                continue;
            }

            for(int ci = ai+1; ci < bi; ci++) {

                int a = nums[ai]; 
                int b = nums[bi]; 
                int c = nums[ci];

                if(0 == a + b + c) {
                    ret[*returnSize] = entry_alloc(a, b, c);
                    *returnSize = *returnSize + 1;
                    break;
                }

                if(0 < a + b + c)
                {
                    break;
                }

            }
        }
    }

    return ret;
}

