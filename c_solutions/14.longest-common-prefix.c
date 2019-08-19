/*
 * @lc app=leetcode id=14 lang=c
 *
 * [14] Longest Common Prefix
 *
 * https://leetcode.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (33.04%)
 * Total Accepted:    416.3K
 * Total Submissions: 1.3M
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * Write a function to find the longest common prefix string amongst an array
 * of strings.
 * 
 * If there is no common prefix, return an empty string "".
 * 
 * Example 1:
 * 
 * 
 * Input: ["flower","flow","flight"]
 * Output: "fl"
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: ["dog","racecar","car"]
 * Output: ""
 * Explanation: There is no common prefix among the input strings.
 * 
 * 
 * Note:
 * 
 * All given inputs are in lowercase letters a-z.
 * 
 */

int prefix_count(char* s1, char* s2) {
    int count = 0;
    int index = 0;

    if((s2 == 0) || (s1 == 0) ) {
        return count;
    } 

    while(1){
        if(s1[index] == '\0') break;
        if(s2[index] == '\0') break;
        if(s1[index] != s2[index]) break;

        index ++;
    }

    count = index;

    return count;
}

char sz[] = "";
char* longestCommonPrefix(char** strs, int strsSize) {
    if(strsSize == 0)
    {
        return sz;
    }

    int count = strlen(strs[0]);

    for(int i = 1; i < strsSize; i++) {
        int c = prefix_count(strs[0], strs[i]);

        if( c < count ) {
            count = c;
        }

        if( count == 0) {
            break;
        }
    }

    strs[0][count] = '\0';
    return strs[0];
}

