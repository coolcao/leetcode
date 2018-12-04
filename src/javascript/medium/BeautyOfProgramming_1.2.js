'use strict';

/**
 * 中国象棋问题
 * 假设在棋盘上只有“将”和“帅”二子，为叙述方便，我们用A表示将，B表示帅。
 * A,B二子被限制在己方的3x3的格子里运动。图：http://7xt3oh.com2.z0.glb.clouddn.com/blog/20131024121938859.png
 * A被正方形{d10,f10,d8,f8}包围，二B被{d3,f3,d1,f1}包围。
 * 每一步，A，B分别可以横向或纵向移动一格，但不能沿对角线移动。另外，A不能面对B，也就是说，A和B不能处于同一纵向直线上。
 * 请写出一个程序，输出A,B所有合法的位置。要求在代码中只能使用一个字节存储变量。
 * 位置简化图：http://7xt3oh.com2.z0.glb.clouddn.com/blog/20131024122735703.png
 */

const fn1 = function() {
    const LMASK = 240;  //左掩码，对应的二进制'11110000'
    const RMASK = 15;   //右掩码，对应的二进制'00001111'
    const GRIDW = 3;    //格子宽度

    let p = 0;          //用于存储左右两个值的变量

    //设置右侧4位值
    const setRight = function(p, n) {
        return p & LMASK | n;
    }

    const getRight = function(p) {
        return p & RMASK;
    }

    //设置左侧4位值
    const setLeft = function(p, n) {
        return (p & RMASK) | (n << 4);
    }
    const getLeft = function(p) {
        return (p & LMASK) >> 4;
    }



    for (p = setLeft(p, 1); getLeft(p) <= GRIDW * GRIDW; p = setLeft(p, getLeft(p) + 1)) {
        for (p = setRight(p, 1); getRight(p) <= GRIDW * GRIDW; p = setRight(p, getRight(p) + 1)) {
            if (getLeft(p) % GRIDW != getRight(p) % GRIDW) {
                console.log(`A = ${getLeft(p)},B = ${getRight(p)}`);
            }
        }
    }
}

const fn2 = function() {
    let a = 81;
    while (a-- > 0) {
        if(((a/9)>>>0) % 3 != a%9%3){
            console.log(`${((a/9)>>>0) + 1} ,${a%9 + 1} `);
        }
    }
}

