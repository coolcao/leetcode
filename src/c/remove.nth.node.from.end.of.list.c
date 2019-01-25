/**

Given a linked list, remove the n-th node from the end of list and return its
head.

Example:

Given linked list: 1->2->3->4->5, and n = 2.

After removing the second node from the end, the linked list becomes 1->2->3->5.
Note:

Given n will always be valid.

Follow up:

Could you do this in one pass?

========
中等
========

这个题目很容易理解，而且逻辑也不难，但是确被标记为中等，主要是因为题目中多了个限制：
只遍历依次链表，完成删除节点操作。

其实按照正常的逻辑来讲，如果要完成这个操作，需要遍历两次链表，第一次遍历，统计链表中元素的个数，第二次再根据个数和要删除的位置删除元素即可。

可是如果要求只遍历依次链表，那么就要求，要么在统计链表个数时，直接删除元素。或者统计完元素个数后，能够根据要删除元素的位置能够直接定位到要删除的元素，然后删除。

很显然，如果是第一种情况，在遍历统计链表中元素个数时，由于没法知道具体元素个数，很难根据倒序的位置去确定元素的位置，所以第一种方式可能不大好实现。

但是第二种，是有方法可以实现的。

我们可以在第一次遍历统计元素个数时，顺便将每个位置对应的元素“缓存”下来。这样我们根据元素总个数以及要删除元素的倒序位置能确定要删除元素的具体位置。

然后由于缓存下来了每个位置的元素，那么是可以直接定位到对应元素的，再删除元素就满足了只遍历依次链表的条件。

要实现O(1)的缓存实现，哈希表就是一种方案，我们可以以位置作为键，元素作为值，放置到map中。

代码里的map使用的是 uthash.h 这个类库。

 */

#include <stdio.h>
#include <stdlib.h>
#include "../common/uthash.h"
struct ListNode {
  int val;
  struct ListNode* next;
};
struct _map {
  int idx;
  struct ListNode* node;
  UT_hash_handle hh;
};
typedef struct _map* Map;
typedef struct ListNode* Node;
struct ListNode* removeNthFromEnd(struct ListNode* head, int n) {
  if (head == NULL || n < 1) return NULL;
  Map map = NULL;
  int nodes = 0;
  Node node = head;
  int idx = 1;
  while (node) {
    Map m = (Map)malloc(sizeof *m);
    m->idx = idx;
    m->node = node;
    HASH_ADD_INT(map, idx, m);
    node = node->next;
    idx += 1;
    nodes += 1;
  }

  if (n > nodes) {
    return NULL;
  }

  Map node2remove = NULL;  // 要删除的节点
  idx = nodes + 1 - n;     // 要删除节点的位置
  HASH_FIND_INT(map, &idx, node2remove);

  if (idx == 1) {  //如果要删除的节点是头节点，则直接返回头节点下一个节点
    return head->next;
  }

  Map preNode = NULL;  //要删除节点的上一个节点
  idx = nodes - n;
  HASH_FIND_INT(map, &idx, preNode);
  if (preNode != NULL) {
    preNode->node->next = node2remove->node->next;
    free(node2remove);
  }
  return head;
}

void printNode(Node node) {
  while (node != NULL) {
    printf("%d ", node->val);
    node = node->next;
  }
  printf("\n");
}

Node createListFromArray(int* nums, int size) {
  if (size == 0) return NULL;
  Node head = (Node)malloc(sizeof(*head));
  head->val = nums[0];

  Node current = head;
  int idx = 1;
  while(idx < size){
    Node node = (Node) malloc(sizeof *node);
    node->val = nums[idx];
    node->next = NULL;

    current->next = node;
    current = node;
    idx += 1;
  }
  

  return head;
}

int main(int argc, char const* argv[]) {

  #define LEN 5
  int nums[LEN] = {5,4,8,2,1};
  Node head = createListFromArray(nums, LEN);

  printNode(head);


  Node result = removeNthFromEnd(head, 2);
  printNode(result);
  return 0;
}
