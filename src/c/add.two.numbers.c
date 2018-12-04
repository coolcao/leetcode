#include <stdio.h>
#include <stdlib.h>
struct ListNode {
  int val;
  struct ListNode* next;
};

struct ListNode* makeNode(int val, struct ListNode* next) {
  struct ListNode* node = malloc(sizeof *node);
  node->val = val;
  node->next = next;
  return node;
}

struct ListNode* addTwoNumbers(struct ListNode* l1, struct ListNode* l2) {
  int carry = 0;
  struct ListNode* head = NULL;
  struct ListNode* current = NULL;
  while (l1 != NULL && l2 != NULL) {
    int sum = l1->val + l2->val + carry;
    if (sum > 9) {
      carry = sum / 10;
      sum = sum % 10;
    } else {
      carry = 0;
    }

    struct ListNode* node = makeNode(sum, NULL);
    if (current) {
      current->next = node;
      current = node;
    } else {
      head = node;
      current = node;
    }
    l1 = l1->next;
    l2 = l2->next;
  }
  while (l1) {
    int sum = l1->val + carry;
    if (sum > 9) {
      carry = sum / 10;
      sum = sum % 10;
    } else {
      carry = 0;
    }

    struct ListNode* node = makeNode(sum, NULL);
    current->next = node;
    current = node;
    l1 = l1->next;
  }
  while (l2) {
    int sum = l2->val + carry;
    if (sum > 9) {
      carry = sum / 10;
      sum = sum % 10;
    } else {
      carry = 0;
    }

    struct ListNode* node = makeNode(sum, NULL);
    current->next = node;
    current = node;
    l2 = l2->next;
  }

  if (carry > 0) {
    struct ListNode* node = makeNode(carry, NULL);
    current->next = node;
    current = node;
  }

  return head;
}

void printNode(struct ListNode* node) {
  struct ListNode* head = node;
  while (head) {
    printf("%d->", head->val);
    head = head->next;
  }
  printf("\n");
}

int main(int argc, char const* argv[]) {
  struct ListNode* list1 = makeNode(1, NULL);
  struct ListNode* list11 = makeNode(4, NULL);
  struct ListNode* list12 = makeNode(9, NULL);
  struct ListNode* list13 = makeNode(7, NULL);
  struct ListNode* list14 = makeNode(6, NULL);

  list13->next = list14;
  list12->next = list13;
  list11->next = list12;
  list1->next = list11;

  struct ListNode* list2 = makeNode(9, NULL);
  struct ListNode* list21 = makeNode(1, NULL);
  struct ListNode* list22 = makeNode(6, NULL);
  struct ListNode* list23 = makeNode(2, NULL);
  list22->next = list23;
  list21->next = list22;
  list2->next = list21;

  printNode(list1);
  printNode(list2);

  struct ListNode* result = addTwoNumbers(list1, list2);
  printNode(result);

  return 0;
}
