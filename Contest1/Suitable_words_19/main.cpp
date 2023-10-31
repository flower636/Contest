#include <iostream>
#include <vector>
#include <algorithm>
#include <set>
using namespace std;

int main(){
    vector <string> strings;
    set <string> strings_dup;

    while (true) {
        string s;
        cin >> s;

        if (s == "end") break;

        bool exists = false;
        for (int i = 0; i < strings.size(); i++) {
            if (strings[i] == s) {
                exists = true;
                strings_dup.insert(s);
                break;
            }
        }
        if (exists) continue ;

        strings.push_back(s);
    }
    for (auto &s : strings_dup) {
        cout << s << " ";
    }

    return 0;
}
