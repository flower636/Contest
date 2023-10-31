#include <iostream>
#include <algorithm>
using namespace std;

int main()
{
    int n;
    string s = "";
    cin >> n;
    while (n > 0){
        s.push_back((n - 1) % 26 + 'A');
        n = (n - 1) / 26;
    }
    reverse(s.begin(), s.end());
    cout << s;
}