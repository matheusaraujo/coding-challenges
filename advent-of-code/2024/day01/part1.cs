namespace PuzzleBox;

public static class Part1
{
    public static Object Solve(List<string> puzzleInput)
    {
        var (left, right) = Helpers.ParseInput(puzzleInput);
        return left.Zip(right, (l, r) => Math.Abs(l - r)).Sum().ToString();
    }
}
