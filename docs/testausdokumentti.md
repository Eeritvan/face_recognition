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
- QR-hajotelma (householder) palauttaa oikeat Q ja R matriisit jotka on varmistettu oikeaksi laskimella. Testit myös varmistaa, että Q * R muodostaa alkuperäisen matriisin takaisin.
- QR-algoritmi palauttaa oikeat ominaisarvot ja ominaisvektorit kokonais- ja liukuluvuilla. Tulokset on laskettu syötteillä etukäteen laskimella ja funktion tulosta verrataan näihin joilla varmistetaan, että ne on oikein.

"pää funktioiden" testit:
- eigenface funktio palauttaa oikeat eigenface vektorit ja kuvien keskiarvon. Tulokset on laskettu laskimella ja testit vertaa funktion syötettä oikeisiin tuloksiin.
- eigenface projektointi testattu palauttavan oikeat projektio matriisit. Tulokset on laskettu laskimella ja funktion syötettä verrataan niihin.
- lähimmän kuvan ja euklidisen etäisyyden laskeva funktio todettu toimivaksi vertaamalla funktion palauttavia arvoja laskimella laskettuihin tiedettyihin tuloksiin. 
- lähimmän kuvan ja euklidisen etäisyyden laskeva palauttaa etäisyyden 0 ja saman kuvan takaisin kun kuva löytyy jo harjoitusdatasta. 

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