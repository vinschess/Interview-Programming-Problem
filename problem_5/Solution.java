import java.io.*;
import java.util.*;

public class Solution {
  public int decodeCount(String str_digit) {
    int count_2nd_last = 1;
    int count_last = 1;
    int count_curr = 1;
    int n = str_digit.length();
    for (int i = 2; i <= n; i++) {
      count_curr = 0;

      if (str_digit.charAt(i-1) > '0') {
        count_curr = count_last;
      }

      if (str_digit.charAt(i-2) == '1' || (str_digit.charAt(i-2) == '2' && str_digit.charAt(i-1) < '7')) {
        count_curr += count_2nd_last;
      }

      count_2nd_last = count_last;
      count_last = count_curr;
    }

    return count_curr;

  }

  public static void main(String[] args) {
    Solution s = new Solution();
    try (BufferedReader reader = new BufferedReader(new FileReader("test_cases.txt"))) {
      String line = reader.readLine();
      while (line != null) {
        String[] contents = line.split("\\s");
        String str_digit = contents[0];
        int ans = Integer.parseInt(contents[1]);
        if (s.decodeCount(str_digit) == ans) {
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
