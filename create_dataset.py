import sys
import random

# 特にデータセットに制限がなかったため、適当そうな値に設定
MAX_ENEMY_NUM = 20
MAX_ENEMY_LIST_SIZE = 20

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python3 create_dataset.py ${dataset_num}")
        exit(1)
    
    if not sys.argv[1].isdecimal():
        print("input number for dataset_num")
        exit(1)
    
    dataset_num = int(sys.argv[1])
    print(dataset_num)
    for i in range(dataset_num):
        enemy_list_size = random.randint(1, MAX_ENEMY_LIST_SIZE)
        special_limit = random.randint(1, enemy_list_size)
        enemy_list = [random.randint(1, MAX_ENEMY_NUM) for i in range(enemy_list_size)]
        print(special_limit)
        for i, enemy_num in enumerate(enemy_list):
            if i == len(enemy_list) - 1:
                print(enemy_num)
            else:
                print(enemy_num, end=" ")