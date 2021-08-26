public class collatz {
    public static void main(String args[]){
        System.out.println("MainNumber, Parts");
        int numbers = Integer.parseInt(args[0]);
        for (int i =1;i<=numbers;i++){
            CalculateCollatz(i);
        }
    }
    public static void CalculateCollatz(int number){
        System.out.print(number);
        System.out.print(",");
        if (number == 1 || number == 2){
            System.out.print("\n");
        }else{
            if (number % 2 == 0){
                int final_number = number / 2;
                CalculateCollatz(final_number);
            }else{
                int final_number = number * 3 + 1;
                CalculateCollatz(final_number);
            }
        }
    }
}