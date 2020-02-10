#include <string>
#include <regex.h>
class Regexp {
  regex_t                       m_vid_pattern;

 public:
  int init() {
    int t_ret = regcomp(&m_vid_pattern,
        "^[01]/[0-9]{1,4}/[0-9a-z]{32}\\.[0-9a-z]{1,4}$",
        REG_EXTENDED | REG_ICASE | REG_NOSUB);
    if (0 != t_ret) {
      printf("regcomp failed! ret[%d]");
      exit(-1);
    }
  }

  bool is_vid_valid(const std::string& vid) {
    // static boost::cregex t_reg = 
    //  boost::cregex::compile("^[01]/[0-9]{1,4}/[0-9a-z]{32}\.[0-9a-z]{1,4}$");
    // return boost::regex_match(vid.c_str(), t_reg);
    regmatch_t pmatch[1];
    const size_t nmatch = 1;
    int t_ret = regexec(&m_vid_pattern, vid.c_str(), nmatch,
        pmatch, 0);
    return t_ret == 0;
  }

};
