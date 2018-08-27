import java.io.*;
import java.util.*;

public class Solution {
  public int segregate(int[] arr) {
    int n = arr.length;
    int j = 0;
    for (int i = 0; i < n; i++) {
      if (arr[i] <= 0) {
        int temp = arr[i];
        arr[i] = arr[j];
        arr[j] = temp;
        j++;
      }
    }
    return j;
  }

  public int findMissingPos(int[] arr) {
    int size = arr.length;
    int neg_size = segregate(arr);
    int pos_size = size - neg_size;
    for (int i = neg_size; i < size; i++) {
      int index = Math.abs(arr[i]) - 1 + neg_size;
      if (index < size && arr[index] > 0) {
        arr[index] = -arr[index];
      }
    }

    for (int i = neg_size; i < size; i++) {
      if (arr[i] > 0) {
        return i - neg_size + 1;
      }
    }
    return pos_size + 1;
  }

  public static void main(String[] args) {
    Solution s = new Solution();
    try (BufferedReader reader = new BufferedReader(new FileReader("test_cases.txt"))) {
      String line = reader.readLine();
      while (line != null) {
        String[] contents = line.split("\\s");
        String[] string_arr = contents[0].split(",");
        int[] arr = Arrays.stream(string_arr).mapToInt(Integer::parseInt).toArray();
        int ans = Integer.parseInt(contents[1]);
        if (s.findMissingPos(arr) == ans) {
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
