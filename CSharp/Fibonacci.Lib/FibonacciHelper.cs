namespace Fibonacci.Lib
{
    internal static class FibonacciHelper
    {
        public static int GetNextNumber(int x, int y)
        {
            var nextNumber = x + y;
            if (nextNumber == 0)
            {
                return 1;
            }

            return nextNumber;
        }

        public static (int, int) GetPriorValues(IEnumerable<int> numbers)
        {
            var maxIndex = numbers.Count() - 1;
            return (GetValueOrDefault(maxIndex, numbers), GetValueOrDefault(maxIndex - 1, numbers));
        }

        private static int GetValueOrDefault(int index, IEnumerable<int> numbers)
        {
            if (index >= 0)
            {
                return numbers.ToArray()[index];
            }

            return 0;
        }
    }
}
