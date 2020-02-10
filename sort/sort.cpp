
bool sort_method(const T1& p1, const T1& p2) {
  if (p1.usage() > p2.usage()) {
    return true;
  } else if (p1.usage() == p2.usage()) {
    return p1.time() > p2.time();
  } else {
    return false;
  }
}

// example
std::sort(t_vec_p.begin(), t_vec_p.end(), sort_method);

// update_time
bool sort_favbase_update_time(const T1& p1, const T1& p2) {
  if (p1.update_time > p2.update_time) {
    return true;
  } else if (p1.update_time == p2.update_time) {
    return p1.create_time > p2.create_time;
  } else {
    return false;
  }
}

bool sort_favbase_create_time(const T1& p1, const T1& p2) {
  if (p1.create_time > p2.create_time) {
    return true;
  } else if (p1.create_time == p2.create_time) {
    return p1.update_time > p2.update_time;
  } else {
    return false;
  }
}
