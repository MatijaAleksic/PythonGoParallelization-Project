# Python - Paralelizacija(Python&Go)

## Double Pendulum Simulation

### Problem
Problem koji rešavamo jeste simulacija duplog klatna. Problem ćemo proširiti sa dodavanjem dodatnih duplih klatna koji se simultano simuliraju kako bi mogli meriti 
performanse paralelizacije. Broj klatana koji ćemo pokušati da simuliramo simultano je 30, ali ovaj broj nije fiksan jel zavisi od performanse racunara na kojem radimo eksperiment kako bi dobili znacajnije rezultate prilikom eksperimenata za paralelizaciju programa i jakog i slabog skaliranja. Simulacija ce trajati nekoliko minuta koja ce biti po jedinici 30 frejmova po sekundi.

![double-pendulum](https://user-images.githubusercontent.com/34009136/131214128-3c4bb9fe-5db6-439f-8f46-fc1c1d026256.gif)

### Resenje
Sekvencijalna i paralelizovana verzija rešenja ce biti odradjene u programskim jezicima Python i Go.
Paralelizacija ce biti odradjena koristeci **multiprocessing** paketa u Python-u i **goroutine** paketa u programskom jeziku GO.

### Skaliranje
Jako skaliranje ce biti sprovedeno sa fiksnim brojem klatna koji ce biti rasporedjeni na procese paralelnog programiranja i meriti njihovo ubrzanje.

Slabo skaliranje ce biti sprovedeno sa dodavanjem dodatnih klatna koji ce ravnomerno biti rasporedjeni na procese paralelnog programiranja i meriti njihovo ubrzanje.

### Vizualizacija rešenja
Svaka iteracija simulacije biće čuvana u poseban fajl kao pozicije klatna u svakom monentu, nakon cega ce vizuelizacija biti odradjeno na Pharo okruzenju kao animacija.
