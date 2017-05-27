'use strict';

class TreeNode {
    constructor(val) {
        this.val = val;
        this.left = this.right = null;
    }
}

const maxDepth = function(root) {
    if(!root.val){
        return 0;
    }
    root.deepth = 1;
    const q = [root];
    while (q.length > 0) {

        let node = q.shift();
        if (node.left) {
            node.left.deepth = node.deepth + 1;
            q.push(node.left);
        }
        if (node.right) {
            node.right.deepth = node.deepth + 1;
            q.push(node.right);
        }

        if (q.length === 0) {
            return node.deepth;
        }
    }
}

const root = {
    "left": {
        "left": {
            "val": "sdlfk"
        },
        "right": {
            "right": {
                "left": {
                    "val": "dkjdslk"
                },
                "right": {
                    "val": "sldfk"
                },
                "val": "asdfk"
            },
            "val": "aaa"
        },
        "val": "l1"
    },
    "right": {
        "left": {
            "left": {
                "val": "dksjl"
            },
            "right": {
                "left": {
                    "val": "dskkd"
                },
                "right": {
                    "val": "sdlfk"
                },
                "val": "kdsf"
            },
            "val": "dks"
        },
        "val": "klsdfj"
    },
    "val": "root"
};

const node = {
    val: 'dsk'
}

console.log(maxDepth({}));