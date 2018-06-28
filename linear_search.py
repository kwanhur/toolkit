#!/usr/bin/env python
# _*_ coding:utf-8 _*_


def find(nums, n):
    '''
    线性遍历是否找到元素n
    '''
    if not nums:
        return False
    for v in nums:
        if v == n:
            return True
    return False


if __name__ == '__main__':
    l = [0,2,4,6,2,6,7]
    print find(l, 10)
    print find(l, 6)

