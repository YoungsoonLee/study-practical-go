package main


"AU-L37YDU"
d07fa03e-da5e-ec11-8f8f-000d3ad217e6

75001b87-e15e-ec11-8f8f-000d3ad217e6
78001b87-e15e-ec11-8f8f-000d3ad217e6

"AU-L37YDU"로 로그 조사.
1. 2021-12-16T17:40:00.833-08:00 이 시간 쯤에 GSCV -> D365로 CREATE 호출, 성공. 이때는 loyalty 없음
2. 2021-12-16T18:32:15.486-08:00 이 시간에 GSCV -> D365로 UPDATE 호출, 성공. 
	- contact는 DOB, gender,mobilephone 업데이트 되고	
	- loyalty가 2개가 create 되고, 그 결과로 아래 각각의 loyalty ID를 D365로 부터 받음
		75001b87-e15e-ec11-8f8f-000d3ad217e6
		78001b87-e15e-ec11-8f8f-000d3ad217e6
 
그 이후  GSCV -> D365로 삭제 요청한 로그가 없음.

noti
1. delete noti가 있음. (D365 -> GSCV)
2021-12-16T18:43:01.069-08:00: 75001b87-e15e-ec11-8f8f-000d3ad217e6
2021-12-16T18:43:05.360-08:00 : 78001b87-e15e-ec11-8f8f-000d3ad217e6


계속 D365 -> GSCV 업데이트 실패
그래서 Dynamo DB 업데이트 실패...




3.105.235.25
13.237.17.207 //uat