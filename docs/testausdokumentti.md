[![codecov](https://codecov.io/gh/Eeritvan/face_recognition/graph/badge.svg?token=VZZML0709G)](https://codecov.io/gh/Eeritvan/face_recognition)

## yksikkötestit
Testit on sijoitettu testattavan koodin kanssa samoihin hakemistoihin.

Matriisioperaatioiden testit:
- Kaikki matriisioperaatiot on todettu toimivaksi kokonaisluvuilla vertaamalla tulosta oikeaan vastaukseen
- Kaikki matriisioperaatiot on todettu toimivaksi liukuluvuilla vertaamalla tulosta oikeaan vastaukseen
- Matriisitulo, yhteen- ja vähennyslasku testattu palauttavan oikean virheen kun matriisien koot ovat väärät 
- matriisitulo, skalaaritulo ja identiteettimatriisi on testattu palauttavan oikean arvon kun kerrotaan 0
- matriisitulo, skalaaritulo, yhteen- ja vähennyslasku on testattu toimiviksi "isoilla" 92x112 matriiseilla
- kovarianssimatriisin laskeva funktio palauttaa oikean tuloksen kun syöte on kokonais- tai liukuluku
- eigenvectorit sorttaava funktio järjestää vektorit oikein ominaisarvojen mukaan.

Kuvanlataus testit:
- kuvan lataaminen palauttaa oikean virheen kun kuvaa / tiedostoa ei löydy
- kuvan "litistäminen" on testattu toimivaksi kokonais- ja liukuluvuilla
- kuvien keskiarvon laskeva funktio palauttaa virheen kun matriisien koot eivät vastaa toisiaan
- kuvien keskiarvon laskea funktio toimii kokonais- ja liukuluvuilla.

QR-algoritmin testit:
- householderin vektorin laskeva funktio palauttaa oikeat arvot eri sarakkeilla sekä toimii positiivisilla että negatiivisilla luvuilla. Testit vertaavat tulosta odotettuun tulokseen
- R-matriisin ja Q-matriisin päivitys householderin vektorilla testataan vertaamalla tulosta odotettuun matriisiin eri syötteillä
- qr-hajotelma palauttaa oikeat Q- ja R-matriisit, jotka on varmistettu vertaamalla laskettuihin arvoihin. Testit tarkistavat myös, että Q on ortogonaalinen (QᵗQ = I) ja että Q*R palauttaa alkuperäisen matriisin
- qr-algoritmi palauttaa oikeat ominaisarvot ja ominaisvektorit sekä kokonais- että liukuluvuilla. Tuloksia verrataan etukäteen laskettuihin tuloksiin

"pää funktioiden" testit:
- eigenface funktio palauttaa oikeat eigenface vektorit ja kuvien keskiarvon. Tulokset on laskettu laskimella ja testit vertaa funktion syötettä näihin tuloksiin
- eigenface projektointi testattu palauttavan oikeat projektio matriisit. Tulokset on laskettu laskimella ja funktion syötettä verrataan niihin
- lähimmän kuvan ja euklidisen etäisyyden laskeva funktio todettu toimivaksi vertaamalla funktion palauttavia arvoja laskimella laskettuihin tiedettyihin tuloksiin
- lähimmän kuvan ja euklidisen etäisyyden laskeva palauttaa etäisyyden 0 ja saman kuvan takaisin kun kuva löytyy jo harjoitusdatasta

## integraatiotestit
Sovellus testaa että algoritmi palauttaa oikeita tuloksia etukäteen validoilulla ja testatulla datalla. Integraatiotesteillä testaan, että:
- jos testikuva on jo harjoitusdatassa kuva on 100% varmuudella kasvo
- jos kuvaa ei ole datassa niin kuvan todennäköisyys olla kasvo on alle 100%
- ohjelma osaa palauttaa oikean eniten samanlaisen kasvon indeksin takaisin
- liian suuri k (eigenfaces määrä) arvo palauttaa virheen
- ohjelma palauttaa validin tuloksen kun treenausdataa on paljon (8 täyttä data settiä = 80 kuvaa)

## testien suorittaminen
Testit voi suorittaa Go:n omalla testauskirjastolla. Sen pitäisi tulla Go:n asennuksen mukana joten sitä varten ei tarvitse ladata mitään. Testit eivät testaa main tai cli paketteja sillä ne sisältävät lähinnä käyttöliittymän koodia joka eivät liity algoritmin toimintaan.

```bash
cd src
go test $(cat testdirs.txt)

## testikattavuuden voi tulostaa: 
go test $(cat testdirs.txt) -cover

## testikattavuuden ja testit voi myös tulostaa make avulla juurikansiosta: 
make test
```
