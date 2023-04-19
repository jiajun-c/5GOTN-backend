# 并查集
class UnionFind:
    def __init__(self, n):
        self.parent = [i for i in range(n)]
        self.rank = [1 for _ in range(n)]

    def find(self, p):
        '''查找祖先'''
        if self.parent[p] != p:
            self.parent[p] = self.find(self.parent[p])
        return self.parent[p]

    def union(self, a, b):
        '''合并两点所在集合'''
        rootA = self.find(a)
        rootB = self.find(b)
        if rootA != rootB:
            if self.rank[rootA] < self.rank[rootB]:
                self.parent[rootA] = rootB
            else:
                self.parent[rootB] = rootA
                if self.rank[rootA] == self.rank[rootB]:
                    self.rank[rootA] += 1
        return 

    def connected(self, a, b):
        '''判断两点是否同类'''
        return self.find(a) == self.find(b)
    
    def groups(self):
        '''返回所有集合'''
        r = range(len(self.parent))
        return [[j for j in r if self.connected(j, i)] for i in r if i == self.parent[i]]
    
    def groups_big(self):
        '''返回所有元素个数大于1的集合'''
        r = range(len(self.parent))
        return [g for g in [[j for j in r if self.connected(j, i)] for i in r if i == self.parent[i]] if len(g) > 1]

