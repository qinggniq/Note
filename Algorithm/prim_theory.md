# 主定理
## 支配理论
假设有递回关系式
$$ T(n) = a \; T\!\left(\frac{n}{b}\right) + f(n) a \geq 1 , b > 1 $$
其中，$n$为问题规模，$a$为递回的子问题数量，$n/b$为每个子问题的规模（假设每个子问题的规模基本一样），$f(n)$为递回以外进行的计算工作。

## 情形一
如果存在常数$\epsilon > 0$，有
$$f(n) = O\left( n^{\log_b (a) - \epsilon} \right)（多项式地小于）$$  
$$T(n) = \Theta\left( n^{\log_b a} \right)$$

## 情形二
如果存在常数k ≥0，有
$$  f(n) = \Theta\left( n^{\log_b a} \log^{k} n \right)$$
则
$$ T(n) = \Theta\left( n^{\log_b a} \log^{k+1} n \right) $$
## 情形三
如果存在常数
$$ \epsilon > 0; $$ 
$$f(n) = \Omega\left( n^{\log_b (a) + \epsilon} \right)（多项式地大于）$$

同时存在常数 $c < 1$ 以及充分大的$n$，满足

$$a f\left( \frac{n}{b} \right) \le c f(n)$$
则
$$T\left(n \right) = \Theta \left(f \left(n \right) \right)$$
