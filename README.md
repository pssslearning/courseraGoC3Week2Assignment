# courseraGoC3Week2Assignment-

Assigment Week #2 - Concurrency in Go - CONCURRENCY BASICS - Course https://www.coursera.org/learn/golang-concurrency/home/week/2


Assigment Week #2 - Concurrency in Go - CONCURRENCY BASICS  

- [Assignment](https://www.coursera.org/learn/golang-concurrency/peer/RAJ0V/module-2-activity)


- [Course: Concurrency in Go](https://www.coursera.org/learn/golang-concurrency/home/welcome)
  
## Instructions

The goal of this activity is to explore race conditions by creating and running two simultaneous goroutines.

Write two goroutines which have a race condition when executed concurrently. Explain what the race condition is and how it can occur.

## My assignment

Source code at `racecondition/racecondition.go`

## Sample compilation and first test execution

```sh
devel1@vbxdeb10mate:~$ cd $GOPATH/src/github.com/pssslearning/courseraGoC3Week2Assignment/racecondition/
devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning/courseraGoC3Week2Assignment/racecondition$ go build racecondition.go 
```

### 1st. try

```sh
devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning/courseraGoC3Week2Assignment/racecondition$ ./racecondition 

============================================================================================================
WELCOME !!!
This program aims to show an example of the appearance of a RACE CONDITION. What does it means?

RACE CONDITION DEFINITION:
A 'RACE CONDITION' (a.k.a 'RACE HAZARD') in software systems is the condition where the system's substantive 
behavior is dependent on the sequence or timing of interleaving actions.
Race conditions arise in software when an application depends on the sequence or timing of processes or threads
for it to operate properly:
  ==> Interleaving is non-deterministic, consequently the outcome depends on non-deterministic ordering.
  ==> Races occur due to communication (access to shared data/variables) among concurrently running threads
============================================================================================================

Press ENTER to continue.

============================================================================================================
LET'S SHOW THAT WITH AN EXAMPLE !!!
This program aims to show an example of the emergence of a RACE CONDITION when more than one thread is being
run concurrently and having simultaneous read and write access to a zone of shared (variable) data.
In the proposed example, a situation is simulated where there is an initial balance of an account and also a
list of pending transactions to be applied on. The process to be followed is to increase the position in the
list (cursor), take the value of the correspondig movement in the list and update the current account balance
with its value. When the entire list of pending movements has been traversed, the resulting final balance is
to be shown.
If you carry out the test to execute this program successively, you will be able to observe that when the
RACE CONDITION (badly managed) occurs, different results will be obtained each time: 
--> UNDETERMINISTIC OUTCOME
============================================================================================================

Press ENTER to continue.

============================================================================================================
STAGE #01:
First it will be shown a completely correct behaviour when a single thread is used (the main Goroutine).
Once applied all movements the resulting balance should be [120.30].
============================================================================================================

Press ENTER to continue.
=========================================================
Initial Balance [300.00].
=========================================================

		--> Entering Goroutine [MAIN Goroutine]. Theoretical initial Account Balance[300.00] (2019-11-30 11:07:27.628778641)

[MAIN Goroutine].	Theroretical Current Balance[300.00].	Account Movement[#0][40.50].	Updated Balance[340.50] (2019-11-30 11:07:27.628975818)
[MAIN Goroutine].	Theroretical Current Balance[340.50].	Account Movement[#1][-20.72].	Updated Balance[319.78] (2019-11-30 11:07:27.629020597)
[MAIN Goroutine].	Theroretical Current Balance[319.78].	Account Movement[#2][89.01].	Updated Balance[408.79] (2019-11-30 11:07:27.629030376)
[MAIN Goroutine].	Theroretical Current Balance[408.79].	Account Movement[#3][60.22].	Updated Balance[469.01] (2019-11-30 11:07:27.629043614)
[MAIN Goroutine].	Theroretical Current Balance[469.01].	Account Movement[#4][-96.02].	Updated Balance[372.99] (2019-11-30 11:07:27.629052565)
[MAIN Goroutine].	Theroretical Current Balance[372.99].	Account Movement[#5][7.42].	Updated Balance[380.41] (2019-11-30 11:07:27.629060352)
[MAIN Goroutine].	Theroretical Current Balance[380.41].	Account Movement[#6][73.79].	Updated Balance[454.20] (2019-11-30 11:07:27.629068473)
[MAIN Goroutine].	Theroretical Current Balance[454.20].	Account Movement[#7][96.24].	Updated Balance[550.44] (2019-11-30 11:07:27.629076461)
[MAIN Goroutine].	Theroretical Current Balance[550.44].	Account Movement[#8][76.81].	Updated Balance[627.25] (2019-11-30 11:07:27.629089586)
[MAIN Goroutine].	Theroretical Current Balance[627.25].	Account Movement[#9][-32.98].	Updated Balance[594.27] (2019-11-30 11:07:27.629097769)
[MAIN Goroutine].	Theroretical Current Balance[594.27].	Account Movement[#10][24.76].	Updated Balance[619.03] (2019-11-30 11:07:27.629110551)
[MAIN Goroutine].	Theroretical Current Balance[619.03].	Account Movement[#11][17.48].	Updated Balance[636.51] (2019-11-30 11:07:27.629119386)
[MAIN Goroutine].	Theroretical Current Balance[636.51].	Account Movement[#12][-11.50].	Updated Balance[625.01] (2019-11-30 11:07:27.629127210)
[MAIN Goroutine].	Theroretical Current Balance[625.01].	Account Movement[#13][94.34].	Updated Balance[719.35] (2019-11-30 11:07:27.629172648)
[MAIN Goroutine].	Theroretical Current Balance[719.35].	Account Movement[#14][43.06].	Updated Balance[762.41] (2019-11-30 11:07:27.629182635)
[MAIN Goroutine].	Theroretical Current Balance[762.41].	Account Movement[#15][56.37].	Updated Balance[818.78] (2019-11-30 11:07:27.629190394)
[MAIN Goroutine].	Theroretical Current Balance[818.78].	Account Movement[#16][-29.41].	Updated Balance[789.37] (2019-11-30 11:07:27.629198210)
[MAIN Goroutine].	Theroretical Current Balance[789.37].	Account Movement[#17][-51.32].	Updated Balance[738.05] (2019-11-30 11:07:27.629206230)
[MAIN Goroutine].	Theroretical Current Balance[738.05].	Account Movement[#18][-53.99].	Updated Balance[684.06] (2019-11-30 11:07:27.629218646)
[MAIN Goroutine].	Theroretical Current Balance[684.06].	Account Movement[#19][-41.77].	Updated Balance[642.29] (2019-11-30 11:07:27.629234354)
[MAIN Goroutine].	Theroretical Current Balance[642.29].	Account Movement[#20][-33.04].	Updated Balance[609.25] (2019-11-30 11:07:27.629276387)
[MAIN Goroutine].	Theroretical Current Balance[609.25].	Account Movement[#21][-49.18].	Updated Balance[560.07] (2019-11-30 11:07:27.629306167)
[MAIN Goroutine].	Theroretical Current Balance[560.07].	Account Movement[#22][84.89].	Updated Balance[644.96] (2019-11-30 11:07:27.629316226)
[MAIN Goroutine].	Theroretical Current Balance[644.96].	Account Movement[#23][-49.91].	Updated Balance[595.05] (2019-11-30 11:07:27.629324070)
[MAIN Goroutine].	Theroretical Current Balance[595.05].	Account Movement[#24][34.41].	Updated Balance[629.46] (2019-11-30 11:07:27.629331778)
[MAIN Goroutine].	Theroretical Current Balance[629.46].	Account Movement[#25][-8.27].	Updated Balance[621.19] (2019-11-30 11:07:27.629339721)
[MAIN Goroutine].	Theroretical Current Balance[621.19].	Account Movement[#26][-83.34].	Updated Balance[537.85] (2019-11-30 11:07:27.629352455)
[MAIN Goroutine].	Theroretical Current Balance[537.85].	Account Movement[#27][-84.29].	Updated Balance[453.56] (2019-11-30 11:07:27.629360626)
[MAIN Goroutine].	Theroretical Current Balance[453.56].	Account Movement[#28][1.71].	Updated Balance[455.27] (2019-11-30 11:07:27.629373472)
[MAIN Goroutine].	Theroretical Current Balance[455.27].	Account Movement[#29][-75.72].	Updated Balance[379.55] (2019-11-30 11:07:27.629382708)
[MAIN Goroutine].	Theroretical Current Balance[379.55].	Account Movement[#30][-73.64].	Updated Balance[305.91] (2019-11-30 11:07:27.629390659)
[MAIN Goroutine].	Theroretical Current Balance[305.91].	Account Movement[#31][-97.00].	Updated Balance[208.91] (2019-11-30 11:07:27.629404265)
[MAIN Goroutine].	Theroretical Current Balance[208.91].	Account Movement[#32][-94.03].	Updated Balance[114.88] (2019-11-30 11:07:27.629413134)
[MAIN Goroutine].	Theroretical Current Balance[114.88].	Account Movement[#33][-15.48].	Updated Balance[99.40] (2019-11-30 11:07:27.629420744)
[MAIN Goroutine].	Theroretical Current Balance[99.40].	Account Movement[#34][0.95].	Updated Balance[100.35] (2019-11-30 11:07:27.629433255)
[MAIN Goroutine].	Theroretical Current Balance[100.35].	Account Movement[#35][18.02].	Updated Balance[118.37] (2019-11-30 11:07:27.629442408)
[MAIN Goroutine].	Theroretical Current Balance[118.37].	Account Movement[#36][70.39].	Updated Balance[188.76] (2019-11-30 11:07:27.629450318)
[MAIN Goroutine].	Theroretical Current Balance[188.76].	Account Movement[#37][-6.15].	Updated Balance[182.61] (2019-11-30 11:07:27.629458240)
[MAIN Goroutine].	Theroretical Current Balance[182.61].	Account Movement[#38][-80.81].	Updated Balance[101.80] (2019-11-30 11:07:27.629473739)
[MAIN Goroutine].	Theroretical Current Balance[101.80].	Account Movement[#39][11.00].	Updated Balance[112.80] (2019-11-30 11:07:27.629517690)
[MAIN Goroutine].	Theroretical Current Balance[112.80].	Account Movement[#40][-3.26].	Updated Balance[109.54] (2019-11-30 11:07:27.629587037)
[MAIN Goroutine].	Theroretical Current Balance[109.54].	Account Movement[#41][81.66].	Updated Balance[191.20] (2019-11-30 11:07:27.629622052)
[MAIN Goroutine].	Theroretical Current Balance[191.20].	Account Movement[#42][58.06].	Updated Balance[249.26] (2019-11-30 11:07:27.629631873)
[MAIN Goroutine].	Theroretical Current Balance[249.26].	Account Movement[#43][96.79].	Updated Balance[346.05] (2019-11-30 11:07:27.629640129)
[MAIN Goroutine].	Theroretical Current Balance[346.05].	Account Movement[#44][-97.54].	Updated Balance[248.51] (2019-11-30 11:07:27.629653953)
[MAIN Goroutine].	Theroretical Current Balance[248.51].	Account Movement[#45][-97.97].	Updated Balance[150.54] (2019-11-30 11:07:27.629662871)
[MAIN Goroutine].	Theroretical Current Balance[150.54].	Account Movement[#46][45.18].	Updated Balance[195.72] (2019-11-30 11:07:27.629670584)
[MAIN Goroutine].	Theroretical Current Balance[195.72].	Account Movement[#47][4.75].	Updated Balance[200.47] (2019-11-30 11:07:27.629692009)
[MAIN Goroutine].	Theroretical Current Balance[200.47].	Account Movement[#48][-84.80].	Updated Balance[115.67] (2019-11-30 11:07:27.629726205)
[MAIN Goroutine].	Theroretical Current Balance[115.67].	Account Movement[#49][75.03].	Updated Balance[190.70] (2019-11-30 11:07:27.629736472)
[MAIN Goroutine].	Theroretical Current Balance[190.70].	Account Movement[#50][31.05].	Updated Balance[221.75] (2019-11-30 11:07:27.629744894)
[MAIN Goroutine].	Theroretical Current Balance[221.75].	Account Movement[#51][-83.56].	Updated Balance[138.19] (2019-11-30 11:07:27.629758673)
[MAIN Goroutine].	Theroretical Current Balance[138.19].	Account Movement[#52][15.40].	Updated Balance[153.59] (2019-11-30 11:07:27.629768064)
[MAIN Goroutine].	Theroretical Current Balance[153.59].	Account Movement[#53][-43.92].	Updated Balance[109.67] (2019-11-30 11:07:27.629776295)
[MAIN Goroutine].	Theroretical Current Balance[109.67].	Account Movement[#54][16.62].	Updated Balance[126.29] (2019-11-30 11:07:27.629806085)
[MAIN Goroutine].	Theroretical Current Balance[126.29].	Account Movement[#55][66.13].	Updated Balance[192.42] (2019-11-30 11:07:27.629815698)
[MAIN Goroutine].	Theroretical Current Balance[192.42].	Account Movement[#56][38.67].	Updated Balance[231.09] (2019-11-30 11:07:27.629824246)
[MAIN Goroutine].	Theroretical Current Balance[231.09].	Account Movement[#57][-80.34].	Updated Balance[150.75] (2019-11-30 11:07:27.629837864)
[MAIN Goroutine].	Theroretical Current Balance[150.75].	Account Movement[#58][21.61].	Updated Balance[172.36] (2019-11-30 11:07:27.629852754)
[MAIN Goroutine].	Theroretical Current Balance[172.36].	Account Movement[#59][42.01].	Updated Balance[214.37] (2019-11-30 11:07:27.629861649)
[MAIN Goroutine].	Theroretical Current Balance[214.37].	Account Movement[#60][-94.07].	Updated Balance[120.30] (2019-11-30 11:07:27.629875426)

		--> Exiting from Goroutine [MAIN Goroutine]. Theoretical resulting Account Balance[120.30] (2019-11-30 11:07:27.629885934)

=========================================================
Resulting Balance [120.30].
=========================================================
Press ENTER to continue.

============================================================================================================
STAGE #02:
Now we will reset both values: the Cursor (position) of the list of pending movements and the value of
the initial Balance to their starting values
Then two threads (Goroutines) will be launched and execute concurrently creating a RACE CONDITION as:
1. The two threads will access concurrently trying to increase the shared value of the cursor of the
   movement list and updating the shared value of the balance with the corresponding movement of the list
2. Running this way final will drive to a totally different final Balance result each time
 
Note: now the main thread will be waiting sleeping during a reasonable time to permit all other threads
to succesfully conclude its tasks. DISCLAIMER: There are better way to sync. threads but they have not
shown/explained at this point in the course.
============================================================================================================

Press ENTER to continue.

=========================================================
RESETTING Position[0] and Initial Account Balance[300.00]
=========================================================

=========================================================
Initial Balance [300.00].
=========================================================
Main Thread is currently Awaiting [10] seconds to let all goroutines to safely end ...

		--> Entering Goroutine [Goroutine #01]. Theoretical initial Account Balance[300.00] (2019-11-30 11:07:28.892796362)

[Goroutine #01].	Theroretical Current Balance[300.00].	Account Movement[#0][40.50].	Updated Balance[340.50] (2019-11-30 11:07:28.892839634)

		--> Entering Goroutine [Goroutine #02]. Theoretical initial Account Balance[300.00] (2019-11-30 11:07:28.892819166)

[Goroutine #02].	Theroretical Current Balance[340.50].	Account Movement[#1][-20.72].	Updated Balance[319.78] (2019-11-30 11:07:28.892906507)
[Goroutine #02].	Theroretical Current Balance[319.78].	Account Movement[#2][89.01].	Updated Balance[408.79] (2019-11-30 11:07:28.903423237)
[Goroutine #01].	Theroretical Current Balance[408.79].	Account Movement[#3][60.22].	Updated Balance[469.01] (2019-11-30 11:07:28.952648122)
[Goroutine #02].	Theroretical Current Balance[469.01].	Account Movement[#4][-96.02].	Updated Balance[372.99] (2019-11-30 11:07:28.956165723)
[Goroutine #02].	Theroretical Current Balance[372.99].	Account Movement[#5][7.42].	Updated Balance[380.41] (2019-11-30 11:07:28.999943524)
[Goroutine #01].	Theroretical Current Balance[372.99].	Account Movement[#5][7.42].	Updated Balance[380.41] (2019-11-30 11:07:29.000092831)
[Goroutine #01].	Theroretical Current Balance[380.41].	Account Movement[#7][96.24].	Updated Balance[476.65] (2019-11-30 11:07:29.021175645)
[Goroutine #02].	Theroretical Current Balance[476.65].	Account Movement[#8][76.81].	Updated Balance[553.46] (2019-11-30 11:07:29.063441719)
[Goroutine #01].	Theroretical Current Balance[553.46].	Account Movement[#9][-32.98].	Updated Balance[520.48] (2019-11-30 11:07:29.065276709)
[Goroutine #02].	Theroretical Current Balance[520.48].	Account Movement[#10][24.76].	Updated Balance[545.24] (2019-11-30 11:07:29.068141318)
[Goroutine #01].	Theroretical Current Balance[545.24].	Account Movement[#11][17.48].	Updated Balance[562.72] (2019-11-30 11:07:29.114071990)
[Goroutine #02].	Theroretical Current Balance[562.72].	Account Movement[#12][-11.50].	Updated Balance[551.22] (2019-11-30 11:07:29.130913644)
[Goroutine #01].	Theroretical Current Balance[551.22].	Account Movement[#13][94.34].	Updated Balance[645.56] (2019-11-30 11:07:29.153981312)
[Goroutine #02].	Theroretical Current Balance[645.56].	Account Movement[#14][43.06].	Updated Balance[688.62] (2019-11-30 11:07:29.185768237)
[Goroutine #01].	Theroretical Current Balance[688.62].	Account Movement[#15][56.37].	Updated Balance[744.99] (2019-11-30 11:07:29.207059353)
[Goroutine #02].	Theroretical Current Balance[744.99].	Account Movement[#16][-29.41].	Updated Balance[715.58] (2019-11-30 11:07:29.208212876)
[Goroutine #01].	Theroretical Current Balance[715.58].	Account Movement[#17][-51.32].	Updated Balance[664.26] (2019-11-30 11:07:29.242731243)
[Goroutine #02].	Theroretical Current Balance[664.26].	Account Movement[#18][-53.99].	Updated Balance[610.27] (2019-11-30 11:07:29.251559568)
[Goroutine #02].	Theroretical Current Balance[610.27].	Account Movement[#19][-41.77].	Updated Balance[568.50] (2019-11-30 11:07:29.261358493)
[Goroutine #02].	Theroretical Current Balance[568.50].	Account Movement[#20][-33.04].	Updated Balance[535.46] (2019-11-30 11:07:29.286717627)
[Goroutine #01].	Theroretical Current Balance[535.46].	Account Movement[#21][-49.18].	Updated Balance[486.28] (2019-11-30 11:07:29.290369213)
[Goroutine #01].	Theroretical Current Balance[486.28].	Account Movement[#22][84.89].	Updated Balance[571.17] (2019-11-30 11:07:29.322252757)
[Goroutine #02].	Theroretical Current Balance[571.17].	Account Movement[#23][-49.91].	Updated Balance[521.26] (2019-11-30 11:07:29.340648268)
[Goroutine #02].	Theroretical Current Balance[521.26].	Account Movement[#24][34.41].	Updated Balance[555.67] (2019-11-30 11:07:29.374469126)
[Goroutine #01].	Theroretical Current Balance[555.67].	Account Movement[#25][-8.27].	Updated Balance[547.40] (2019-11-30 11:07:29.381013816)
[Goroutine #02].	Theroretical Current Balance[547.40].	Account Movement[#26][-83.34].	Updated Balance[464.06] (2019-11-30 11:07:29.386064644)
[Goroutine #01].	Theroretical Current Balance[464.06].	Account Movement[#27][-84.29].	Updated Balance[379.77] (2019-11-30 11:07:29.450072147)
[Goroutine #02].	Theroretical Current Balance[379.77].	Account Movement[#28][1.71].	Updated Balance[381.48] (2019-11-30 11:07:29.457974662)
[Goroutine #02].	Theroretical Current Balance[381.48].	Account Movement[#29][-75.72].	Updated Balance[305.76] (2019-11-30 11:07:29.465461236)
[Goroutine #02].	Theroretical Current Balance[305.76].	Account Movement[#30][-73.64].	Updated Balance[232.12] (2019-11-30 11:07:29.478254064)
[Goroutine #02].	Theroretical Current Balance[232.12].	Account Movement[#31][-97.00].	Updated Balance[135.12] (2019-11-30 11:07:29.499927883)
[Goroutine #01].	Theroretical Current Balance[135.12].	Account Movement[#32][-94.03].	Updated Balance[41.09] (2019-11-30 11:07:29.507357265)
[Goroutine #02].	Theroretical Current Balance[41.09].	Account Movement[#33][-15.48].	Updated Balance[25.61] (2019-11-30 11:07:29.524017585)
[Goroutine #02].	Theroretical Current Balance[25.61].	Account Movement[#34][0.95].	Updated Balance[26.56] (2019-11-30 11:07:29.569332721)
[Goroutine #01].	Theroretical Current Balance[26.56].	Account Movement[#35][18.02].	Updated Balance[44.58] (2019-11-30 11:07:29.573277010)
[Goroutine #02].	Theroretical Current Balance[44.58].	Account Movement[#36][70.39].	Updated Balance[114.97] (2019-11-30 11:07:29.619349627)
[Goroutine #01].	Theroretical Current Balance[114.97].	Account Movement[#37][-6.15].	Updated Balance[108.82] (2019-11-30 11:07:29.642432014)
[Goroutine #02].	Theroretical Current Balance[108.82].	Account Movement[#38][-80.81].	Updated Balance[28.01] (2019-11-30 11:07:29.669407428)
[Goroutine #01].	Theroretical Current Balance[28.01].	Account Movement[#39][11.00].	Updated Balance[39.01] (2019-11-30 11:07:29.724983097)
[Goroutine #02].	Theroretical Current Balance[39.01].	Account Movement[#40][-3.26].	Updated Balance[35.75] (2019-11-30 11:07:29.727772256)
[Goroutine #01].	Theroretical Current Balance[35.75].	Account Movement[#41][81.66].	Updated Balance[117.41] (2019-11-30 11:07:29.790346515)
[Goroutine #02].	Theroretical Current Balance[117.41].	Account Movement[#42][58.06].	Updated Balance[175.47] (2019-11-30 11:07:29.791586816)
[Goroutine #02].	Theroretical Current Balance[175.47].	Account Movement[#43][96.79].	Updated Balance[272.26] (2019-11-30 11:07:29.802427333)
[Goroutine #02].	Theroretical Current Balance[272.26].	Account Movement[#44][-97.54].	Updated Balance[174.72] (2019-11-30 11:07:29.818349946)
[Goroutine #01].	Theroretical Current Balance[174.72].	Account Movement[#45][-97.97].	Updated Balance[76.75] (2019-11-30 11:07:29.837716244)
[Goroutine #02].	Theroretical Current Balance[76.75].	Account Movement[#46][45.18].	Updated Balance[121.93] (2019-11-30 11:07:29.837864068)
[Goroutine #02].	Theroretical Current Balance[121.93].	Account Movement[#47][4.75].	Updated Balance[126.68] (2019-11-30 11:07:29.887092220)
[Goroutine #01].	Theroretical Current Balance[126.68].	Account Movement[#48][-84.80].	Updated Balance[41.88] (2019-11-30 11:07:29.888363187)
[Goroutine #02].	Theroretical Current Balance[41.88].	Account Movement[#49][75.03].	Updated Balance[116.91] (2019-11-30 11:07:29.936860410)
[Goroutine #01].	Theroretical Current Balance[41.88].	Account Movement[#49][75.03].	Updated Balance[116.91] (2019-11-30 11:07:29.936851052)
[Goroutine #01].	Theroretical Current Balance[116.91].	Account Movement[#51][-83.56].	Updated Balance[33.35] (2019-11-30 11:07:29.954277460)
[Goroutine #02].	Theroretical Current Balance[33.35].	Account Movement[#52][15.40].	Updated Balance[48.75] (2019-11-30 11:07:29.958072904)
[Goroutine #02].	Theroretical Current Balance[48.75].	Account Movement[#53][-43.92].	Updated Balance[4.83] (2019-11-30 11:07:29.994043756)
[Goroutine #01].	Theroretical Current Balance[4.83].	Account Movement[#54][16.62].	Updated Balance[21.45] (2019-11-30 11:07:30.035990419)
[Goroutine #02].	Theroretical Current Balance[21.45].	Account Movement[#55][66.13].	Updated Balance[87.58] (2019-11-30 11:07:30.064359422)
[Goroutine #01].	Theroretical Current Balance[87.58].	Account Movement[#56][38.67].	Updated Balance[126.25] (2019-11-30 11:07:30.097780541)
[Goroutine #02].	Theroretical Current Balance[126.25].	Account Movement[#57][-80.34].	Updated Balance[45.91] (2019-11-30 11:07:30.120352505)
[Goroutine #02].	Theroretical Current Balance[45.91].	Account Movement[#58][21.61].	Updated Balance[67.52] (2019-11-30 11:07:30.125834522)
[Goroutine #01].	Theroretical Current Balance[67.52].	Account Movement[#59][42.01].	Updated Balance[109.53] (2019-11-30 11:07:30.157819395)
[Goroutine #02].	Theroretical Current Balance[109.53].	Account Movement[#60][-94.07].	Updated Balance[15.46] (2019-11-30 11:07:30.190717676)

		--> Exiting from Goroutine [Goroutine #01]. Theoretical resulting Account Balance[109.53] (2019-11-30 11:07:30.200466544)


		--> Exiting from Goroutine [Goroutine #02]. Theoretical resulting Account Balance[15.46] (2019-11-30 11:07:30.258268387)

=========================================================
Resulting Balance [15.46].
=========================================================
Press ENTER to continue.

============================================================================================================
STAGE #03:
Each thread has been writing in a log its activity showing at every moment in time the values that 'see'
and update the shared variables: Cursor and Current balance.
Now it's time to see how they were behaving individually.
============================================================================================================

Press ENTER to continue.

========================================================================
Showing now the event's history registered at Goroutine [Goroutine #01]
========================================================================
------------------------------------------------------------------------------------------------------------------------------------------------------
		--> Entering Goroutine [Goroutine #01]. Theoretical initial Account Balance[300.00] (2019-11-30 11:07:28.892796362)
------------------------------------------------------------------------------------------------------------------------------------------------------
[Goroutine #01].	Theroretical Current Balance[300.00].	Account Movement[#0][40.50].	Updated Balance[340.50] (2019-11-30 11:07:28.892839634)
[Goroutine #01].	Theroretical Current Balance[408.79].	Account Movement[#3][60.22].	Updated Balance[469.01] (2019-11-30 11:07:28.952648122)
[Goroutine #01].	Theroretical Current Balance[372.99].	Account Movement[#5][7.42].	Updated Balance[380.41] (2019-11-30 11:07:29.000092831)
[Goroutine #01].	Theroretical Current Balance[380.41].	Account Movement[#7][96.24].	Updated Balance[476.65] (2019-11-30 11:07:29.021175645)
[Goroutine #01].	Theroretical Current Balance[553.46].	Account Movement[#9][-32.98].	Updated Balance[520.48] (2019-11-30 11:07:29.065276709)
[Goroutine #01].	Theroretical Current Balance[545.24].	Account Movement[#11][17.48].	Updated Balance[562.72] (2019-11-30 11:07:29.114071990)
[Goroutine #01].	Theroretical Current Balance[551.22].	Account Movement[#13][94.34].	Updated Balance[645.56] (2019-11-30 11:07:29.153981312)
[Goroutine #01].	Theroretical Current Balance[688.62].	Account Movement[#15][56.37].	Updated Balance[744.99] (2019-11-30 11:07:29.207059353)
[Goroutine #01].	Theroretical Current Balance[715.58].	Account Movement[#17][-51.32].	Updated Balance[664.26] (2019-11-30 11:07:29.242731243)
[Goroutine #01].	Theroretical Current Balance[535.46].	Account Movement[#21][-49.18].	Updated Balance[486.28] (2019-11-30 11:07:29.290369213)
[Goroutine #01].	Theroretical Current Balance[486.28].	Account Movement[#22][84.89].	Updated Balance[571.17] (2019-11-30 11:07:29.322252757)
[Goroutine #01].	Theroretical Current Balance[555.67].	Account Movement[#25][-8.27].	Updated Balance[547.40] (2019-11-30 11:07:29.381013816)
[Goroutine #01].	Theroretical Current Balance[464.06].	Account Movement[#27][-84.29].	Updated Balance[379.77] (2019-11-30 11:07:29.450072147)
[Goroutine #01].	Theroretical Current Balance[135.12].	Account Movement[#32][-94.03].	Updated Balance[41.09] (2019-11-30 11:07:29.507357265)
[Goroutine #01].	Theroretical Current Balance[26.56].	Account Movement[#35][18.02].	Updated Balance[44.58] (2019-11-30 11:07:29.573277010)
[Goroutine #01].	Theroretical Current Balance[114.97].	Account Movement[#37][-6.15].	Updated Balance[108.82] (2019-11-30 11:07:29.642432014)
[Goroutine #01].	Theroretical Current Balance[28.01].	Account Movement[#39][11.00].	Updated Balance[39.01] (2019-11-30 11:07:29.724983097)
[Goroutine #01].	Theroretical Current Balance[35.75].	Account Movement[#41][81.66].	Updated Balance[117.41] (2019-11-30 11:07:29.790346515)
[Goroutine #01].	Theroretical Current Balance[174.72].	Account Movement[#45][-97.97].	Updated Balance[76.75] (2019-11-30 11:07:29.837716244)
[Goroutine #01].	Theroretical Current Balance[126.68].	Account Movement[#48][-84.80].	Updated Balance[41.88] (2019-11-30 11:07:29.888363187)
[Goroutine #01].	Theroretical Current Balance[41.88].	Account Movement[#49][75.03].	Updated Balance[116.91] (2019-11-30 11:07:29.936851052)
[Goroutine #01].	Theroretical Current Balance[116.91].	Account Movement[#51][-83.56].	Updated Balance[33.35] (2019-11-30 11:07:29.954277460)
[Goroutine #01].	Theroretical Current Balance[4.83].	Account Movement[#54][16.62].	Updated Balance[21.45] (2019-11-30 11:07:30.035990419)
[Goroutine #01].	Theroretical Current Balance[87.58].	Account Movement[#56][38.67].	Updated Balance[126.25] (2019-11-30 11:07:30.097780541)
[Goroutine #01].	Theroretical Current Balance[67.52].	Account Movement[#59][42.01].	Updated Balance[109.53] (2019-11-30 11:07:30.157819395)
------------------------------------------------------------------------------------------------------------------------------------------------------
		--> Exiting from Goroutine [Goroutine #01]. Theoretical resulting Account Balance[109.53] (2019-11-30 11:07:30.200466544)
------------------------------------------------------------------------------------------------------------------------------------------------------
Press ENTER to continue.

========================================================================
Showing now the event's history registered at Goroutine [Goroutine #02]
========================================================================
------------------------------------------------------------------------------------------------------------------------------------------------------
		--> Entering Goroutine [Goroutine #02]. Theoretical initial Account Balance[300.00] (2019-11-30 11:07:28.892819166)
------------------------------------------------------------------------------------------------------------------------------------------------------
[Goroutine #02].	Theroretical Current Balance[340.50].	Account Movement[#1][-20.72].	Updated Balance[319.78] (2019-11-30 11:07:28.892906507)
[Goroutine #02].	Theroretical Current Balance[319.78].	Account Movement[#2][89.01].	Updated Balance[408.79] (2019-11-30 11:07:28.903423237)
[Goroutine #02].	Theroretical Current Balance[469.01].	Account Movement[#4][-96.02].	Updated Balance[372.99] (2019-11-30 11:07:28.956165723)
[Goroutine #02].	Theroretical Current Balance[372.99].	Account Movement[#5][7.42].	Updated Balance[380.41] (2019-11-30 11:07:28.999943524)
[Goroutine #02].	Theroretical Current Balance[476.65].	Account Movement[#8][76.81].	Updated Balance[553.46] (2019-11-30 11:07:29.063441719)
[Goroutine #02].	Theroretical Current Balance[520.48].	Account Movement[#10][24.76].	Updated Balance[545.24] (2019-11-30 11:07:29.068141318)
[Goroutine #02].	Theroretical Current Balance[562.72].	Account Movement[#12][-11.50].	Updated Balance[551.22] (2019-11-30 11:07:29.130913644)
[Goroutine #02].	Theroretical Current Balance[645.56].	Account Movement[#14][43.06].	Updated Balance[688.62] (2019-11-30 11:07:29.185768237)
[Goroutine #02].	Theroretical Current Balance[744.99].	Account Movement[#16][-29.41].	Updated Balance[715.58] (2019-11-30 11:07:29.208212876)
[Goroutine #02].	Theroretical Current Balance[664.26].	Account Movement[#18][-53.99].	Updated Balance[610.27] (2019-11-30 11:07:29.251559568)
[Goroutine #02].	Theroretical Current Balance[610.27].	Account Movement[#19][-41.77].	Updated Balance[568.50] (2019-11-30 11:07:29.261358493)
[Goroutine #02].	Theroretical Current Balance[568.50].	Account Movement[#20][-33.04].	Updated Balance[535.46] (2019-11-30 11:07:29.286717627)
[Goroutine #02].	Theroretical Current Balance[571.17].	Account Movement[#23][-49.91].	Updated Balance[521.26] (2019-11-30 11:07:29.340648268)
[Goroutine #02].	Theroretical Current Balance[521.26].	Account Movement[#24][34.41].	Updated Balance[555.67] (2019-11-30 11:07:29.374469126)
[Goroutine #02].	Theroretical Current Balance[547.40].	Account Movement[#26][-83.34].	Updated Balance[464.06] (2019-11-30 11:07:29.386064644)
[Goroutine #02].	Theroretical Current Balance[379.77].	Account Movement[#28][1.71].	Updated Balance[381.48] (2019-11-30 11:07:29.457974662)
[Goroutine #02].	Theroretical Current Balance[381.48].	Account Movement[#29][-75.72].	Updated Balance[305.76] (2019-11-30 11:07:29.465461236)
[Goroutine #02].	Theroretical Current Balance[305.76].	Account Movement[#30][-73.64].	Updated Balance[232.12] (2019-11-30 11:07:29.478254064)
[Goroutine #02].	Theroretical Current Balance[232.12].	Account Movement[#31][-97.00].	Updated Balance[135.12] (2019-11-30 11:07:29.499927883)
[Goroutine #02].	Theroretical Current Balance[41.09].	Account Movement[#33][-15.48].	Updated Balance[25.61] (2019-11-30 11:07:29.524017585)
[Goroutine #02].	Theroretical Current Balance[25.61].	Account Movement[#34][0.95].	Updated Balance[26.56] (2019-11-30 11:07:29.569332721)
[Goroutine #02].	Theroretical Current Balance[44.58].	Account Movement[#36][70.39].	Updated Balance[114.97] (2019-11-30 11:07:29.619349627)
[Goroutine #02].	Theroretical Current Balance[108.82].	Account Movement[#38][-80.81].	Updated Balance[28.01] (2019-11-30 11:07:29.669407428)
[Goroutine #02].	Theroretical Current Balance[39.01].	Account Movement[#40][-3.26].	Updated Balance[35.75] (2019-11-30 11:07:29.727772256)
[Goroutine #02].	Theroretical Current Balance[117.41].	Account Movement[#42][58.06].	Updated Balance[175.47] (2019-11-30 11:07:29.791586816)
[Goroutine #02].	Theroretical Current Balance[175.47].	Account Movement[#43][96.79].	Updated Balance[272.26] (2019-11-30 11:07:29.802427333)
[Goroutine #02].	Theroretical Current Balance[272.26].	Account Movement[#44][-97.54].	Updated Balance[174.72] (2019-11-30 11:07:29.818349946)
[Goroutine #02].	Theroretical Current Balance[76.75].	Account Movement[#46][45.18].	Updated Balance[121.93] (2019-11-30 11:07:29.837864068)
[Goroutine #02].	Theroretical Current Balance[121.93].	Account Movement[#47][4.75].	Updated Balance[126.68] (2019-11-30 11:07:29.887092220)
[Goroutine #02].	Theroretical Current Balance[41.88].	Account Movement[#49][75.03].	Updated Balance[116.91] (2019-11-30 11:07:29.936860410)
[Goroutine #02].	Theroretical Current Balance[33.35].	Account Movement[#52][15.40].	Updated Balance[48.75] (2019-11-30 11:07:29.958072904)
[Goroutine #02].	Theroretical Current Balance[48.75].	Account Movement[#53][-43.92].	Updated Balance[4.83] (2019-11-30 11:07:29.994043756)
[Goroutine #02].	Theroretical Current Balance[21.45].	Account Movement[#55][66.13].	Updated Balance[87.58] (2019-11-30 11:07:30.064359422)
[Goroutine #02].	Theroretical Current Balance[126.25].	Account Movement[#57][-80.34].	Updated Balance[45.91] (2019-11-30 11:07:30.120352505)
[Goroutine #02].	Theroretical Current Balance[45.91].	Account Movement[#58][21.61].	Updated Balance[67.52] (2019-11-30 11:07:30.125834522)
[Goroutine #02].	Theroretical Current Balance[109.53].	Account Movement[#60][-94.07].	Updated Balance[15.46] (2019-11-30 11:07:30.190717676)
------------------------------------------------------------------------------------------------------------------------------------------------------
		--> Exiting from Goroutine [Goroutine #02]. Theoretical resulting Account Balance[15.46] (2019-11-30 11:07:30.258268387)
------------------------------------------------------------------------------------------------------------------------------------------------------
=========================================================
Resulting Balance [15.46].
=========================================================
Press ENTER to continue.
GOODBYE !!!

```

### 2nd. try

```sh
devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning/courseraGoC3Week2Assignment/racecondition$ ./racecondition 


============================================================================================================
WELCOME !!!
This program aims to show an example of the appearance of a RACE CONDITION. What does it means?

RACE CONDITION DEFINITION:
A 'RACE CONDITION' (a.k.a 'RACE HAZARD') in software systems is the condition where the system's substantive 
behavior is dependent on the sequence or timing of interleaving actions.
Race conditions arise in software when an application depends on the sequence or timing of processes or threads
for it to operate properly:
  ==> Interleaving is non-deterministic, consequently the outcome depends on non-deterministic ordering.
  ==> Races occur due to communication (access to shared data/variables) among concurrently running threads
============================================================================================================

Press ENTER to continue.

============================================================================================================
LET'S SHOW THAT WITH AN EXAMPLE !!!
This program aims to show an example of the emergence of a RACE CONDITION when more than one thread is being
run concurrently and having simultaneous read and write access to a zone of shared (variable) data.
In the proposed example, a situation is simulated where there is an initial balance of an account and also a
list of pending transactions to be applied on. The process to be followed is to increase the position in the
list (cursor), take the value of the correspondig movement in the list and update the current account balance
with its value. When the entire list of pending movements has been traversed, the resulting final balance is
to be shown.
If you carry out the test to execute this program successively, you will be able to observe that when the
RACE CONDITION (badly managed) occurs, different results will be obtained each time: 
--> UNDETERMINISTIC OUTCOME
============================================================================================================

Press ENTER to continue.

============================================================================================================
STAGE #01:
First it will be shown a completely correct behaviour when a single thread is used (the main Goroutine).
Once applied all movements the resulting balance should be [120.30].
============================================================================================================

Press ENTER to continue.
=========================================================
Initial Balance [300.00].
=========================================================

		--> Entering Goroutine [MAIN Goroutine]. Theoretical initial Account Balance[300.00] (2019-11-30 11:13:43.988324051)

[MAIN Goroutine].	Theroretical Current Balance[300.00].	Account Movement[#0][40.50].	Updated Balance[340.50] (2019-11-30 11:13:43.988467186)
[MAIN Goroutine].	Theroretical Current Balance[340.50].	Account Movement[#1][-20.72].	Updated Balance[319.78] (2019-11-30 11:13:43.988498920)
[MAIN Goroutine].	Theroretical Current Balance[319.78].	Account Movement[#2][89.01].	Updated Balance[408.79] (2019-11-30 11:13:43.988527811)
[MAIN Goroutine].	Theroretical Current Balance[408.79].	Account Movement[#3][60.22].	Updated Balance[469.01] (2019-11-30 11:13:43.988556095)
[MAIN Goroutine].	Theroretical Current Balance[469.01].	Account Movement[#4][-96.02].	Updated Balance[372.99] (2019-11-30 11:13:43.988584381)
[MAIN Goroutine].	Theroretical Current Balance[372.99].	Account Movement[#5][7.42].	Updated Balance[380.41] (2019-11-30 11:13:43.988612685)
[MAIN Goroutine].	Theroretical Current Balance[380.41].	Account Movement[#6][73.79].	Updated Balance[454.20] (2019-11-30 11:13:43.988678736)
[MAIN Goroutine].	Theroretical Current Balance[454.20].	Account Movement[#7][96.24].	Updated Balance[550.44] (2019-11-30 11:13:43.988711308)
[MAIN Goroutine].	Theroretical Current Balance[550.44].	Account Movement[#8][76.81].	Updated Balance[627.25] (2019-11-30 11:13:43.988741052)
[MAIN Goroutine].	Theroretical Current Balance[627.25].	Account Movement[#9][-32.98].	Updated Balance[594.27] (2019-11-30 11:13:43.988769426)
[MAIN Goroutine].	Theroretical Current Balance[594.27].	Account Movement[#10][24.76].	Updated Balance[619.03] (2019-11-30 11:13:43.988800093)
[MAIN Goroutine].	Theroretical Current Balance[619.03].	Account Movement[#11][17.48].	Updated Balance[636.51] (2019-11-30 11:13:43.988828401)
[MAIN Goroutine].	Theroretical Current Balance[636.51].	Account Movement[#12][-11.50].	Updated Balance[625.01] (2019-11-30 11:13:43.988912760)
[MAIN Goroutine].	Theroretical Current Balance[625.01].	Account Movement[#13][94.34].	Updated Balance[719.35] (2019-11-30 11:13:43.988952624)
[MAIN Goroutine].	Theroretical Current Balance[719.35].	Account Movement[#14][43.06].	Updated Balance[762.41] (2019-11-30 11:13:43.988982953)
[MAIN Goroutine].	Theroretical Current Balance[762.41].	Account Movement[#15][56.37].	Updated Balance[818.78] (2019-11-30 11:13:43.989015687)
[MAIN Goroutine].	Theroretical Current Balance[818.78].	Account Movement[#16][-29.41].	Updated Balance[789.37] (2019-11-30 11:13:43.989048326)
[MAIN Goroutine].	Theroretical Current Balance[789.37].	Account Movement[#17][-51.32].	Updated Balance[738.05] (2019-11-30 11:13:43.989085764)
[MAIN Goroutine].	Theroretical Current Balance[738.05].	Account Movement[#18][-53.99].	Updated Balance[684.06] (2019-11-30 11:13:43.989116591)
[MAIN Goroutine].	Theroretical Current Balance[684.06].	Account Movement[#19][-41.77].	Updated Balance[642.29] (2019-11-30 11:13:43.989149203)
[MAIN Goroutine].	Theroretical Current Balance[642.29].	Account Movement[#20][-33.04].	Updated Balance[609.25] (2019-11-30 11:13:43.989214458)
[MAIN Goroutine].	Theroretical Current Balance[609.25].	Account Movement[#21][-49.18].	Updated Balance[560.07] (2019-11-30 11:13:43.989249516)
[MAIN Goroutine].	Theroretical Current Balance[560.07].	Account Movement[#22][84.89].	Updated Balance[644.96] (2019-11-30 11:13:43.989282701)
[MAIN Goroutine].	Theroretical Current Balance[644.96].	Account Movement[#23][-49.91].	Updated Balance[595.05] (2019-11-30 11:13:43.989315756)
[MAIN Goroutine].	Theroretical Current Balance[595.05].	Account Movement[#24][34.41].	Updated Balance[629.46] (2019-11-30 11:13:43.989348325)
[MAIN Goroutine].	Theroretical Current Balance[629.46].	Account Movement[#25][-8.27].	Updated Balance[621.19] (2019-11-30 11:13:43.989383918)
[MAIN Goroutine].	Theroretical Current Balance[621.19].	Account Movement[#26][-83.34].	Updated Balance[537.85] (2019-11-30 11:13:43.989413847)
[MAIN Goroutine].	Theroretical Current Balance[537.85].	Account Movement[#27][-84.29].	Updated Balance[453.56] (2019-11-30 11:13:43.989443022)
[MAIN Goroutine].	Theroretical Current Balance[453.56].	Account Movement[#28][1.71].	Updated Balance[455.27] (2019-11-30 11:13:43.989508282)
[MAIN Goroutine].	Theroretical Current Balance[455.27].	Account Movement[#29][-75.72].	Updated Balance[379.55] (2019-11-30 11:13:43.989541486)
[MAIN Goroutine].	Theroretical Current Balance[379.55].	Account Movement[#30][-73.64].	Updated Balance[305.91] (2019-11-30 11:13:43.989574222)
[MAIN Goroutine].	Theroretical Current Balance[305.91].	Account Movement[#31][-97.00].	Updated Balance[208.91] (2019-11-30 11:13:43.989631887)
[MAIN Goroutine].	Theroretical Current Balance[208.91].	Account Movement[#32][-94.03].	Updated Balance[114.88] (2019-11-30 11:13:43.989674829)
[MAIN Goroutine].	Theroretical Current Balance[114.88].	Account Movement[#33][-15.48].	Updated Balance[99.40] (2019-11-30 11:13:43.989742844)
[MAIN Goroutine].	Theroretical Current Balance[99.40].	Account Movement[#34][0.95].	Updated Balance[100.35] (2019-11-30 11:13:43.989781070)
[MAIN Goroutine].	Theroretical Current Balance[100.35].	Account Movement[#35][18.02].	Updated Balance[118.37] (2019-11-30 11:13:43.989818574)
[MAIN Goroutine].	Theroretical Current Balance[118.37].	Account Movement[#36][70.39].	Updated Balance[188.76] (2019-11-30 11:13:43.989849035)
[MAIN Goroutine].	Theroretical Current Balance[188.76].	Account Movement[#37][-6.15].	Updated Balance[182.61] (2019-11-30 11:13:43.989881398)
[MAIN Goroutine].	Theroretical Current Balance[182.61].	Account Movement[#38][-80.81].	Updated Balance[101.80] (2019-11-30 11:13:43.989914413)
[MAIN Goroutine].	Theroretical Current Balance[101.80].	Account Movement[#39][11.00].	Updated Balance[112.80] (2019-11-30 11:13:43.989951040)
[MAIN Goroutine].	Theroretical Current Balance[112.80].	Account Movement[#40][-3.26].	Updated Balance[109.54] (2019-11-30 11:13:43.989980955)
[MAIN Goroutine].	Theroretical Current Balance[109.54].	Account Movement[#41][81.66].	Updated Balance[191.20] (2019-11-30 11:13:43.990045111)
[MAIN Goroutine].	Theroretical Current Balance[191.20].	Account Movement[#42][58.06].	Updated Balance[249.26] (2019-11-30 11:13:43.990109921)
[MAIN Goroutine].	Theroretical Current Balance[249.26].	Account Movement[#43][96.79].	Updated Balance[346.05] (2019-11-30 11:13:43.990144793)
[MAIN Goroutine].	Theroretical Current Balance[346.05].	Account Movement[#44][-97.54].	Updated Balance[248.51] (2019-11-30 11:13:43.990177305)
[MAIN Goroutine].	Theroretical Current Balance[248.51].	Account Movement[#45][-97.97].	Updated Balance[150.54] (2019-11-30 11:13:43.990241059)
[MAIN Goroutine].	Theroretical Current Balance[150.54].	Account Movement[#46][45.18].	Updated Balance[195.72] (2019-11-30 11:13:43.990307733)
[MAIN Goroutine].	Theroretical Current Balance[195.72].	Account Movement[#47][4.75].	Updated Balance[200.47] (2019-11-30 11:13:43.990342205)
[MAIN Goroutine].	Theroretical Current Balance[200.47].	Account Movement[#48][-84.80].	Updated Balance[115.67] (2019-11-30 11:13:43.990408049)
[MAIN Goroutine].	Theroretical Current Balance[115.67].	Account Movement[#49][75.03].	Updated Balance[190.70] (2019-11-30 11:13:43.990443454)
[MAIN Goroutine].	Theroretical Current Balance[190.70].	Account Movement[#50][31.05].	Updated Balance[221.75] (2019-11-30 11:13:43.990509474)
[MAIN Goroutine].	Theroretical Current Balance[221.75].	Account Movement[#51][-83.56].	Updated Balance[138.19] (2019-11-30 11:13:43.990544497)
[MAIN Goroutine].	Theroretical Current Balance[138.19].	Account Movement[#52][15.40].	Updated Balance[153.59] (2019-11-30 11:13:43.990612624)
[MAIN Goroutine].	Theroretical Current Balance[153.59].	Account Movement[#53][-43.92].	Updated Balance[109.67] (2019-11-30 11:13:43.990647728)
[MAIN Goroutine].	Theroretical Current Balance[109.67].	Account Movement[#54][16.62].	Updated Balance[126.29] (2019-11-30 11:13:43.990679664)
[MAIN Goroutine].	Theroretical Current Balance[126.29].	Account Movement[#55][66.13].	Updated Balance[192.42] (2019-11-30 11:13:43.990747310)
[MAIN Goroutine].	Theroretical Current Balance[192.42].	Account Movement[#56][38.67].	Updated Balance[231.09] (2019-11-30 11:13:43.990782324)
[MAIN Goroutine].	Theroretical Current Balance[231.09].	Account Movement[#57][-80.34].	Updated Balance[150.75] (2019-11-30 11:13:43.990850520)
[MAIN Goroutine].	Theroretical Current Balance[150.75].	Account Movement[#58][21.61].	Updated Balance[172.36] (2019-11-30 11:13:43.990885381)
[MAIN Goroutine].	Theroretical Current Balance[172.36].	Account Movement[#59][42.01].	Updated Balance[214.37] (2019-11-30 11:13:43.990917859)
[MAIN Goroutine].	Theroretical Current Balance[214.37].	Account Movement[#60][-94.07].	Updated Balance[120.30] (2019-11-30 11:13:43.990949846)

		--> Exiting from Goroutine [MAIN Goroutine]. Theoretical resulting Account Balance[120.30] (2019-11-30 11:13:43.990982120)

=========================================================
Resulting Balance [120.30].
=========================================================
Press ENTER to continue.

============================================================================================================
STAGE #02:
Now we will reset both values: the Cursor (position) of the list of pending movements and the value of
the initial Balance to their starting values
Then two threads (Goroutines) will be launched and execute concurrently creating a RACE CONDITION as:
1. The two threads will access concurrently trying to increase the shared value of the cursor of the
   movement list and updating the shared value of the balance with the corresponding movement of the list
2. Running this way final will drive to a totally different final Balance result each time
 
Note: now the main thread will be waiting sleeping during a reasonable time to permit all other threads
to succesfully conclude its tasks. DISCLAIMER: There are better way to sync. threads but they have not
shown/explained at this point in the course.
============================================================================================================

Press ENTER to continue.

=========================================================
RESETTING Position[0] and Initial Account Balance[300.00]
=========================================================

=========================================================
Initial Balance [300.00].
=========================================================
Main Thread is currently Awaiting [10] seconds to let all goroutines to safely end ...

		--> Entering Goroutine [Goroutine #02]. Theoretical initial Account Balance[300.00] (2019-11-30 11:13:45.008287376)

[Goroutine #02].	Theroretical Current Balance[300.00].	Account Movement[#0][40.50].	Updated Balance[340.50] (2019-11-30 11:13:45.008379869)

		--> Entering Goroutine [Goroutine #01]. Theoretical initial Account Balance[300.00] (2019-11-30 11:13:45.008310779)

[Goroutine #01].	Theroretical Current Balance[340.50].	Account Movement[#1][-20.72].	Updated Balance[319.78] (2019-11-30 11:13:45.008515657)
[Goroutine #02].	Theroretical Current Balance[319.78].	Account Movement[#2][89.01].	Updated Balance[408.79] (2019-11-30 11:13:45.021234819)
[Goroutine #01].	Theroretical Current Balance[408.79].	Account Movement[#3][60.22].	Updated Balance[469.01] (2019-11-30 11:13:45.041177791)
[Goroutine #02].	Theroretical Current Balance[469.01].	Account Movement[#4][-96.02].	Updated Balance[372.99] (2019-11-30 11:13:45.046574976)
[Goroutine #02].	Theroretical Current Balance[372.99].	Account Movement[#5][7.42].	Updated Balance[380.41] (2019-11-30 11:13:45.048937349)
[Goroutine #01].	Theroretical Current Balance[380.41].	Account Movement[#6][73.79].	Updated Balance[454.20] (2019-11-30 11:13:45.070669817)
[Goroutine #02].	Theroretical Current Balance[454.20].	Account Movement[#7][96.24].	Updated Balance[550.44] (2019-11-30 11:13:45.084660670)
[Goroutine #02].	Theroretical Current Balance[550.44].	Account Movement[#8][76.81].	Updated Balance[627.25] (2019-11-30 11:13:45.091703141)
[Goroutine #02].	Theroretical Current Balance[627.25].	Account Movement[#9][-32.98].	Updated Balance[594.27] (2019-11-30 11:13:45.108439819)
[Goroutine #01].	Theroretical Current Balance[594.27].	Account Movement[#10][24.76].	Updated Balance[619.03] (2019-11-30 11:13:45.111828978)
[Goroutine #02].	Theroretical Current Balance[619.03].	Account Movement[#11][17.48].	Updated Balance[636.51] (2019-11-30 11:13:45.124684868)
[Goroutine #02].	Theroretical Current Balance[636.51].	Account Movement[#12][-11.50].	Updated Balance[625.01] (2019-11-30 11:13:45.152052925)
[Goroutine #01].	Theroretical Current Balance[625.01].	Account Movement[#13][94.34].	Updated Balance[719.35] (2019-11-30 11:13:45.153303216)
[Goroutine #02].	Theroretical Current Balance[719.35].	Account Movement[#14][43.06].	Updated Balance[762.41] (2019-11-30 11:13:45.163880034)
[Goroutine #02].	Theroretical Current Balance[762.41].	Account Movement[#15][56.37].	Updated Balance[818.78] (2019-11-30 11:13:45.176587768)
[Goroutine #01].	Theroretical Current Balance[818.78].	Account Movement[#16][-29.41].	Updated Balance[789.37] (2019-11-30 11:13:45.181326061)
[Goroutine #01].	Theroretical Current Balance[789.37].	Account Movement[#17][-51.32].	Updated Balance[738.05] (2019-11-30 11:13:45.207462119)
[Goroutine #02].	Theroretical Current Balance[738.05].	Account Movement[#18][-53.99].	Updated Balance[684.06] (2019-11-30 11:13:45.208852353)
[Goroutine #01].	Theroretical Current Balance[684.06].	Account Movement[#19][-41.77].	Updated Balance[642.29] (2019-11-30 11:13:45.222815468)
[Goroutine #01].	Theroretical Current Balance[642.29].	Account Movement[#20][-33.04].	Updated Balance[609.25] (2019-11-30 11:13:45.229321105)
[Goroutine #02].	Theroretical Current Balance[609.25].	Account Movement[#21][-49.18].	Updated Balance[560.07] (2019-11-30 11:13:45.257373758)
[Goroutine #01].	Theroretical Current Balance[560.07].	Account Movement[#22][84.89].	Updated Balance[644.96] (2019-11-30 11:13:45.271111502)
[Goroutine #02].	Theroretical Current Balance[644.96].	Account Movement[#23][-49.91].	Updated Balance[595.05] (2019-11-30 11:13:45.282791464)
[Goroutine #01].	Theroretical Current Balance[595.05].	Account Movement[#24][34.41].	Updated Balance[629.46] (2019-11-30 11:13:45.305535404)
[Goroutine #02].	Theroretical Current Balance[629.46].	Account Movement[#25][-8.27].	Updated Balance[621.19] (2019-11-30 11:13:45.309466293)
[Goroutine #02].	Theroretical Current Balance[621.19].	Account Movement[#26][-83.34].	Updated Balance[537.85] (2019-11-30 11:13:45.332475366)
[Goroutine #02].	Theroretical Current Balance[537.85].	Account Movement[#27][-84.29].	Updated Balance[453.56] (2019-11-30 11:13:45.338766251)
[Goroutine #01].	Theroretical Current Balance[453.56].	Account Movement[#28][1.71].	Updated Balance[455.27] (2019-11-30 11:13:45.341677559)
[Goroutine #02].	Theroretical Current Balance[455.27].	Account Movement[#29][-75.72].	Updated Balance[379.55] (2019-11-30 11:13:45.359167659)
[Goroutine #02].	Theroretical Current Balance[379.55].	Account Movement[#30][-73.64].	Updated Balance[305.91] (2019-11-30 11:13:45.371937282)
[Goroutine #01].	Theroretical Current Balance[379.55].	Account Movement[#30][-73.64].	Updated Balance[305.91] (2019-11-30 11:13:45.372141835)
[Goroutine #01].	Theroretical Current Balance[305.91].	Account Movement[#32][-94.03].	Updated Balance[211.88] (2019-11-30 11:13:45.382745787)
[Goroutine #01].	Theroretical Current Balance[211.88].	Account Movement[#33][-15.48].	Updated Balance[196.40] (2019-11-30 11:13:45.407053125)
[Goroutine #02].	Theroretical Current Balance[211.88].	Account Movement[#33][-15.48].	Updated Balance[196.40] (2019-11-30 11:13:45.407116780)
[Goroutine #02].	Theroretical Current Balance[196.40].	Account Movement[#35][18.02].	Updated Balance[214.42] (2019-11-30 11:13:45.440213882)
[Goroutine #01].	Theroretical Current Balance[214.42].	Account Movement[#36][70.39].	Updated Balance[284.81] (2019-11-30 11:13:45.452049694)
[Goroutine #02].	Theroretical Current Balance[284.81].	Account Movement[#37][-6.15].	Updated Balance[278.66] (2019-11-30 11:13:45.460612789)
[Goroutine #01].	Theroretical Current Balance[278.66].	Account Movement[#38][-80.81].	Updated Balance[197.85] (2019-11-30 11:13:45.462246554)
[Goroutine #02].	Theroretical Current Balance[197.85].	Account Movement[#39][11.00].	Updated Balance[208.85] (2019-11-30 11:13:45.487111057)
[Goroutine #01].	Theroretical Current Balance[208.85].	Account Movement[#40][-3.26].	Updated Balance[205.59] (2019-11-30 11:13:45.487526247)
[Goroutine #02].	Theroretical Current Balance[205.59].	Account Movement[#41][81.66].	Updated Balance[287.25] (2019-11-30 11:13:45.524344665)
[Goroutine #01].	Theroretical Current Balance[287.25].	Account Movement[#42][58.06].	Updated Balance[345.31] (2019-11-30 11:13:45.524665906)
[Goroutine #02].	Theroretical Current Balance[345.31].	Account Movement[#43][96.79].	Updated Balance[442.10] (2019-11-30 11:13:45.530251462)
[Goroutine #02].	Theroretical Current Balance[442.10].	Account Movement[#44][-97.54].	Updated Balance[344.56] (2019-11-30 11:13:45.554149060)
[Goroutine #01].	Theroretical Current Balance[344.56].	Account Movement[#45][-97.97].	Updated Balance[246.59] (2019-11-30 11:13:45.557071974)
[Goroutine #02].	Theroretical Current Balance[246.59].	Account Movement[#46][45.18].	Updated Balance[291.77] (2019-11-30 11:13:45.563218224)
[Goroutine #02].	Theroretical Current Balance[291.77].	Account Movement[#47][4.75].	Updated Balance[296.52] (2019-11-30 11:13:45.579135121)
[Goroutine #01].	Theroretical Current Balance[296.52].	Account Movement[#48][-84.80].	Updated Balance[211.72] (2019-11-30 11:13:45.579475880)
[Goroutine #01].	Theroretical Current Balance[211.72].	Account Movement[#49][75.03].	Updated Balance[286.75] (2019-11-30 11:13:45.598295586)
[Goroutine #01].	Theroretical Current Balance[286.75].	Account Movement[#50][31.05].	Updated Balance[317.80] (2019-11-30 11:13:45.621737927)
[Goroutine #02].	Theroretical Current Balance[317.80].	Account Movement[#51][-83.56].	Updated Balance[234.24] (2019-11-30 11:13:45.623143434)
[Goroutine #01].	Theroretical Current Balance[234.24].	Account Movement[#52][15.40].	Updated Balance[249.64] (2019-11-30 11:13:45.640941853)
[Goroutine #02].	Theroretical Current Balance[249.64].	Account Movement[#53][-43.92].	Updated Balance[205.72] (2019-11-30 11:13:45.644350941)
[Goroutine #02].	Theroretical Current Balance[205.72].	Account Movement[#54][16.62].	Updated Balance[222.34] (2019-11-30 11:13:45.654053943)
[Goroutine #02].	Theroretical Current Balance[222.34].	Account Movement[#55][66.13].	Updated Balance[288.47] (2019-11-30 11:13:45.681431544)
[Goroutine #02].	Theroretical Current Balance[288.47].	Account Movement[#56][38.67].	Updated Balance[327.14] (2019-11-30 11:13:45.687510562)
[Goroutine #01].	Theroretical Current Balance[288.47].	Account Movement[#56][38.67].	Updated Balance[327.14] (2019-11-30 11:13:45.687592788)
[Goroutine #01].	Theroretical Current Balance[327.14].	Account Movement[#58][21.61].	Updated Balance[348.75] (2019-11-30 11:13:45.690452198)
[Goroutine #02].	Theroretical Current Balance[348.75].	Account Movement[#59][42.01].	Updated Balance[390.76] (2019-11-30 11:13:45.719692217)
[Goroutine #01].	Theroretical Current Balance[390.76].	Account Movement[#60][-94.07].	Updated Balance[296.69] (2019-11-30 11:13:45.731315285)

		--> Exiting from Goroutine [Goroutine #01]. Theoretical resulting Account Balance[296.69] (2019-11-30 11:13:45.746332445)


		--> Exiting from Goroutine [Goroutine #02]. Theoretical resulting Account Balance[390.76] (2019-11-30 11:13:45.767003525)

=========================================================
Resulting Balance [296.69].
=========================================================
Press ENTER to continue.

============================================================================================================
STAGE #03:
Each thread has been writing in a log its activity showing at every moment in time the values that 'see'
and update the shared variables: Cursor and Current balance.
Now it's time to see how they were behaving individually.
============================================================================================================

Press ENTER to continue.

========================================================================
Showing now the event's history registered at Goroutine [Goroutine #01]
========================================================================
------------------------------------------------------------------------------------------------------------------------------------------------------
		--> Entering Goroutine [Goroutine #01]. Theoretical initial Account Balance[300.00] (2019-11-30 11:13:45.008310779)
------------------------------------------------------------------------------------------------------------------------------------------------------
[Goroutine #01].	Theroretical Current Balance[340.50].	Account Movement[#1][-20.72].	Updated Balance[319.78] (2019-11-30 11:13:45.008515657)
[Goroutine #01].	Theroretical Current Balance[408.79].	Account Movement[#3][60.22].	Updated Balance[469.01] (2019-11-30 11:13:45.041177791)
[Goroutine #01].	Theroretical Current Balance[380.41].	Account Movement[#6][73.79].	Updated Balance[454.20] (2019-11-30 11:13:45.070669817)
[Goroutine #01].	Theroretical Current Balance[594.27].	Account Movement[#10][24.76].	Updated Balance[619.03] (2019-11-30 11:13:45.111828978)
[Goroutine #01].	Theroretical Current Balance[625.01].	Account Movement[#13][94.34].	Updated Balance[719.35] (2019-11-30 11:13:45.153303216)
[Goroutine #01].	Theroretical Current Balance[818.78].	Account Movement[#16][-29.41].	Updated Balance[789.37] (2019-11-30 11:13:45.181326061)
[Goroutine #01].	Theroretical Current Balance[789.37].	Account Movement[#17][-51.32].	Updated Balance[738.05] (2019-11-30 11:13:45.207462119)
[Goroutine #01].	Theroretical Current Balance[684.06].	Account Movement[#19][-41.77].	Updated Balance[642.29] (2019-11-30 11:13:45.222815468)
[Goroutine #01].	Theroretical Current Balance[642.29].	Account Movement[#20][-33.04].	Updated Balance[609.25] (2019-11-30 11:13:45.229321105)
[Goroutine #01].	Theroretical Current Balance[560.07].	Account Movement[#22][84.89].	Updated Balance[644.96] (2019-11-30 11:13:45.271111502)
[Goroutine #01].	Theroretical Current Balance[595.05].	Account Movement[#24][34.41].	Updated Balance[629.46] (2019-11-30 11:13:45.305535404)
[Goroutine #01].	Theroretical Current Balance[453.56].	Account Movement[#28][1.71].	Updated Balance[455.27] (2019-11-30 11:13:45.341677559)
[Goroutine #01].	Theroretical Current Balance[379.55].	Account Movement[#30][-73.64].	Updated Balance[305.91] (2019-11-30 11:13:45.372141835)
[Goroutine #01].	Theroretical Current Balance[305.91].	Account Movement[#32][-94.03].	Updated Balance[211.88] (2019-11-30 11:13:45.382745787)
[Goroutine #01].	Theroretical Current Balance[211.88].	Account Movement[#33][-15.48].	Updated Balance[196.40] (2019-11-30 11:13:45.407053125)
[Goroutine #01].	Theroretical Current Balance[214.42].	Account Movement[#36][70.39].	Updated Balance[284.81] (2019-11-30 11:13:45.452049694)
[Goroutine #01].	Theroretical Current Balance[278.66].	Account Movement[#38][-80.81].	Updated Balance[197.85] (2019-11-30 11:13:45.462246554)
[Goroutine #01].	Theroretical Current Balance[208.85].	Account Movement[#40][-3.26].	Updated Balance[205.59] (2019-11-30 11:13:45.487526247)
[Goroutine #01].	Theroretical Current Balance[287.25].	Account Movement[#42][58.06].	Updated Balance[345.31] (2019-11-30 11:13:45.524665906)
[Goroutine #01].	Theroretical Current Balance[344.56].	Account Movement[#45][-97.97].	Updated Balance[246.59] (2019-11-30 11:13:45.557071974)
[Goroutine #01].	Theroretical Current Balance[296.52].	Account Movement[#48][-84.80].	Updated Balance[211.72] (2019-11-30 11:13:45.579475880)
[Goroutine #01].	Theroretical Current Balance[211.72].	Account Movement[#49][75.03].	Updated Balance[286.75] (2019-11-30 11:13:45.598295586)
[Goroutine #01].	Theroretical Current Balance[286.75].	Account Movement[#50][31.05].	Updated Balance[317.80] (2019-11-30 11:13:45.621737927)
[Goroutine #01].	Theroretical Current Balance[234.24].	Account Movement[#52][15.40].	Updated Balance[249.64] (2019-11-30 11:13:45.640941853)
[Goroutine #01].	Theroretical Current Balance[288.47].	Account Movement[#56][38.67].	Updated Balance[327.14] (2019-11-30 11:13:45.687592788)
[Goroutine #01].	Theroretical Current Balance[327.14].	Account Movement[#58][21.61].	Updated Balance[348.75] (2019-11-30 11:13:45.690452198)
[Goroutine #01].	Theroretical Current Balance[390.76].	Account Movement[#60][-94.07].	Updated Balance[296.69] (2019-11-30 11:13:45.731315285)
------------------------------------------------------------------------------------------------------------------------------------------------------
		--> Exiting from Goroutine [Goroutine #01]. Theoretical resulting Account Balance[296.69] (2019-11-30 11:13:45.746332445)
------------------------------------------------------------------------------------------------------------------------------------------------------
Press ENTER to continue.

========================================================================
Showing now the event's history registered at Goroutine [Goroutine #02]
========================================================================
------------------------------------------------------------------------------------------------------------------------------------------------------
		--> Entering Goroutine [Goroutine #02]. Theoretical initial Account Balance[300.00] (2019-11-30 11:13:45.008287376)
------------------------------------------------------------------------------------------------------------------------------------------------------
[Goroutine #02].	Theroretical Current Balance[300.00].	Account Movement[#0][40.50].	Updated Balance[340.50] (2019-11-30 11:13:45.008379869)
[Goroutine #02].	Theroretical Current Balance[319.78].	Account Movement[#2][89.01].	Updated Balance[408.79] (2019-11-30 11:13:45.021234819)
[Goroutine #02].	Theroretical Current Balance[469.01].	Account Movement[#4][-96.02].	Updated Balance[372.99] (2019-11-30 11:13:45.046574976)
[Goroutine #02].	Theroretical Current Balance[372.99].	Account Movement[#5][7.42].	Updated Balance[380.41] (2019-11-30 11:13:45.048937349)
[Goroutine #02].	Theroretical Current Balance[454.20].	Account Movement[#7][96.24].	Updated Balance[550.44] (2019-11-30 11:13:45.084660670)
[Goroutine #02].	Theroretical Current Balance[550.44].	Account Movement[#8][76.81].	Updated Balance[627.25] (2019-11-30 11:13:45.091703141)
[Goroutine #02].	Theroretical Current Balance[627.25].	Account Movement[#9][-32.98].	Updated Balance[594.27] (2019-11-30 11:13:45.108439819)
[Goroutine #02].	Theroretical Current Balance[619.03].	Account Movement[#11][17.48].	Updated Balance[636.51] (2019-11-30 11:13:45.124684868)
[Goroutine #02].	Theroretical Current Balance[636.51].	Account Movement[#12][-11.50].	Updated Balance[625.01] (2019-11-30 11:13:45.152052925)
[Goroutine #02].	Theroretical Current Balance[719.35].	Account Movement[#14][43.06].	Updated Balance[762.41] (2019-11-30 11:13:45.163880034)
[Goroutine #02].	Theroretical Current Balance[762.41].	Account Movement[#15][56.37].	Updated Balance[818.78] (2019-11-30 11:13:45.176587768)
[Goroutine #02].	Theroretical Current Balance[738.05].	Account Movement[#18][-53.99].	Updated Balance[684.06] (2019-11-30 11:13:45.208852353)
[Goroutine #02].	Theroretical Current Balance[609.25].	Account Movement[#21][-49.18].	Updated Balance[560.07] (2019-11-30 11:13:45.257373758)
[Goroutine #02].	Theroretical Current Balance[644.96].	Account Movement[#23][-49.91].	Updated Balance[595.05] (2019-11-30 11:13:45.282791464)
[Goroutine #02].	Theroretical Current Balance[629.46].	Account Movement[#25][-8.27].	Updated Balance[621.19] (2019-11-30 11:13:45.309466293)
[Goroutine #02].	Theroretical Current Balance[621.19].	Account Movement[#26][-83.34].	Updated Balance[537.85] (2019-11-30 11:13:45.332475366)
[Goroutine #02].	Theroretical Current Balance[537.85].	Account Movement[#27][-84.29].	Updated Balance[453.56] (2019-11-30 11:13:45.338766251)
[Goroutine #02].	Theroretical Current Balance[455.27].	Account Movement[#29][-75.72].	Updated Balance[379.55] (2019-11-30 11:13:45.359167659)
[Goroutine #02].	Theroretical Current Balance[379.55].	Account Movement[#30][-73.64].	Updated Balance[305.91] (2019-11-30 11:13:45.371937282)
[Goroutine #02].	Theroretical Current Balance[211.88].	Account Movement[#33][-15.48].	Updated Balance[196.40] (2019-11-30 11:13:45.407116780)
[Goroutine #02].	Theroretical Current Balance[196.40].	Account Movement[#35][18.02].	Updated Balance[214.42] (2019-11-30 11:13:45.440213882)
[Goroutine #02].	Theroretical Current Balance[284.81].	Account Movement[#37][-6.15].	Updated Balance[278.66] (2019-11-30 11:13:45.460612789)
[Goroutine #02].	Theroretical Current Balance[197.85].	Account Movement[#39][11.00].	Updated Balance[208.85] (2019-11-30 11:13:45.487111057)
[Goroutine #02].	Theroretical Current Balance[205.59].	Account Movement[#41][81.66].	Updated Balance[287.25] (2019-11-30 11:13:45.524344665)
[Goroutine #02].	Theroretical Current Balance[345.31].	Account Movement[#43][96.79].	Updated Balance[442.10] (2019-11-30 11:13:45.530251462)
[Goroutine #02].	Theroretical Current Balance[442.10].	Account Movement[#44][-97.54].	Updated Balance[344.56] (2019-11-30 11:13:45.554149060)
[Goroutine #02].	Theroretical Current Balance[246.59].	Account Movement[#46][45.18].	Updated Balance[291.77] (2019-11-30 11:13:45.563218224)
[Goroutine #02].	Theroretical Current Balance[291.77].	Account Movement[#47][4.75].	Updated Balance[296.52] (2019-11-30 11:13:45.579135121)
[Goroutine #02].	Theroretical Current Balance[317.80].	Account Movement[#51][-83.56].	Updated Balance[234.24] (2019-11-30 11:13:45.623143434)
[Goroutine #02].	Theroretical Current Balance[249.64].	Account Movement[#53][-43.92].	Updated Balance[205.72] (2019-11-30 11:13:45.644350941)
[Goroutine #02].	Theroretical Current Balance[205.72].	Account Movement[#54][16.62].	Updated Balance[222.34] (2019-11-30 11:13:45.654053943)
[Goroutine #02].	Theroretical Current Balance[222.34].	Account Movement[#55][66.13].	Updated Balance[288.47] (2019-11-30 11:13:45.681431544)
[Goroutine #02].	Theroretical Current Balance[288.47].	Account Movement[#56][38.67].	Updated Balance[327.14] (2019-11-30 11:13:45.687510562)
[Goroutine #02].	Theroretical Current Balance[348.75].	Account Movement[#59][42.01].	Updated Balance[390.76] (2019-11-30 11:13:45.719692217)
------------------------------------------------------------------------------------------------------------------------------------------------------
		--> Exiting from Goroutine [Goroutine #02]. Theoretical resulting Account Balance[390.76] (2019-11-30 11:13:45.767003525)
------------------------------------------------------------------------------------------------------------------------------------------------------
=========================================================
Resulting Balance [296.69].
=========================================================
Press ENTER to continue.
GOODBYE !!!

```

### 3rd. Try

```sh
devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning/courseraGoC3Week2Assignment/racecondition$ ./racecondition 

============================================================================================================
WELCOME !!!
This program aims to show an example of the appearance of a RACE CONDITION. What does it means?

RACE CONDITION DEFINITION:
A 'RACE CONDITION' (a.k.a 'RACE HAZARD') in software systems is the condition where the system's substantive 
behavior is dependent on the sequence or timing of interleaving actions.
Race conditions arise in software when an application depends on the sequence or timing of processes or threads
for it to operate properly:
  ==> Interleaving is non-deterministic, consequently the outcome depends on non-deterministic ordering.
  ==> Races occur due to communication (access to shared data/variables) among concurrently running threads
============================================================================================================

Press ENTER to continue.

============================================================================================================
LET'S SHOW THAT WITH AN EXAMPLE !!!
This program aims to show an example of the emergence of a RACE CONDITION when more than one thread is being
run concurrently and having simultaneous read and write access to a zone of shared (variable) data.
In the proposed example, a situation is simulated where there is an initial balance of an account and also a
list of pending transactions to be applied on. The process to be followed is to increase the position in the
list (cursor), take the value of the correspondig movement in the list and update the current account balance
with its value. When the entire list of pending movements has been traversed, the resulting final balance is
to be shown.
If you carry out the test to execute this program successively, you will be able to observe that when the
RACE CONDITION (badly managed) occurs, different results will be obtained each time: 
--> UNDETERMINISTIC OUTCOME
============================================================================================================

Press ENTER to continue.

============================================================================================================
STAGE #01:
First it will be shown a completely correct behaviour when a single thread is used (the main Goroutine).
Once applied all movements the resulting balance should be [120.30].
============================================================================================================

Press ENTER to continue.
=========================================================
Initial Balance [300.00].
=========================================================

		--> Entering Goroutine [MAIN Goroutine]. Theoretical initial Account Balance[300.00] (2019-11-30 11:16:09.365991077)

[MAIN Goroutine].	Theroretical Current Balance[300.00].	Account Movement[#0][40.50].	Updated Balance[340.50] (2019-11-30 11:16:09.366156615)
[MAIN Goroutine].	Theroretical Current Balance[340.50].	Account Movement[#1][-20.72].	Updated Balance[319.78] (2019-11-30 11:16:09.366195766)
[MAIN Goroutine].	Theroretical Current Balance[319.78].	Account Movement[#2][89.01].	Updated Balance[408.79] (2019-11-30 11:16:09.366205956)
[MAIN Goroutine].	Theroretical Current Balance[408.79].	Account Movement[#3][60.22].	Updated Balance[469.01] (2019-11-30 11:16:09.366213790)
[MAIN Goroutine].	Theroretical Current Balance[469.01].	Account Movement[#4][-96.02].	Updated Balance[372.99] (2019-11-30 11:16:09.366221128)
[MAIN Goroutine].	Theroretical Current Balance[372.99].	Account Movement[#5][7.42].	Updated Balance[380.41] (2019-11-30 11:16:09.366228508)
[MAIN Goroutine].	Theroretical Current Balance[380.41].	Account Movement[#6][73.79].	Updated Balance[454.20] (2019-11-30 11:16:09.366242270)
[MAIN Goroutine].	Theroretical Current Balance[454.20].	Account Movement[#7][96.24].	Updated Balance[550.44] (2019-11-30 11:16:09.366249770)
[MAIN Goroutine].	Theroretical Current Balance[550.44].	Account Movement[#8][76.81].	Updated Balance[627.25] (2019-11-30 11:16:09.366257247)
[MAIN Goroutine].	Theroretical Current Balance[627.25].	Account Movement[#9][-32.98].	Updated Balance[594.27] (2019-11-30 11:16:09.366264453)
[MAIN Goroutine].	Theroretical Current Balance[594.27].	Account Movement[#10][24.76].	Updated Balance[619.03] (2019-11-30 11:16:09.366276675)
[MAIN Goroutine].	Theroretical Current Balance[619.03].	Account Movement[#11][17.48].	Updated Balance[636.51] (2019-11-30 11:16:09.366284254)
[MAIN Goroutine].	Theroretical Current Balance[636.51].	Account Movement[#12][-11.50].	Updated Balance[625.01] (2019-11-30 11:16:09.366291451)
[MAIN Goroutine].	Theroretical Current Balance[625.01].	Account Movement[#13][94.34].	Updated Balance[719.35] (2019-11-30 11:16:09.366304330)
[MAIN Goroutine].	Theroretical Current Balance[719.35].	Account Movement[#14][43.06].	Updated Balance[762.41] (2019-11-30 11:16:09.366311913)
[MAIN Goroutine].	Theroretical Current Balance[762.41].	Account Movement[#15][56.37].	Updated Balance[818.78] (2019-11-30 11:16:09.366319217)
[MAIN Goroutine].	Theroretical Current Balance[818.78].	Account Movement[#16][-29.41].	Updated Balance[789.37] (2019-11-30 11:16:09.366326390)
[MAIN Goroutine].	Theroretical Current Balance[789.37].	Account Movement[#17][-51.32].	Updated Balance[738.05] (2019-11-30 11:16:09.366333828)
[MAIN Goroutine].	Theroretical Current Balance[738.05].	Account Movement[#18][-53.99].	Updated Balance[684.06] (2019-11-30 11:16:09.366341028)
[MAIN Goroutine].	Theroretical Current Balance[684.06].	Account Movement[#19][-41.77].	Updated Balance[642.29] (2019-11-30 11:16:09.366385680)
[MAIN Goroutine].	Theroretical Current Balance[642.29].	Account Movement[#20][-33.04].	Updated Balance[609.25] (2019-11-30 11:16:09.366398758)
[MAIN Goroutine].	Theroretical Current Balance[609.25].	Account Movement[#21][-49.18].	Updated Balance[560.07] (2019-11-30 11:16:09.366406363)
[MAIN Goroutine].	Theroretical Current Balance[560.07].	Account Movement[#22][84.89].	Updated Balance[644.96] (2019-11-30 11:16:09.366413702)
[MAIN Goroutine].	Theroretical Current Balance[644.96].	Account Movement[#23][-49.91].	Updated Balance[595.05] (2019-11-30 11:16:09.366441912)
[MAIN Goroutine].	Theroretical Current Balance[595.05].	Account Movement[#24][34.41].	Updated Balance[629.46] (2019-11-30 11:16:09.366473311)
[MAIN Goroutine].	Theroretical Current Balance[629.46].	Account Movement[#25][-8.27].	Updated Balance[621.19] (2019-11-30 11:16:09.366481089)
[MAIN Goroutine].	Theroretical Current Balance[621.19].	Account Movement[#26][-83.34].	Updated Balance[537.85] (2019-11-30 11:16:09.366488874)
[MAIN Goroutine].	Theroretical Current Balance[537.85].	Account Movement[#27][-84.29].	Updated Balance[453.56] (2019-11-30 11:16:09.366496544)
[MAIN Goroutine].	Theroretical Current Balance[453.56].	Account Movement[#28][1.71].	Updated Balance[455.27] (2019-11-30 11:16:09.366504224)
[MAIN Goroutine].	Theroretical Current Balance[455.27].	Account Movement[#29][-75.72].	Updated Balance[379.55] (2019-11-30 11:16:09.366518916)
[MAIN Goroutine].	Theroretical Current Balance[379.55].	Account Movement[#30][-73.64].	Updated Balance[305.91] (2019-11-30 11:16:09.366527585)
[MAIN Goroutine].	Theroretical Current Balance[305.91].	Account Movement[#31][-97.00].	Updated Balance[208.91] (2019-11-30 11:16:09.366540723)
[MAIN Goroutine].	Theroretical Current Balance[208.91].	Account Movement[#32][-94.03].	Updated Balance[114.88] (2019-11-30 11:16:09.366548805)
[MAIN Goroutine].	Theroretical Current Balance[114.88].	Account Movement[#33][-15.48].	Updated Balance[99.40] (2019-11-30 11:16:09.366556558)
[MAIN Goroutine].	Theroretical Current Balance[99.40].	Account Movement[#34][0.95].	Updated Balance[100.35] (2019-11-30 11:16:09.366570057)
[MAIN Goroutine].	Theroretical Current Balance[100.35].	Account Movement[#35][18.02].	Updated Balance[118.37] (2019-11-30 11:16:09.366578299)
[MAIN Goroutine].	Theroretical Current Balance[118.37].	Account Movement[#36][70.39].	Updated Balance[188.76] (2019-11-30 11:16:09.366585998)
[MAIN Goroutine].	Theroretical Current Balance[188.76].	Account Movement[#37][-6.15].	Updated Balance[182.61] (2019-11-30 11:16:09.366599777)
[MAIN Goroutine].	Theroretical Current Balance[182.61].	Account Movement[#38][-80.81].	Updated Balance[101.80] (2019-11-30 11:16:09.366608434)
[MAIN Goroutine].	Theroretical Current Balance[101.80].	Account Movement[#39][11.00].	Updated Balance[112.80] (2019-11-30 11:16:09.366616268)
[MAIN Goroutine].	Theroretical Current Balance[112.80].	Account Movement[#40][-3.26].	Updated Balance[109.54] (2019-11-30 11:16:09.366624123)
[MAIN Goroutine].	Theroretical Current Balance[109.54].	Account Movement[#41][81.66].	Updated Balance[191.20] (2019-11-30 11:16:09.366662149)
[MAIN Goroutine].	Theroretical Current Balance[191.20].	Account Movement[#42][58.06].	Updated Balance[249.26] (2019-11-30 11:16:09.366709841)
[MAIN Goroutine].	Theroretical Current Balance[249.26].	Account Movement[#43][96.79].	Updated Balance[346.05] (2019-11-30 11:16:09.366720066)
[MAIN Goroutine].	Theroretical Current Balance[346.05].	Account Movement[#44][-97.54].	Updated Balance[248.51] (2019-11-30 11:16:09.366728729)
[MAIN Goroutine].	Theroretical Current Balance[248.51].	Account Movement[#45][-97.97].	Updated Balance[150.54] (2019-11-30 11:16:09.366736331)
[MAIN Goroutine].	Theroretical Current Balance[150.54].	Account Movement[#46][45.18].	Updated Balance[195.72] (2019-11-30 11:16:09.366771604)
[MAIN Goroutine].	Theroretical Current Balance[195.72].	Account Movement[#47][4.75].	Updated Balance[200.47] (2019-11-30 11:16:09.366815174)
[MAIN Goroutine].	Theroretical Current Balance[200.47].	Account Movement[#48][-84.80].	Updated Balance[115.67] (2019-11-30 11:16:09.366846334)
[MAIN Goroutine].	Theroretical Current Balance[115.67].	Account Movement[#49][75.03].	Updated Balance[190.70] (2019-11-30 11:16:09.366855461)
[MAIN Goroutine].	Theroretical Current Balance[190.70].	Account Movement[#50][31.05].	Updated Balance[221.75] (2019-11-30 11:16:09.366863992)
[MAIN Goroutine].	Theroretical Current Balance[221.75].	Account Movement[#51][-83.56].	Updated Balance[138.19] (2019-11-30 11:16:09.366873248)
[MAIN Goroutine].	Theroretical Current Balance[138.19].	Account Movement[#52][15.40].	Updated Balance[153.59] (2019-11-30 11:16:09.366944970)
[MAIN Goroutine].	Theroretical Current Balance[153.59].	Account Movement[#53][-43.92].	Updated Balance[109.67] (2019-11-30 11:16:09.366992857)
[MAIN Goroutine].	Theroretical Current Balance[109.67].	Account Movement[#54][16.62].	Updated Balance[126.29] (2019-11-30 11:16:09.367016766)
[MAIN Goroutine].	Theroretical Current Balance[126.29].	Account Movement[#55][66.13].	Updated Balance[192.42] (2019-11-30 11:16:09.367026145)
[MAIN Goroutine].	Theroretical Current Balance[192.42].	Account Movement[#56][38.67].	Updated Balance[231.09] (2019-11-30 11:16:09.367033644)
[MAIN Goroutine].	Theroretical Current Balance[231.09].	Account Movement[#57][-80.34].	Updated Balance[150.75] (2019-11-30 11:16:09.367046241)
[MAIN Goroutine].	Theroretical Current Balance[150.75].	Account Movement[#58][21.61].	Updated Balance[172.36] (2019-11-30 11:16:09.367060630)
[MAIN Goroutine].	Theroretical Current Balance[172.36].	Account Movement[#59][42.01].	Updated Balance[214.37] (2019-11-30 11:16:09.367074220)
[MAIN Goroutine].	Theroretical Current Balance[214.37].	Account Movement[#60][-94.07].	Updated Balance[120.30] (2019-11-30 11:16:09.367109286)

		--> Exiting from Goroutine [MAIN Goroutine]. Theoretical resulting Account Balance[120.30] (2019-11-30 11:16:09.367119483)

=========================================================
Resulting Balance [120.30].
=========================================================
Press ENTER to continue.

============================================================================================================
STAGE #02:
Now we will reset both values: the Cursor (position) of the list of pending movements and the value of
the initial Balance to their starting values
Then two threads (Goroutines) will be launched and execute concurrently creating a RACE CONDITION as:
1. The two threads will access concurrently trying to increase the shared value of the cursor of the
   movement list and updating the shared value of the balance with the corresponding movement of the list
2. Running this way final will drive to a totally different final Balance result each time
 
Note: now the main thread will be waiting sleeping during a reasonable time to permit all other threads
to succesfully conclude its tasks. DISCLAIMER: There are better way to sync. threads but they have not
shown/explained at this point in the course.
============================================================================================================

Press ENTER to continue.

=========================================================
RESETTING Position[0] and Initial Account Balance[300.00]
=========================================================

=========================================================
Initial Balance [300.00].
=========================================================
Main Thread is currently Awaiting [10] seconds to let all goroutines to safely end ...

		--> Entering Goroutine [Goroutine #02]. Theoretical initial Account Balance[300.00] (2019-11-30 11:16:11.082184280)

[Goroutine #02].	Theroretical Current Balance[300.00].	Account Movement[#0][40.50].	Updated Balance[340.50] (2019-11-30 11:16:11.082221645)

		--> Entering Goroutine [Goroutine #01]. Theoretical initial Account Balance[300.00] (2019-11-30 11:16:11.082143093)

[Goroutine #01].	Theroretical Current Balance[340.50].	Account Movement[#1][-20.72].	Updated Balance[319.78] (2019-11-30 11:16:11.082245627)
[Goroutine #02].	Theroretical Current Balance[319.78].	Account Movement[#2][89.01].	Updated Balance[408.79] (2019-11-30 11:16:11.126270693)
[Goroutine #01].	Theroretical Current Balance[319.78].	Account Movement[#2][89.01].	Updated Balance[408.79] (2019-11-30 11:16:11.126271460)
[Goroutine #01].	Theroretical Current Balance[408.79].	Account Movement[#4][-96.02].	Updated Balance[312.77] (2019-11-30 11:16:11.142320947)
[Goroutine #02].	Theroretical Current Balance[312.77].	Account Movement[#5][7.42].	Updated Balance[320.19] (2019-11-30 11:16:11.155129324)
[Goroutine #01].	Theroretical Current Balance[320.19].	Account Movement[#6][73.79].	Updated Balance[393.98] (2019-11-30 11:16:11.183561703)
[Goroutine #02].	Theroretical Current Balance[393.98].	Account Movement[#7][96.24].	Updated Balance[490.22] (2019-11-30 11:16:11.202049660)
[Goroutine #02].	Theroretical Current Balance[490.22].	Account Movement[#8][76.81].	Updated Balance[567.03] (2019-11-30 11:16:11.203997763)
[Goroutine #01].	Theroretical Current Balance[567.03].	Account Movement[#9][-32.98].	Updated Balance[534.05] (2019-11-30 11:16:11.225345159)
[Goroutine #02].	Theroretical Current Balance[534.05].	Account Movement[#10][24.76].	Updated Balance[558.81] (2019-11-30 11:16:11.231289347)
[Goroutine #02].	Theroretical Current Balance[558.81].	Account Movement[#11][17.48].	Updated Balance[576.29] (2019-11-30 11:16:11.242679723)
[Goroutine #02].	Theroretical Current Balance[576.29].	Account Movement[#12][-11.50].	Updated Balance[564.79] (2019-11-30 11:16:11.246424995)
[Goroutine #01].	Theroretical Current Balance[564.79].	Account Movement[#13][94.34].	Updated Balance[659.13] (2019-11-30 11:16:11.251223947)
[Goroutine #01].	Theroretical Current Balance[659.13].	Account Movement[#14][43.06].	Updated Balance[702.19] (2019-11-30 11:16:11.281787955)
[Goroutine #02].	Theroretical Current Balance[702.19].	Account Movement[#15][56.37].	Updated Balance[758.56] (2019-11-30 11:16:11.287042677)
[Goroutine #02].	Theroretical Current Balance[758.56].	Account Movement[#16][-29.41].	Updated Balance[729.15] (2019-11-30 11:16:11.306664201)
[Goroutine #01].	Theroretical Current Balance[729.15].	Account Movement[#17][-51.32].	Updated Balance[677.83] (2019-11-30 11:16:11.307318457)
[Goroutine #02].	Theroretical Current Balance[677.83].	Account Movement[#18][-53.99].	Updated Balance[623.84] (2019-11-30 11:16:11.312389463)
[Goroutine #02].	Theroretical Current Balance[623.84].	Account Movement[#19][-41.77].	Updated Balance[582.07] (2019-11-30 11:16:11.321730982)
[Goroutine #01].	Theroretical Current Balance[582.07].	Account Movement[#20][-33.04].	Updated Balance[549.03] (2019-11-30 11:16:11.343255777)
[Goroutine #02].	Theroretical Current Balance[549.03].	Account Movement[#21][-49.18].	Updated Balance[499.85] (2019-11-30 11:16:11.350252212)
[Goroutine #02].	Theroretical Current Balance[499.85].	Account Movement[#22][84.89].	Updated Balance[584.74] (2019-11-30 11:16:11.359537245)
[Goroutine #01].	Theroretical Current Balance[584.74].	Account Movement[#23][-49.91].	Updated Balance[534.83] (2019-11-30 11:16:11.366501989)
[Goroutine #02].	Theroretical Current Balance[534.83].	Account Movement[#24][34.41].	Updated Balance[569.24] (2019-11-30 11:16:11.375078745)
[Goroutine #02].	Theroretical Current Balance[569.24].	Account Movement[#25][-8.27].	Updated Balance[560.97] (2019-11-30 11:16:11.383425768)
[Goroutine #01].	Theroretical Current Balance[560.97].	Account Movement[#26][-83.34].	Updated Balance[477.63] (2019-11-30 11:16:11.410146250)
[Goroutine #02].	Theroretical Current Balance[477.63].	Account Movement[#27][-84.29].	Updated Balance[393.34] (2019-11-30 11:16:11.421407006)
[Goroutine #01].	Theroretical Current Balance[393.34].	Account Movement[#28][1.71].	Updated Balance[395.05] (2019-11-30 11:16:11.432499688)
[Goroutine #01].	Theroretical Current Balance[395.05].	Account Movement[#29][-75.72].	Updated Balance[319.33] (2019-11-30 11:16:11.446308452)
[Goroutine #01].	Theroretical Current Balance[319.33].	Account Movement[#30][-73.64].	Updated Balance[245.69] (2019-11-30 11:16:11.455419987)
[Goroutine #02].	Theroretical Current Balance[245.69].	Account Movement[#31][-97.00].	Updated Balance[148.69] (2019-11-30 11:16:11.470752192)
[Goroutine #01].	Theroretical Current Balance[148.69].	Account Movement[#32][-94.03].	Updated Balance[54.66] (2019-11-30 11:16:11.492305036)
[Goroutine #02].	Theroretical Current Balance[54.66].	Account Movement[#33][-15.48].	Updated Balance[39.18] (2019-11-30 11:16:11.500018828)
[Goroutine #02].	Theroretical Current Balance[39.18].	Account Movement[#34][0.95].	Updated Balance[40.13] (2019-11-30 11:16:11.507534704)
[Goroutine #02].	Theroretical Current Balance[40.13].	Account Movement[#35][18.02].	Updated Balance[58.15] (2019-11-30 11:16:11.517171068)
[Goroutine #02].	Theroretical Current Balance[58.15].	Account Movement[#36][70.39].	Updated Balance[128.54] (2019-11-30 11:16:11.527912734)
[Goroutine #01].	Theroretical Current Balance[128.54].	Account Movement[#37][-6.15].	Updated Balance[122.39] (2019-11-30 11:16:11.535045325)
[Goroutine #02].	Theroretical Current Balance[122.39].	Account Movement[#38][-80.81].	Updated Balance[41.58] (2019-11-30 11:16:11.551298038)
[Goroutine #01].	Theroretical Current Balance[41.58].	Account Movement[#39][11.00].	Updated Balance[52.58] (2019-11-30 11:16:11.565400760)
[Goroutine #02].	Theroretical Current Balance[52.58].	Account Movement[#40][-3.26].	Updated Balance[49.32] (2019-11-30 11:16:11.567577068)
[Goroutine #01].	Theroretical Current Balance[49.32].	Account Movement[#41][81.66].	Updated Balance[130.98] (2019-11-30 11:16:11.604414056)
[Goroutine #02].	Theroretical Current Balance[130.98].	Account Movement[#42][58.06].	Updated Balance[189.04] (2019-11-30 11:16:11.604737525)
[Goroutine #01].	Theroretical Current Balance[189.04].	Account Movement[#43][96.79].	Updated Balance[285.83] (2019-11-30 11:16:11.612490152)
[Goroutine #02].	Theroretical Current Balance[285.83].	Account Movement[#44][-97.54].	Updated Balance[188.29] (2019-11-30 11:16:11.632756692)
[Goroutine #02].	Theroretical Current Balance[188.29].	Account Movement[#45][-97.97].	Updated Balance[90.32] (2019-11-30 11:16:11.636573297)
[Goroutine #01].	Theroretical Current Balance[90.32].	Account Movement[#46][45.18].	Updated Balance[135.50] (2019-11-30 11:16:11.654705004)
[Goroutine #01].	Theroretical Current Balance[135.50].	Account Movement[#47][4.75].	Updated Balance[140.25] (2019-11-30 11:16:11.664916435)
[Goroutine #01].	Theroretical Current Balance[140.25].	Account Movement[#48][-84.80].	Updated Balance[55.45] (2019-11-30 11:16:11.679060808)
[Goroutine #02].	Theroretical Current Balance[55.45].	Account Movement[#49][75.03].	Updated Balance[130.48] (2019-11-30 11:16:11.684322049)
[Goroutine #01].	Theroretical Current Balance[130.48].	Account Movement[#50][31.05].	Updated Balance[161.53] (2019-11-30 11:16:11.703498454)
[Goroutine #01].	Theroretical Current Balance[161.53].	Account Movement[#51][-83.56].	Updated Balance[77.97] (2019-11-30 11:16:11.713302384)
[Goroutine #02].	Theroretical Current Balance[77.97].	Account Movement[#52][15.40].	Updated Balance[93.37] (2019-11-30 11:16:11.713690605)
[Goroutine #02].	Theroretical Current Balance[93.37].	Account Movement[#53][-43.92].	Updated Balance[49.45] (2019-11-30 11:16:11.720130210)
[Goroutine #01].	Theroretical Current Balance[49.45].	Account Movement[#54][16.62].	Updated Balance[66.07] (2019-11-30 11:16:11.724963020)
[Goroutine #02].	Theroretical Current Balance[66.07].	Account Movement[#55][66.13].	Updated Balance[132.20] (2019-11-30 11:16:11.739621094)
[Goroutine #02].	Theroretical Current Balance[132.20].	Account Movement[#56][38.67].	Updated Balance[170.87] (2019-11-30 11:16:11.742017931)
[Goroutine #02].	Theroretical Current Balance[170.87].	Account Movement[#57][-80.34].	Updated Balance[90.53] (2019-11-30 11:16:11.775150215)
[Goroutine #01].	Theroretical Current Balance[170.87].	Account Movement[#57][-80.34].	Updated Balance[90.53] (2019-11-30 11:16:11.775198728)
[Goroutine #02].	Theroretical Current Balance[90.53].	Account Movement[#59][42.01].	Updated Balance[132.54] (2019-11-30 11:16:11.778854096)
[Goroutine #02].	Theroretical Current Balance[132.54].	Account Movement[#60][-94.07].	Updated Balance[38.47] (2019-11-30 11:16:11.811088854)

		--> Exiting from Goroutine [Goroutine #01]. Theoretical resulting Account Balance[90.53] (2019-11-30 11:16:11.815929402)


		--> Exiting from Goroutine [Goroutine #02]. Theoretical resulting Account Balance[38.47] (2019-11-30 11:16:11.820382072)

=========================================================
Resulting Balance [38.47].
=========================================================
Press ENTER to continue.

============================================================================================================
STAGE #03:
Each thread has been writing in a log its activity showing at every moment in time the values that 'see'
and update the shared variables: Cursor and Current balance.
Now it's time to see how they were behaving individually.
============================================================================================================

Press ENTER to continue.

========================================================================
Showing now the event's history registered at Goroutine [Goroutine #01]
========================================================================
------------------------------------------------------------------------------------------------------------------------------------------------------
		--> Entering Goroutine [Goroutine #01]. Theoretical initial Account Balance[300.00] (2019-11-30 11:16:11.082143093)
------------------------------------------------------------------------------------------------------------------------------------------------------
[Goroutine #01].	Theroretical Current Balance[340.50].	Account Movement[#1][-20.72].	Updated Balance[319.78] (2019-11-30 11:16:11.082245627)
[Goroutine #01].	Theroretical Current Balance[319.78].	Account Movement[#2][89.01].	Updated Balance[408.79] (2019-11-30 11:16:11.126271460)
[Goroutine #01].	Theroretical Current Balance[408.79].	Account Movement[#4][-96.02].	Updated Balance[312.77] (2019-11-30 11:16:11.142320947)
[Goroutine #01].	Theroretical Current Balance[320.19].	Account Movement[#6][73.79].	Updated Balance[393.98] (2019-11-30 11:16:11.183561703)
[Goroutine #01].	Theroretical Current Balance[567.03].	Account Movement[#9][-32.98].	Updated Balance[534.05] (2019-11-30 11:16:11.225345159)
[Goroutine #01].	Theroretical Current Balance[564.79].	Account Movement[#13][94.34].	Updated Balance[659.13] (2019-11-30 11:16:11.251223947)
[Goroutine #01].	Theroretical Current Balance[659.13].	Account Movement[#14][43.06].	Updated Balance[702.19] (2019-11-30 11:16:11.281787955)
[Goroutine #01].	Theroretical Current Balance[729.15].	Account Movement[#17][-51.32].	Updated Balance[677.83] (2019-11-30 11:16:11.307318457)
[Goroutine #01].	Theroretical Current Balance[582.07].	Account Movement[#20][-33.04].	Updated Balance[549.03] (2019-11-30 11:16:11.343255777)
[Goroutine #01].	Theroretical Current Balance[584.74].	Account Movement[#23][-49.91].	Updated Balance[534.83] (2019-11-30 11:16:11.366501989)
[Goroutine #01].	Theroretical Current Balance[560.97].	Account Movement[#26][-83.34].	Updated Balance[477.63] (2019-11-30 11:16:11.410146250)
[Goroutine #01].	Theroretical Current Balance[393.34].	Account Movement[#28][1.71].	Updated Balance[395.05] (2019-11-30 11:16:11.432499688)
[Goroutine #01].	Theroretical Current Balance[395.05].	Account Movement[#29][-75.72].	Updated Balance[319.33] (2019-11-30 11:16:11.446308452)
[Goroutine #01].	Theroretical Current Balance[319.33].	Account Movement[#30][-73.64].	Updated Balance[245.69] (2019-11-30 11:16:11.455419987)
[Goroutine #01].	Theroretical Current Balance[148.69].	Account Movement[#32][-94.03].	Updated Balance[54.66] (2019-11-30 11:16:11.492305036)
[Goroutine #01].	Theroretical Current Balance[128.54].	Account Movement[#37][-6.15].	Updated Balance[122.39] (2019-11-30 11:16:11.535045325)
[Goroutine #01].	Theroretical Current Balance[41.58].	Account Movement[#39][11.00].	Updated Balance[52.58] (2019-11-30 11:16:11.565400760)
[Goroutine #01].	Theroretical Current Balance[49.32].	Account Movement[#41][81.66].	Updated Balance[130.98] (2019-11-30 11:16:11.604414056)
[Goroutine #01].	Theroretical Current Balance[189.04].	Account Movement[#43][96.79].	Updated Balance[285.83] (2019-11-30 11:16:11.612490152)
[Goroutine #01].	Theroretical Current Balance[90.32].	Account Movement[#46][45.18].	Updated Balance[135.50] (2019-11-30 11:16:11.654705004)
[Goroutine #01].	Theroretical Current Balance[135.50].	Account Movement[#47][4.75].	Updated Balance[140.25] (2019-11-30 11:16:11.664916435)
[Goroutine #01].	Theroretical Current Balance[140.25].	Account Movement[#48][-84.80].	Updated Balance[55.45] (2019-11-30 11:16:11.679060808)
[Goroutine #01].	Theroretical Current Balance[130.48].	Account Movement[#50][31.05].	Updated Balance[161.53] (2019-11-30 11:16:11.703498454)
[Goroutine #01].	Theroretical Current Balance[161.53].	Account Movement[#51][-83.56].	Updated Balance[77.97] (2019-11-30 11:16:11.713302384)
[Goroutine #01].	Theroretical Current Balance[49.45].	Account Movement[#54][16.62].	Updated Balance[66.07] (2019-11-30 11:16:11.724963020)
[Goroutine #01].	Theroretical Current Balance[170.87].	Account Movement[#57][-80.34].	Updated Balance[90.53] (2019-11-30 11:16:11.775198728)
------------------------------------------------------------------------------------------------------------------------------------------------------
		--> Exiting from Goroutine [Goroutine #01]. Theoretical resulting Account Balance[90.53] (2019-11-30 11:16:11.815929402)
------------------------------------------------------------------------------------------------------------------------------------------------------
Press ENTER to continue.

========================================================================
Showing now the event's history registered at Goroutine [Goroutine #02]
========================================================================
------------------------------------------------------------------------------------------------------------------------------------------------------
		--> Entering Goroutine [Goroutine #02]. Theoretical initial Account Balance[300.00] (2019-11-30 11:16:11.082184280)
------------------------------------------------------------------------------------------------------------------------------------------------------
[Goroutine #02].	Theroretical Current Balance[300.00].	Account Movement[#0][40.50].	Updated Balance[340.50] (2019-11-30 11:16:11.082221645)
[Goroutine #02].	Theroretical Current Balance[319.78].	Account Movement[#2][89.01].	Updated Balance[408.79] (2019-11-30 11:16:11.126270693)
[Goroutine #02].	Theroretical Current Balance[312.77].	Account Movement[#5][7.42].	Updated Balance[320.19] (2019-11-30 11:16:11.155129324)
[Goroutine #02].	Theroretical Current Balance[393.98].	Account Movement[#7][96.24].	Updated Balance[490.22] (2019-11-30 11:16:11.202049660)
[Goroutine #02].	Theroretical Current Balance[490.22].	Account Movement[#8][76.81].	Updated Balance[567.03] (2019-11-30 11:16:11.203997763)
[Goroutine #02].	Theroretical Current Balance[534.05].	Account Movement[#10][24.76].	Updated Balance[558.81] (2019-11-30 11:16:11.231289347)
[Goroutine #02].	Theroretical Current Balance[558.81].	Account Movement[#11][17.48].	Updated Balance[576.29] (2019-11-30 11:16:11.242679723)
[Goroutine #02].	Theroretical Current Balance[576.29].	Account Movement[#12][-11.50].	Updated Balance[564.79] (2019-11-30 11:16:11.246424995)
[Goroutine #02].	Theroretical Current Balance[702.19].	Account Movement[#15][56.37].	Updated Balance[758.56] (2019-11-30 11:16:11.287042677)
[Goroutine #02].	Theroretical Current Balance[758.56].	Account Movement[#16][-29.41].	Updated Balance[729.15] (2019-11-30 11:16:11.306664201)
[Goroutine #02].	Theroretical Current Balance[677.83].	Account Movement[#18][-53.99].	Updated Balance[623.84] (2019-11-30 11:16:11.312389463)
[Goroutine #02].	Theroretical Current Balance[623.84].	Account Movement[#19][-41.77].	Updated Balance[582.07] (2019-11-30 11:16:11.321730982)
[Goroutine #02].	Theroretical Current Balance[549.03].	Account Movement[#21][-49.18].	Updated Balance[499.85] (2019-11-30 11:16:11.350252212)
[Goroutine #02].	Theroretical Current Balance[499.85].	Account Movement[#22][84.89].	Updated Balance[584.74] (2019-11-30 11:16:11.359537245)
[Goroutine #02].	Theroretical Current Balance[534.83].	Account Movement[#24][34.41].	Updated Balance[569.24] (2019-11-30 11:16:11.375078745)
[Goroutine #02].	Theroretical Current Balance[569.24].	Account Movement[#25][-8.27].	Updated Balance[560.97] (2019-11-30 11:16:11.383425768)
[Goroutine #02].	Theroretical Current Balance[477.63].	Account Movement[#27][-84.29].	Updated Balance[393.34] (2019-11-30 11:16:11.421407006)
[Goroutine #02].	Theroretical Current Balance[245.69].	Account Movement[#31][-97.00].	Updated Balance[148.69] (2019-11-30 11:16:11.470752192)
[Goroutine #02].	Theroretical Current Balance[54.66].	Account Movement[#33][-15.48].	Updated Balance[39.18] (2019-11-30 11:16:11.500018828)
[Goroutine #02].	Theroretical Current Balance[39.18].	Account Movement[#34][0.95].	Updated Balance[40.13] (2019-11-30 11:16:11.507534704)
[Goroutine #02].	Theroretical Current Balance[40.13].	Account Movement[#35][18.02].	Updated Balance[58.15] (2019-11-30 11:16:11.517171068)
[Goroutine #02].	Theroretical Current Balance[58.15].	Account Movement[#36][70.39].	Updated Balance[128.54] (2019-11-30 11:16:11.527912734)
[Goroutine #02].	Theroretical Current Balance[122.39].	Account Movement[#38][-80.81].	Updated Balance[41.58] (2019-11-30 11:16:11.551298038)
[Goroutine #02].	Theroretical Current Balance[52.58].	Account Movement[#40][-3.26].	Updated Balance[49.32] (2019-11-30 11:16:11.567577068)
[Goroutine #02].	Theroretical Current Balance[130.98].	Account Movement[#42][58.06].	Updated Balance[189.04] (2019-11-30 11:16:11.604737525)
[Goroutine #02].	Theroretical Current Balance[285.83].	Account Movement[#44][-97.54].	Updated Balance[188.29] (2019-11-30 11:16:11.632756692)
[Goroutine #02].	Theroretical Current Balance[188.29].	Account Movement[#45][-97.97].	Updated Balance[90.32] (2019-11-30 11:16:11.636573297)
[Goroutine #02].	Theroretical Current Balance[55.45].	Account Movement[#49][75.03].	Updated Balance[130.48] (2019-11-30 11:16:11.684322049)
[Goroutine #02].	Theroretical Current Balance[77.97].	Account Movement[#52][15.40].	Updated Balance[93.37] (2019-11-30 11:16:11.713690605)
[Goroutine #02].	Theroretical Current Balance[93.37].	Account Movement[#53][-43.92].	Updated Balance[49.45] (2019-11-30 11:16:11.720130210)
[Goroutine #02].	Theroretical Current Balance[66.07].	Account Movement[#55][66.13].	Updated Balance[132.20] (2019-11-30 11:16:11.739621094)
[Goroutine #02].	Theroretical Current Balance[132.20].	Account Movement[#56][38.67].	Updated Balance[170.87] (2019-11-30 11:16:11.742017931)
[Goroutine #02].	Theroretical Current Balance[170.87].	Account Movement[#57][-80.34].	Updated Balance[90.53] (2019-11-30 11:16:11.775150215)
[Goroutine #02].	Theroretical Current Balance[90.53].	Account Movement[#59][42.01].	Updated Balance[132.54] (2019-11-30 11:16:11.778854096)
[Goroutine #02].	Theroretical Current Balance[132.54].	Account Movement[#60][-94.07].	Updated Balance[38.47] (2019-11-30 11:16:11.811088854)
------------------------------------------------------------------------------------------------------------------------------------------------------
		--> Exiting from Goroutine [Goroutine #02]. Theoretical resulting Account Balance[38.47] (2019-11-30 11:16:11.820382072)
------------------------------------------------------------------------------------------------------------------------------------------------------
=========================================================
Resulting Balance [38.47].
=========================================================
Press ENTER to continue.
GOODBYE !!!
```

