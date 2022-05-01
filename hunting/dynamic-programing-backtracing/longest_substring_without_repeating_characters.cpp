// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
//
//
#include<iostream>
using namespace std;


class Solution {
  public:
    int lengthOfLongestSubstring(string s) {
      if (s.size() == 0) return 0;
      map<char, int> tm_charpos;
      map<char, int>::iterator t_cpitr;

      vector<int> tv_maxlen(s.size());
      int max = 0;
      for (int i = 0, size = s.size(); i < size; ++i) {
        if (i == 0) {

          tv_maxlen[i] = 1;
          tm_charpos[s[i]] = i;
          if (tv_maxlen[i] > max) max = tv_maxlen[i];
          continue;
        }
        t_cpitr = tm_charpos.find(s[i]);
        if (t_cpitr == tm_charpos.end()) {
          tv_maxlen[i] = tv_maxlen[i-1] + 1;
          tm_charpos[s[i]] = i;
          if (tv_maxlen[i] > max) max = tv_maxlen[i];
          continue;
        }
        tv_maxlen[i] = i - t_cpitr->second > tv_maxlen[i-1] + 1? tv_maxlen[i-1] + 1 : i - t_cpitr->second;
        tm_charpos[s[i]] = i;
        if (tv_maxlen[i] > max) max = tv_maxlen[i];
      }
      return max;
    }
};
