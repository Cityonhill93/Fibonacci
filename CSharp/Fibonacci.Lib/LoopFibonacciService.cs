namespace Fibonacci.Lib
{
    public class LoopFibonacciService : IFibonacciService
    {
        public IEnumerable<int> GetNumbers(int count)
        {
            var numbers = new List<int>();
            for (var i = 0; i < count; i++)
            {
                if (i != 0)
                {
                    var priorValues = FibonacciHelper.GetPriorValues(numbers);
                    var nextValue = FibonacciHelper.GetNextNumber(priorValues.Item1, priorValues.Item2);

                    numbers.Add(nextValue);

                }
                else
                {
                    numbers.Add(0);
                }

            }

            return numbers;
        }
    }
}
