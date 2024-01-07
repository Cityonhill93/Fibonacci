using Fibonacci.Lib;

const int NumberCount = 10;

Console.WriteLine("Welcome to the Fibonacci app!");

PrintNumbers(NumberCount, new LoopFibonacciService());
PrintNumbers(NumberCount, new RecursiveFibonacciService());

_ = Console.ReadLine();

void PrintNumbers(int count, IFibonacciService service)
{
    Console.WriteLine($"Starting the {service.GetType().Name}...");

    var numbers = service.GetNumbers(count);
    foreach (var number in numbers)
    {
        Console.WriteLine(number);
    }

    Console.WriteLine($"End of the {service.GetType().Name}...");
}
