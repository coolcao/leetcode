#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "../common/uthash.h"

int isPalindromic(char* s, int start, int end) {
  int length = end - start + 1;
  if (length < 0) return 0;
  int r = 1;
  for (int i = start; i < start + length / 2; i++) {
    char c1 = s[i];
    char c2 = s[start + end - i];
    if (c1 != c2) {
      r = 0;
      break;
    }
  }
  return r;
}

char* longestPalindrome(char* s) {
  int length = strlen(s);
  int max = 0;
  int len = 0;
  int maxStart = 0;
  int maxEnd = 0;
  for (int i = 0; i < length; i++) {
    for (int j = 0; j < length; j++) {
      if (isPalindromic(s, i, j)) {
        len = j - i;
        if (len > max) {
          max = len;
          maxStart = i;
          maxEnd = j;
        }
      }
    }
  }

  printf("%d,%d,%d\n", max, maxStart, maxEnd);

  char* result;
  result = (char*)malloc(sizeof(char*) * (max + 1));

  for (int i = 0; i <= max; i++) {
    result[i] = s[maxStart + i];
  }
  result[max + 1] = '\0';

  return result;
}

typedef struct map_struct* Map;

struct map_struct {
  int key;           /* we'll use this field as the key */
  UT_hash_handle hh; /* makes this structure hashable */
};

Map map = NULL;

void add_num(int key) {
  Map _map;
  _map = (Map)malloc(sizeof(*map));
  _map->key = key;
  HASH_ADD_INT(map, key, _map);
}
Map find_num(int key) {
  Map _map;
  HASH_FIND_INT(map, &key, _map);
  return _map;
}
int contains(int key) {
  Map _map;
  HASH_FIND_INT(map, &key, _map);
  return _map != NULL;
}
void clear() {
  Map m = map;
  map = NULL;
  free(m);
}

char* longestPalindrome2(char* s) {
  int length = strlen(s);
  int max = 0;
  int len = 0;
  int maxStart = 0;
  int maxEnd = 0;
  for (int i = 0; i < length - 1; i++) {
    clear();
    add_num(s[i]);
    for (int j = i + 1; j < length; j++) {
      if (contains(s[j])) {
        if (isPalindromic(s, i, j)) {
          len = j - i;
          if (len > max) {
            max = len;
            maxStart = i;
            maxEnd = j;
          }
        }
      } else {
        continue;
      }
    }
  }

  char* result;
  result = (char*)malloc(sizeof(char*) * (max + 1));

  for (int i = 0; i <= max; i++) {
    result[i] = s[maxStart + i];
  }
  result[max + 1] = '\0';
  return result;
}

char* longestPalindrome3(char* s) {
  int length = strlen(s);
  int max = 0;
  int maxStart = 0;
  int maxEnd = 0;

  for (int i = 0; i < length - 1; i++) {
    int len = 0;
    for (int j = length - 1; j > i; j--) {
      len = j - i;
      if (len > max) {
        if (isPalindromic(s, i, j)) {
          max = len;
          maxStart = i;
          maxEnd = j;
        }
      }
    }
  }

  printf("%d,%d,%d\n", max, maxStart, maxEnd);
  char* result;
  result = (char*)malloc(sizeof(char*) * (max + 1));

  for (int i = 0; i <= max; i++) {
    result[i] = s[maxStart + i];
  }
  result[max + 1] = '\0';
  return result;
}

// 直接寻找回文子串
// 分两种情况，子串长度为奇数和偶数分别寻找
char* longestPalindrome4(char* s) {
  int length = strlen(s);
  if (length<=1) {
    return s;
  }
  int max = 1;
  int maxStart = 0;
  int maxEnd = 0;

  for (int i=0;i<length;i++) {
    int start;
    int end;
    // 分两种情况
    // 一种是回文串是偶数个的
    for (start=i,end=i+1;start>=0 && end<length;start--,end++) {
      if (s[start] == s[end]) {
        if (start == 0 || end == length - 1) {
          if (max < end - start + 1 ) {
            max = end - start + 1;
            maxStart = start;
            maxEnd = end;
          }
        }
        continue;
      } else {
        if (max < end - start) {
          max = end - start - 1;
          maxStart = start + 1;
          maxEnd = end + 1;
        }
        break;
      }
    }
    // 一种是回文串是奇数个的
    for (start=i-1,end=i+1;start>=0 && end<length;start--,end++) { 
      if (s[start] == s[end]) {
        if (start == 0 || end == length - 1) {
          if (max < end - start + 1 ) {
            max = end - start + 1;
            maxStart = start;
            maxEnd = end;
          }
        }
        continue;
      } else {
        if (max < end - start) {
          max = end - start - 1;
          maxStart = start + 1;
          maxEnd = end + 1;
        }
        break;
      }
    }

  }

  printf("%d,%d,%d\n", max, maxStart, maxEnd);
  char* result = NULL;
  result = (char*)malloc(sizeof(char*) * (max + 1));

  for (int i = 0; i < max; i++) {
    // printf("result[%d] = s[%d+%d] = %c\n", i, maxStart, i, s[maxStart + i]);
    result[i] = s[maxStart + i];
  }

  result[max] = '\0';
  return result; 
}

char* longestPalindrome5 (char* s) {
  int length = strlen(s);
  int left = (length - 1)/2;
  int right = left + 1;

  int max = 1;
  int maxStart = 0;
  int maxEnd = 0;

  while (left>=0 || right<=length) {
    int start, end;
  	if (left>=0) {
      for (start=left,end=left+1;start>=0 && end<length;start--,end++) {
        if (s[start] == s[end]) {
          if (start == 0 || end == length - 1) {
            if (max < end - start + 1 ) {
              max = end - start + 1;
              maxStart = start;
              maxEnd = end;
            }
          }
          continue;
        } else {
          if (max < end - start) {
            max = end - start - 1;
            maxStart = start + 1;
            maxEnd = end + 1;
          }
          break;
        }
      }
      for (start=left-1,end=left+1;start>=0 && end<length;start--,end++) {
        if (s[start] == s[end]) {
          if (start == 0 || end == length - 1) {
            if (max < end - start + 1 ) {
              max = end - start + 1;
              maxStart = start;
              maxEnd = end;
            }
          }
          continue;
        } else {
          if (max < end - start) {
            max = end - start - 1;
            maxStart = start + 1;
            maxEnd = end + 1;
          }
          break;
        }
      }
  	}
  	if (right<=length) {
      for (start=right,end=right+1;start>=0 && end<length;start--,end++) {
        if (s[start] == s[end]) {
          if (start == 0 || end == length - 1) {
            if (max < end - start + 1 ) {
              max = end - start + 1;
              maxStart = start;
              maxEnd = end;
            }
          }
          continue;
        } else {
          if (max < end - start) {
            max = end - start - 1;
            maxStart = start + 1;
            maxEnd = end + 1;
          }
          break;
        }
      }
      for (start=right-1,end=right+1;start>=0 && end<length;start--,end++) {
        if (s[start] == s[end]) {
          if (start == 0 || end == length - 1) {
            if (max < end - start + 1 ) {
              max = end - start + 1;
              maxStart = start;
              maxEnd = end;
            }
          }
          continue;
        } else {
          if (max < end - start) {
            max = end - start - 1;
            maxStart = start + 1;
            maxEnd = end + 1;
          }
          break;
        }
      }
  	}
    left -= 1;
    right += 1;
  }


  printf("%d,%d,%d\n", max, maxStart, maxEnd);
  char* result = NULL;
  result = (char*)malloc(sizeof(char*) * (max + 1));

  for (int i = 0; i < max; i++) {
    // printf("result[%d] = s[%d+%d] = %c\n", i, maxStart, i, s[maxStart + i]);
    result[i] = s[maxStart + i];
  }

  result[max] = '\0';
  return result;


}

int main(int argc, char const* argv[]) {
//  char* s = "jjjjjjjjjkkkkkkkddkksffjjksdjfkdjfksjdkfjkdjfkkkkkkkkkkkkllllslsllsllslslsllslsdkfksdjfsdkfskskskskkskkskkkksksssfjfjfjfjfjjfjfjfdkddkdkkdkdkdkdkkdkdkddkkkdkkdkkdkkdkdkdkdkdkdkkdkdkdkdkdkkdkdkdkdkkdkdkdkdkdkdkkdkkkkdddkkdkkdddkdkkd";
//   char *s = "dabebac";
   char *s = "aa";
  char* r4 = longestPalindrome4(s);
  char* r5 = longestPalindrome5(s);
  printf("%s\n", r4);
  printf("%s\n", r5);

  //	add_num('a');
  //	add_num('b');
  //	add_num('c');
  //
  //	printf("c exist?%d\n", contains('c'));
  //	printf("d exist?%d\n", contains('d'));

  return 0;
}
