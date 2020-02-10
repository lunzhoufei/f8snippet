
if (likely(strncmp(request_method, "POST", 4) == 0)) {
    g_enable_compress = 1;
    int len_content = 0;
    char *p = getenv("CONTENT_LENGTH");
    if (p != NULL) {
        len_content = atoi(p);
    }

    LOG(DEBUG, "content length:%d", len_content);

    if(len_content == 0)
    {
        len_content = MAX_LEN_CONTENT-1;
    }

    if (len_content > MAX_LEN_CONTENT-1) 
    {
        LOG(ERROR, "ERROR:http pkg lenth:[%d] too big",len_content);
        len_content = MAX_LEN_CONTENT-1;
    }

    cgi_param = &g_cgi_param_xml;
    cgi_param->reset();

    std::string & content = cgi_param->content;
    content.resize(len_content);

    int pos=0, n=0;
    while((n = read(0, (void *) (&content[pos]), len_content-pos))>0) 
    {
        pos+=n;
    }
    content.resize(pos);
    LOG(DEBUG, "content read:%d", int(content.size()));
    LOG(DEBUG, "post data is:<%s>", content.c_str());
}
