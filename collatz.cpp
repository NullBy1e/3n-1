#include <iostream>
#include <stdlib.h>

using namespace std;

void calculateCollatz(int number){
    cout << to_string(number) << ",";
    if (number == 1 || number == 2 ){
        cout << "\n";
    }else{
        if (number % 2 == 0){
            int final_number = number / 2;
            calculateCollatz(final_number);
        }else{
            int final_number = number * 3 + 1;
            calculateCollatz(final_number);
        }
    }
}

int main(int argc, char *argv[]) {
    cout << "MainNumber, Parts" << "\n";
    char *p;
    int numbers = strtol(argv[1], &p, 10);
    for (int i = 1; i <= numbers;i++){
        calculateCollatz(i);
    }
    return 0;
}