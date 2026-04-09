#include <climits>
#include <vector>
#include <algorithm>
using namespace std;

// 1.Check a value in or not in a binary 
struct node{
  int val;
  node* left;
  node* right;
};

int in(node* root, int target){
    while(root){
        if(root->val == target) return 1;
        else if(root->val < target) root = root->right;
        else root = root->left;
    }
    return 0;
}

// 2. Find the maximum sum of a contiguous subarray 
// that does not exceed a given limit.
int maxSum(vector<int>& nums, int limit) {
    int max_sum = 0;
    int current_sum = 0;
    int left = 0;

    for (int right = 0; right < nums.size(); right++) {
        current_sum += nums[right];

        while (current_sum > limit) {
            current_sum -= nums[left];
            left++;
        }

        max_sum = max(max_sum, current_sum);
    }

    return max_sum;
}

unsigned f(unsigned n,unsigned b, unsigned p[]){
    unsigned rb = 0,maxrb = 0;
    for(int i=0,j=0;i<n;i++){
        rb += p[i];
        while(rb > b && j <= i){
            rb -= p[j++];
        }
        if(rb > maxrb) maxrb = rb;
    }
    return maxrb;
}


// 3. Calculate the parity bit of a number (odd or even).
int parity(unsigned n) {
    int result = 0;
    while (n) {
        result ^= (n & 1);
        n >>= 1;
    }
    return result;
}
// this problem is like first calculate the number of 1 bit.
// second is the number of 1 bit is odd or even, we can use this. OR(1^1)->0 ,means even

// 4. Find the maximum value in a binary tree.
struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

int findMax(TreeNode* root) {
    if (!root) return INT_MIN; // Return minimum integer if the tree is empty
    int leftMax = findMax(root->left);
    int rightMax = findMax(root->right);
    return max(root->val, max(leftMax, rightMax));
}

// 5.Question:
/*
You are given the following C++ class that maintains a simple in-memory counter 
and allows incrementing and retrieving the current value. The class is not thread-safe.*/
#include <chrono>
#include <thread>

class Counter {
private:
    int count;

public:
    Counter() : count(0) {}

    void increment() {
        // Simulate non-atomic operation
        int current = count;
        std::this_thread::sleep_for(std::chrono::microseconds(100));
        count = current + 1;
    }

    int get_value() const {
        return count;
    }
};
// Your task:
// Modify this class to be thread-safe using appropriate synchronization mechanisms in C++.
// You may use any C++ standard library features (C++11 or later). 
// Multiple threads will call increment() and get_value() concurrently.

// Answer:
#include <chrono>
#include <thread>
#include <mutex>

class Counter_solved {
private:
    int count;
    mutable std::mutex mtx; // Mutex to protect access to count

public:
    Counter_solved() : count(0) {}

    void increment() {
        std::lock_guard<std::mutex> lock(mtx); // Lock the mutex for the duration of this function
        int current = count;
        std::this_thread::sleep_for(std::chrono::microseconds(100));
        count = current + 1;
    }

    int get_value() const {
        std::lock_guard<std::mutex> lock(mtx); // Lock the mutex to safely read the value
        return count;
    }
};

// 6.Multi-threaded Account Transfer with Thread Safety  
// ---> see multi-threadAccountTransfer.cpp


// 7. Hand write priority_queue(Construct,Destruct,pop,push,top) in C++.
// --->see handWritePriorityQueue.cpp
