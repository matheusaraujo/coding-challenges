import java.util.*;

public class helpers {
  public static List<List<Integer>> parseInput(final List<String> puzzleInput) {
    final List<Integer> left = new ArrayList<>();
    final List<Integer> right = new ArrayList<>();

    for (final String line : puzzleInput) {
      final String[] parts = line.split("   ");
      left.add(Integer.parseInt(parts[0]));
      right.add(Integer.parseInt(parts[1]));
    }

    Collections.sort(left);
    Collections.sort(right);

    final List<List<Integer>> result = new ArrayList<>();
    result.add(left);
    result.add(right);
    return result;
  }
}
