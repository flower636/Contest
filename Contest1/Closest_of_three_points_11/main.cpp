#include <iostream>
#include <cmath>
int main(){
int A, B, C, d, p;
std::cin >> A >> B >> C;
d = abs(A - B);
p = abs(A - C);
if(d < p){
std::cout << "B" << " " << d;
}
else{
std::cout << "C" << " " << p;
}
}