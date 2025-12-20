namespace PuzzleBox;

public static class Part2
{
    public static string Solve(List<string> puzzleInput)
    {
        int floor = 0;

        for (int i = 0; i < puzzleInput[0].Length; i++)
        {
            floor += puzzleInput[0][i] == '(' ? 1 : -1;
            if (floor == -1)
                return (i + 1).ToString();
        }

        return "0";
    }
}
