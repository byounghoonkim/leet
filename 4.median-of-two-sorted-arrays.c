/*
 * [4] Median of Two Sorted Arrays
 *
 * https://leetcode.com/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (24.31%)
 * Total Accepted:    329.4K
 * Total Submissions: 1.4M
 * Testcase Example:  '[1,3]\n[2]'
 *
 * There are two sorted arrays nums1 and nums2 of size m and n respectively.
 * 
 * Find the median of the two sorted arrays. The overall run time complexity
 * should be O(log (m+n)).
 * 
 * You may assume nums1 and nums2 cannot be both empty.
 * 
 * Example 1:
 * 
 * 
 * nums1 = [1, 3]
 * nums2 = [2]
 * 
 * The median is 2.0
 * 
 * 
 * Example 2:
 * 
 * 
 * nums1 = [1, 2]
 * nums2 = [3, 4]
 * 
 * The median is (2 + 3)/2 = 2.5
 * 
 * 
 */

double findMedianSortedArrays(int* nums1, int nums1Size, int* nums2, int nums2Size) {

    int m1, m2 = 0;
    int v1, v2 = 0;
    int i1, i2 = 0;
    int count = 0;

    if ((nums1Size + nums2Size) % 2) {
        m1 = m2 = (nums1Size + nums2Size) / 2 + 1;
    }
    else {
        m1 = (nums1Size + nums2Size) / 2;
        m2 = m1 + 1;
    }

    while(true) {
        count ++;

        if((i1 < nums1Size) && (i2 < nums2Size)) {
            if(nums1[i1] < nums2[i2]) {
                i1++;
                if(count == m1) {
                    v1 = nums1[i1-1];
                }
                if(count == m2) {
                    v2 = nums1[i1-1];
                    break;
                }
            } else {
                i2++;
                if(count == m1) {
                    v1 = nums2[i2-1];
                }
                if(count == m2) {
                    v2 = nums2[i2-1];
                    break;
                }
            }
        } else if(i1 < nums1Size) {
                i1++;
                if(count == m1) {
                    v1 = nums1[i1-1];
                }
                if(count == m2) {
                    v2 = nums1[i1-1];
                    break;
                }
        } else if(i2 < nums2Size) {
                i2++;
                if(count == m1) {
                    v1 = nums2[i2-1];
                }
                if(count == m2) {
                    v2 = nums2[i2-1];
                    break;
                }
        } else {
            return 0.0;
        }
    }
    

    return (v1 + v2) / 2.0;
}

