# määrittelydokumentti
opinto-ohjelma: tietojenkäsittelytieteen kandidaatti (TKT)

## aihe
ydin: koneoppiminen

Eigenface-algoritmilla toteutettu kasvojentunnistusohjelma. Käytän aineistona AT&T "The Database of Faces" tietokantaa joka sisältää 40 henkilöstä kymmenen 92x112 kokoista mustavalkokuvaa. 

Projektissa käytetään PCA, kovarianssimatriiseja ja ominaisvektoreita joilla kuvien pääpiirteet voidaan säilyttää ja käyttää kasvojentunnistukseen. Ominaisarvot ja ominaisvektorit voidaan selvittää esim. power iteration tai QR-algoritmillä (päätän hiukan myöhemmin kumman toteutan). Projekti hyödyntää paljon matriisilaskentaa joiden operaatioille toteutan algoritmit itse. Näistä algoritmeista kenties hitain aikavaativuudeltaan on matriisitulo jonka vaativuus on O(n^3).

## ohjelmointi
Toteutan ohjelman go-kielellä. Pyrin käyttämään mahdollisimman paljon go:n standardikirjastoja yksinkertaisuuden vuoksi. Pystyn vertaisarvioimaan myös muita ohjelmointikieliä.

## dokumentaatio
Kirjoitan dokumentaation ja viikkopalautukset suomeksi mutta koodin, testit ja commit-viestit englanniksi

## lähteet (voi muuttua)
https://en.wikipedia.org/wiki/Eigenface
https://en.wikipedia.org/wiki/Eigenvalue_algorithm
https://www.cl.cam.ac.uk/research/dtg/attarchive/facedatabase.html