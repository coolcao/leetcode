'use strict';

/**
 * You are playing the following Nim Game with your friend: There is a heap of stones on the table, each time one of you take turns to remove 1 to 3 stones. The one who removes the last stone will be the winner. You will take the first turn to remove the stones.

    Both of you are very clever and have optimal strategies for the game. Write a function to determine whether you can win the game given the number of stones in the heap.

    For example, if there are 4 stones in the heap, then you will never win the game: no matter 1, 2, or 3 stones you remove, the last stone will always be removed by your friend.

    Credits:
    Special thanks to @jianchao.li.fighter for adding this problem and creating all test cases.
 */

/**
 * 这个有点智力游戏的意思，和编程关系不是很大。
 * 首先，题目中给了一个4个石头的情况，如果只有4个石头，谁先开始，谁必输。这是题中给出的例子。
 * 因为，4个石头，不管最先开始的拿几个，1个，2个，3个，第二个人肯定能收尾赢。
 * 因此，我们这里，判断能不能赢的关键就在于，我们如果把最后四个留给对方，则必赢。
 * 那么，如何判断我们能不能把最后4个留给对方呢？
 * 因为每次拿，我们都是只能拿1，2，3个，因此，不管对方拿几个，我们只需要凑4个即可。
 * 比如，对方拿3个，我们拿1个，对方拿2个，我们也拿2个。
 * 好了，现在问题来了，由于我们是先手，会不会被对方“反套路”呢？我们拿1个，对方拿3个，我们拿2个对方拿2个，对方按照我们拿的出牌？
 * 当然会，所以我们判断的就是，石子数量是不是4的倍数，如果是4的整数倍，那么，那么很明显，我们会被对方套路进来，最后4个肯定是我们的，必输。
 * 但是，如果不是4的整数倍，还有余数，那么好办了，我们第一次就拿这个余数，后续的拿，根据对方拿的凑4即可。
 * 所以，最后就是判断，石子总数是不是4的倍数即可。如果是4的倍数，我们必输。如果不是，那么我们必赢。
 */

/**
 * @param {number} n
 * @return {boolean}
 */
const canWinNim = function(n) {
    return n % 4 !== 0;
};

