# justify
Create nicely formatted text outputs where lines have uniform width, and margins are aligned.

### Example:
The `fmt` tool transforms the text this way:
```
Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed pretium,
velit vel rhoncus tempus, urna sem pellentesque urna, id laoreet mauris
lorem eget purus. Donec maximus urna et odio semper, nec fermentum enim
congue. Vivamus eleifend fringilla ligula ut tincidunt. Vestibulum ante
ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae;
Nullam egestas justo a turpis ullamcorper dapibus. Nunc dictum ligula
a risus pellentesque, et blandit ligula convallis.

Proin tempor justo et ante pulvinar rhoncus. Curabitur sed efficitur
elit. Donec sem mauris, malesuada id iaculis ac, aliquet vel enim. Donec
semper volutpat metus, nec laoreet est. Fusce vestibulum facilisis dui,
ut sollicitudin eros vestibulum in. Quisque lacinia orci in volutpat
dignissim.
```

With `justify` will return
```
Lorem  ipsum  dolor  sit  amet, consectetur adipiscing elit.
Sed   pretium,   velit   vel   rhoncus   tempus,   urna  sem
pellentesque  urna,  id  laoreet  mauris  lorem  eget purus.
Donec  maximus  urna  et  odio  semper,  nec  fermentum enim
congue.  Vivamus  eleifend  fringilla  ligula  ut tincidunt.
Vestibulum  ante  ipsum  primis  in  faucibus orci luctus et
ultrices  posuere  cubilia  curae;  Nullam  egestas  justo a
turpis  ullamcorper  dapibus.  Nunc  dictum  ligula  a risus
pellentesque, et blandit ligula convallis.

Proin  tempor  justo et ante pulvinar rhoncus. Curabitur sed
efficitur  elit.  Donec sem mauris, malesuada id iaculis ac,
aliquet  vel  enim. Donec semper volutpat metus, nec laoreet
est.  Fusce  vestibulum  facilisis dui, ut sollicitudin eros
vestibulum  in.  Quisque lacinia orci in volutpat dignissim.
```
