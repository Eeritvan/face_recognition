# määrittelydokumentti
opinto-ohjelma: tietojenkäsittelytieteen kandidaatti (TKT)

## aihe
ydin: koneoppiminen

Eigenface-algoritmilla toteutettu kasvojentunnistusohjelma. Käytän aineistona AT&T "The Database of Faces" tietokantaa joka sisältää 40 henkilöstä kymmenen 92x112 kokoista mustavalkokuvaa. 

Projektissa käytetään PCA, kovarianssimatriiseja ja ominaisvektoreita joilla kuvien pääpiirteet voidaan säilyttää ja käyttää kasvojentunnistukseen. Ominaisarvot ja ominaisvektorit voidaan selvittää esim. power iteration tai QR-algoritmillä. Projekti hyödyntää paljon matriisilaskentaa joiden operaatioille toteutan algoritmit itse. Aikavaativuudeltaan esimerkiksi kovarianssimatriisin on O(m^2*n), missä n on kuvan pikselien määrä (92×112) ja m on treenauskuvien lukumäärä. QR-hajotelman aikavaativuus on O(m^3) m×m kokoisille kovarianssimatriiseille, missä n on treenauskuvien määrä. Myös matriisitulon aikavaativuus on O(n^3) kun matriisin koko on n×n.

## ohjelmointi
Toteutan ohjelman go-kielellä. Pyrin käyttämään mahdollisimman paljon go:n standardikirjastoja yksinkertaisuuden vuoksi. Pystyn vertaisarvioimaan myös muita ohjelmointikieliä.

## dokumentaatio
Kirjoitan dokumentaation ja viikkopalautukset suomeksi mutta koodin, testit ja commit-viestit englanniksi

## lähteet (voi muuttua)
- https://en.wikipedia.org/wiki/Eigenface
- https://en.wikipedia.org/wiki/Eigenvalue_algorithm
- https://en.wikipedia.org/wiki/QR_algorithm
- https://km.pcz.pl/konferencja/dokumenty/MMFT2017/Caban%20L.pdf
- https://www.geeksforgeeks.org/ml-face-recognition-using-eigenfaces-pca-algorithm/
- https://www.cl.cam.ac.uk/research/dtg/attarchive/facedatabase.html
- https://www.youtube.com/watch?v=n0zDgkbFyQk