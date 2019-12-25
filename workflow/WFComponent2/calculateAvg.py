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

currentModule = "module1"

workflowNumber = sys.argv[1]

if workflowNumber == "1":
	orderOfModules = userScript.orderOfModules1
	classMarks = userScript.classMarks1
	marks = userScript.marks1

@python_app
def classAverage ():
    classMarks.append(marks)
    average = sum(classMarks) / len(classMarks)
    return average

print(classAverage().result())

print("Module Completed : Calculate Class Average")
