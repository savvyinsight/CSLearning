#include <iostream>
#include <vector>
using namespace std;

string solve(int n, int x, vector<int>& arr) {
    int fixed_odd = 0;   // boxes that always give odd
    int fixed_even = 0;  // boxes that always give even
    int flex = 0;        // boxes that can give odd or even
    
    // Split into boxes
    for (int i = 0; i < n; ) {
        if (i + 1 < n) {
            // Two-element box
            int a = arr[i], b = arr[i + 1];
            if (a % 2 == 1 && b % 2 == 1) {
                fixed_odd++;      // (odd, odd) -> fixed odd
            } else if (a % 2 == 0 && b % 2 == 0) {
                fixed_even++;     // (even, even) -> fixed even
            } else {
                flex++;           // (odd, even) or (even, odd) -> flexible
            }
            i += 2;
        } else {
            // Single-element box (last element when n is odd)
            if (arr[i] % 2 == 1) {
                fixed_odd++;      // single odd -> fixed odd
            } else {
                fixed_even++;     // single even -> fixed even
            }
            i++;
        }
    }
    
    // Check if we can get an odd sum
    // We need to choose exactly x boxes
    // Let t = number of boxes chosen that can contribute odd (fixed_odd + flex chosen)
    // t must be odd, and we must have enough fixed_even to fill the rest
    
    // t can range from max(0, x - fixed_even) to min(x, fixed_odd + flex)
    int min_t = max(0, x - fixed_even);
    int max_t = min(x, fixed_odd + flex);
    
    // Check if there exists an odd t in [min_t, max_t]
    for (int t = min_t; t <= max_t; t++) {
        if (t % 2 == 1) {
            return "Yes";
        }
    }
    
    return "No";
}

int main() {
    int t;
    cin >> t;
    
    while (t--) {
        int n, x;
        cin >> n >> x;
        vector<int> arr(n);
        for (int i = 0; i < n; i++) {
            cin >> arr[i];
        }
        cout << solve(n, x, arr) << endl;
    }
    
    return 0;
}