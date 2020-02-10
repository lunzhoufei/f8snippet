#include <iostream>
#include <vector>
using namespace std;

int partition(vector<int>& list, const int& start, const int& end) {
  if (start >= end) return 0;
  int t_pivot = list[start];
  int t_start = start + 1;
  int t_end = end;
  int t_position = start;
  bool find_min = true;

  while (t_start <= t_end) {
    if (find_min) {
      while (list[t_end] >= t_pivot) --t_end;
      if (t_start > t_end) break;
      list[t_position] = list[t_end];
      find_min = false;
      t_position = t_end;
    } else {
      while (list[t_start] <= t_pivot) ++t_start;
      if (t_start > t_end) break;
      list[t_position] = list[t_start];
      find_min = true;
      t_position = t_start;
    }
  }
  list[t_position] = t_pivot;
  return t_position;
}

int quick_sort(vector<int>& list, const int& start, const int& end) {
  if (start < end) {
    int t_pivot = partition(list, start, end);
    quick_sort(list, start, t_pivot - 1);
    quick_sort(list, t_pivot + 1, end);
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

  quick_sort(tv, 0, tv.size() - 1);
  for (size_t i = 0, size = tv.size(); i < size; ++i) {
    printf("%d\n", tv[i]);
  }
  return 0;
}

