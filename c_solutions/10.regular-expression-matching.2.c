/*
 * [10] Regular Expression Matching
 *
 * https://leetcode.com/problems/regular-expression-matching/description/
 *
 * algorithms
 * Hard (24.53%)
 * Total Accepted:    247.5K
 * Total Submissions: 1M
 * Testcase Example:  '"aa"\n"a"'
 *
 * Given an input string (s) and a pattern (p), implement regular expression
 * matching with support for '.' and '*'.
 * 
 * 
 * '.' Matches any single character.
 * '*' Matches zero or more of the preceding element.
 * 
 * 
 * The matching should cover the entire input string (not partial).
 * 
 * Note:
 * 
 * 
 * s could be empty and contains only lowercase letters a-z.
 * p could be empty and contains only lowercase letters a-z, and characters
 * like . or *.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input:
 * s = "aa"
 * p = "a"
 * Output: false
 * Explanation: "a" does not match the entire string "aa".
 * 
 * 
 * Example 2:
 * 
 * 
 * Input:
 * s = "aa"
 * p = "a*"
 * Output: true
 * Explanation: '*' means zero or more of the precedeng element, 'a'.
 * Therefore, by repeating 'a' once, it becomes "aa".
 * 
 * 
 * Example 3:
 * 
 * 
 * Input:
 * s = "ab"
 * p = ".*"
 * Output: true
 * Explanation: ".*" means "zero or more (*) of any character (.)".
 * 
 * 
 * Example 4:
 * 
 * 
 * Input:
 * s = "aab"
 * p = "c*a*b"
 * Output: true
 * Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore
 * it matches "aab".
 * 
 * 
 * Example 5:
 * 
 * 
 * Input:
 * s = "mississippi"
 * p = "mis*is*p*."
 * Output: false
 * 
 * 
 */

bool matchChar(char s, char p)
{
    if(s == p)
    {
        return true;
    }

    if(p == '.')
    {
        return true;
    }

    return false;
}



bool isMatch(char* s, char* p) {

    int i = 0;
    int j = 0;

    if(p[j] == '\0')
    {
        return false;
    }


    while(s[i] != '\0')
    {
        if(p[j+1] == '*')
        {
            if(isMatch(&s[i], &p[j+2]))
            {
                return true;
            }

            if(true == matchChar(s[i], p[j]))
            {
                i++;
                continue;
            }
            else 
            {
                j+=2;
                continue;

            }
        }

        if(false == matchChar(s[i], p[j]))
        {
            return false;
        }

        i++;
        j++;
    }

    if(p[j] == '*')
    {
        j++;
    }

    if(p[j+1] == '*')
    {
        j+=2;
    }

    if(p[j] != '\0')
    {
        return false;
    }


    return true;
    
}
