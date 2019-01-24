/**
Given a string containing digits from 2-9 inclusive, return all possible letter
combinations that the number could represent.

A mapping of digit to letters (just like on the telephone buttons) is given
below. Note that 1 does not map to any letters.



Example:

Input: "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
Note:

Although the above answer is in lexicographical order, your answer could be in
any order you want.

====
Medium
====

*/

/*
这个题目的题意很简单，就是手机九宫格按键的排列组合问题。
手机的九宫格按键，如下所示：

|------|------|------|
|1     |2  abc|3  def|
|------|------|------|
|4  ghi|5  jkl|6  mno|
|------|------|------|
|7 pqrs|8  tuv|9 wxyz|
|------|------|------|

输入一个数字字符串，输出组成的英文字符串的所有可能。
例如：
输入: "23"
输出: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].

题意很简单，但是难度确是中等难度。

先来分析一下这个题，看例子中的，如果只有两个数字，"23"，那么我们可以指定两个数字去标识每个键所处的位置，然后依次切换每个键所代表的字母。

```
for (int i=0;i<3;i++) {
  for (int j=0;j<3;j++) {
    str = str2[i]+str3[j];
    result.push(str);
  }
}
```

如果有三个字符，那么我们就用3个数字去标识。但问题是，我们并不能预先知道要输入的字符串有几个，而且是哪几个数字（不同的数字，表示的字母个数也不一样）。
所以我们通过标识每个数字键的位置，并循环遍历所有可能性的方式是行不通的。

我们换一种思路去标识每一个数字键的索引位置，看看能不能找到规律：

首先，根据输入的数字键，我们能够很快的得到所有组合的可能性的个数，例如，输入"23"那么就有 3*3=9 种可能。

我们需要两个数字标识，那么可以将这两个数字放到数组中： [0,0] 表示取2键的第0个字符a，3键的第0个字符d，最后组合成就是 "ad"。
再下一个组合的话，应该是 [1, 0] 表示取2键的第1个字符b，3键的第0个字符d，组合起来就是 "bd"。
依次便是 [2, 0], [0, 1], [1, 1], [2, 1], [0, 2], [1, 2], [2, 2]...

看出规律没，是不是感觉非常像我们在计算机基础课程中学到的二进制的算法。不过这里可能不是一个单纯的3进制或4进制，因为不同的键有不同的字符表示个数。

但是，其逻辑原理确是一致的，我们将每个数字键代表的字符，以及其字符个数都表示出来，然后依次按照进制计算的逻辑进行往上累加计算即可。

我们抽象出一个结构体来表示所需用到的数据。
其中，arr数组用于标识每个数字键当前所遍历到的字符位置。
arrLen数组用于表示每个数字键上字符的个数。
digitIdx数组用于标识每个数字字符和其代表的字符串的索引映射。（因为这里我抽象数字和字符串映射时是从0开始，而手机键盘上是2数字键才开始有字符表示）。

好了，现在的问题就是，如何实现一个加法运算，实现类似的进制转换。

代码如下，直接看代码吧。

*/




#include <stdio.h>
#include <string.h>
#include <stdlib.h>
typedef struct cus *CustomStruct;
struct cus {
  int* arr;
  int* arrLen;
  int* digitIdx;
  int arrSize;
  int totalSize;
};

void printStruct(CustomStruct cs) {
  printf("-------------\n");
  printf("arr:");
  for (int i=0;i<cs->arrSize;i++) {
    printf("%d", cs->arr[i]);
  }
  printf("\n");
  printf("arrLen:");
  for (int i=0;i<cs->arrSize;i++) {
    printf("%d", cs->arrLen[i]);
  }
  printf("\n");
  printf("digitIdx:");
  for (int i=0;i<cs->arrSize;i++) {
    printf("%d", cs->digitIdx[i]);
  }
  printf("\n");
  printf("totalSize:%d\n", cs->totalSize);
  printf("===============\n");
}
void add(CustomStruct cs) {
  int* arr = cs->arr;
  int* arrLen = cs->arrLen;
  int i=0;
  // printf("%d %d %d %d\n", arr[0], arr[1], arr[2], arr[3]);
  arr[i] += 1;
  while(arr[i] >= arrLen[i] && i<cs->arrSize){
    arr[i] = 0;
    i += 1;
    arr[i] += 1;
  }
  // printf("%d %d %d %d\n", arr[0], arr[1], arr[2], arr[3]);
  // printf("================\n");
}

char** letterCombinations(char* digits, int* returnSize) {
  char* digitsMap[8] = {"abc", "def",  "ghi", "jkl",
                        "mno", "pqrs", "tuv", "wxyz"};
  int lengths[8] = {3, 3, 3, 3, 3, 4, 3, 4};

  int length = strlen(digits);  // 数字的数量
  if (length == 0) return NULL;

  CustomStruct cs = malloc(sizeof *cs);
  cs->arrSize = length;
  
  int tmpArr[length];
  int tmpArrLen[length];
  int tmpDigitIdx[length];

  char** result = NULL;
  int size = 1;     // 记录总的结果数
  for (int i = 0; i < length; i++) {
    int idx = digits[i] - 48 - 2;
    int letterSize = lengths[idx];
    size *= letterSize;
    tmpArr[i] = 0;
    tmpArrLen[i] = letterSize;
    tmpDigitIdx[i] = idx;
  }

  *returnSize = size;

  cs->totalSize = size;
  cs->arr = tmpArr;
  cs->arrLen = tmpArrLen;
  cs->digitIdx = tmpDigitIdx;

  // printStruct(cs);
  char** returnResult = (char**)malloc((sizeof(char**))*size);

  for (int i=0;i<size;i++) {
    char* tmpStr = (char*)malloc(sizeof(char*)*(length+1));
    for (int j=0;j<length;j++) {
      char* currentDigits = digitsMap[cs->digitIdx[j]];
      // printf("%c", currentDigits[cs->arr[j]]);
      tmpStr[j] = currentDigits[cs->arr[j]];
      if (j == length-1) {
        tmpStr[j + 1] = '\0';
      }
    }
    returnResult[i] = tmpStr;
    add(cs);
  }
  
  return returnResult;
}

int main(int argc, char const* argv[]) {
  char* str = "69784";
  int returnSize = 0;
  char** result = letterCombinations(str, &returnSize);
  // printf("%d\n", returnSize);

  for (int i=0;i<returnSize;i++) {
    printf("%d: %s\n", i+1, result[i]);
  }

  return 0;
}
