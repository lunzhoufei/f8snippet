
namespace tools {

  string format_time(const time_t& time) {
    char buf[64] = {0};
    strftime(buf, sizeof(buf), "%Y年%m月%d日 %H:%M:%S", localtime(&time));
    return string(buf);
  }

int recored_time() {
    struct timeval t_begin;
    struct timeval t_end;
    gettimeofday(&t_begin, NULL);
    /// TODO
    gettimeofday(&t_end, NULL);
    float t_cost = 1000000 * (t_end.tv_sec - t_begin.tv_sec) +
        (t_end.tv_usec - t_begin.tv_usec);
    printf("DDDDD->judge one fiend! time[%f ms]", t_cost / 1000);
    return 0;
}

time_t strtotime(char* const date,char* const format="%Y%m%d%H%M%S") {
    struct tm tm;
    strptime(date,format, &tm) ;
    time_t ft=mktime(&tm);
    return ft;
}

string   timetodate(time_t const timer) {
    struct tm *l=localtime(&timer);

    char buf[128];
    snprintf(buf,sizeof(buf),"%04d-%02d-%02d %02d:%02d:%02d",
            l->tm_year+1900, l->tm_mon+1, l->tm_mday, l->tm_hour,
            l->tm_min, l->tm_sec);
    string s(buf);
    return s;
}

void unixtimestamp2ymd() {
  time_t t_current = time(NULL);
  tm* t_cur = localtime(&t_current);
  int t_cur_year = t_cur->tm_year + 1900;
  int t_cur_month = t_cur->tm_mon + 1;
  int t_cur_day = t_cur->tm_mday;
}

#ifndef _TM_DEFINED
struct tm {
    int tm_sec;     /* 秒 – 取值区间为[0,59] */
    int tm_min;     /* 分 - 取值区间为[0,59] */
    int tm_hour;    /* 时 - 取值区间为[0,23] */
    int tm_mday;    /* 一个月中的日期 - 取值区间为[1,31] */
    int tm_mon;     /* 月份（从一月开始，0代表一月） - 取值区间为[0,11] */
    int tm_year;    /* 年份，其值等于实际年份减去1900 */
    int tm_wday;    /* 星期 – 取值区间为[0,6]，其中0代表星期天，1代表星期一，以此类推 */
    int tm_yday;    /* 从每年的1月1日开始的天数 – 取值区间为[0,365]，其中0代表1月1日，1代表1月2日，以此类推 */
    int tm_isdst;   /* 夏令时标识符，实行夏令时的时候，tm_isdst为正。不实行夏令时的进候，tm_isdst为0；不了解情况时，tm_isdst()为负。*/
        };
#define _TM_DEFINED
#endif}

}

/*---------------------------------------------------------------------------*/

#include <string>

using std::string;

/******************************************************************************
 *                              TimeShot
 *****************************************************************************/

class TimeShot {
  public:
    float get_interval(const string& start_pos, const string& end_pos);
    int shot(const string& pos);
    int clear();
  private:
    boost::unordered_map<string, struct timeval> m_st;
};

TimeShot::TimeShot() {
}

TimeShot::~TimeShot() {
}

int TimeShot::clear() {
  m_st.clear();
  return 0;
}

int TimeShot::shot(const string& pos) {
  int t_ret = 0;
  struct timeval t_t;
  gettimeofday(&t_t, NULL);
  m_st.insert(make_pair(pos, t_t));
  return t_ret;
}

float TimeShot::get_interval(const string& start, const string& end) {
  int t_ret = 0;
  struct timeval t_start;
  struct timeval t_end;
  boost::unordered_map<string, struct timeval>::iterator t_itr;
  t_itr = m_st.find(start);
  if (t_itr == m_st.end()) {
    return -1.0;
  } else {
    t_start = t_itr->second;
  }

  t_itr = m_st.find(end);
  if (t_itr == m_st.end()) {
    return -1.0;
  } else {
    t_end = t_itr->second;
  }
  return (1000 * (t_end.tv_sec - t_start.tv_sec) +
      (t_end.tv_usec - t_start.tv_usec) / 1000);
}
