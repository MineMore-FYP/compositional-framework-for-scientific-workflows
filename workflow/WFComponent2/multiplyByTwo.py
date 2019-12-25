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

currentModule = "multiplyByTwo"

orderOfModules = userScript.orderOfModules
studentMarks = userScript.studentMarks
classMarks = userScript.classMarks
outputLocation = userScript.outputLocation

@python_app
def multiplyByTwo (x):
    multiplyByTwo = x*2
    return multiplyByTwo

mulTwoArray = []
for i in studentMarks:
	mulTwoArray.append(multiplyByTwo(i).result())

with open(outputLocation+currentModule+'.csv', 'w') as f:
    for k in mulTwoArray:
        f.write("%s\n" % k)

print("Module Completed : Multiply Marks by 2 Complete")

