
/******************************************************************************
 *                                anti-csrf token
 *****************************************************************************/


/******************************************************************************
 *                                anti-xss filter
 *****************************************************************************/
// 常用的有html注入和js注入
// 提交表单的时候:
// 1. jsonp 过滤
// 2. 正则表达过滤

#include <regex.h>
  regex_t                       m_vid_pattern;

int init() {
  t_ret = regcomp(&m_vid_pattern,
      "^[01]/[0-9]{1,4}/[0-9a-z]{32}\\.[0-9a-z]{1,4}$",
      REG_EXTENDED | REG_ICASE | REG_NOSUB);
  if (0 != t_ret) {
    API_Error_Log(LM_ERROR, "regcomp failed! ret[%d]");
    exit(-1);
  }
}

bool MyCGI::is_vid_valid(const string& vid) {
  // boost 1.6 is unavailable
  // static boost::cregex t_reg = 
  //  boost::cregex::compile("^[01]/[0-9]{1,4}/[0-9a-z]{32}\.[0-9a-z]{1,4}$");
  // return boost::regex_match(vid.c_str(), t_reg);
  regmatch_t pmatch[1];
  const size_t nmatch = 1;
  int t_ret = regexec(&m_vid_pattern, vid.c_str(), nmatch, pmatch, 0);
  return t_ret == 0;
}



