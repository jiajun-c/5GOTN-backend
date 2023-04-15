from collections import defaultdict
import os

def Mining_FreSeqPatterns(Seqs,minsp,minpatternlen=2,maxpatternlen=10,maxgap=100):
    '''序列规则挖掘'''
    itemToidx = defaultdict(int)
    idxToitem = defaultdict(str)
    with open("./spmf/input.txt","w+") as fp:
        idx = 1
        for i in range(len(Seqs)):
            line = ""
            for j in range(len(Seqs[i])):
                for item in Seqs[i][j]:
                    if itemToidx[item] == 0:
                        itemToidx[item] = idx
                        idxToitem[idx] = item
                        idx += 1
                    line = line + str(itemToidx[item]) + " " 
                line = line + "-1 "
            line += "-2\n"
            fp.write(line)
    fp.close()
    
    # Run VGEN
    # os.system( f'java -jar ./spmf/spmf.jar run VGEN ./spmf/input.txt ./spmf/output.txt {minsp} {maxpatternlen} {maxgap}')
    # Run SPAM
    os.system( f'java -jar ./spmf/spmf.jar run SPAM ./spmf/input.txt ./spmf/output.txt {minsp} {minpatternlen} {maxpatternlen} {maxgap}')

    res = []
    # Read the output file line by line
    outFile = open("./spmf/output.txt",'r', encoding='utf-8')
    for string in outFile:
        line = string.strip('\n').split(' ')
        line[-2]='-2'
        line = list(map(int,line))
        seq = [[]]
        for i in range(len(line)):
            if line[i]==-2:
                seq.pop()
                res.append([seq,line[i+1]])
                break
            elif line[i] == -1:
                seq.append([])
            else:
                seq[-1].append(idxToitem[line[i]])
    outFile.close()
    return res

def Mining_Seqrules(Seqs,minsp,minconf,max_antecedent_len,max_consequent_len):
    '''频繁序列挖掘'''
    itemToidx = defaultdict(int)
    idxToitem = defaultdict(str)
    with open("./spmf/input.txt","w+") as fp:
        idx = 1
        for i in range(len(Seqs)):
            line = ""
            for j in range(len(Seqs[i])):
                for item in Seqs[i][j]:
                    if itemToidx[item] == 0:
                        itemToidx[item] = idx
                        idxToitem[idx] = item
                        idx += 1
                    line = line + str(itemToidx[item]) + " " 
                line = line + "-1 "
            line += "-2\n"
            fp.write(line)
    fp.close()
    # Run RuleGrowth
    os.system( f'java -jar ./spmf/spmf.jar run RuleGrowth ./spmf/input.txt ./spmf/output.txt {minsp} {minconf} {max_antecedent_len} {max_consequent_len}')

    res = []
    # Read the output file line by line
    outFile = open("./spmf/output.txt",'r', encoding='utf-8')
    for string in outFile:
        rule = [[],[],0,0]
        line = string.strip('\n').split(' ')
        rule[0] = list(map(lambda x:idxToitem[int(x)],line[0].split(',')))
        rule[1] = list(map(lambda x:idxToitem[int(x)],line[2].split(',')))
        rule[2] = int(line[4])
        rule[3] = float(line[6])
        res.append(rule)
    outFile.close()
    return res

def Mining_FreItemsets(Itemsets,minsp,minpatternlen=1,maxpatternlen=10):
    '''频繁项挖掘'''
    with open("./spmf/input.txt","w+") as fp:
        for itemset in Itemsets:
            line = ' '.join(list(map(str,itemset)))
            fp.write(line+'\n')
    fp.close()

    # Run FPClose
    os.system( f'java -jar ./spmf/spmf.jar run FPClose ./spmf/input.txt ./spmf/output.txt {minsp} {maxpatternlen}')

    res = []
    outFile = open("./spmf/output.txt",'r', encoding='utf-8')
    for string in outFile:
        line = string.strip('\n').split(' ')
        itemset = []
        for x in line[:-2]:
            itemset.append(int(x))
        res.append([itemset,int(line[-1])])
    outFile.close()
    return res

def Mining_associationrules(Itemsets,minsp,minconf,max_antecedent_len,max_consequent_len):
    '''关联规则挖掘'''
    with open("./spmf/input.txt","w+") as fp:
        for itemset in Itemsets:
            line = ' '.join(list(map(str,itemset)))
            fp.write(line+'\n')
    fp.close()

    # Run FPGrowth_association_rules
    os.system( f'java -jar ./spmf/spmf.jar run FPGrowth_association_rules ./spmf/input.txt ./spmf/output.txt {minsp} {minconf} {max_antecedent_len} {max_consequent_len}')

    res = []
    outFile = open("./spmf/output.txt",'r', encoding='utf-8')
    for string in outFile:
        line = string.strip('\n').split(' ')
        rule = [[],[],0,0]
        cur = 0
        for i in range(len(line)):
            if line[i] == '==>':
                cur = 1
            elif line[i] == '#SUP:':
                rule[2] = int(line[i+1])
                rule[3] = float(line[i+3])
                break
            else:
                rule[cur].append(int(line[i]))
        res.append(rule)
    outFile.close()
    return res