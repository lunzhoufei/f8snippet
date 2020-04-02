
size_t seed_hash(const uint64_t& seed, const string& key) {
  const char* p_str = key.c_str();
  size_t hash = 0;
  size_t i;
  for (i = 0; *p_str; i++) {
    hash = hash * seed + (*p_str++);  // BKDRHash算法
    if (*p_str) {
      if ((i & 1) == 0) {
        hash ^= ((hash << 7) ^ (*p_str++) ^ (hash >> 3));
      } else {
        hash ^= (~((hash << 11) ^ (*p_str++) ^ (hash >> 5)));
      }
    }
  }
  return (hash & 0x7FFFFFFF);
}
