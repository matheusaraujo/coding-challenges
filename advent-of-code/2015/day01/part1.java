import java.util.List;

public class part1 {
  public static String solve(final List<String> puzzleInput) {
    return Integer.toString(count(puzzleInput.get(0), '(') - count(puzzleInput.get(0), ')'));
  }

  public static int count(final String input, final char target) {
    int count = 0;
    for (int i = 0; i < input.length(); i++) {
      if (input.charAt(i) == target) {
        count++;
      }
    }
    return count;
  }
}
