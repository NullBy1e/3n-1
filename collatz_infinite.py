import sys


def main():
    print("MainNumber, Parts")
    counter = 1
    while True:
        calculateCollatz(counter)
        counter += 1


def calculateCollatz(number):
    sys.stdout.write(str(number))
    if number == 1:
        sys.stdout.write("\n")
    else:
        sys.stdout.write(",")
        if number % 2 == 0:
            final_number = number / 2
            calculateCollatz(final_number)
        else:
            final_number = number * 3 + 1
            calculateCollatz(final_number)


main()
