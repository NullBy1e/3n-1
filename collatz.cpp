#include <iostream>
#include <stdlib.h>

using namespace std;

void calculateCollatz(long number){
    cout << to_string(number) << ",";
    if (number == 1 || number == 2 ){
        cout << "\n";
    }else{
        if (number % 2 == 0){
            double long final_number = number / 2;
            calculateCollatz(final_number);
        }else{
            double long final_number = number * 3 + 1;
            calculateCollatz(final_number);
        }
    }
    return;
}

int main(int argc, char *argv[]) {
    cout << "MainNumber, Parts" << "\n";
    char *p;
    if (argc < 2){
        cerr << "pass number to argv\n";
        return 1;
    }
    int numbers = strtol(argv[1], &p, 10);
    for (double long i = 1; i <= numbers;i++){
        calculateCollatz(i);
    }
    return 0;
}
