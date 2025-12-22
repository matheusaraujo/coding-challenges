import java.util.List;

public class part2 {
  private static final char OPEN = '(';

  public static Object solve(final List<String> puzzleInput) {
    int floor = 0;
    int i = 0;

    for (; i < puzzleInput.get(0).length(); i++) {
      if (puzzleInput.get(0).charAt(i) == OPEN) {
        floor++;
      } else {
        floor--;
      }

      if (floor == -1) {
        break;
      }
    }
    return Integer.toString(floor == -1 ? i + 1 : 0);
  }
}
