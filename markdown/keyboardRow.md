```
/**
 * Given a List of words, return the words that can be typed using letters of alphabet on only one row's of American keyboard like the image below.
 * Example 1:
 * Input: ["Hello", "Alaska", "Dad", "Peace"]
 * Output: ["Alaska", "Dad"]
 * Note:
 * You may use one character in the keyboard more than once.
 * You may assume the input string will only contain letters of alphabet.
 */
```

给定一个单词列表，返回能够使用alphabet键盘在一行打印出来的单词。alphabet键盘如下图：
![键盘](http://7xt3oh.com2.z0.glb.clouddn.com/blog/keyboard.png)

例如：
输入`["Hello","alaska","Dad","Peace"]`，输出：`["Alaska","Dad"]`。

注意：
你可以多次使用键盘中的字符。假设输出的字符串将只包含`alphabet`键盘字符。

### 解析

这个问题难易程度为简单，我们可以遍历每个单词中的所有字符，如果所有字符在一行上，那么说明这个单词满足条件，输出，否则不满足条件。

那，如何判断一个单词是否只由一行键打印呢？我们可以将所有字符按照键盘位置，编制一个成一个字符数组，总共26个字符，分3段，第一段0-10为第一行，第二段10-19为第二行，第三段19-26为第三行。我们只需判断一个单词中所有字符均在其中的一段即可。

### 代码

```
const findWords = function(words) {
    //取下不取上
    //0-10
    //10-19
    //19-25
    let chars = ['q','w','e','r','t','y','u','i','o','p','a','s','d','f','g','h','j','k','l','z','x','c','v','b','n','m'];
    let map = new Map();
    for(let i=0;i<chars.length;i++){
        map.set(chars[i],i);
    }

    let isOneRow = function (word) {
        word = word.toLowerCase();
        let start = 0;
        let end = 26;
        let length = word.length;
        let index = map.get(word[0]);
        if(index>=0 && index <10){
            start = 0;
            end = 10;
        }else if(index >=10 && index < 19){
            start = 10;
            end = 19;
        }else if(index >= 19 && index < 26){
            start = 19;
            end = 26;
        }
        for(let i=1;i<length;i++){
            index = map.get(word[i]);
            if(!(index >= start && index < end)){
                return false;
            }
        }
        return true;
    }

    return words.filter(isOneRow);

};
```