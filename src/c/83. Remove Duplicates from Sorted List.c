/*
Given a sorted linked list, delete all duplicates such that each element appear
only once.

Example 1:

Input: 1->1->2
Output: 1->2
Example 2:

Input: 1->1->2->3->3
Output: 1->2->3
*/

#include <stdio.h>
#include <stdlib.h>

struct ListNode {
  int val;
  struct ListNode *next;
};

struct ListNode *makeNode(int val) {
  struct ListNode *node = malloc(sizeof *node);
  node->val = val;
  node->next = NULL;
  return node;
}
struct ListNode *makeListFromArray(int *array, int size) {
  if (size == 0) return NULL;
  struct ListNode *head = makeNode(array[0]);
  struct ListNode *current = head;
  for (int i = 1; i < size; i++) {
    current->next = makeNode(array[i]);
    current = current->next;
  }
  return head;
}
void printList(struct ListNode *head) {
  if (head == NULL) {
    printf("NULL\n");
    return;
  };
  printf("%d->", head->val);
  printList(head->next);
}

struct ListNode *deleteDuplicates(struct ListNode *head) {
  if (head == NULL) {
    return NULL;
  }

  struct ListNode *current = head;
  int diffVal = head->val;

  while (current->next != NULL) {
    if (current->next->val != diffVal) {
      current = current->next;
      diffVal = current->val;
      continue;
    }
    struct ListNode *next = current->next->next;
    current->next = next;
  }

  return head;
}

int main(int argc, char const *argv[]) {
#define LEN 5
  int a[LEN] = {1,2,2,2,4};
  struct ListNode *head = makeListFromArray(a, LEN);
  printList(head);
  deleteDuplicates(head);
  printList(head);
  return 0;
}
