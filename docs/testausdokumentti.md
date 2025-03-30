[![codecov](https://codecov.io/gh/Eeritvan/face_recognition/graph/badge.svg?token=VZZML0709G)](https://codecov.io/gh/Eeritvan/face_recognition)

## yksikkötestit
Ohjelma sisältää testejä. Testit on sijoitettu koodin kanssa samoihin hakemistoihin. #todo

## testien suorittaminen
Testit voi suorittaa Go:n omalla testauskirjastolla. Sen pitäisi tulla Go:n asennuksen mukana joten sitä varten ei tarvitse ladata mitään. 

```bash
cd src
go test ./...

## testikattavuuden voi tulostaa: 
go test ./... -cover
```