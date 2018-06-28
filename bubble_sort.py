#!/usr/bin/env python
# _*_ coding:utf-8 _*_


def bsort_recursion(els):
    ret = []
    if len(els) == 1:
        return els

    index = len(els)-1
    
    while index-1 >= 0:
        if els[index] < els[index-1]:
            els[index], els[index-1]= els[i-1],els[index]
        index -= 1
    
    ret.append(els[0])
    ret.extend(bsort_recursion(els[1:]))
    return ret

def bsort(els):
    loops = len(els) - 1
    while loops > 0:
        loops -= 1
        
        i = len(els) - 1
        while i-1 > 0:
            if els[i] < els[i-1]:
                els[i],els[i-1] = els[i-1], els[i]
            i -= 1
    return els


if __name__ == '__main__':
    l = [0,1,2,3,4,2,1,8,9,3,5]
    print bsort(l)
    print bsort_recursion(l)

    
