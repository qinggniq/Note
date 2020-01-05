# 接雨水

## 题面

给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/22/rainwatertrap.png)

上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢 Marcos 贡献此图。

示例:

<pre>输入: [0,1,0,2,1,0,1,3,2,1,2,1]
输出: 6

## 找左边最大值和右边最大值

```c++
class Solution {
public:
    int trap(vector<int>& height) {
        const int n = height.size();
        vector<int> leftMax(n + 1, 0), rightMax(n + 1, 0);
        for (int i = 1; i < n; ++i) {
            leftMax[i] = max(leftMax[i - 1], height[i-1]);
        }
        for (int i = n - 2; i >= 0; --i) {
            rightMax[i] = max(rightMax[i + 1], height[i+1]);
        }
        int ans = 0;
        for (int i = 0; i < n; ++i) {
            ans += max(0, min(leftMax[i], rightMax[i]) - height[i]);
        }
        return ans;
    }
};
```

## 单调栈（单调减）

```c++
class Solution {
public:
    int trap(vector<int>& height) {
        const int n = height.size();
        stack<int> st;
        int ans = 0;
        for (int i = 0; i < n; ++i) {
            while (!st.empty() && height[st.top()] <= height[i]) {
                int topHeight = height[st.top()];st.pop();
                if (st.empty()) break;
                int distance = i - st.top() - 1;
                ans = ans + distance * min(height[st.top()] - topHeight, height[i] - topHeight);
            }
            st.push(i);
        }
        return ans;
    }
};
```

找到离自己最近的比自己高的两个边，然后算横条。

## 双指针

既然根据的是左边最高和右边最低，那么我们先确定高的一边，然后移动低的一边，低的一边的值可以求出，执导低的一边有不再低，换着来一遍。

```c++
class Solution {
public:
    int trap(vector<int>& height) {
        const int n = height.size();
        int l = 0, r = n - 1, leftMax = 0, rightMax = 0, ans = 0;
        while (l < r) {
            if (height[l] < height[r]) {
                if (height[l] > leftMax) {
                    leftMax = height[l];
                }else{
                    ans += leftMax - height[l];
                }
                l++;
            }else{
                if (height[r] > rightMax) {
                    rightMax = height[r];
                }else{
                    ans += rightMax - height[r];
                }
                r--;
            }
        }
        return ans;
    }
};
```

