#include <iostream>
#include <string>
using namespace std;


string pwd()
{
  char pwd[256] = {0};
  getcwd(pwd, sizeof(pwd));
  return string(pwd);
}



int main (int argc, char const* argv[])
{
  cout<<pwd()<<endl;
  return 0;
}
