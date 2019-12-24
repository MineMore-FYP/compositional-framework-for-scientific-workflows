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

currentModule = "module2"

workflowNumber = sys.argv[1]

if workflowNumber == "1":
	orderOfModules = userScript.orderOfModules1
	world = userScript.world1
elif workflowNumber == "2":
	orderOfModules = userScript.orderOfModules2
	world = userScript.world2

@python_app
def helloFunc ():
    text = world + " World!"
    return text

print(helloFunc().result())
