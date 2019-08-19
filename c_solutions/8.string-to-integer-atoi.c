/*
 * [8] String to Integer (atoi)
 *
 * https://leetcode.com/problems/string-to-integer-atoi/description/
 *
 * algorithms
 * Medium (14.20%)
 * Total Accepted:    287.4K
 * Total Submissions: 2M
 * Testcase Example:  '"42"'
 *
 * Implement atoi which converts a string to an integer.
 * 
 * The function first discards as many whitespace characters as necessary until
 * the first non-whitespace character is found. Then, starting from this
 * character, takes an optional initial plus or minus sign followed by as many
 * numerical digits as possible, and interprets them as a numerical value.
 * 
 * The string can contain additional characters after those that form the
 * integral number, which are ignored and have no effect on the behavior of
 * this function.
 * 
 * If the first sequence of non-whitespace characters in str is not a valid
 * integral number, or if no such sequence exists because either str is empty
 * or it contains only whitespace characters, no conversion is performed.
 * 
 * If no valid conversion could be performed, a zero value is returned.
 * 
 * Note:
 * 
 * 
 * Only the space character ' ' is considered as whitespace character.
 * Assume we are dealing with an environment which could only store integers
 * within the 32-bit signed integer range: [−231,  231 − 1]. If the numerical
 * value is out of the range of representable values, INT_MAX (231 − 1) or
 * INT_MIN (−231) is returned.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: "42"
 * Output: 42
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: "   -42"
 * Output: -42
 * Explanation: The first non-whitespace character is '-', which is the minus
 * sign.
 * Then take as many numerical digits as possible, which gets 42.
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: "4193 with words"
 * Output: 4193
 * Explanation: Conversion stops at digit '3' as the next character is not a
 * numerical digit.
 * 
 * 
 * Example 4:
 * 
 * 
 * Input: "words and 987"
 * Output: 0
 * Explanation: The first non-whitespace character is 'w', which is not a
 * numerical 
 * digit or a +/- sign. Therefore no valid conversion could be performed.
 * 
 * Example 5:
 * 
 * 
 * Input: "-91283472332"
 * Output: -2147483648
 * Explanation: The number "-91283472332" is out of the range of a 32-bit
 * signed integer.
 * Thefore INT_MIN (−231) is returned.
 * 
 */

bool is_sign(char c)
{
    if(c == '+') return true;
    if(c == '-') return true;

    return false;
}

int to_sign(char c)
{
    if(c == '+') return 1;
    if(c == '-') return -1;
    return 1;
}

bool is_digit(char c)
{
    if((c >= '0') && (c <= '9'))
        return true;
    return false;
}

int to_digit(char c)
{
    return (int)c - '0';
}


int myAtoi(char* str) {
    int i = 0;
    int ret = 0;
    int sign = 1;

    // Skip space;
    while(str[i] == ' ')
    {
        i++;
    }

    if(is_sign(str[i]))
    {
        sign = to_sign(str[i]);
        i++;
    }

    while(true == is_digit(str[i]))
    {
        int d = to_digit(str[i]);

        if (sign > 0) 
        {
            if(INT_MAX / 10 < ret)
            {
                ret = INT_MAX;
                break;
            }

            if((INT_MAX / 10 == ret) && (INT_MAX % 10 < d))
            {
                ret = INT_MAX;
                break;
            }
        }
        else
        {
            if(INT_MIN / 10 > ret)
            {
                ret = INT_MIN;
                break;
            }

            if((INT_MIN / 10 == ret) && (INT_MIN % 10 > sign * d))
            {
                ret = INT_MIN;
                break;
            }
        }

        ret = (ret * 10) + (sign * d);
        i++;
    }

    return ret;
}
