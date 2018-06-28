#! /usr/bin/env python
# _*_ coding:utf-8 _*_

def find(nums, n):
    '''
    在数组里找出是否存在数字n，二分查找
    '''
    if not nums:
        return False
    if len(nums) == 1:
        return nums[0] == n

    while True:
        index = len(nums)/2
        if index == 0:
            return nums[index] == n

        if nums[index] > n:
            nums = nums[:index]
            continue
        if nums[index] < n:
            nums = nums[index:]
            continue
        return nums[index] == n

if __name__ == '__main__':
    l = [0,1,2,3,4,6,3,2,1,9]
    print find(l, 5)
    print find(l, 6)
        
