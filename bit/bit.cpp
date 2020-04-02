
#define SET_FLAG(flag,b) (flag = flag|(0x1<<(b)))
#define CLEAR_THE_FLAG(flag,b)(flag = flag & ~(0x<<(b)))
#define CLEAR_FLAG(flag) (flag=0)
#define TEST_FLAG(flag,b) (flag&(0x1<<(b)))
