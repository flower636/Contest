#include <iostream>
#include <cmath>
int main(){
    double pi;
    pi = sqrt(12.0)*(1.0 - (1.0/(3.0*pow(3.0,1))) + (1.0/(5.0*pow(3.0,2))) - (1.0/(7.0*pow(3.0,3))) + (1.0/(9.0*pow(3.0,4))) - (1.0/(11.0*pow(3.0,5))));
    std::cout << pi;
    return 0;
}