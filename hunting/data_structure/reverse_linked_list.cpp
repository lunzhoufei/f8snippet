// https://leetcode-cn.com/problems/reverse-linked-list/submissions/
//
//


/**
 * Definition for singly-linked list.
 * struct ListNode {
 *    int val;
 *    ListNode *next;
 *    ListNode(int x) : val(x), next(NULL) {}
 * };
 */

class Solution {
  public:
    ListNode* reverseList(ListNode* head) {
      ListNode* p_cur = head;
      ListNode* p_ret = NULL;
      while (p_cur != NULL) {
        ListNode* p_processing = p_cur;
        p_cur = p_cur->next;
        p_processing ->next = p_ret;
        p_ret = p_processing;
      }
      return p_ret;
    }
};
