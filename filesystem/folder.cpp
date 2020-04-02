#include <sys/stat.h>
#include <unistd.h>
#include <string.h>
#include <string>
#include <vector>
#include <utility>
#include <map>

namespace tools {
string get_file_base(const string& path) {
  string::size_type t_pos = path.rfind("/");
  return string(path, 0, t_pos);
}

int check_and_create_folder(const string& folder) {
  int t_ret = 0;
  struct stat t_st;
  stat(folder.c_str(), &t_st);
  bool isdir = S_ISDIR(t_st.st_mode);
  if (!isdir) {
    t_ret = mkdir(folder.c_str(), 0777);
  }
  return t_ret;
}
}  // namespace tools
