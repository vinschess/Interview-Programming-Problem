public class Solution {
  public float sqrt(float val) {
    float sqrt_val = val / 2;
    float prev = 0.0f;
    while(sqrt_val-prev >= 0.001 || prev-sqrt_val >= 0.001) {
      prev = sqrt_val;
      sqrt_val -= (sqrt_val*sqrt_val - val) / (2*sqrt_val);
    }
    return Math.round(sqrt_val * 100) / 100f;
  }

  public static void main(String[] args) {
    Solution s = new Solution();
    System.out.println(s.sqrt(1));
    System.out.println(s.sqrt(16));
    System.out.println(s.sqrt(25));
    System.out.println(s.sqrt(100));
    System.out.println(s.sqrt(101));
    System.out.println(s.sqrt(176));
  }
}
