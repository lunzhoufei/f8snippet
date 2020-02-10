
#include "curl.h"

static size_t write_callback(char* ptr, size_t size, size_t nmemb,
    void* userdata) {
  string* pstr = (string*)userdata;  // NOLINT
  pstr->append(ptr, size * nmemb);
  return size * nmemb;
}

int TestMgz::fetch_pic_data(const string& url,
    MusicImageUpload::UploadParams* p_param) {
  if (!p_param) return 0;
  int t_ret = 0;
  CURL *curl = curl_easy_init();
  if (!curl) {
    printf("curl easy_init failed!\n");
    return -99;
  }

  // init
  do {
    curl_easy_setopt(curl, CURLOPT_URL, url.c_str());
    curl_easy_setopt(curl, CURLOPT_TIMEOUT, 5000);
    /* curl_easy_setopt(curl, CURLOPT_FOLLOWLOCATION, 1); */
    curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, write_callback);
    curl_easy_setopt(curl, CURLOPT_WRITEDATA, &p_param->pic_data);
    /* curl_easy_setopt(curl, CURLOPT_HTTPHEADER, plist); */

    CURLcode rcode = curl_easy_perform(curl);
    if (rcode != CURLE_OK) {
      printf("curl_easy_perform failed! rcode[%d]\n", rcode);
      t_ret = -101;
      break;
    }

    long http_code;
    rcode = curl_easy_getinfo(curl, CURLINFO_RESPONSE_CODE, &http_code);
    if (rcode != CURLE_OK) {
      printf("curl_easy_getinfo failed!\n");
      t_ret =  -201;
      break;
    }
    if (http_code != 200) {
      printf("curl_easy_getinfo ret http_code != 200\n");
      t_ret = -301;
      break;
    }
  } while(0);

  curl_easy_cleanup(curl);
  return t_ret;
}
