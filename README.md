# bus-eta-hk

KMB Only at the moment.


Query | Description | Required | Example
--- | --- | --- | ---
latlng | Lat Lng | Y | 22.336540,114.160980
r | Scan Range in Km. Max 1.0. Default 0.5 | N | 0.5
route | Bus Route | N | 905
dir | Required field when `route` is set | N | `I` or `O`


```
{URL}/stops?latlng=22.336540,114.160980&r=0.5&route=905&dir=I
```