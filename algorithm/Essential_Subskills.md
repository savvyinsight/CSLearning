# Subskills, template or pattern

# Linked List

## 🔧 **Essential Linked List Subskills**

### 1. **Fast-Slow Pointer (Tortoise & Hare)**
```cpp
//Find middle
ListNode* slow = head, *fast = head;
while(fast && fast->next) {
    slow = slow->next;
    fast = fast->next->next;
}
//slow is in the middle node
```
**Use cases**: Find middle, detect cycles, find cycle entry point

### 2. **Linked List Reversal**
```cpp
ListNode* reverseList(ListNode* head) {
    ListNode* prev = nullptr;
    ListNode* curr = head;
    while(curr) {
        ListNode* nextTemp = curr->next;
        curr->next = prev;
        prev = curr;
        curr = nextTemp;
    }
    return prev;
}
```

### 3. **Dummy Head Pattern**
```cpp
ListNode* dummy = new ListNode(0);
dummy->next = head;
ListNode* curr = dummy;
// ... operations
return dummy->next;
```
**Use cases**: When head might change, deletion operations

## 🎯 **Common Patterns & Templates**

### 4. **Cycle Detection**
```cpp
bool hasCycle(ListNode* head) {
    ListNode* slow = head, *fast = head;
    while(fast && fast->next) {
        slow = slow->next;
        fast = fast->next->next;
        if(slow == fast) return true;
    }
    return false;
}
```

### 5. **Merge Two Sorted Lists**
```cpp
ListNode* mergeTwoLists(ListNode* l1, ListNode* l2) {
    ListNode dummy(0);
    ListNode* tail = &dummy;
    
    while(l1 && l2) {
        if(l1->val <= l2->val) {
            tail->next = l1;
            l1 = l1->next;
        } else {
            tail->next = l2;
            l2 = l2->next;
        }
        tail = tail->next;
    }
    tail->next = l1 ? l1 : l2;
    return dummy.next;
}
```

### 6. **Two-Pass Techniques**
```cpp
// First pass: get length
// Second pass: do operation at specific position
```

### 7. **K-pointer Technique**
```cpp
ListNode* p1 = head, *p2 = head;
// Move p2 k steps ahead, then move both until p2 reaches end
```

## 🚀 **Advanced Patterns**

### 8. **In-place Reordering**
```cpp
// Example: L0 → L1 → L2 → ... → Ln-1 → Ln 
// Reorder to: L0 → Ln → L1 → Ln-1 → L2 → Ln-2 → ...
// Pattern: Find middle + Reverse second half + Merge alternately
```

### 9. **Node Deletion without Head Reference**
```cpp
void deleteNode(ListNode* node) {
    node->val = node->next->val;
    node->next = node->next->next;
}
```

### 10. **Add Two Numbers**
```cpp
// Carry-based addition with two lists
```

## 📋 **Problem-Specific Templates**

### **For "K-group Reversal":**
1. Count length
2. Process in groups of K
3. Reverse each group
4. Connect groups properly

### **For "Remove Nth from End":**
1. Move fast pointer N steps ahead
2. Move both until fast reaches end
3. Slow will be at node before target

### **For "Intersection of Two Lists":**
1. Get lengths of both lists
2. Move longer list's pointer forward by difference
3. Move both until they meet

## 💡 **Mastery Tips:**

1. **Practice these patterns individually** until they become muscle memory
2. **Combine patterns** - most medium/hard problems use 2+ of these subskills
3. **Draw diagrams** - essential for pointer manipulation
4. **Test edge cases**: empty list, single node, two nodes, even/odd lengths

## 🎯 **Recommended Practice Order:**

1. Reverse Linked List
2. Merge Two Sorted Lists  
3. Detect Cycle
4. Remove Nth From End
5. Palindrome (which you just mastered!)
6. Add Two Numbers
7. Reorder List
8. Reverse Nodes in k-Group



# Dynamic Programming

Dive into Dynamic Programming patterns and subskills. DP is all about recognizing patterns and building solutions from subproblems.

## 🧠 **Core DP Mindset**
**Key Principle**: Break down complex problems into overlapping subproblems and store results to avoid recomputation.

## 🔧 **Essential DP Subskills**

### 1. **State Definition**
```cpp
// How to define your DP state
dp[i] = best solution for first i elements
dp[i][j] = best solution for subproblem from i to j
dp[i][j] = best solution with first i items and capacity j
```

### 2. **Base Case Identification**
```cpp
// Smallest subproblems that we know the answer to
dp[0] = 0;
dp[1] = 1;
dp[i][0] = 1;  // empty subset
```

### 3. **Recurrence Relation**
```cpp
// How to build larger solutions from smaller ones
dp[i] = max(dp[i-1], dp[i-2] + nums[i]);
dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]];
```

## 🎯 **Major DP Patterns & Templates**

### **Pattern 1: Fibonacci-style (1D Linear)**
```cpp
int dp[n+1];
dp[0] = 0; dp[1] = 1;
for(int i = 2; i <= n; i++) {
    dp[i] = dp[i-1] + dp[i-2];
}
```
**Problems**: Climbing Stairs, Fibonacci, House Robber

### **Pattern 2: 0/1 Knapsack**
```cpp
vector<vector<int>> dp(n+1, vector<int>(capacity+1, 0));
for(int i = 1; i <= n; i++) {
    for(int w = 0; w <= capacity; w++) {
        if(weight[i-1] <= w) {
            dp[i][w] = max(dp[i-1][w], 
                          dp[i-1][w - weight[i-1]] + value[i-1]);
        } else {
            dp[i][w] = dp[i-1][w];
        }
    }
}
```
**Problems**: Subset Sum, Partition Equal Subset Sum, Target Sum

### **Pattern 3: Unbounded Knapsack**
```cpp
vector<int> dp(amount+1, 0);
dp[0] = 1;
for(int coin : coins) {
    for(int i = coin; i <= amount; i++) {
        dp[i] += dp[i - coin];
    }
}
```
**Problems**: Coin Change, Coin Change II, Perfect Squares

### **Pattern 4: Longest Common Subsequence (LCS)**
```cpp
vector<vector<int>> dp(m+1, vector<int>(n+1, 0));
for(int i = 1; i <= m; i++) {
    for(int j = 1; j <= n; j++) {
        if(text1[i-1] == text2[j-1]) {
            dp[i][j] = dp[i-1][j-1] + 1;
        } else {
            dp[i][j] = max(dp[i-1][j], dp[i][j-1]);
        }
    }
}
```
**Problems**: LCS, Edit Distance, Delete Operation for Two Strings

### **Pattern 5: Matrix Chain Multiplication**
```cpp
vector<vector<int>> dp(n, vector<int>(n, 0));
for(int len = 2; len < n; len++) {
    for(int i = 1; i < n-len+1; i++) {
        int j = i+len-1;
        dp[i][j] = INT_MAX;
        for(int k = i; k < j; k++) {
            dp[i][j] = min(dp[i][j], 
                          dp[i][k] + dp[k+1][j] + 
                          matrix[i-1]*matrix[k]*matrix[j]);
        }
    }
}
```
**Problems**: Burst Balloons, Minimum Cost Tree From Leaf Values

### **Pattern 6: Interval DP**
```cpp
vector<vector<int>> dp(n, vector<int>(n, 0));
for(int len = 1; len <= n; len++) {
    for(int i = 0; i <= n-len; i++) {
        int j = i + len - 1;
        // Fill dp[i][j] based on smaller intervals
    }
}
```
**Problems**: Palindrome Partitioning, Stone Game, Guess Number Higher/Lower

## 🚀 **Advanced DP Patterns**

### **Pattern 7: State Machine DP**
```cpp
vector<vector<int>> dp(n, vector<int>(2, 0));
dp[0][0] = 0;           // not holding stock
dp[0][1] = -prices[0];  // holding stock

for(int i = 1; i < n; i++) {
    dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i]);
    dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i]);
}
```
**Problems**: Best Time to Buy/Sell Stock, Maximum Subarray

### **Pattern 8: Digit DP**
```cpp
// Count numbers with certain properties in range [L,R]
int dp[pos][tight][state] = count of numbers with:
- pos: current digit position
- tight: whether we're tight to upper bound
- state: additional state (sum, product, etc.)
```

### **Pattern 9: Bitmask DP**
```cpp
vector<int> dp(1 << n, INT_MAX);
dp[0] = 0;
for(int mask = 0; mask < (1 << n); mask++) {
    for(int i = 0; i < n; i++) {
        if(!(mask & (1 << i))) {
            dp[mask | (1 << i)] = min(dp[mask | (1 << i)], 
                                     dp[mask] + cost[i]);
        }
    }
}
```
**Problems**: Traveling Salesman, Assigning Tasks

## 📋 **DP Problem-Solving Framework**

### **Step-by-Step Template:**
```cpp
bool canSolveWithDP(problem) {
    // 1. Check for optimal substructure
    // 2. Check for overlapping subproblems
    // 3. Can we define states and transitions?
}

vector<int> dp(n);
// 1. Define state meaning
// 2. Initialize base cases  
// 3. Define recurrence relation
// 4. Determine iteration order
// 5. Return final answer
```

## 🔍 **How to Recognize DP Problems**

**Look for these clues:**
- "Maximum", "Minimum", "Longest", "Shortest"
- "Count the number of ways"
- "Is it possible to..."
- Decisions that affect future choices
- Problems that can be broken into similar smaller problems

## 💡 **Mastery Tips:**

1. **Start with 1D DP** before moving to 2D
2. **Draw the state transition diagram** for complex problems
3. **Practice pattern recognition** - most DP problems fit known patterns
4. **Learn space optimization** (convert 2D→1D when possible)
5. **Memorize the core 5-6 patterns** - they cover 80% of problems

## 🎯 **Recommended Learning Order:**

1. Fibonacci/Climbing Stairs
2. 0/1 Knapsack
3. Longest Common Subsequence
4. Coin Change
5. House Robber
6. Unique Paths
7. Edit Distance
8. Matrix Chain Multiplication

Would you like me to dive deeper into any specific DP pattern or show example problems for each category?
## Essential DP Patterns to Learn

1. **Subset Sum** - what we're learning now
   
2. **Knapsack** - similar but with values and weights
   
3. **Longest Increasing Subsequence**
   
4. **Coin Change** - counting ways vs minimum coins
   
5. **Matrix Chain Multiplication**
   
6. **Edit Distance**
   

## Practice Progression

Start with these LC problems in order:

1. **LC 416** - Partition Equal Subset Sum (current problem)
   
2. **LC 494** - Target Sum (subset sum variation)
   
3. **LC 518** - Coin Change II (counting subsets)
   
4. **LC 1049** - Last Stone Weight II (another partition problem)