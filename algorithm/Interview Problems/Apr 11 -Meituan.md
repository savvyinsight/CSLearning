# Metuan
[Meituan](./Apr%2011%20-Meituan.cpp)

1. selection problem:
   1. 

## 1.Description
input multi group number. for every group, first line input n,x . n is the number of numbers. x is the number of selected box.  second line is numbers
there are n numbers : a1,a2,a3,a4,...,an.
split into boxs : (a1,a2),(a3,a4)...    ; if n odd, the last number as a one group itself.

Now, select x box, then select one number from each box, then add these number together. whether exists sum is odd situation, if yes, print "Yes", else "No".

example:
t -> group number

5 3
1 2 3 4 5
> output:
> "Yes"
5 2
2 2 2 2 2
> "No"
>


Thinking process:
need to choose sum of odd combination.
1. odd + odd ->even
2. odd + even ->odd    ->this is what we wanted
3. even + even ->even


group boxs into three catigory, three queue: a:(odd,odd),b:(odd,even),c:(even,even)

choose x numbers, that sum of is odd number.


### "Feasible range" thinking

The key insight you're missing: Instead of checking every possible combination (which is exponential), we find the range of possible values for t (number of odd-contributing boxes we can select).

Why range works:

    We can choose any number of fixed_odd boxes from 0 to fixed_odd

    We can choose any number of flex boxes from 0 to flex

    So t (odd-contributing boxes chosen) can be ANY integer from 0 to fixed_odd + flex? Not exactly - we also need exactly x total boxes!

Let me derive the range formula step by step:

Let:

    t = number of chosen boxes that can produce odd (from fixed_odd + flex)

    Then x - t boxes must come from fixed_even

Constraints:

    0 ≤ t ≤ fixed_odd + flex (can't pick more odd-capable boxes than exist)

    0 ≤ x - t ≤ fixed_even (can't pick more even boxes than exist)

From constraint 2: x - fixed_even ≤ t ≤ x

Combine with constraint 1:
```
max(0, x - fixed_even) ≤ t ≤ min(x, fixed_odd + flex)
```
This is the critical formula you need to internalize. 



## 2.Description




第二题是构造一个字符串,给了n,k,最长不重复子串数量为k,长为n,看能构造出来不  ababab
第三题就是一个旅游价值问题，给了i对应去哪个城市重复到达价值就+i*次数，样例过了但超时可能要dp吧最后查询o(1)复杂度?
听说美团没hc了，试试水吧。