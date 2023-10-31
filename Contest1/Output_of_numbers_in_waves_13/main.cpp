#include <iostream>
using namespace std;
int main(){
    int number_count = 1;
    int n, max_wave_size = 2, current_wave_size = 1;
    cin >> n;
    bool is_wave_grow = true;
    for (int a = 1; a <= n; a++){
        cout << a << " ";
        if (number_count == current_wave_size){
            cout << endl;
            current_wave_size +=  is_wave_grow ? 1: -1;
            number_count = 0;
        }
        if (current_wave_size == max_wave_size || current_wave_size == 1){
            is_wave_grow = !is_wave_grow;
            if(current_wave_size == max_wave_size){
            max_wave_size++;}
        }
        number_count++;
    }
    return 0;
}