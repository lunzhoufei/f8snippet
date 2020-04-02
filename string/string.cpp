
  ostringstream oss;
  uint64_t timeDelta = currentTime - i_time;
  if (follow == 1) {
    oss << "已关注";
    style1 = "#00000099";
    style2 = "";
  /* } else if (t1_year == t2_year && t1_month == t2_month && t1_day == t2_day) { */
  } else if (timeDelta < 12 * 60 * 60) {
    oss << "最新推荐";
    style1 = "#489571";
    style2 = "";
  } else if (readcnt >= 50000) {
    oss << "热门推荐";
    style1 = "#B27324";
    style2 = "#D8D8D8";
  } else if (readcnt < 50000) {
    oss << "{time}";
    style1 = "#00000099";
    style2 = "";
  }
  details = oss.str();
  return;
