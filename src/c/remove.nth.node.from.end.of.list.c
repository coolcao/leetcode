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
