[![codecov](https://codecov.io/gh/Eeritvan/face_recognition/graph/badge.svg?token=VZZML0709G)](https://codecov.io/gh/Eeritvan/face_recognition)

## yksikkötestit
Ohjelma sisältää testejä. Testit on sijoitettu koodin kanssa samoihin hakemistoihin.

Matriisioperaatioiden testit:
- Kaikki matriisioperaatiot on todettu toimivaksi kokonaisluvuilla vertaamalla tulosta todelliseen vastaukseen
- Kaikki matriisioperaatiot on todettu toimivaksi liukuluvuilla vertaamalla tulosta todelliseen vastaukseen
- matriisitulo, skalaaritulo, yhteen- ja vähennyslasku on testattu toimiviksi "isoilla" 92x112 matriiseilla
- matriisitulo, skalaaritulo ja identiteettimatriisi on testattu palauttavan oikean arvon kun kerrotaan 0.

Kuvanlataus testit:
- kuvan lataaminen palauttaa oikean virheen kun kuvaa / tiedostoa ei löydy
- kuvan "litistäminen" on testattu toimivaksi kokonais- ja liukuluvuilla
- kuvien keskiarvon laskeva funktio palauttaa oikean virheen kun matriisien koot eivät vastaa toisiaan
- kuvien keskiarvon laskea funktio toimii kokonais- ja liukuluvuilla.

QR-algoritmin testit:
- Householderin vektorin laskeva funktio palauttaa oikeat arvot eri sarakkeilla sekä toimii positiivisilla että negatiivisilla luvuilla
- normivektorin laskeminen testattu toimivan positiivisilla ja negatiivisilla kokonaisluvuilla.
- normivektorin laskeminen testattu palauttavan 0 kun vektori on nollavektori.
#todo

## testien suorittaminen
Testit voi suorittaa Go:n omalla testauskirjastolla. Sen pitäisi tulla Go:n asennuksen mukana joten sitä varten ei tarvitse ladata mitään. 

```bash
cd src
go test ./...

## testikattavuuden voi tulostaa: 
go test ./... -cover

## testikattavuuden ja testit voi myös tulostaa make avulla juurikansiosta: 
make test
```