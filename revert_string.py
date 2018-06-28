#! /usr/bin/env python
# _*_ coding:utf-8 _*_


def revert(s):
    '''
    非递归
    '''
    if not s or len(s) == 1:
        return s

    i = -1
    n = -len(s)

    ret = ''
    while 1:
        ret += s[i]

        i -= 1
        if i >= n:
            continue
        else:
            break
    return ret

def revert_by_recursion(s):
    '''
    递归
    '''
    ret = ''

    if not s:
        return ''
    if len(s) == 1:
        return s
    else:
        ret = s[-1] + revert_by_recursion(s[0:len(s)-1])
    return ret

if __name__ == '__main__':
    str1 = "kwanhur"
    print revert(str1)
    print revert_by_recursion(str1)
