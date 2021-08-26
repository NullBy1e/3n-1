using System;

namespace _3n_1
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("MainNumber, Parts");
            int number = Int32.Parse(args[0]);
            for (int i=1;i<=number;i++){
                CalculateCollatz(i);
            }
        }
        static void CalculateCollatz(int number){
            Console.Write(number);
            Console.Write(",");
            if (number == 1 ||number ==2){
                Console.Write("\n");
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
}
