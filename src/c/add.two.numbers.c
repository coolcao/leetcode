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

// 采用递归的形式进行计算
struct ListNode* add(struct ListNode* l1, struct ListNode* l2, int carry) {
  struct ListNode* head = NULL;
  if (l1 && l2) {
    int sum = l1->val + l2->val + carry;
    if (sum > 9) {
      carry = sum / 10;
      sum = sum % 10;
    } else {
      carry = 0;
    }
    head = makeNode(sum, NULL);
    head->next = add(l1->next, l2->next, carry);
  } else if (l1) {
    int sum = l1->val + carry;
    if (sum > 9) {
      carry = sum / 10;
      sum = sum % 10;
    } else {
      carry = 0;
    }
    head = makeNode(sum, NULL);
    head->next = add(l1->next, NULL, carry);
  } else if (l2) {
    int sum = l2->val + carry;
    if (sum > 9) {
      carry = sum / 10;
      sum = sum % 10;
    } else {
      carry = 0;
    }
    head = makeNode(sum, NULL);
    head->next = add(l2->next, NULL, carry);
  } else if (carry > 0) {
    head = makeNode(carry, NULL);
  }
  return head;
}
struct ListNode* addTwoNumbers2(struct ListNode* l1, struct ListNode* l2) {
  int carry = 0;
  return add(l1, l2, carry);
}

void printNode(struct ListNode* node) {
  struct ListNode* head = node;
  printf("[");
  while (head) {
    if (head->next) {
      printf("%d,", head->val);
    } else {
      printf("%d", head->val);
    }
    head = head->next;
  }
  printf("]\n");
}

struct ListNode* createListFromArray(int* a, int length) {
  struct ListNode* head = NULL;
  struct ListNode* current = NULL;
  for (int i = 0; i < length; i++) {
    struct ListNode* node = makeNode(a[i], NULL);
    if (current) {
      current->next = node;
      current = node;
    } else {
      current = node;
      head = node;
    }
  }
  return head;
}

int main(int argc, char const* argv[]) {
  #define length1 7
  #define length2 6
  int a[length1] = {3, 2, 4, 6, 1, 4, 7};
  int b[length2] = {4, 1, 7, 5, 9, 1};
  struct ListNode* l1 = createListFromArray(a, length1);
  struct ListNode* l2 = createListFromArray(b, length2);

  printNode(l1);
  printNode(l2);

  struct ListNode* result = addTwoNumbers2(l1, l2);
  printNode(result);

  return 0;
}
