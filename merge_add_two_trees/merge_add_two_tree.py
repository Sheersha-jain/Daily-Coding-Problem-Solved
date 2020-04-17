"""
Write a program to merge two binary trees. Each node in the new tree should
hold a value equal to the sum of the values of the corresponding nodes of
the input trees.If only one input tree has a node in a given position, the
corresponding node in the new tree should match that input node.
"""


class Tree:

    def __init__(self, data):
        self.val = data
        self.left = None
        self.right = None


def merge_tree(t1, t2):
    if t1 and t2 is None:
        return
    if t1 is None:
        return t2
    if t2 is None:
        return t1
    if t2 is not None:
        t1.val += t2.val
        t1.left = merge_tree(t1.left, t2.left)
        t1.right = merge_tree(t1.right, t2.right)
    return t1
