import parsl
import os
from parsl.app.app import python_app, bash_app
from parsl.configs.local_threads import config

parsl.load(config)

import sys
import os,sys,inspect
currentdir = os.path.dirname(os.path.abspath(inspect.getfile(inspect.currentframe())))
parentdir = os.path.dirname(currentdir)
sys.path.insert(0,parentdir)
import userScript

classMarks = userScript.classMarks
outputLocation = userScript.outputLocation

mulTwoArray = []

# open file and read the content in a list
with open(outputLocation+'multiplyByTwo.txt', 'r') as f:
    for line in f:
        # remove linebreak which is the last character of the string
        mark = line[:-1]
        # add item to the list
        mulTwoArray.append(mark)

studentTotal = 0
for j in mulTwoArray:
	studentTotal += int(j)

classMarks.append(studentTotal)

@python_app
def classAverage ():
    average = sum(classMarks) / len(classMarks)
    return average

print(classAverage().result())

print("Module Completed : Calculate Class Average")

