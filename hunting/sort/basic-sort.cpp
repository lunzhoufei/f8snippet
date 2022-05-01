#include<iostream>
#include<vector>
using namespace std;



// o(n2)
void bubble_sort(vector<int>* v_data) {
    if (v_data == NULL || v_data->empty()) {
        return;
    }
    // FIXME: for (int i = 0, size = v_data->size() - 1; i < size; ++i) {
    for (int i = 0, size = v_data->size(); i < size; ++i) {
        for (int k = 0; k < size - 1 - i; ++k) {
            if (v_data->at(k)  > v_data->at(k + 1)) {
                int tmp = v_data->at(k);
                (*v_data)[k] = v_data->at(k + 1);
                (*v_data)[k + 1] = tmp;
            }
        }
    }
}


// 
void select_sort(vector<int>* v_data) {
    if (v_data == NULL || v_data->empty()) {
        return;
    }
    for (int i = 0, size = v_data->size(); i < size; ++i) {
        int min = v_data->at(i);
        int min_index = i;
        for (int k = i + 1; k < size; ++k) {
            if (v_data->at(k) >= min) {
                continue;
            }
            min = v_data->at(k);
            min_index = k;

        }
        if (min_index == i) {
            continue;
        }
        (*v_data)[min_index] = v_data->at(i);
        (*v_data)[i] = min;
    }
}

// TODO
/* void insert_sort(vector<int>* v_data) { */
/*     if (v_data == NULL || v_data->empty()) { */
/*         return; */
/*     } */
/*     for (int i = 0, size = v_data->size(); i < size; ++i) { */
/*         for (int k = i; k >= 0; --k) { */
/*             if (v_data->at(k) < ) */

/*         } */
/*     } */
/* } */


void print(vector<int> v_data, string prefix) {
    std::cout << prefix + " output:";
    for (int i = 0, size = v_data.size(); i < size; ++i) {
        std::cout << " " << v_data[i];
    }
    std::cout << std::endl;
}


int main(int argc, char *argv[])
{

    vector<int> input{8, 3, 4, 5, 9, 83, 45, 28, 99, 234, 555, 34, 9};

    vector<int> bubble_input(input);
    bubble_sort(&bubble_input);
    print(bubble_input, "bubble_sort");

    vector<int> select_input(input);
    select_sort(&select_input);
    print(select_input, "select_sort");
    
    return 0;
}

