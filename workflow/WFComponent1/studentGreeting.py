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

currentModule = "studentGreeting"

orderOfModules = userScript.orderOfModules
studentName = userScript.studentName

@python_app
def helloFunc ():
    text = "Hello " + studentName
    return text

print(helloFunc().result())
