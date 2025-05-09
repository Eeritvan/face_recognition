## yleisrakenne
Ohjelma on kasvojentunnistus algoritmi käyttäen eigenface-menetelmää.

Ohjelma koostuu useammasta eri moduulista tai go-paketista (eri hakemistot src-alla) jotka vastaavat jostakin toiminnallisuudesta. Esimerkiksi "image" vastaa kuvien lataamisesta, "matrix" eri matriisioperaatioista ja "qr" qr-algoritmin toteutuksesta. 

Eigenface-menetelmä toteutetaan aluksi valitsemalla jokin määrä harjoitusdataa ja laskemalla kaikkien harjoitus kuvien keskiarvo. Tämän jälkeen keskiarvoa voidaan käyttää laskemaan kuinka paljon jokainen harjoitusdatan kuva eroaa keskiarvosta ja muodostaa siitä matriisi jossa korostuu treenaus kuvien uniikit ja erikoispiirteet. Tätä voidaan taas käyttää kovarianssimatriisin laskemiseen jolla saadaan ominaisarvot sekä ominaisvektorit selvitettyä QR-algoritmilla ja householderin reflektiolla. Tämän jälkeen ominaisvektoreista voidaan valita k suurinta ja laskea projektio ominaisavaruuteen. Kun testikuvat myös projektoidaan samaan avaruuteen voidaan laskea testidatan ja harjoitusdatan etäisyys josta voidaan luoda arvio onko kuva kasvo vai ei tai miten todennäköisesti se on kasvo.

Tämä prosessi käytännössä vastaa pääkomponenttianalyysiä (PCA), jossa lasketaan ne pääkomponentit, jotka parhaiten kuvaavat datan varianssia. PCA auttaa siis tunnistamaan ne ominaispiirteet, jotka ovat merkityksellisimpiä kasvojentunnistuksen kannalta.

## aika vaativuudet
- matriisitulo = O(n**3), missä n on n * n matriisin sarakkeiden/rivien määrä
- QR-algoritmi = O(k * n**3), missä n on kovarianssimatriisin koko ja k laskemiseen tarvittavien iteraatioiden määrä (max 1000)
- Kuvien projektointi = O(n * b * k), missä m on kuvien määrä, b on sarakkeiden määrä jokaisessa kuvassa ja k on eigenfaces määrä
- lähimmän osuman etsiminen = O(n * d), missä n on kuvien määrä ja d on projektoidujen kuvien dimensio

## puutteet ja parannusehdotukset
Ohjelma nyt vaati etukäteen, että käyttäjä tuntee harjoitusdatan jotenkin erityisesti jos haluaa testata juuri tietyillä kuvilla ja syötteillä. Varmasti jokin graafinen käyttöliittymä mistä voi nähdä valittavat kuvat ja esimerkiksi näyttää mikä on samankaltaisin kuva treenausdatassa. Toinen hyvä parannus olisi antaa mahdollisuus lisätä itse harjoitusdataa tai testikuvaksi omia kuvia. Olisi myös mielenkiintoista nähdä mitä tapahtuu jos treenausdataan lisää muutamia ei-kasvo kuvia ja miten ne vaikuttavat tuloksiin tai miten algoritmi reagoi kun yrittää vertailla jotakin ei-kasvoa nykyiseen treenausdataan. 

## AI käyttö
Käytin claude 3.5 sonnet mallia keksimään joillekin muuttujille ja funktioille parempia nimiä sekä kirjoittamaan joillekin funktioille parempia kommentteja. Käytin myös gemini 2.5 pro mallia optimoimaan ja tehostamaan QR-algoritmia ja Householderin reflektiota kurssin lopussa.

## lähteet
- https://en.wikipedia.org/wiki/Eigenface
- https://en.wikipedia.org/wiki/QR_algorithm
https://en.wikipedia.org/wiki/QR_decomposition#Using_Householder_reflections
- https://km.pcz.pl/konferencja/dokumenty/MMFT2017/Caban%20L.pdf
- https://www.geeksforgeeks.org/ml-face-recognition-using-eigenfaces-pca-algorithm/
- https://www.youtube.com/watch?v=n0zDgkbFyQk
- https://www.youtube.com/watch?v=McHW221J3UM

data:
- https://www.cl.cam.ac.uk/research/dtg/attarchive/facedatabase.html
