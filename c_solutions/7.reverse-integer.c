/*
 * [7] Reverse Integer
 *
 * https://leetcode.com/problems/reverse-integer/description/
 *
 * algorithms
 * Easy (24.55%)
 * Total Accepted:    522.5K
 * Total Submissions: 2.1M
 * Testcase Example:  '123'
 *
 * Given a 32-bit signed integer, reverse digits of an integer.
 * 
 * Example 1:
 * 
 * 
 * Input: 123
 * Output: 321
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: -123
 * Output: -321
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: 120
 * Output: 21
 * 
 * 
 * Note:
 * Assume we are dealing with an environment which could only store integers
 * within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of
 * this problem, assume that your function returns 0 when the reversed integer
 * overflows.
 * 
 */

bool can_overflow(int ret)
{
    #define INT_MAX ((1<<31)-1)
    #define INT_MIN -(1<<31)

    if (ret != 0)
    {
        #define INT_MAX ((1<<31)-1)
        #define INT_MIN -(1<<31)
        if (ret > 0)
        {
            if ( INT_MAX / 10  < ret )
            {
                return true;
            }
        }
        else 
        {
            if ( INT_MIN / 10  > ret )
            {
                return true;
            }
        }
    }

    return false;
}

int reverse(int x) {
    int x1 = x;
    int ret = 0;

    while(true)
    {
        int xmod = x1 % 10;
        x1 /= 10;

        if (( xmod == 0) && (x1 == 0) )
        {
            break;
        }

        if(can_overflow(ret))
        {
            return 0;
        }

        ret = ret * 10 + xmod;

    }

    return ret;

}
