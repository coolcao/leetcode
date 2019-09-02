## 题目

The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of
rows like this: (you may want to display this pattern in a fixed font for better
legibility)

```
P   A   H   N
A P L S I I G
Y   I   R
```

And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number
of rows:

string convert(string s, int numRows);

#### Example 1:

Input: s = "PAYPALISHIRING", numRows = 3

Output: "PAHNAPLSIIGYIR"

#### Example 2:

Input: s = "PAYPALISHIRING", numRows = 4

Output: "PINALSIGYAHRPI"

Explanation:
```
P     I    N
A   L S  I G
Y A   H R
P     I
```

## 解析
这个题目的关键就是，理解题目中所说的 ZigZag 是个啥？

从题目中的两个例子可以看出，所谓的 ZigZag 字符串就是一个 Z字形的字符串，该题目的要求就是输入一个字符串，然后将该字符串转换成Z字形字符串，然后按行输出该字符串。

所谓Z字形即如下：

![z](https://img001-10042971.cos.ap-shanghai.myqcloud.com/blog/ZigZag_Conversion_-_LeetCode1.png)

我们以第二个例子说明，就是按照上图中的Z字形（叫倒N字形更合适）路径书写原字符串，然后再按照行的顺序输出字符串，即 `PINALSIGYAHRPI`。

题目说明就是这样，那该如何解决呢？

按照题目的逻辑，我们应该先构造Z字形字串。

这怎么构造呢？我们先找一下这里面的规律，我们把每个字符的索引标上：

![z](https://img001-10042971.cos.ap-shanghai.myqcloud.com/blog/ZigZag_Conversion_-_LeetCode2.png)

我们按照行观察，每行的差是多少。记住，此时numRows为 4 ，即分了4行：

```
第1行： 6-0=6, 12-6=6
第2行： 5-1=4, 7-5=2, 11-7=4,13-12=2
第3行： 4-2=2, 8-4=4, 10-8=2
第4行： 9-3=6
```

从上面看，好像没啥规律呀，第一行都是6，第二行 4,2,4,2， 第三行 2,4,2，第四行 6.

可能是数据有点少，我们再增加一点数据，把字符隐藏，只留下标：

![z](https://img001-10042971.cos.ap-shanghai.myqcloud.com/blog/ZigZag_Conversion_-_LeetCode3.png)

```
第一行： 6,6,6
第二行： 4,2,4,2,4,2
第三行： 2,4,2,4,2,4
第四行： 6,6,6
```

好像有点规律了，第一行和最后一行都是6, 第二行是4,2交替，第三行是2,4交替。

我们从对比图和总结的差，可以发现，第一行和最后一行“对称”，中间两行“对称”，这里对称打的引号，是指有点相似。

我们把numRows调到5，再试试看。


![z](https://img001-10042971.cos.ap-shanghai.myqcloud.com/blog/ZigZag_Conversion_-_LeetCode4.png)

```
第一行： 8,8,8,8
第二行： 6,2,6,2,6,2,6,2
第三行： 4,4,4,4,4,4,4,4
第四行： 2,6,2,6,2,6,2,6
第五行： 8,8,8,8
```

这样看来，规律越来越明显了，每一行都是两个数交替，第一行可以看作：8和0，第二行 6,2，第三行：4和4，第四行：2和6，第五行：0和8。而每行的这两个数的和都相同。上面这个例子是8。再上面一个例子是6.

当numRows为 4 时，和为 6
当numRows为 5 时，和为 8

好像是当 numRows为 n 时， 和为 `(n-1)*2`

是的，就是这样，如果你不信，可以再调一下numRows试试看，就是这样的。

找到这个规律，我们再看看每一行的这两个数有啥规律，第一行 8,0，第二行 6,2，第三行：4,4，第四行：2,6，第五行：0，8。

好象是第一个数每行递减2，第二个数每行递增2.

是的，就是这样，我就不画图了，你不信可以自己画图试试，多找两个numRows试试。


总结一下规律，这里每一行的和为 `(n-1)*2`。而且每一行的下标都交替的加固定的两个值。

好了，有了这个规律，我们就可以从这个规律来构造Z字形字符串，然后再按行输出即可。

## 代码
```c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
char* convert(char* s, int numRows) {
  if (numRows == 1) return s;
  int steps[numRows];
  for (int i = 0; i < numRows; i++) {
    steps[i] = (numRows - 1 - i) * 2;
  }
  int length = strlen(s);
  char* result = (char*)malloc(sizeof(char*) * (length + 1));
  int nums = 0;
  for (int i = 0; i < numRows; i++) {
    int idx = i;
    int p1 = steps[i];
    int p2 = steps[numRows - 1 - i];
    int flag = 0;
    while (idx < length) {
      printf("%c,", s[idx]);
      char c = s[idx];
      if (flag % 2 == 0) {
        if (p1 == 0) {
          idx += p2;
        } else {
          idx += p1;
        }
      } else {
        if (p2 == 0) {
          idx += p1;
        } else {
          idx += p2;
        }
      }
      result[nums] = c;
      flag += 1;
      nums += 1;
    }
  }
  result[length] = '\0';
  return result;
}
```

上面是完整的c的代码。

首先构造一个steps，表示每一行的第一个元素的下标。

由于数组的下标是从0开始，所以应用上面规律得到每行的第一个元素的下标应该是： `(numRows - 1 - i)*2`

然后根据规律构造每一行的元素，条件就是，下标要小于整个字符串的长度。

构造时，p1,p2为每行交替加的值。
由于要交替加，因此这里用flag标识奇偶数，用来标识加p1还是p2。

中间有个判断p1,p2等于0时的判断，因为这里第一行和最后一行是个特别的行，需要特别对待。

构造完成后，直接按照顺序将构造的下标对应的字符拼成新的结果字符串result返回即可。
