/**

Merge two sorted linked lists and return it as a new list. The new list should
be made by splicing together the nodes of the first two lists.

Example:

Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4

====
简单
====

 */

#include <stdio.h>
#include <stdlib.h>
struct ListNode {
  int val;
  struct ListNode* next;
};
typedef struct ListNode* Node;

struct ListNode* mergeTwoLists(struct ListNode* l1, struct ListNode* l2) {

  if (l1 == NULL && l2 == NULL) return NULL;

  Node node = (Node)malloc(sizeof *node);
  Node current = node;
  Node ln1 = l1, ln2 = l2;


  while (ln1 && ln2) {
    Node n = (Node)malloc(sizeof *n);
    if (ln1->val > ln2->val) {
      n->val = ln2->val;
      n->next = NULL;

      current->next = n;
      current = n;
      ln2 = ln2->next;

    } else {
      n->val = ln1->val;
      n->next = NULL;

      current->next = n;
      current = n;
      ln1 = ln1->next;
    }
  }

  while (ln1) {
    Node n = (Node)malloc(sizeof *n);
    n->val = ln1->val;
    n->next = NULL;

    current->next = n;
    current = n;
    ln1 = ln1->next;
  }

  while (ln2) {
    Node n = (Node)malloc(sizeof *n);
    n->val = ln2->val;
    n->next = NULL;

    current->next = n;
    current = n;
    ln2 = ln2->next;
  }

  Node head = node->next;
  free(node);
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
  head->next = NULL;

  Node current = head;
  int idx = 1;
  while (idx < size) {
    Node node = (Node)malloc(sizeof *node);
    node->val = nums[idx];
    node->next = NULL;

    current->next = node;
    current = node;
    idx += 1;
  }

  return head;
}

int main(int argc, char const* argv[]) {
#define LEN1 1
#define LEN2 1
  int nums1[LEN1] = {4};
  int nums2[LEN2] = {2};
  Node l1 = createListFromArray(nums1, LEN1);
  printNode(l1);
  Node l2 = createListFromArray(nums2, LEN2);
  printNode(l2);
  Node l3 = mergeTwoLists(l1, l2);
  printNode(l3);
  return 0;
}
