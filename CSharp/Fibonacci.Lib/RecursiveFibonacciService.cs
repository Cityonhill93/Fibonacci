namespace Fibonacci.Lib
{
    public class RecursiveFibonacciService : IFibonacciService
    {
        public IEnumerable<int> GetNumbers(int count) => GetNumbersRecursive(count, new List<int>());

        private List<int> GetNumbersRecursive(int count, IEnumerable<int> numbers)
        {
            var outNumbers = numbers.ToList();
            if (numbers.Any())
            {
                var priorValues = FibonacciHelper.GetPriorValues(numbers);
                var nextValue = FibonacciHelper.GetNextNumber(priorValues.Item1, priorValues.Item2);
                outNumbers.Add(nextValue);
            }
            else
            {
                outNumbers.Add(0);
            }

            if (outNumbers.Count < count)
            {
                outNumbers = GetNumbersRecursive(count, outNumbers);
            }

            return outNumbers;
        }
    }
}
