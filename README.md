# SciFlow
## A Compositional Framework for Scientific Workflows
Let us explore the possibilities provided by SciFlow with a simple example. 

Mrs. Branson, Central High School’s Mathematics teacher wants to create a portal where her students could inquire whether they performed well during a semester or not. Student performance is calculated based on 10 different assignments given throughout the semester and marks are given out of 50 for each. 

Mrs. Branson has the following data with her;
* Student Name. 
*Eg : “John”*
* List of student’s marks for the 10 assignments, out of 50. 
*Eg : [23,45,26,47,28,38,46,43,32,44]*
* The total assignment marks for the rest of the students, out of 1000. *Eg : [927,734,358,212,938,603,912,592,657,556,740,625,495,864,346,623,824,543,951,591,833,340,690,544,797,818,847,791,727,652,856,651,686,964,765,732,662,674,714]*

The SciFlow framework could be easily used to construct a workflow in order to achieve Mrs. Branson’s vision. 

![Simple Workflow](/images/simpleWorkflow.png)

1. The student would be greeted first, by printing “Hello John!”
2. Thereafter, each of his marks would be multiplied by two, using parallel execution.
3. Then the class average would be calculated.
4. Based on a comparison between John’s total marks and the class average, the system would tell John whether he had an amazing/terrible semester

This example emphasises several key features provided by SciFlow;
* **Dynamic execution** - The control thread would make a decision on the path to take, going forward. In the example provided, the control thread would compare John’s mark with the class average and adjust the path accordingly.
* **Implicit Parallelism** - Multiplying John’s marks is performed parallely. Each mark in the array is handled parallely using the Parsl library. 
* **High Performance Computing** - Mrs. Branson could easily configure the systems to run on her machine’s threads, cores or even a cluster, by simply loading a pre-written configuration file.

All this could be achieved making minor changes to provided files.
* The user variables such as, the student’s name, should be included in *“userScript.py”*
* Individual workflow modules go inside the *“workflow”* folder. You can break the workflow into several components if you wish. Do so by adding a subfolder within the workflow folder. 
* Load the pre-written configurations using *“parslConfig.py”*
* Where appropriate, you can parallelize individual modules. Take a look at *“/workflow/WFComponent2/multiplyByTwo.py”* to get some guidance on how to do this. A Parsl parallelizable python function is written as multiplyByTwo(x) with the @python_app annotation. This function is called on the studentMarks array, executing parallely.
* The workflow is constructed in the main function of the *“controlThread.go”* file. The commands array is captured from the “userScript.py” file, where we assign a number to each workflow module. The control thread executes the module based on this numbering.

SciFlow can be used easily for more complex scenario, our website includes projects for;
* A dynamic Data Analytics workflow 
* A Particle Swarm Optimization solution to the Travelling Salesman Problem

For further information on how to use the SciFlow framework, make sure to check out our user guide at https://sciflow-fyp.github.io
