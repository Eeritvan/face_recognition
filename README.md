# eigenface
[![checking code](https://github.com/Eeritvan/face_recognition/actions/workflows/main.yml/badge.svg)](https://github.com/Eeritvan/face_recognition/actions/workflows/main.yml) [![codecov](https://codecov.io/gh/Eeritvan/face_recognition/graph/badge.svg?token=VZZML0709G)](https://codecov.io/gh/Eeritvan/face_recognition)
## dokumentaatio
- [määrittelydokumentti](docs/maarittelydokumentti.md)
- [testausdokumentti](docs/testausdokumentti.md)
- [toteutusdokumentti](docs/toteutusdokumentti.md)

## viikkoraportit
- [viikko 1](docs/viikkopalautukset/viikko1.md)
- [viikko 2](docs/viikkopalautukset/viikko2.md)
- [viikko 3](docs/viikkopalautukset/viikko3.md)
- [viikko 4](docs/viikkopalautukset/viikko4.md)
- [viikko 5](docs/viikkopalautukset/viikko5.md)
- [viikko 6](docs/viikkopalautukset/viikko6.md)


## Käyttöohje
### esivaatimukset
- Go 1.24.1 tai uudempi
- Make (vapaaehtoinen)

### lataaminen
```bash
git clone https://github.com/Eeritvan/face_recognition.git
cd face_recognition
```

## ajaminen
Sovellusta pystyy käyttämään kahdella tavalla:
- suoraan komentoriviltä
- terminaalissa toimivalla käyttöliittymällä


### suoraan komentoriviltä
argumentit näkee alhaalta.
```bash
#Käyttämällä make juurikansiosta:
make ARGS="..."


# tai käyttämällä Go:
cd src
go run . (ARGS...)
```
### terminaalissa toimivalla 
```bash
#Käyttämällä make juurikansiosta:
make


# tai käyttämällä Go:
cd src
go run . 
```

## toiminnot
ohjelmassa voi asettaa joitakin asetuksia kuten:
- valita mitä kuva settejä käytetään treenausdatana. Liian paljon treenausdataa ei kuitenkaan paranna tulosta vaan voi johtaa ylimääräisen kohinan tai turhien yksityiskohtien ylikorostumiseen.
- säätää eigenfaces määrä joka käytännössä tarkoittaa kuinka paljon yksityiskohtia jokaisesta treenaisdatan kuvasta säilytetään. Yleensä pienemmät arvot antavat parempia tuloksia, sillä suuremmilla arvoilla kohina ja turhat yksityiskohdat voi ylikorostua
- valita mitä kuvaa käytetään testidatana ja jota verrataan treenausdataan
- valita kuinka monta harjoituskuvaa jokaisesta treenausdatan setistä valitaan.

#### Argumentit komentorivi tilalle:
Kaikki toiminnot myös ohjelmassa näkee käyttämällä "-h" argumenttia.

- `-h` näyttää terminaalissa kaikki asetukset, vaihtoehdot ja esimerkkejä
- `-t` näyttää kuinka kauan algoritmissä kestää eri vaiheiden suorittamiseen
- `-k <num>` antaa valita kuinka monta eigenface kuvaa algoritmi käyttää. Vakioasetus on 5
- `-s <num num>` antaa valita testattavan kuvan itse. Ensimmäinen numero valitsee setin / henkilön (1-40) ja toinen numero mitä kuvaa setistä käytetään (1-10). Vakiona ohjelma ohjelma arpoo jonkin kuvan.
- `-d <num ...>` antaa valita käytettävän treenausdatan setit (esim. 1 2 5). Vakiona ohjelma arpoo kaksi settiä joita algoritmi käyttää.
- `-i <num>` antaa valita ladattavien kuvien määrän jokaisesta datasetitstä joissa jokaisessa on 10 kuvaa. i voi olla 1-10. Oletuksena i on 10 eli kaikki kuvat käytetään.

> huom!<br>
> käytettävien kuvien määrä kannattaa olla enintään 15 sillä algoritmi on muuten melko hidas