#include <iostream>
#include <vector>
using namespace std;

int bubble_sort(vector<int>& iv) {
  if (iv.empty()) return 0;
  int max_pos = 0;
  for (size_t i = 0, size = iv.size(); i < size; ++i) {
    for (size_t j = 0; j < size - i; ++j) {
      if (iv[j] > iv[max_pos]) {
        int temp = iv[max_pos];
        iv[max_pos]

        max_pos = j;
      }
    }
  }
  return 0;
}

