
uint64_t get_random_interval(const uint64_t& fixed, const uint64_t& random);

uint64_t get_random_interval(const uint64_t& fixed,
    const uint64_t& random) {
  static uint64_t t_seed = time(NULL);
  srand(t_seed);
  t_seed = rand();
  /* *p = t_seed % (60 * 3); */
  return fixed + t_seed % (random);
}


