

func CalculateCollatz(number: Int){
    print(number, terminator:"")
    print(",", terminator:"")
    if (number == 1 || number == 2){
        print("\n", terminator:"")
    }else{
        if (number % 2 == 0){
            let final_number = number / 2;
            CalculateCollatz(number: final_number);
        }else{
            let final_number = number * 3 + 1;
            CalculateCollatz(number: final_number);
        }
    }
}

func main(){
    let arguments = CommandLine.arguments
    let numbers_raw = arguments[1]
    let numbers = Int(numbers_raw)!
    var i = 1
    while numbers >= i{
        CalculateCollatz(number: i)
        i += 1
    }
}

main()