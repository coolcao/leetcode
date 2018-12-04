// #include <stdio.h>
// #include "../common/utils.h"
// int* twoSum(int* nums, int numsSize, int target) {
//   static int result[2] = {-1, -1};
//   for (int i = 0; i < numsSize - 1; i++) {
//     for (int j = i + 1; j < numsSize; j++) {
//       if (nums[i] + nums[j] == target) {
//         result[0] = i;
//         result[1] = j;
//         return result;
//       }
//     }
//   }
//   return result;
// }


#include<stdio.h>
#include <stdlib.h>
#include <string.h>
#define LEN 5
#define DEFAULT 512

typedef struct _node *Node;

struct _node {
  int idx;
  int value;
};

typedef struct _entry *Entry;

struct _entry {
  int key;
  Node node;
  Entry next;
};

typedef struct _map *HashMap;

struct _map {
  int size; // 标识当前map元素个数
  int usage;// 标识数组已使用个数
  int length;// 标识数组长度
  Entry *datas;// 数据
};

Node makeNode(int idx, int value) {
  Node node = malloc(sizeof *node);
  node->value = value;
  node->idx = idx;
  return node;
}

Entry makeEntry(int key, int idx, int value) {
  Entry entry = malloc(sizeof *entry);
  entry->key = key;
  entry->node = makeNode(idx, value);
  entry->next = NULL;
  return entry;
}

HashMap initHashMap() {
  HashMap map = malloc(sizeof *map);
  map->size = 0;
  map->usage = 0;
  map->length = DEFAULT;
  map->datas = (Entry *)malloc(sizeof(Entry) * map->length);
  for (int i = 0; i < map->length; i++) {
    map->datas[i] = 0;
  }
  return map;
}
Entry findEntry(Entry entry, int key) {
  if (!entry) return NULL;
  Entry current = entry;
  while (current && key != current->key) {
    current = current->next;
  }
  return current;
}
int findPosition(HashMap map, int key) {
  if (!map) return -1;
  if (key<0) key = -key;
  int pos = key % map->length;
  return pos;
}

void putToHashMap(HashMap map, int key, int idx,int value) {
  if (!map) return;
  int pos = findPosition(map, key);
  Entry entry = makeEntry(key, idx, value);
  if (map->datas[pos] == NULL) {
    map->datas[pos] = entry;
    map->usage += 1;
    map->size += 1;
  } else {
    // 这里需要检查key值是否已存在，如果已存在，那么需要覆盖当前的值
    Entry existEntry = findEntry(map->datas[pos], key);
    if (existEntry) {
      // 如果key已存在，直接覆盖value
      existEntry->node = makeNode(idx, value);
    } else {
      // 如果不存在该key，则将该节点插入到链表
      entry->next = map->datas[pos];
      map->datas[pos] = entry;
      map->size += 1;
    }
  }
}

Node getFromHashMap(HashMap map, int key) {
  if (!map) return NULL;
  int pos = findPosition(map, key);
  Entry start = map->datas[pos];
  if (start == NULL) return NULL;
  while (key != start->key && start->next != NULL) {
    start = start->next;
  }
  if (key == start->key) return start->node;
  return NULL;
}

int* twoSum(int* nums, int numsSize, int target) {
  static int result[2] = {-1, -1};
  HashMap map = initHashMap();
  for (int i=0;i<numsSize;i++) {
    int num = nums[i];
    int sub = target - num;
    // 差已存在，说明有值
    Node node = getFromHashMap(map, sub);
    if (node != NULL) {
      result[0] = node->idx;
      result[1] = i;
      break;
    } else {
      putToHashMap(map, num, i, num);
    }

  }
  return result;
}


int main(int argc, char const* argv[]) {
  int nums[LEN] = {-1, 1, 3, 4, 5};
  int target = 4;
  int* result = twoSum(nums, LEN, target);
  printf("%d,%d\n", result[0], result[1]);

  return 0;
}
