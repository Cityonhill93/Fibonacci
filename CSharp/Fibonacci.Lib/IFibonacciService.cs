namespace Fibonacci.Lib
{
    public interface IFibonacciService
    {
        IEnumerable<int> GetNumbers(int count);
    }
}
