/*
 * [6] ZigZag Conversion
 *
 * https://leetcode.com/problems/zigzag-conversion/description/
 *
 * algorithms
 * Medium (29.17%)
 * Total Accepted:    253.5K
 * Total Submissions: 868.7K
 * Testcase Example:  '"PAYPALISHIRING"\n3'
 *
 * The string "PAYPALISHIRING" is written in a zigzag pattern on a given number
 * of rows like this: (you may want to display this pattern in a fixed font for
 * better legibility)
 * 
 * 
 * P   A   H   N
 * A P L S I I G
 * Y   I   R
 * 
 * 
 * And then read line by line: "PAHNAPLSIIGYIR"
 * 
 * Write the code that will take a string and make this conversion given a
 * number of rows:
 * 
 * 
 * string convert(string s, int numRows);
 * 
 * Example 1:
 * 
 * 
 * Input: s = "PAYPALISHIRING", numRows = 3
 * Output: "PAHNAPLSIIGYIR"
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: s = "PAYPALISHIRING", numRows = 4
 * Output:    "PINALSIGYAHRPI"
 * Explanation:
 * 
 * P     I    N
 * A   L S  I G
 * Y A   H R
 * P     I
 * 
 */

char* convert(char* s, int numRows) {

    if(numRows == 1)
    {
        return s;
    }

    int len = strlen(s); 
    char* ret = malloc(sizeof(char) * (len + 1));
    ret [len] = '\0';
    int ret_index = 0;

    int jump_size = 2 * (numRows - 1);

    for(int i = 0 ; i < numRows ; i++) 
    {
        for(int j = i; j < len ; j += jump_size)
        {
            if (j >= len)
            {
                break;
            }

            ret[ret_index] = s[j];
            ret_index++;

            if( (0 < i) && (numRows -1 > i) )
            {
                int onemore = j - i + jump_size - i;
                if (onemore >= len)
                {
                    break;
                }
                ret[ret_index] = s[onemore];
                ret_index++;
            }
        }
    }
    
    return ret;
}
