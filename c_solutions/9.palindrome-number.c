/*
 * [9] Palindrome Number
 *
 * https://leetcode.com/problems/palindrome-number/description/
 *
 * algorithms
 * Easy (39.52%)
 * Total Accepted:    432.2K
 * Total Submissions: 1.1M
 * Testcase Example:  '121'
 *
 * Determine whether an integer is a palindrome. An integer is a palindrome
 * when it reads the same backward as forward.
 * 
 * Example 1:
 * 
 * 
 * Input: 121
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: -121
 * Output: false
 * Explanation: From left to right, it reads -121. From right to left, it
 * becomes 121-. Therefore it is not a palindrome.
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: 10
 * Output: false
 * Explanation: Reads 01 from right to left. Therefore it is not a
 * palindrome.
 * 
 * 
 * Follow up:
 * 
 * Coud you solve it without converting the integer to a string?
 * 
 */
bool isPalindrome(int x) {

    if (x < 0)
        return false;

    if (x == 0)
        return true;

    int x1 = x;
    int reverse = 0;


    while((0 != x1 / 10) || ( 0 != x1 % 10))
    {

        int d = x1 % 10;
        x1 = x1 / 10;
        reverse = reverse * 10 + d;
    }

    if( 0 == x - reverse)
        return true;
    
    return false;
    
}
