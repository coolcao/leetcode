### 题目说明
> Given an array of integers, return indices of the two numbers such that they add up to a specific target.
>
>  You may assume that each input would have exactly one solution, and you may not use the same element twice.
>
>  Example:
>
>  Given nums = [2, 7, 11, 15], target = 9,
>
>  Because nums[0] + nums[1] = 2 + 7 = 9,
>
>  return [0, 1].
>
> ----
> 给定一个整数数组和一个指定的target值，返回数组中两个元素的和为target值的两个元素的索引。
>
> 假设任何一个输入都只有一个解，相同的元素不能使用两次。
>
> 例如
>
> 给定一个数组 `nums=[2,7,11,15],target=9`
>
> 因为`nums[0] + nums[1] = 2 + 7 = 9`
>
> 所以，返回 `[0,1]`

### 难易程度：简单

这个题目一眼上来，可能最简单的就是，遍历两次，两次的和为target即找到了相应的index值。

但是题目中，明确指定了，每个元素只能使用一次。

对于这个指定，可能有点小误解，就是怎么算使用一次？其实这里的意思是，每个元素只能遍历一遍。

那么，好了，如果只能遍历一遍，而又需要求两个数的和，那只能将遍历的值缓存起来，然后从缓存中取值进行计算即可。

### 代码
```javascript
let twoSum = function(nums, target) {
    let result = [];
    let map = new Map();
    for (let i = 0; i < nums.length; i++) {
        let value = nums[i];
        let an = target - value;
        let diff = map.get(value);
        if(diff){
            if(diff.diff === value){
                result.push(diff.index,i);
                break;
            }
        }else{
            map.set(an,{index:i,diff:an});
        }
    }
    return result;
}
```

这里，我遍历到每个元素时，先检查缓存中有没有其值，如果有并且值相等，则直接将缓存值的index和当前index返回。如果没有，则将该值的差值缓存起来。
