/*
 * @lc app=leetcode id=16 lang=c
 *
 * [16] 3Sum Closest
 *
 * https://leetcode.com/problems/3sum-closest/description/
 *
 * algorithms
 * Medium (41.12%)
 * Total Accepted:    292.2K
 * Total Submissions: 709.5K
 * Testcase Example:  '[-1,2,1,-4]\n1'
 *
 * Given an array nums of n integers and an integer target, find three integers
 * in nums such that the sum is closest to target. Return the sum of the three
 * integers. You may assume that each input would have exactly one solution.
 * 
 * Example:
 * 
 * 
 * Given array nums = [-1, 2, 1, -4], and target = 1.
 * 
 * The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
 * 
 * 
 */
int threeSumClosest(int* nums, int numsSize, int target) {
	int ret = 0;
	int diff = 0;

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
            //continue;
        }

        for(int bi = numsSize -1; bi > ai; bi--) {
            if(bi < numsSize - 1 && nums[bi] == nums[bi+1])
            {
                //continue;
            }

            for(int ci = ai+1; ci < bi; ci++) {
				//printf("a:%d b:%d c:%d\n", ai, bi, ci);

                int a = nums[ai]; 
                int b = nums[bi]; 
                int c = nums[ci];

                //printf("t: %d, abc: %d", target, a + b + c);
				if( target == a+b+c) {
                    return target;
                }

				if(diff == 0)
				{
                	diff = abs(a + b + c - target );
					ret = a + b + c;
				}

                if( abs(a + b + c - target ) < diff )
				{
					ret = a + b + c;
                	diff = abs(a + b + c - target );
				}
            }
        }
    }


    return ret;
}

