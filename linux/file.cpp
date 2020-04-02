
#include <fstream>
    fstream fs_(file.c_str());

    string t_line = "";
    int t_cnt = 0;
  for(; getline(fs_, t_line, '\n'); ++t_cnt)
  {
    uint64_t t_uin;
    t_ret = build_value(t_line, t_val, t_uin);
    if(t_ret != 0) continue;
#if MOREWORK == 1
    more_work(t_uin, phase_); /// 多余的工作回调哦
#endif
    get_key(t_uin, t_key);
    t_ret = set(t_key, t_val);
    if(t_ret == 0)
    {
      ++t_ok;
      if(t_ok % 10000 == 0)
      {
        printlog(LM_ERROR, "already load [%dw] lines", t_ok/10000);
      }
    }
  }

/*---------------------------------------------------------------------------*/

int read_in() {
  std::string m_file_path = "";
  std::vector<uint64_t> m_v_uin;
  if (!m_file_path.empty()) {
    fstream t_fs(m_file_path.c_str());
    string t_line = "";
    int t_cnt = 0;
    for (; getline(t_fs, t_line, '\n'); ++t_cnt) {
      m_v_uin.push_back(atoll(t_line.c_str()));
    }
  }
}

int write_out() {
  ofstream                  m_failed_log;
  char buf[1024] = {0};
  snprintf(buf, sizeof(buf), "./cnf/%s.log", argv[0]);
  m_failed_log.open(buf, ios::in);
  m_failed_log << static_cast<int>(status) << "\t" << tid << endl;

  m_failed_log.flush();
  m_failed_log.close();
}





  int get_file_size(const string& path);

int UgcMvSessionOpt::get_file_size(const string& path) {
  int filesize = -1;  
  FILE *fp;  
  fp = fopen(path.c_str(), "r");  
  if(fp == NULL)  
    return filesize;  
  fseek(fp, 0L, SEEK_END);  
  filesize = ftell(fp);  
  fclose(fp);  
  return filesize;  
}





