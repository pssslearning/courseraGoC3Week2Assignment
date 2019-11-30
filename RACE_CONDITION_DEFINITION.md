## Concurrency in Go - CONCURRENCY BASICS - Assigment Week #2 


- [Assignment](https://www.coursera.org/learn/golang-concurrency/peer/RAJ0V/module-2-activity)


- [Course: Concurrency in Go](https://www.coursera.org/learn/golang-concurrency/home/welcome)
  
## RACE CONDITION DEFINITION:


A **`RACE CONDITION`** (a.k.a `RACE HAZARD`) in software systems is ***the condition where the system's substantive 
behavior is dependent on the sequence or timing of interleaving actions***.

Race conditions arise in software when an application depends on the sequence or timing of processes or threads
for it to operate properly:

- Interleaving is non-deterministic, consequently the outcome depends on non-deterministic ordering.
- Races occur due to communication (access to shared data/variables) among concurrently running threads
