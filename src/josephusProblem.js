'use strict';

/**
 * 约瑟夫问题：
 *     在犹太罗马战争期间，他们41名犹太反抗者困在了罗马人包围的洞穴中。
 *     这些反抗者宁愿自杀，也不愿被活捉，于是决定围成一个圆圈，并沿着圆圈每隔两个人杀死一个人，直到剩下最后一个人为止。
 * 输入：人数
 * 输出：最后存活者的编号
 */


class Node {
    constructor(element) {
        this.element = element;
        this.next = null;
    }
    toString() {
        return this.element;
    }
    getData() {
        return this.element;
    }
}

const start = Symbol.for('start');
const end = Symbol.for('end');
const index = Symbol.for('index');

class LinkedList {
    constructor(count, _start = 1) {
        this[start] = null;
        this[end] = null;
        this[index] = null;
        if (count === 0) {
            return null;
        } else {
            this[start] = this[end] = this[index] = new Node(1);
            for (let i = 2; i <= count; i++) {
                const node = new Node(i);
                if (i === _start) {
                    this[index] = node;
                }
                let current = this[end];
                current.next = node;
                this[end] = node;
                if (i === count) {
                    this[end].next = this[start];
                }
            }
        }

    }
    sha() {
        while (this[index].next) {
            if (this[index].next.getData() === this[index].getData()) {
                break;
            } else {
                let current = this[index];
                let todelete = this[index].next;
                let deletenext = todelete.next;
                current.next = deletenext;
                this[index] = current.next;
            }
        }
    }
    toString() {
        let current = this[index];
        let result = [];
        while (current) {
            result.push(current.getData());
            if (current === current.next) {
                break;
            }
            current = current.next;
        }
        return result.join();
    }
}

// const ll = new LinkedList(10);
// ll.sha();
// console.log(ll.toString());


const sha = (n) => {
    const array = [];
    for (let i = 1; i <= n; i++) {
        array.push(i);
    }

    let index = 1;
    while (array.length > 1) {
        array.splice(index, 1);
        index++;
        if (index >= array.length) {
            index -= array.length;
        }
    }
    return array.join();

}

const calsha = (n) => {
    let i = 1;
    let pow2 ;
    while ((pow2 = Math.pow(2, i) - 1) < n) {
        i++
    }
    return pow2 - 2*(pow2-n);
}


console.log(sha(255));
console.log(calsha(255));