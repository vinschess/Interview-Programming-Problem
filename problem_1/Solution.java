import java.io.*;
import java.util.*;

public class Solution {

  public List<Pair> getPairs(int[] arr, int sum_val) {
    HashSet<Integer> set = new HashSet<Integer>();
    List<Pair> ans = new ArrayList<Pair>();
    for(int val : arr) {
      int diff = sum_val - val;
      if (set.contains(diff)) {
        ans.add(new Pair(diff, val));
      }
      set.add(val);
    }
    return ans;
  }

  public static void main(String[] args) {
    try (BufferedReader reader = new BufferedReader(new FileReader("test_cases.txt"))) {
      String line = reader.readLine();
      while (line != null) {
        String[] contents = line.split("\\s");
        String[] string_arr = contents[0].split(",");
        int[] arr = Arrays.stream(string_arr).mapToInt(Integer::parseInt).toArray();
        int k = Integer.parseInt(contents[1]);
        List<Pair> pairs = new Solution().getPairs(arr, k);
        if (!pairs.isEmpty()){
          System.out.printf("Following are the pairs which sum is equal to %d: %s.\n", k, pairs);
        } else {
          System.out.printf("No pair is present which sum is euqal to %d.\n", k);
        }
        line = reader.readLine();
      }

    } catch (IOException ex) {
      System.out.println(ex);
    }

  }
}

class Pair {
  int a, b;
  public Pair(int a, int b) {
    this.a = a;
    this.b = b;
  }

  @Override
  public String toString() {
    return String.format("(%d, %d)", this.a, this.b);
  }
}
