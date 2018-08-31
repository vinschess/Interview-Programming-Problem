import java.io.*;
import java.util.*;

public class Solution {
  public int largestNonAdjSum (int[] arr) {
    int first = 0;
    int next = 0;
    for (int i : arr) {
      int temp = Math.max(first, next);
      first = next + i;
      next = temp;
    }
    return Math.max(first, next);
  }

  public static void main(String[] args) {
    Solution s = new Solution();
    try(BufferedReader reader = new BufferedReader(new FileReader("test_cases.txt"))) {
      String line = reader.readLine();
      while (line != null) {
        String[] content = line.split("\\s");
        String[] string_arr = content[0].split(",");
        int[] arr = Arrays.stream(string_arr).mapToInt(Integer::parseInt).toArray();
        int ans = Integer.parseInt(content[1]);
        if (s.largestNonAdjSum(arr) == ans) {
          System.out.println("Passed");
        } else {
          System.out.println("Failed");
        }
        line = reader.readLine();
      }
    } catch (IOException ex) {
      System.out.println(ex);
    }
  }
}
