
namespace tools {
string get_file_base(const string& path) {
  string::size_type t_pos = path.rfind("/");
  return string(path, 0, t_pos);
}

// mkdir -p 的功能;
int check_and_create_folder(const string& path) {
  int t_ret = 0;
  struct stat t_st;
  std::stack<string> t_stack;
  string t_path = get_file_base(path);
  while(!t_path.empty()) {
    struct stat t_st;
    /* stat(t_path.c_str(), &t_st); */
    bool isdir = stat(t_path.c_str(), &t_st) == 0 && S_ISDIR(t_st.st_mode);
    if (!isdir) {
      t_stack.push(t_path);
    }
    t_path = get_file_base(t_path);
  }

  while(!t_stack.empty()) {
    string t_folder = t_stack.top();
    stat(t_folder.c_str(), &t_st);
    bool isdir = stat(t_folder.c_str(), &t_st) == 0 && S_ISDIR(t_st.st_mode);
    if (!isdir) {
      t_ret = mkdir(t_folder.c_str(), 0777);
      API_Error_Log(LM_ERROR, "mkdir path[%s] ret[%d]",
          t_path.c_str(), t_ret);
    }
    t_stack.pop();
  }
  return t_ret;
}

}

int UgcMvSessionOpt::concate_file(ugc_mv::UploadSession* p_session) {
  if (!p_session) {
    return -1;
  }
  if (p_session->path().empty()) {
    return -2;
  }
  map<int, string> t_m_idx_path;
  get_shards_absolute_path(p_session, &t_m_idx_path);
  string t_entire_file = m_path_prefix + p_session->path();
  if (t_m_idx_path.size() != p_session->sharding_cnt()) {
    return -3;
  }

  // XXX(lunzhoufei): should check if exist first?
  map<int, string>::iterator t_itr = t_m_idx_path.begin();
  std::ofstream of_dst(t_entire_file.c_str(), std::ios_base::binary);
  for (; t_itr != t_m_idx_path.end(); ++t_itr) {
    std::ifstream if_shard(t_itr->second.c_str(), std::ios_base::binary);
    if (!if_shard.good()) {
      printf("%s not found\n", t_itr->second.c_str());
    }
    of_dst << if_shard.rdbuf();
  }

  of_dst.flush();
  of_dst.close();
  return 0;
}
