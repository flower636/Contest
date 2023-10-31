#include <iostream>
#include <cmath>
int main(){
float a, b, d;
double c;
std::cin >> a >> b;
d = pow(a, 2) + pow(b, 2);
c = sqrt(d);
std::cout << c;
return 0;
}