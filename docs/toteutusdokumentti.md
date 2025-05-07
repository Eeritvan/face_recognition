## yleisrakenne
Ohjelma on kasvojentunnistus algoritmi käyttäen eigenface-menetelmää.

Ohjelma koostuu useammasta eri moduulista tai go-paketista (eri hakemistot src-alla) jotka vastaavat jostakin toiminnallisuudesta. Esimerkiksi "image" vastaa kuvien lataamisesta, "matrix" eri matriisioperaatioista ja "qr" qr-algoritmin toteutuksesta. 

Eigenface-menetelmä toteutetaan aluksi valitsemalla jokin määrä harjoitusdataa ja laskemalla kaikkien harjoitus kuvien keskiarvo. Tämän jälkeen keskiarvoa voidaan käyttää laskemaan kuinka paljon jokainen harjoitusdatan kuva eroaa keskiarvosta ja muodostaa siitä matriisi. Tätä voidaan taas käyttää kovarianssimatriisin laskemiseen jolla saadaan ominaisarvot sekä ominaisvektorit selvitettyä QR-algoritmilla. Tämän jälkeen ominaisvektoreista voidaan valita k suurinta ja laskea projektio ominaisavaruuteen. Kun testikuvat myös projektoidaan samaan avaruuteen voidaan laskea testidatan ja harjoitusdatan etäisyys josta voidaan luoda arvio onko kuva kasvo vai ei.

Tämä prosessi käytännössä vastaa pääkomponenttianalyysiä (PCA), jossa lasketaan ne pääkomponentit, jotka parhaiten kuvaavat datan varianssia. PCA auttaa siis tunnistamaan ne ominaispiirteet, jotka ovat merkityksellisimpiä kasvojentunnistuksen kannalta.

## aika vaativuudet
- matriisitulo = O(n**3)
- QR-algoritmi = O(k * n**3)
- Kuvien projektointi = O(m * k * b) 
- lähimmän osuman etsiminen = O(k * n)

n = koko
k = eigenfaces

## puutteet ja parannusehdotukset
#todo

## AI käyttö
Olen käyttänyt claude 3.5 sonnet mallia keksimään joillekin muuttujille ja funktioille parempia nimiä

## lähteet
- https://en.wikipedia.org/wiki/Eigenface
- https://en.wikipedia.org/wiki/QR_algorithm
- https://km.pcz.pl/konferencja/dokumenty/MMFT2017/Caban%20L.pdf
- https://www.geeksforgeeks.org/ml-face-recognition-using-eigenfaces-pca-algorithm/
- https://www.youtube.com/watch?v=n0zDgkbFyQk
- https://www.youtube.com/watch?v=McHW221J3UM

data:
- https://www.cl.cam.ac.uk/research/dtg/attarchive/facedatabase.html