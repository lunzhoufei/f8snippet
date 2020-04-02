
  pthread_cond_init(&m_cond, NULL);
  m_mutex = (pthread_mutex_t*)malloc(sizeof(pthread_mutex_t));
  pthread_mutex_init(m_mutex, NULL);
  m_deq_sys_msg.clear();
