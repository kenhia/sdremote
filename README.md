# sdremote

Stream Deck remote launcher

Playing with an idea a colleague had for making use of the Stream Deck
connected to a Windows machine (where it's supported nicely) to launch
things on my Linux daily driver where Steam Deck is not supported.

From the Linux side, I just build (if needed) then start the sdremote server:
```bash
λ  go build .
λ  ./sdremote 
```

From the Windows side where I'm running the Stream Deck, I'm using
[BarRaider's API Ninja plugin](https://barraider.com/#plugins). I was
initially using GET, but found it much easier to pass URLs and other things
that have characters that needed to be encoded with POST...no encoding needed.
