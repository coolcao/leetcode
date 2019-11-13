struct ListNode {
  int val;
  struct ListNode *next;
};

#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

bool hasCycle(struct ListNode *head) {
  if (head == NULL) {
    return false;
  }

  struct ListNode *slow = head;
  struct ListNode *fast = head->next;

  while (true) {
      if (fast == NULL) {
          return false;
      }
      if (slow == fast) {
          return true;
      }
      slow = slow->next;
      fast = fast->next;
      if (fast == NULL) {
          return false;
      }
      fast = fast->next;
  }
}
