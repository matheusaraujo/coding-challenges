namespace PuzzleBox;

public static class Part1
{
    public static Object Solve(List<string> puzzleInput)
    {
        return (puzzleInput[0].Count('(') - puzzleInput[0].Count(')')).ToString();
    }

    static int Count(this string input, char target)
    {
        var count = 0;
        foreach (char ch in input)
            if (ch == target)
                count++;
        return count;
    }
}
