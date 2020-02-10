
#define MD5_KEY                 "MD5SALAT"

using namespace qmcgi;
int verify_anti_hotlinking_key(string& code)
{
    int t_ret = 0;
    if(code.size() < 5)
    {
        printlog(LM_ERROR, "code size error, len[%zu]", code.size());
        return -1;
    }
    string t_md5_content = m_from_mid ? 
        m_songmid + string(MD5_KEY) : m_songid_str + string(MD5_KEY);

    char t_md5value[128];
    int ai_case = 0;
    isd_md5_str( (unsigned char*)(t_md5_content.c_str()), t_md5_content.size(), 
            t_md5value, ai_case);
    string server_code = string(t_md5value, 0, 5);
    string client_code = string(lowerCase_Ex(code), 0, 5);
    t_ret = (server_code == client_code) ? 0 : -2;
    if(t_ret != 0)
    {
        printlog(LM_ERROR, "content[%s] client_key[%s] server_key[%s]", 
                t_md5_content.c_str(), code.c_str(), server_code.c_str());
    }
    return t_ret;
}
