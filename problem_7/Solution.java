import java.io.*;
import java.util.*;

public class Solution {
  // Where m represents total ways of steps at a time
  public int stepsCount(int n, int m) {
    int[] arr = new int[n+1];
    arr[0] = arr[1] = 1;
    for (int i = 2; i <= n; i++) {
      arr[i] = 0;
      int j = 1;
      while (j<=m && j<=i) {
        arr[i] += arr[i-j];
        j++;
      }
    }
    return arr[n];
  }

  public static void main(String[] args) {
    Solution s = new Solution();
    try(BufferedReader reader = new BufferedReader(new FileReader("test_cases.txt"))) {
      String line;
      while ((line = reader.readLine()) != null) {
        String[] contents = line.split("\\s");
        int steps = Integer.parseInt(contents[0]);
        int ans = Integer.parseInt(contents[1]);
        if (s.stepsCount(steps, 2) == ans) {
          System.out.println("Passed");
        } else {
          System.out.println("Failed");
        }
      }
    } catch (IOException ex) {
      System.out.println(ex);
    }
  }
}
