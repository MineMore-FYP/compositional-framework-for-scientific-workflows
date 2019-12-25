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

currentModule = "terrible"

orderOfModules = userScript.orderOfModules
durationType = userScript.durationType

@python_app
def terrible ():
    text = "You had a terrible "+ durationType
    return text

print(terrible().result())

print("Module Completed : Terrible :'(")
