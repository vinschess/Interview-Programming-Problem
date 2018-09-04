import java.io.*;
import java.util.*;

public class Solution {
  public boolean isValid(int[] count_arr, int k) {
    int val = 0;
    for (int count : count_arr) {
      if (count > 0)
        val += 1;
    }
    return (k >= val);
  }

  public boolean validateString(String str, int k) {
    int[] count_arr = new int[26];
    int unique_char = 0;
    for (char c : str.toCharArray()) {
      if (count_arr[c - 'a'] > 0)
        unique_char += 1;
      count_arr[c - 'a'] += 1;
    }
    return (unique_char >= k);

  }

  public int longestSubstrKUniq(String str, int k) {
    if (!validateString(str, k)) {
      System.out.println("Failure here");
      return -1;
    }

    char[] s = str.toCharArray();

    int curr_start = 0;
    int curr_end = 0;

    int max_window_size = 1;
    int max_window_start = 0;

    int[] count_arr = new int[26];
    count_arr[s[0] - 'a'] = 1;

    for (int i = 1; i < str.length(); i++) {
      count_arr[s[i] - 'a'] += 1;
      curr_end += 1;

      while (!isValid(count_arr, k)) {
        count_arr[s[curr_start] - 'a'] -= 1;
        curr_start += 1;
      }

      if (curr_end - curr_start + 1 > max_window_size) {
        max_window_size = curr_end - curr_start + 1;
        max_window_start = curr_start;
      }
    }
    // System.out.println("Max substring: " + str.substring(max_window_start, max_window_size + max_window_start));

    return max_window_size;
  }

  public static void main(String[] args) {
    Solution s = new Solution();
    try(BufferedReader reader = new BufferedReader(new FileReader("test_cases.txt"))) {
      String line;
      while ((line =reader.readLine()) != null) {
        String[] contents = line.split("\\s");
        String str = contents[0];
        int k = Integer.parseInt(contents[1]);
        int ans = Integer.parseInt(contents[2]);

        if (s.longestSubstrKUniq(str, k) == ans) {
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
