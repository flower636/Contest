#include <iostream>
using namespace std;

int main()
{
    string s_start, s_finish;
    cin >> s_start;
    int i, k = 1;
    for (i = 1; i <= s_start.length(); i++)
    {
        if (s_start[i] == s_start[i - 1])
        {
            k++;
        }
        if (s_start[i] != s_start[i - 1])
        {
            if (k > 1)
            {
                cout << s_start[i - 1] << k;
                k = 1;
            }
            else if (k = 1)
            {
                cout << s_start[i - 1];
            }
        }
    }
    return 0;
}
