
#include <iostream>
#include <queue>
#include <vector>
#include <map>
#include <string>
using namespace std;

class Node {
    public:
        Node* left;
        Node* right;
        int value;
};


// inorder 
void inorder_iterate_tree(Node* root, std::vector<Node*>* result) {
    if (root == NULL) {
        return;
    }
    inorder_iterate_tree(root->left, result);
    result->push_back(root);
    inorder_iterate_tree(root->left, result);
}

// preorder
void preorder_iterate_tree(Node* root, std::vector<Node*>* result) {
    if (root == NULL) {
        return;
    }
    result->push_back(root);
    preorder_iterate_tree(root->left, result);
    preorder_iterate_tree(root->left, result);
}

// postorder
void postorder_iterate_tree(Node* root, std::vector<Node*>* result) {
    if (root == NULL) {
        return;
    }
    postorder_iterate_tree(root->left, result);
    postorder_iterate_tree(root->left, result);
    result->push_back(root);
}


bool is_sorted_tree(vector<Node*> data) {
    if (data.empty()) {
        return true;
    }
    for (int i = 0, size = data.size() - 1; i < size; ++i) {
        if (data[i + 1]->value < data[i]->value) {
            return false;
        }
    }
    return true;
}


int main(int argc, char *argv[])
{
    Node* broot = new Node();
    vector<Node*> tv_middle_sort;

    // 判断是否为搜索二叉树， 搜索二叉树为：所有节点的左孩子< 当前节点<右孩子
    inorder_iterate_tree(broot, &tv_middle_sort);
    bool is_sorted = is_sorted_tree(tv_middle_sort);

    return 0;
}

