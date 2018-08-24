import java.io.*;
import java.util.*;

public class Solution {
  public int[] productArray(int[] arr) {
    int n = arr.length;
    int[] prod_arr = new int[n];
    Arrays.fill(prod_arr, 1);

    int temp = 1;
    for(int i = 0; i < n; i++) {
      prod_arr[i] = temp;
      temp *= arr[i];
    }

    temp = 1;
    for(int i = n - 1; i >= 0; i--) {
      prod_arr[i] *= temp;
      temp *= arr[i];
    }
    return prod_arr;
  }

  public static void main(String[] args) {
    Solution s = new Solution();
    try (BufferedReader reader = new BufferedReader(new FileReader("test_cases.txt"))) {
      String line = reader.readLine();
      while (line != null) {
        String[] string_arr = line.split(",");
        int[] arr = Arrays.stream(string_arr).mapToInt(Integer::parseInt).toArray();
        System.out.printf("Product array: %s\n", Arrays.toString(s.productArray(arr)));
        line = reader.readLine();
      }
    } catch (IOException ex) {
      System.out.println(ex);
    }
  }
}
