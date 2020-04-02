

/*---------------------------------------------------------------------------*/

#include <signal.h>
#include <execinfo.h>

#define MAX_FRAME_DEEP 32
void get_backtrace(int signal, int fd)
{
    char proc_name[512] = {0};
    readlink("/proc/self/exe", proc_name, sizeof(proc_name));
    //强制\0结尾
    proc_name[sizeof(proc_name) - 1] = '\0';

    char _line[512] = {0};
    int ret = snprintf(_line, sizeof(_line), 
            "=====process(%d) catch signal %d @%d=====\n'%s'\n", 
            getpid(), signal, (int)time(NULL), proc_name);
    //多个堆栈的分割符m
    write(fd, _line, ret);

    //帧地址数组
    void* frame[MAX_FRAME_DEEP];
    //reals <= MAX_FRAME_DEEP
    int reals = backtrace(frame, MAX_FRAME_DEEP);
    char** sysbomls = backtrace_symbols(frame, reals);
    if (sysbomls)
    {
        backtrace_symbols_fd(frame, reals, fd);
    }
}

int _target_fd = 1;
void core_sigal_handle(int signal)
{
    get_backtrace(signal, _target_fd);

    //打印已经完成,该怎样还是怎样
    struct sigaction act;
    act.sa_flags = SA_RESTART;
    act.sa_handler = SIG_DFL;
    sigemptyset(&act.sa_mask);
    sigaction(signal, &act, NULL);
    kill(getpid(), signal);
}

#define TARGET_FILE "/data/log/core.log"
void __attribute((constructor)) core_init(void)
{
    int fd = open(TARGET_FILE, O_RDWR|O_CREAT|O_APPEND, 0644);
    if (fd > 0)
    {
        _target_fd = fd;
    }

    //默认处理动作是core的一些信号
    struct sigaction act, oact;
    act.sa_flags = SA_RESTART;   
    act.sa_handler = core_sigal_handle;
    sigemptyset(&act.sa_mask);
    sigaction(SIGABRT, &act, &oact);
    sigaction(SIGBUS, &act, &oact);
    sigaction(SIGFPE, &act, &oact);
    sigaction(SIGILL, &act, &oact);
    sigaction(SIGSEGV, &act, &oact);
    sigaction(SIGSYS, &act, &oact);
}
