#include <iostream>
#include <vector>
#include "stdio.h"
using namespace std;

int select_sort(vector<int>& iv) {
  if (iv.empty()) return 0;
  for (size_t i = 0, size = iv.size(); i < size; ++i) {
    int max_pos = 0;
    for (size_t j = 0; j < size - i; ++j) {
      if (iv[j] > iv[max_pos]) {
        max_pos = j;
      }
    }
    int temp = iv[size - i - 1];
    iv[size - i - 1] = iv[max_pos];
    iv[max_pos] = temp;
  }
  return 0;
}

int main() {
  vector<int> tv;
  tv.push_back(9);
  tv.push_back(19);
  tv.push_back(4);
  tv.push_back(819);
  tv.push_back(11);
  tv.push_back(12);
  tv.push_back(1991);
  tv.push_back(39);
  tv.push_back(79);
  tv.push_back(921);
  tv.push_back(3);

  select_sort(tv);
  for (size_t i = 0, size = tv.size(); i < size; ++i) {
    printf("%d\n", tv[i]);
  }
}
