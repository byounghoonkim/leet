/*
 * [12] Integer to Roman
 *
 * https://leetcode.com/problems/integer-to-roman/description/
 *
 * algorithms
 * Medium (48.37%)
 * Total Accepted:    179.4K
 * Total Submissions: 370.8K
 * Testcase Example:  '3'
 *
 * Roman numerals are represented by seven different symbols: I, V, X, L, C, D
 * and M.
 * 
 * 
 * Symbol       Value
 * I             1
 * V             5
 * X             10
 * L             50
 * C             100
 * D             500
 * M             1000
 * 
 * For example, two is written as II in Roman numeral, just two one's added
 * together. Twelve is written as, XII, which is simply X + II. The number
 * twenty seven is written as XXVII, which is XX + V + II.
 * 
 * Roman numerals are usually written largest to smallest from left to right.
 * However, the numeral for four is not IIII. Instead, the number four is
 * written as IV. Because the one is before the five we subtract it making
 * four. The same principle applies to the number nine, which is written as IX.
 * There are six instances where subtraction is used:
 * 
 * 
 * I can be placed before V (5) and X (10) to make 4 and 9. 
 * X can be placed before L (50) and C (100) to make 40 and 90. 
 * C can be placed before D (500) and M (1000) to make 400 and 900.
 * 
 * 
 * Given an integer, convert it to a roman numeral. Input is guaranteed to be
 * within the range from 1 to 3999.
 * 
 * Example 1:
 * 
 * 
 * Input: 3
 * Output: "III"
 * 
 * Example 2:
 * 
 * 
 * Input: 4
 * Output: "IV"
 * 
 * Example 3:
 * 
 * 
 * Input: 9
 * Output: "IX"
 * 
 * Example 4:
 * 
 * 
 * Input: 58
 * Output: "LVIII"
 * Explanation: L = 50, V = 5, III = 3.
 * 
 * 
 * Example 5:
 * 
 * 
 * Input: 1994
 * Output: "MCMXCIV"
 * Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
 * 
 */

void cat(char* ret, int n, char* c, char* c5, char* c10)
{
    if(n < 4)
    {
        for(int i = 0; i < n; i++)
        {
            strcat(ret, c);
        }

    }
    else if(n == 4)
    {
        strcat(ret, c);
        strcat(ret, c5);
    }
    else if(n < 9 )
    {
        strcat(ret, c5);
        for(int i = 5; i < n; i++)
        {
            strcat(ret, c);
        }
    }
    else if(n == 9)
    {
        strcat(ret, c);
        strcat(ret, c10);
    }
}


char* intToRoman(int num) {
    char* ret = malloc(sizeof(char) * 30);
    ret[0] = '\0';

    int m = num / 1000;
    int c = (num % 1000) / 100;
    int x = (num % 100 ) / 10;
    int i = (num % 10  ) / 1;

    cat(ret, m, "M", "", "");
    cat(ret, c, "C", "D", "M");
    cat(ret, x, "X", "L", "C");
    cat(ret, i, "I", "V", "X");

    return ret;
}
