import sys


def main():
    print("MainNumber, Parts")
    numbers = sys.argv[1]
    for i in range(1, int(numbers)+1):
        calculateCollatz(i)


def calculateCollatz(number):
    sys.stdout.write(str(number) + ",")
    if number == 1 or number == 2:
        sys.stdout.write("\n")
    else:
        if number % 2 == 0:
            final_number = number / 2
            calculateCollatz(final_number)
        else:
            final_number = number * 3 + 1
            calculateCollatz(final_number)


main()
