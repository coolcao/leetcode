/*
 * @lc app=leetcode id=24 lang=c
 *
 * [24] Swap Nodes in Pairs
 *
 * https://leetcode.com/problems/swap-nodes-in-pairs/description/
 *
 * algorithms
 * Medium (42.76%)
 * Total Accepted:    274.9K
 * Total Submissions: 642.1K
 * Testcase Example:  '[1,2,3,4]'
 *
 * Given a linked list, swap every two adjacent nodes and return its head.
 *
 * You may not modify the values in the list's nodes, only nodes itself may be
 * changed.
 *
 *
 *
 * Example:
 *
 *
 * Given 1->2->3->4, you should return the list as 2->1->4->3.
 *
 *
 */

#include <stdio.h>
#include <stdlib.h>
struct ListNode {
  int val;
  struct ListNode* next;
};

struct ListNode* makeNode(int val) {
  struct ListNode* node = malloc(sizeof *node);
  node->val = val;
  node->next = NULL;
  return node;
}

struct ListNode* makeListFromArray(int* array, int size) {
  if (size == 0) return NULL;
  struct ListNode* head = makeNode(array[0]);
  struct ListNode* current = head;
  for (int i = 1; i < size; i++) {
    current->next = makeNode(array[i]);
    current = current->next;
  }
  return head;
}

void printList(struct ListNode* head) {
  if (head == NULL) {
    printf("NULL\n");
    return;
  };
  printf("%d->", head->val);
  printList(head->next);
}

// 递归方式
// ----------------- start ---------------
// struct ListNode* swapPairs(struct ListNode* head) {
//   if (head == NULL || head->next == NULL) return head;

//   struct ListNode* node1 = head;
//   struct ListNode* node2 = head->next;

//   struct ListNode* current = node2->next;

//   // node1->next = node2->next;
//   node2->next = node1;

//   node1->next = swapPairs(current);

//   return node2;
// }
// ------------------ end ----------------


// 采用迭代的方式
// ----------- start -------------------
struct ListNode* swap(struct ListNode* head) {
  if (head == NULL || head->next == NULL) return head;

  struct ListNode* node1 = head;
  struct ListNode* node2 = head->next;

  node1->next = node2->next;
  node2->next = node1;

  // node1->next = swapPairs(current);

  return node2;
}
struct ListNode* swapPairs(struct ListNode* head) {
  struct ListNode* current = swap(head);
  struct ListNode* result = current;

  while (current && current->next) {
    current = current->next;
    current->next = swap(current->next);
    current = current->next;
  }

  return result;
}
// -------------- end ------------------

int main(int argc, char const* argv[]) {
#define LEN 10
  int a[LEN] = {8, 4, 5, 6, 2, 1, 8, 9, 3, 10};
  struct ListNode* head = makeListFromArray(a, LEN);
  printList(head);
  struct ListNode* swaped = swapPairs(head);
  printList(swaped);
  return 0;
}
