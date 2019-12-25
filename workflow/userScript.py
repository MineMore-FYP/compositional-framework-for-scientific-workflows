import os
from collections import OrderedDict
import sys

#Workflow
#0
##WFComponent1/studentGreeting.py
#1
##WFComponent2/multiplyByTwo.py
#2
##WFComponent2/calculateAvg.py
#3
##WFComponent3/amazing.py
#4
##WFComponent3/terrible.py

orderOfModules = ["studentGreeting", "multiplyByTwo", "calculateAvg", "amazing", "terrible"]

outputLocation = "/home/mpiuser/Desktop/"


'''#######################		WFComponent1	####################################'''

studentName = "John."


'''#######################		WFComponent2	####################################'''

studentMarks = [23,45,26,47,28,38,46,43,32,44]

classMarks  = [927,734,358,212,938,603,912,592,657,556,740,625,495,864,346,623,824,543,951,591,833,340,690,
544,797,818,847,791,727,652,856,651,686,964,765,732,662,674,714]

'''#######################		WFComponent3	####################################'''

durationType = "Semester"
