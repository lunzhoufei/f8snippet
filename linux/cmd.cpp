
string exe_cmd(const string& cmd) {
  int t_ret = 0;
  string t_result = "";
  t_result.resize(1024);
  FILE* fp = popen(cmd.c_str(), "r");
  fgets(&t_result[0], t_result.capacity(), fp);
  pclose(fp);
  return t_result;
}

string get_md5(const string& file) {
  char buf[128] = {'\0'};
  string cmd = "md5sum ";
  cmd.append(file);
  cmd.append(" | awk '{printf $1}'");
  FILE* fp = popen(cmd.c_str(), "r");
  fgets(buf, sizeof(buf), fp);
  pclose(fp);
  return string(buf);
}
