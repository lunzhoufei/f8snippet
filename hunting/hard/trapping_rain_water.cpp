// https://leetcode-cn.com/problems/trapping-rain-water/



class Solution {
  public:
    int trap(vector<int>& height) {


      vector<int> tv_summit;
      for (size_t i = 0, size = height.size(); i < size; ++i) {
        if ((i == 0 || height[i] > height[i-1]) && (i == size-1 || height[i] > height[i + 1])) {
          tv_summit.push_back(i);
        }
      }

      int total =0;
      for (size_t i = 1, size = tv_summit.size(); i < size; ++i) {
        int weith = height[tv_summit[i]] < height[tv_summit[i - 1]] ? height[tv_summit[i]]: height[tv_summit[i-1]];
        for (int k = tv_summit[i-1] + 1; k < tv_summit[i]; ++k) {
          total += (weith > height[k]) ? weith - height[k] : 0;
        }          
      }
      return total;       
    }
};
