/*
 * [5] Longest Palindromic Substring
 *
 * https://leetcode.com/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (25.70%)
 * Total Accepted:    407.3K
 * Total Submissions: 1.6M
 * Testcase Example:  '"babad"'
 *
 * Given a string s, find the longest palindromic substring in s. You may
 * assume that the maximum length of s is 1000.
 * 
 * Example 1:
 * 
 * 
 * Input: "babad"
 * Output: "bab"
 * Note: "aba" is also a valid answer.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: "cbbd"
 * Output: "bb"
 * 
 * 
 */


// if palindromic, return not zero, it not return 0
int is_palindromic(char* s, int start, int end) 
{
    int si, ei = 0;
    si = start;
    ei = end;

    while((si < ei) && (s[si] == s[ei]))
    {
        si++;
        ei--;
    }

    if( si >= ei )
    {
        return end - start +1 ;
    }

    return 0;
}

char* longestPalindrome(char* s) {
    int l = 0;
    int li = 0;


    for(int i = 0 ; s[i] != '\0' ; i++)
    {
        for(int j = 0 ; s[j] != '\0' ; j++)
        {
            int t = is_palindromic(s, i, j);

            if(l < t)
            {
                l = t;
                li = i;
            }
        }
    }

    s[li + l] = '\0';
    return &s[li];
}
